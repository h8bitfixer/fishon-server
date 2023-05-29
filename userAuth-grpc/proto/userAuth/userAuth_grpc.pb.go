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
	VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error)
	GetUserAccountByPhone(ctx context.Context, in *GetUserAccountByPhoneRequest, opts ...grpc.CallOption) (*GetUserAccountByPhoneResponse, error)
	GetTokenByPhone(ctx context.Context, in *GetTokenByPhoneRequest, opts ...grpc.CallOption) (*GetTokenByPhoneResponse, error)
	GetTokenByUserID(ctx context.Context, in *GetTokenByUserIDRequest, opts ...grpc.CallOption) (*GetTokenByUserIDResponse, error)
	VerifyUserEmailAndPassword(ctx context.Context, in *VerifyUserEmailAndPasswordRequest, opts ...grpc.CallOption) (*GetUserAccountByPhoneResponse, error)
	CreateUserByEmail(ctx context.Context, in *CreateUserByEmailRequest, opts ...grpc.CallOption) (*CreateUserByEmailResponse, error)
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

func (c *userAuthClient) VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*VerifyOTPResponse, error) {
	out := new(VerifyOTPResponse)
	err := c.cc.Invoke(ctx, "/userAuth.UserAuth/VerifyOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthClient) GetUserAccountByPhone(ctx context.Context, in *GetUserAccountByPhoneRequest, opts ...grpc.CallOption) (*GetUserAccountByPhoneResponse, error) {
	out := new(GetUserAccountByPhoneResponse)
	err := c.cc.Invoke(ctx, "/userAuth.UserAuth/GetUserAccountByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthClient) GetTokenByPhone(ctx context.Context, in *GetTokenByPhoneRequest, opts ...grpc.CallOption) (*GetTokenByPhoneResponse, error) {
	out := new(GetTokenByPhoneResponse)
	err := c.cc.Invoke(ctx, "/userAuth.UserAuth/GetTokenByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthClient) GetTokenByUserID(ctx context.Context, in *GetTokenByUserIDRequest, opts ...grpc.CallOption) (*GetTokenByUserIDResponse, error) {
	out := new(GetTokenByUserIDResponse)
	err := c.cc.Invoke(ctx, "/userAuth.UserAuth/GetTokenByUserID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthClient) VerifyUserEmailAndPassword(ctx context.Context, in *VerifyUserEmailAndPasswordRequest, opts ...grpc.CallOption) (*GetUserAccountByPhoneResponse, error) {
	out := new(GetUserAccountByPhoneResponse)
	err := c.cc.Invoke(ctx, "/userAuth.UserAuth/VerifyUserEmailAndPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthClient) CreateUserByEmail(ctx context.Context, in *CreateUserByEmailRequest, opts ...grpc.CallOption) (*CreateUserByEmailResponse, error) {
	out := new(CreateUserByEmailResponse)
	err := c.cc.Invoke(ctx, "/userAuth.UserAuth/CreateUserByEmail", in, out, opts...)
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
	VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error)
	GetUserAccountByPhone(context.Context, *GetUserAccountByPhoneRequest) (*GetUserAccountByPhoneResponse, error)
	GetTokenByPhone(context.Context, *GetTokenByPhoneRequest) (*GetTokenByPhoneResponse, error)
	GetTokenByUserID(context.Context, *GetTokenByUserIDRequest) (*GetTokenByUserIDResponse, error)
	VerifyUserEmailAndPassword(context.Context, *VerifyUserEmailAndPasswordRequest) (*GetUserAccountByPhoneResponse, error)
	CreateUserByEmail(context.Context, *CreateUserByEmailRequest) (*CreateUserByEmailResponse, error)
}

// UnimplementedUserAuthServer should be embedded to have forward compatible implementations.
type UnimplementedUserAuthServer struct {
}

func (UnimplementedUserAuthServer) GetOTP(context.Context, *GetOTPRequest) (*GetOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOTP not implemented")
}
func (UnimplementedUserAuthServer) VerifyOTP(context.Context, *VerifyOTPRequest) (*VerifyOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOTP not implemented")
}
func (UnimplementedUserAuthServer) GetUserAccountByPhone(context.Context, *GetUserAccountByPhoneRequest) (*GetUserAccountByPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAccountByPhone not implemented")
}
func (UnimplementedUserAuthServer) GetTokenByPhone(context.Context, *GetTokenByPhoneRequest) (*GetTokenByPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenByPhone not implemented")
}
func (UnimplementedUserAuthServer) GetTokenByUserID(context.Context, *GetTokenByUserIDRequest) (*GetTokenByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenByUserID not implemented")
}
func (UnimplementedUserAuthServer) VerifyUserEmailAndPassword(context.Context, *VerifyUserEmailAndPasswordRequest) (*GetUserAccountByPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUserEmailAndPassword not implemented")
}
func (UnimplementedUserAuthServer) CreateUserByEmail(context.Context, *CreateUserByEmailRequest) (*CreateUserByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserByEmail not implemented")
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

func _UserAuth_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userAuth.UserAuth/VerifyOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).VerifyOTP(ctx, req.(*VerifyOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuth_GetUserAccountByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAccountByPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).GetUserAccountByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userAuth.UserAuth/GetUserAccountByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).GetUserAccountByPhone(ctx, req.(*GetUserAccountByPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuth_GetTokenByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenByPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).GetTokenByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userAuth.UserAuth/GetTokenByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).GetTokenByPhone(ctx, req.(*GetTokenByPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuth_GetTokenByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).GetTokenByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userAuth.UserAuth/GetTokenByUserID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).GetTokenByUserID(ctx, req.(*GetTokenByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuth_VerifyUserEmailAndPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserEmailAndPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).VerifyUserEmailAndPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userAuth.UserAuth/VerifyUserEmailAndPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).VerifyUserEmailAndPassword(ctx, req.(*VerifyUserEmailAndPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuth_CreateUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthServer).CreateUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userAuth.UserAuth/CreateUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthServer).CreateUserByEmail(ctx, req.(*CreateUserByEmailRequest))
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
		{
			MethodName: "VerifyOTP",
			Handler:    _UserAuth_VerifyOTP_Handler,
		},
		{
			MethodName: "GetUserAccountByPhone",
			Handler:    _UserAuth_GetUserAccountByPhone_Handler,
		},
		{
			MethodName: "GetTokenByPhone",
			Handler:    _UserAuth_GetTokenByPhone_Handler,
		},
		{
			MethodName: "GetTokenByUserID",
			Handler:    _UserAuth_GetTokenByUserID_Handler,
		},
		{
			MethodName: "VerifyUserEmailAndPassword",
			Handler:    _UserAuth_VerifyUserEmailAndPassword_Handler,
		},
		{
			MethodName: "CreateUserByEmail",
			Handler:    _UserAuth_CreateUserByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userAuth/userAuth.proto",
}
