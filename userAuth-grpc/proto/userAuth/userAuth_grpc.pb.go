// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: userAuth/userAuth.proto

package userAuth

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

// UserAuthClient is the client API for UserAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAuthClient interface {
	GetOTP(ctx context.Context, in *GetOTPRequest, opts ...grpc.CallOption) (*GetOTPResponse, error)
}

type userAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAuthClient(cc grpc.ClientConnInterface) UserAuthClient {
	return &userAuthClient{cc}
}

func (c *userAuthClient) GetOTP(ctx context.Context, in *GetOTPRequest, opts ...grpc.CallOption) (*GetOTPResponse, error) {
	out := new(GetOTPResponse)
	err := c.cc.Invoke(ctx, "/userAuth.UserAuth/GetOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAuthServer is the server API for UserAuth service.
// All implementations should embed UnimplementedUserAuthServer
// for forward compatibility
type UserAuthServer interface {
	GetOTP(context.Context, *GetOTPRequest) (*GetOTPResponse, error)
}

// UnimplementedUserAuthServer should be embedded to have forward compatible implementations.
type UnimplementedUserAuthServer struct {
}

func (UnimplementedUserAuthServer) GetOTP(context.Context, *GetOTPRequest) (*GetOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOTP not implemented")
}

// UnsafeUserAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAuthServer will
// result in compilation errors.
type UnsafeUserAuthServer interface {
	mustEmbedUnimplementedUserAuthServer()
}

func RegisterUserAuthServer(s grpc.ServiceRegistrar, srv UserAuthServer) {
	s.RegisterService(&UserAuth_ServiceDesc, srv)
}

func _UserAuth_GetOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).GetOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userAuth.UserAuth/GetOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).GetOTP(ctx, req.(*GetOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAuth_ServiceDesc is the grpc.ServiceDesc for UserAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userAuth.UserAuth",
	HandlerType: (*UserAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOTP",
			Handler:    _UserAuth_GetOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userAuth/userAuth.proto",
}