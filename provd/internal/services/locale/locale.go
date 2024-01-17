// Package locale implements the Locale gRPC service.
package locale

import (
	"context"

	pb "github.com/canonical/ubuntu-desktop-provision/provd/protos"
	"github.com/godbus/dbus/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DbusObject is an abstraction of a dbus object.
type DbusObject interface {
	Call(method string, flags dbus.Flags, args ...interface{}) *dbus.Call
	GetProperty(p string) (dbus.Variant, error)
}

// DbusConnector is the interface to the dbus connection which allow easy mocking.
type DbusConnector interface {
	Object(dest string, path dbus.ObjectPath) DbusObject
}

// Service is the implementation of the User module service.
type Service struct {
	pb.UnimplementedLocaleServiceServer
	Conn   DbusConnector
	Locale DbusObject
}

// New returns a new instance of the Locale service.
func New(Conn DbusConnector) *Service {
    return &Service{
        Conn:   Conn,
        Locale: Conn.Object("org.freedesktop.locale1", "/org/freedesktop/locale1"),
    }
}

func (s *Service) GetLocale(ctx context.Context, req *emptypb.Empty) (*pb.GetLocaleResponse, error) {
	// Validate request
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "received a nil request")
	}
    // Get the locale from the dbus object properties.
    locale, err := s.Locale.GetProperty("org.freedesktop.locale1." + "Locale")
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to get locale: %v", err)
    }
    return &pb.GetLocaleResponse{Locale: locale.String()}, nil
}

func (s *Service) SetLocale(ctx context.Context, req *pb.SetLocaleRequest) (*emptypb.Empty, error) {
    // Validate request
    if req == nil {
        return nil, status.Errorf(codes.InvalidArgument, "received a nil request")
    }
    if req.Locale == "" {
        return nil, status.Errorf(codes.InvalidArgument, "locale must be specified")
    }

	//err = userObject.Call(consts.DbusUserPrefix+".SetPassword", 0, hashed, "").Err
    // Set the locale using the dbus object.
    err := s.Locale.Call("org.freedesktop.locale1." + "SetLocale", 0, []string{req.Locale}, false).Err
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to set locale: %v", err)
    }
    return &emptypb.Empty{}, nil
}
