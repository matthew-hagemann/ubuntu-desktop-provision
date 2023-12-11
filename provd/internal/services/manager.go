// Package services mediates all the business logic of the application via a manager.
package services

import (
	"context"
	"log/slog"

	"github.com/canonical/ubuntu-desktop-provision/provd"
	"github.com/canonical/ubuntu-desktop-provision/provd/internal/services/hello"
	"github.com/canonical/ubuntu-desktop-provision/provd/internal/services/user"
	proto "github.com/canonical/ubuntu-desktop-provision/provd/proto"
	"github.com/ubuntu/decorate"
	"google.golang.org/grpc"
)

// Manager mediates the whole business logic of the application.
type Manager struct {
	helloService hello.Service
	userService  user.Service
}

// NewManager returns a new manager after creating all necessary items for our business logic.
func NewManager(ctx context.Context) (m Manager, err error) {
	defer decorate.OnError(&err, "can't create provd object")

	slog.Debug("Building provd object")

	helloService := hello.Service{}
	userService := user.Service{}

	return Manager{
		helloService: helloService,
		userService:  userService,
	}, nil
}

// RegisterGRPCServices returns a new grpc Server.
func (m Manager) RegisterGRPCServices(ctx context.Context) *grpc.Server {
	slog.Debug("Registering GRPC services")

	grpcServer := grpc.NewServer()

	provd.RegisterHelloWorldServiceServer(grpcServer, &m.helloService)
	proto.RegisterUserServiceServer(grpcServer, &m.userService)
	return grpcServer
}
