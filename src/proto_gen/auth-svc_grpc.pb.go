// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.9.1
// source: auth-svc.proto

package proto_gen

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
	GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Health, error)
	CreateCustomerAnonymousToken(ctx context.Context, in *AnonymousDto, opts ...grpc.CallOption) (*AuthResultDto, error)
	CreateSellerAnonymousToken(ctx context.Context, in *AnonymousDto, opts ...grpc.CallOption) (*AuthResultDto, error)
	CreateAdminAnonymousToken(ctx context.Context, in *AnonymousDto, opts ...grpc.CallOption) (*AuthResultDto, error)
	ValidateToken(ctx context.Context, in *ValidateTokenDto, opts ...grpc.CallOption) (*Subject, error)
	// Customer
	RegisterUserAuth(ctx context.Context, in *RegisterDto, opts ...grpc.CallOption) (*AuthResultDto, error)
	IsUserEmailExists(ctx context.Context, in *UserAuthDto, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	ValidateUserAccount(ctx context.Context, in *RegisterDto, opts ...grpc.CallOption) (*AuthResultDto, error)
	EmailVerificationRequest(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	SendUserEmailVerificationMail(ctx context.Context, in *SendEmailVerificationDto, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	ValidateUserEmailVerification(ctx context.Context, in *TokenDto, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
	EditUserPassword(ctx context.Context, in *EditPasswordDto, opts ...grpc.CallOption) (*wrappers.BoolValue, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateCustomerAnonymousToken(ctx context.Context, in *AnonymousDto, opts ...grpc.CallOption) (*AuthResultDto, error) {
	out := new(AuthResultDto)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/CreateCustomerAnonymousToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateSellerAnonymousToken(ctx context.Context, in *AnonymousDto, opts ...grpc.CallOption) (*AuthResultDto, error) {
	out := new(AuthResultDto)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/CreateSellerAnonymousToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateAdminAnonymousToken(ctx context.Context, in *AnonymousDto, opts ...grpc.CallOption) (*AuthResultDto, error) {
	out := new(AuthResultDto)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/CreateAdminAnonymousToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidateToken(ctx context.Context, in *ValidateTokenDto, opts ...grpc.CallOption) (*Subject, error) {
	out := new(Subject)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RegisterUserAuth(ctx context.Context, in *RegisterDto, opts ...grpc.CallOption) (*AuthResultDto, error) {
	out := new(AuthResultDto)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/RegisterUserAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) IsUserEmailExists(ctx context.Context, in *UserAuthDto, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/IsUserEmailExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidateUserAccount(ctx context.Context, in *RegisterDto, opts ...grpc.CallOption) (*AuthResultDto, error) {
	out := new(AuthResultDto)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/ValidateUserAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) EmailVerificationRequest(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/EmailVerificationRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SendUserEmailVerificationMail(ctx context.Context, in *SendEmailVerificationDto, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/SendUserEmailVerificationMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidateUserEmailVerification(ctx context.Context, in *TokenDto, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/ValidateUserEmailVerification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) EditUserPassword(ctx context.Context, in *EditPasswordDto, opts ...grpc.CallOption) (*wrappers.BoolValue, error) {
	out := new(wrappers.BoolValue)
	err := c.cc.Invoke(ctx, "/authSvc.Auth/EditUserPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	GetHealth(context.Context, *empty.Empty) (*Health, error)
	CreateCustomerAnonymousToken(context.Context, *AnonymousDto) (*AuthResultDto, error)
	CreateSellerAnonymousToken(context.Context, *AnonymousDto) (*AuthResultDto, error)
	CreateAdminAnonymousToken(context.Context, *AnonymousDto) (*AuthResultDto, error)
	ValidateToken(context.Context, *ValidateTokenDto) (*Subject, error)
	// Customer
	RegisterUserAuth(context.Context, *RegisterDto) (*AuthResultDto, error)
	IsUserEmailExists(context.Context, *UserAuthDto) (*wrappers.BoolValue, error)
	ValidateUserAccount(context.Context, *RegisterDto) (*AuthResultDto, error)
	EmailVerificationRequest(context.Context, *wrappers.StringValue) (*wrappers.BoolValue, error)
	SendUserEmailVerificationMail(context.Context, *SendEmailVerificationDto) (*wrappers.BoolValue, error)
	ValidateUserEmailVerification(context.Context, *TokenDto) (*wrappers.BoolValue, error)
	EditUserPassword(context.Context, *EditPasswordDto) (*wrappers.BoolValue, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) GetHealth(context.Context, *empty.Empty) (*Health, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedAuthServer) CreateCustomerAnonymousToken(context.Context, *AnonymousDto) (*AuthResultDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomerAnonymousToken not implemented")
}
func (UnimplementedAuthServer) CreateSellerAnonymousToken(context.Context, *AnonymousDto) (*AuthResultDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSellerAnonymousToken not implemented")
}
func (UnimplementedAuthServer) CreateAdminAnonymousToken(context.Context, *AnonymousDto) (*AuthResultDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdminAnonymousToken not implemented")
}
func (UnimplementedAuthServer) ValidateToken(context.Context, *ValidateTokenDto) (*Subject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedAuthServer) RegisterUserAuth(context.Context, *RegisterDto) (*AuthResultDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUserAuth not implemented")
}
func (UnimplementedAuthServer) IsUserEmailExists(context.Context, *UserAuthDto) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUserEmailExists not implemented")
}
func (UnimplementedAuthServer) ValidateUserAccount(context.Context, *RegisterDto) (*AuthResultDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUserAccount not implemented")
}
func (UnimplementedAuthServer) EmailVerificationRequest(context.Context, *wrappers.StringValue) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailVerificationRequest not implemented")
}
func (UnimplementedAuthServer) SendUserEmailVerificationMail(context.Context, *SendEmailVerificationDto) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserEmailVerificationMail not implemented")
}
func (UnimplementedAuthServer) ValidateUserEmailVerification(context.Context, *TokenDto) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUserEmailVerification not implemented")
}
func (UnimplementedAuthServer) EditUserPassword(context.Context, *EditPasswordDto) (*wrappers.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserPassword not implemented")
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

func _Auth_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetHealth(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateCustomerAnonymousToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnonymousDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateCustomerAnonymousToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/CreateCustomerAnonymousToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateCustomerAnonymousToken(ctx, req.(*AnonymousDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateSellerAnonymousToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnonymousDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateSellerAnonymousToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/CreateSellerAnonymousToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateSellerAnonymousToken(ctx, req.(*AnonymousDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateAdminAnonymousToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnonymousDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateAdminAnonymousToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/CreateAdminAnonymousToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateAdminAnonymousToken(ctx, req.(*AnonymousDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ValidateToken(ctx, req.(*ValidateTokenDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RegisterUserAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RegisterUserAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/RegisterUserAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RegisterUserAuth(ctx, req.(*RegisterDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_IsUserEmailExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).IsUserEmailExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/IsUserEmailExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).IsUserEmailExists(ctx, req.(*UserAuthDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ValidateUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ValidateUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/ValidateUserAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ValidateUserAccount(ctx, req.(*RegisterDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_EmailVerificationRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).EmailVerificationRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/EmailVerificationRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).EmailVerificationRequest(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SendUserEmailVerificationMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailVerificationDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SendUserEmailVerificationMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/SendUserEmailVerificationMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SendUserEmailVerificationMail(ctx, req.(*SendEmailVerificationDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ValidateUserEmailVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ValidateUserEmailVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/ValidateUserEmailVerification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ValidateUserEmailVerification(ctx, req.(*TokenDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_EditUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditPasswordDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).EditUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authSvc.Auth/EditUserPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).EditUserPassword(ctx, req.(*EditPasswordDto))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authSvc.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _Auth_GetHealth_Handler,
		},
		{
			MethodName: "CreateCustomerAnonymousToken",
			Handler:    _Auth_CreateCustomerAnonymousToken_Handler,
		},
		{
			MethodName: "CreateSellerAnonymousToken",
			Handler:    _Auth_CreateSellerAnonymousToken_Handler,
		},
		{
			MethodName: "CreateAdminAnonymousToken",
			Handler:    _Auth_CreateAdminAnonymousToken_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _Auth_ValidateToken_Handler,
		},
		{
			MethodName: "RegisterUserAuth",
			Handler:    _Auth_RegisterUserAuth_Handler,
		},
		{
			MethodName: "IsUserEmailExists",
			Handler:    _Auth_IsUserEmailExists_Handler,
		},
		{
			MethodName: "ValidateUserAccount",
			Handler:    _Auth_ValidateUserAccount_Handler,
		},
		{
			MethodName: "EmailVerificationRequest",
			Handler:    _Auth_EmailVerificationRequest_Handler,
		},
		{
			MethodName: "SendUserEmailVerificationMail",
			Handler:    _Auth_SendUserEmailVerificationMail_Handler,
		},
		{
			MethodName: "ValidateUserEmailVerification",
			Handler:    _Auth_ValidateUserEmailVerification_Handler,
		},
		{
			MethodName: "EditUserPassword",
			Handler:    _Auth_EditUserPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth-svc.proto",
}
