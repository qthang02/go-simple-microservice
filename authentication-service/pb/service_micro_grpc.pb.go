// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: service_micro.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MicroGo_CreateUser_FullMethodName = "/pb.MicroGo/CreateUser"
	MicroGo_LoginUser_FullMethodName  = "/pb.MicroGo/LoginUser"
)

// MicroGoClient is the client API for MicroGo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroGoClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type microGoClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroGoClient(cc grpc.ClientConnInterface) MicroGoClient {
	return &microGoClient{cc}
}

func (c *microGoClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, MicroGo_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microGoClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, MicroGo_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroGoServer is the server API for MicroGo service.
// All implementations must embed UnimplementedMicroGoServer
// for forward compatibility
type MicroGoServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedMicroGoServer()
}

// UnimplementedMicroGoServer must be embedded to have forward compatible implementations.
type UnimplementedMicroGoServer struct {
}

func (UnimplementedMicroGoServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMicroGoServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedMicroGoServer) mustEmbedUnimplementedMicroGoServer() {}

// UnsafeMicroGoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroGoServer will
// result in compilation errors.
type UnsafeMicroGoServer interface {
	mustEmbedUnimplementedMicroGoServer()
}

func RegisterMicroGoServer(s grpc.ServiceRegistrar, srv MicroGoServer) {
	s.RegisterService(&MicroGo_ServiceDesc, srv)
}

func _MicroGo_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroGoServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroGo_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroGoServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroGo_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroGoServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroGo_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroGoServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MicroGo_ServiceDesc is the grpc.ServiceDesc for MicroGo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MicroGo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MicroGo",
	HandlerType: (*MicroGoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _MicroGo_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _MicroGo_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_micro.proto",
}
