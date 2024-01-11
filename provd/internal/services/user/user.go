// Package user implements the User gRPC service.
package user

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"

	pb "github.com/canonical/ubuntu-desktop-provision/provd/protos"
	"github.com/godbus/dbus/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	// DbusAccountsPrefix is the prefix for the Accounts D-Bus interface.
	DbusAccountsPrefix = "org.freedesktop.Accounts"
	// DbusHostnamePrefix is the prefix for the Hostname D-Bus interface.
	DbusHostnamePrefix = "org.freedesktop.hostname1"
	// UsernameMaxLen is the maximum length of a username.
	UsernameMaxLen = 32
	// UsernameRegex is the regex for a valid username.
	UsernameRegex = "^[a-z_][a-z0-9_-]*$"
)

// DbusObject is an abstraction of a dbus object.
type DbusObject interface {
	Call(method string, flags dbus.Flags, args ...interface{}) *dbus.Call
	GetProperty(p string) (dbus.Variant, error)
}

// ObjectFactory is a factory for creating DbusObjects for users.
// This is used to allow mocking of the DbusObjects for testing.
type ObjectFactory interface {
	GetUserObject(userObjectPath dbus.ObjectPath) DbusObject
}

// objectFactoryImpl is the default implementation of UserObjectFactory.
type objectFactoryImpl struct {
	conn *dbus.Conn
}

// Wraps the dbus call for getting a user object.
func (f *objectFactoryImpl) GetUserObject(userObjectPath dbus.ObjectPath) DbusObject {
	userObject := f.conn.Object(DbusAccountsPrefix, userObjectPath)
	return userObject
}

// options are the configurable functional options of the User service.
type options struct {
	accounts    DbusObject
	hostname    DbusObject
	userFactory ObjectFactory
}

// Option type exported for tests.
type Option func(*options)

// Service is the implementation of the User module service.
type Service struct {
	pb.UnimplementedUserServiceServer
	accounts    DbusObject
	hostname    DbusObject
	userFactory ObjectFactory
}

// New returns a new instance of the User service.
func New(bus *dbus.Conn, args ...Option) *Service {
	accounts := bus.Object(DbusAccountsPrefix, "/org/freedesktop/Accounts")
	hostname := bus.Object(DbusHostnamePrefix, "/org/freedesktop/hostname1")
	userFactory := &objectFactoryImpl{conn: bus}

	opts := options{
		accounts:    accounts,
		hostname:    hostname,
		userFactory: userFactory,
	}

	// Apply given options
	for _, f := range args {
		f(&opts)
	}

	return &Service{
		accounts:    opts.accounts,
		hostname:    opts.hostname,
		userFactory: opts.userFactory,
	}
}

const validSaltChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789./"

// generateSalt generates a salt string of the specified length using validSalts characters.
func generateSalt(length int) (string, error) {
	salt := make([]byte, length)
	for i := range salt {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(validSaltChars))))
		if err != nil {
			return "", err
		}
		salt[i] = validSaltChars[index.Int64()]
	}
	return string(salt), nil
}

// HashPassword hashes the given password, returning in the SHA-512 crypt format.
// A salt can be provided for testing purposes.
func HashPassword(password string, testSalt *string) (string, error) {
	if password == "" {
		return "", status.Errorf(codes.InvalidArgument, "received an empty password")
	}
	var salt string
	var err error

	if testSalt != nil {
		salt = *testSalt
	} else {
		salt, err = generateSalt(16)
		if err != nil {
			return "", err
		}
	}
	hasher := sha512.New()
	hasher.Write([]byte(salt + password))
	hashedBytes := hasher.Sum(nil)
	hashedPassword := base64.RawStdEncoding.EncodeToString(hashedBytes)

	return fmt.Sprintf("$6$%s$%s", salt, hashedPassword), nil
}

// CreateUser creates a new user on the system.
func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	// Validate requtest
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "received a nil request")
	}
	user := req.GetUser()
	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, "received a nil user")
	}
	username := user.GetUsername()
	if username == "" {
		return nil, status.Errorf(codes.InvalidArgument, "received an empty username")
	}
	realName := user.GetRealName()
	if realName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "received an empty realName")
	}
	isAdmin := req.GetIsAdmin()
	var accountType int32
	if isAdmin {
		accountType = 1
	} else {
		accountType = 0
	}
	password := user.GetPassword() // allow empty password?
	autologin := user.GetAutoLogin()
	hostname := user.GetHostname()
	if hostname == "" {
		return nil, status.Errorf(codes.InvalidArgument, "received an empty hostname")
	}
	// Create the user
	var userObjectPath dbus.ObjectPath
	call := s.accounts.Call(DbusAccountsPrefix+".CreateUser", 0, username, realName, accountType)

	err := call.Store(&userObjectPath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	hashed, err := HashPassword(password, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate hashed password: %s", err)
	}

	// Set the password for the user
	userObject := s.userFactory.GetUserObject(userObjectPath)
	err = userObject.Call(DbusAccountsPrefix+".User.SetPassword", 0, hashed, "").Err
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set password: %s", err)
	}

	// Set autologin for the user
	err = userObject.Call(DbusAccountsPrefix+".User.SetAutomaticLogin", 0, autologin).Err
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set autologin: %s", err)
	}

	// Set the hostname
	err = s.hostname.Call(DbusHostnamePrefix+".SetStaticHostname", 0, hostname, false).Err
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set hostname: %s", err)
	}

	return &emptypb.Empty{}, nil
}

// ValidateUsername validates the given username. Returns an enum value indicating
// the result of the validation.
func (s *Service) ValidateUsername(ctx context.Context, req *pb.ValidateUsernameRequest) (*pb.ValidateUsernameResponse, error) {
	// Validate request
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "received a nil request")
	}
	username := req.GetUsername()
	if username == "" {
		return &pb.ValidateUsernameResponse{UsernameValidation: pb.UsernameValidation_EMPTY}, nil
	}

	// Check if username uses valid characters
	matched, _ := regexp.MatchString(UsernameRegex, username)
	if !matched {
		return &pb.ValidateUsernameResponse{UsernameValidation: pb.UsernameValidation_INVALID_CHARS}, nil
	}

	// Check if username is too long
	if len(username) > UsernameMaxLen {
		return &pb.ValidateUsernameResponse{UsernameValidation: pb.UsernameValidation_TOO_LONG}, nil
	}

	// Check if username is in reserved list
	// Read line by line to avoid loading the whole file into memory
	file, err := os.Open("reserved-usernames")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error opening reserved usernames file: %v", err)
	}
	defer file.Close()

	isReserved := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ignore comment lines
		if strings.HasPrefix(line, "#") {
			continue
		}
		if line == username {
			isReserved = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, status.Errorf(codes.Internal, "error reading reserved usernames file: %v", err)
	}

	if isReserved {
		return &pb.ValidateUsernameResponse{UsernameValidation: pb.UsernameValidation_SYSTEM_RESERVED}, nil
	}

	// Check if username is already in use
	err = s.accounts.Call(DbusAccountsPrefix+".FindUserByName", 0, username).Err
	if err != nil {
		var dbusError dbus.Error
		if errors.As(err, &dbusError) {
			if dbusError.Name == DbusAccountsPrefix+".Error.Failed" {
				// User not found
				return &pb.ValidateUsernameResponse{UsernameValidation: pb.UsernameValidation_OK}, nil
			}
			// Handle other dbus errors
			return nil, status.Errorf(codes.Internal, "dbus error: %v", dbusError)
		}
		// Unknown error
		return nil, status.Errorf(codes.Internal, "unknown error: %v", err)
	}

	// User found
	return &pb.ValidateUsernameResponse{UsernameValidation: pb.UsernameValidation_ALREADY_IN_USE}, nil
}
