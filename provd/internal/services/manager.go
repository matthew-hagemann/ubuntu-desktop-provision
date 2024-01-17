// Package services mediates all the business logic of the application via a manager.
package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/canonical/ubuntu-desktop-provision/provd"
	"github.com/canonical/ubuntu-desktop-provision/provd/internal/services/hello"
	"github.com/canonical/ubuntu-desktop-provision/provd/internal/services/locale"
	"github.com/canonical/ubuntu-desktop-provision/provd/internal/services/user"
	pb "github.com/canonical/ubuntu-desktop-provision/provd/protos"
	"github.com/godbus/dbus/v5"
	"github.com/ubuntu/decorate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Manager mediates the whole business logic of the application.
type Manager struct {
	helloService hello.Service
	userService  user.Service
    localeService locale.Service
	bus          *dbus.Conn
}

type dbusConnectionAdapter struct {
	*dbus.Conn
}

type dbusConnectionAdapter2 struct {
	*dbus.Conn
}

func (bus dbusConnectionAdapter) Object(iface string, path dbus.ObjectPath) user.DbusObject {
	return bus.Conn.Object(iface, path)
}

func (bus dbusConnectionAdapter2) Object(iface string, path dbus.ObjectPath) locale.DbusObject {
	return bus.Conn.Object(iface, path)
}

// NewManager returns a new manager after creating all necessary items for our business logic.
func NewManager(ctx context.Context) (m *Manager, err error) {
	defer decorate.OnError(&err, "can't create provd object")

	bus, err := dbus.ConnectSystemBus(
		dbus.WithIncomingInterceptor(func(msg *dbus.Message) {
			slog.Debug(fmt.Sprintf("DBUS: %s", msg))
		}))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to connect to system bus: %s", err)
	}

	helloService := hello.Service{}

	userService := user.New(dbusConnectionAdapter{bus})

    localeService := locale.New(dbusConnectionAdapter2{bus})

	return &Manager{
		helloService: helloService,
		userService:  *userService,
        localeService: *localeService,
		bus:          bus,
	}, nil
}

// RegisterGRPCServices returns a new grpc Server.
func (m Manager) RegisterGRPCServices(ctx context.Context) *grpc.Server {
	slog.Debug("Registering GRPC services")

	grpcServer := grpc.NewServer()

	provd.RegisterHelloWorldServiceServer(grpcServer, &m.helloService)
	pb.RegisterUserServiceServer(grpcServer, &m.userService)
    pb.RegisterLocaleServiceServer(grpcServer, &m.localeService)
	return grpcServer
}

// Stop closes the dbus connection.
func (m *Manager) Stop() error {
	slog.Debug("Closing grpc manager and dbus connection")

	return m.bus.Close()
}
