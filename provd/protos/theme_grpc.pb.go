// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.4
// source: protos/theme.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ThemeService_SetTheme_FullMethodName  = "/theme.ThemeService/SetTheme"
	ThemeService_GetTheme_FullMethodName  = "/theme.ThemeService/GetTheme"
	ThemeService_SetAccent_FullMethodName = "/theme.ThemeService/SetAccent"
	ThemeService_GetAccent_FullMethodName = "/theme.ThemeService/GetAccent"
)

// ThemeServiceClient is the client API for ThemeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThemeServiceClient interface {
	SetTheme(ctx context.Context, in *SetThemeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTheme(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetThemeResponse, error)
	SetAccent(ctx context.Context, in *SetAccentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAccent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAccentResponse, error)
}

type themeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThemeServiceClient(cc grpc.ClientConnInterface) ThemeServiceClient {
	return &themeServiceClient{cc}
}

func (c *themeServiceClient) SetTheme(ctx context.Context, in *SetThemeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ThemeService_SetTheme_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themeServiceClient) GetTheme(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetThemeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetThemeResponse)
	err := c.cc.Invoke(ctx, ThemeService_GetTheme_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themeServiceClient) SetAccent(ctx context.Context, in *SetAccentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ThemeService_SetAccent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themeServiceClient) GetAccent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAccentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAccentResponse)
	err := c.cc.Invoke(ctx, ThemeService_GetAccent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThemeServiceServer is the server API for ThemeService service.
// All implementations must embed UnimplementedThemeServiceServer
// for forward compatibility.
type ThemeServiceServer interface {
	SetTheme(context.Context, *SetThemeRequest) (*emptypb.Empty, error)
	GetTheme(context.Context, *emptypb.Empty) (*GetThemeResponse, error)
	SetAccent(context.Context, *SetAccentRequest) (*emptypb.Empty, error)
	GetAccent(context.Context, *emptypb.Empty) (*GetAccentResponse, error)
	mustEmbedUnimplementedThemeServiceServer()
}

// UnimplementedThemeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedThemeServiceServer struct{}

func (UnimplementedThemeServiceServer) SetTheme(context.Context, *SetThemeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTheme not implemented")
}
func (UnimplementedThemeServiceServer) GetTheme(context.Context, *emptypb.Empty) (*GetThemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTheme not implemented")
}
func (UnimplementedThemeServiceServer) SetAccent(context.Context, *SetAccentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccent not implemented")
}
func (UnimplementedThemeServiceServer) GetAccent(context.Context, *emptypb.Empty) (*GetAccentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccent not implemented")
}
func (UnimplementedThemeServiceServer) mustEmbedUnimplementedThemeServiceServer() {}
func (UnimplementedThemeServiceServer) testEmbeddedByValue()                      {}

// UnsafeThemeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThemeServiceServer will
// result in compilation errors.
type UnsafeThemeServiceServer interface {
	mustEmbedUnimplementedThemeServiceServer()
}

func RegisterThemeServiceServer(s grpc.ServiceRegistrar, srv ThemeServiceServer) {
	// If the following call pancis, it indicates UnimplementedThemeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ThemeService_ServiceDesc, srv)
}

func _ThemeService_SetTheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemeServiceServer).SetTheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThemeService_SetTheme_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemeServiceServer).SetTheme(ctx, req.(*SetThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThemeService_GetTheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemeServiceServer).GetTheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThemeService_GetTheme_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemeServiceServer).GetTheme(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThemeService_SetAccent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAccentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemeServiceServer).SetAccent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThemeService_SetAccent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemeServiceServer).SetAccent(ctx, req.(*SetAccentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThemeService_GetAccent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemeServiceServer).GetAccent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThemeService_GetAccent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemeServiceServer).GetAccent(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ThemeService_ServiceDesc is the grpc.ServiceDesc for ThemeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThemeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "theme.ThemeService",
	HandlerType: (*ThemeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetTheme",
			Handler:    _ThemeService_SetTheme_Handler,
		},
		{
			MethodName: "GetTheme",
			Handler:    _ThemeService_GetTheme_Handler,
		},
		{
			MethodName: "SetAccent",
			Handler:    _ThemeService_SetAccent_Handler,
		},
		{
			MethodName: "GetAccent",
			Handler:    _ThemeService_GetAccent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/theme.proto",
}
