// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: internal/uam/service/auth/auth.proto

package auth

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Login(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	VerifyToken(ctx context.Context, in *AuthVerifyTokenRequest, opts ...grpc.CallOption) (*AuthVerifyTokenResponse, error)
	RestEnforce(ctx context.Context, in *AuthEnforceRequest, opts ...grpc.CallOption) (*AuthEnforceResponse, error)
	RbacEnforce(ctx context.Context, in *AuthEnforceRequest, opts ...grpc.CallOption) (*AuthEnforceResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) VerifyToken(ctx context.Context, in *AuthVerifyTokenRequest, opts ...grpc.CallOption) (*AuthVerifyTokenResponse, error) {
	out := new(AuthVerifyTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RestEnforce(ctx context.Context, in *AuthEnforceRequest, opts ...grpc.CallOption) (*AuthEnforceResponse, error) {
	out := new(AuthEnforceResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/RestEnforce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RbacEnforce(ctx context.Context, in *AuthEnforceRequest, opts ...grpc.CallOption) (*AuthEnforceResponse, error) {
	out := new(AuthEnforceResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/RbacEnforce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	Login(context.Context, *AuthRequest) (*AuthResponse, error)
	VerifyToken(context.Context, *AuthVerifyTokenRequest) (*AuthVerifyTokenResponse, error)
	RestEnforce(context.Context, *AuthEnforceRequest) (*AuthEnforceResponse, error)
	RbacEnforce(context.Context, *AuthEnforceRequest) (*AuthEnforceResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Login(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) VerifyToken(context.Context, *AuthVerifyTokenRequest) (*AuthVerifyTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedAuthServer) RestEnforce(context.Context, *AuthEnforceRequest) (*AuthEnforceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestEnforce not implemented")
}
func (UnimplementedAuthServer) RbacEnforce(context.Context, *AuthEnforceRequest) (*AuthEnforceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RbacEnforce not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthVerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifyToken(ctx, req.(*AuthVerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RestEnforce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthEnforceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RestEnforce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/RestEnforce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RestEnforce(ctx, req.(*AuthEnforceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RbacEnforce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthEnforceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RbacEnforce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/RbacEnforce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RbacEnforce(ctx, req.(*AuthEnforceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _Auth_VerifyToken_Handler,
		},
		{
			MethodName: "RestEnforce",
			Handler:    _Auth_RestEnforce_Handler,
		},
		{
			MethodName: "RbacEnforce",
			Handler:    _Auth_RbacEnforce_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/uam/service/auth/auth.proto",
}
