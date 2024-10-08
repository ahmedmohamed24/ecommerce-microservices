// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: protos/v1/admin.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AdminsService_GetAdmin_FullMethodName          = "/proto_admin_v1.AdminsService/GetAdmin"
	AdminsService_UpdateAdmin_FullMethodName       = "/proto_admin_v1.AdminsService/UpdateAdmin"
	AdminsService_GenerateAuthToken_FullMethodName = "/proto_admin_v1.AdminsService/GenerateAuthToken"
	AdminsService_RefreshAuthToken_FullMethodName  = "/proto_admin_v1.AdminsService/RefreshAuthToken"
)

// AdminsServiceClient is the client API for AdminsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminsServiceClient interface {
	GetAdmin(ctx context.Context, in *GetAdminRequest, opts ...grpc.CallOption) (*GetAdminResponse, error)
	UpdateAdmin(ctx context.Context, in *UpdateAdminRequest, opts ...grpc.CallOption) (*UpdateAdminResponse, error)
	GenerateAuthToken(ctx context.Context, in *GenerateAuthTokenRequest, opts ...grpc.CallOption) (*GenerateAuthTokenResponse, error)
	RefreshAuthToken(ctx context.Context, in *RefreshAuthTokenRequest, opts ...grpc.CallOption) (*RefreshAuthTokenResponse, error)
}

type adminsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminsServiceClient(cc grpc.ClientConnInterface) AdminsServiceClient {
	return &adminsServiceClient{cc}
}

func (c *adminsServiceClient) GetAdmin(ctx context.Context, in *GetAdminRequest, opts ...grpc.CallOption) (*GetAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAdminResponse)
	err := c.cc.Invoke(ctx, AdminsService_GetAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminsServiceClient) UpdateAdmin(ctx context.Context, in *UpdateAdminRequest, opts ...grpc.CallOption) (*UpdateAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAdminResponse)
	err := c.cc.Invoke(ctx, AdminsService_UpdateAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminsServiceClient) GenerateAuthToken(ctx context.Context, in *GenerateAuthTokenRequest, opts ...grpc.CallOption) (*GenerateAuthTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateAuthTokenResponse)
	err := c.cc.Invoke(ctx, AdminsService_GenerateAuthToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminsServiceClient) RefreshAuthToken(ctx context.Context, in *RefreshAuthTokenRequest, opts ...grpc.CallOption) (*RefreshAuthTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshAuthTokenResponse)
	err := c.cc.Invoke(ctx, AdminsService_RefreshAuthToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminsServiceServer is the server API for AdminsService service.
// All implementations must embed UnimplementedAdminsServiceServer
// for forward compatibility.
type AdminsServiceServer interface {
	GetAdmin(context.Context, *GetAdminRequest) (*GetAdminResponse, error)
	UpdateAdmin(context.Context, *UpdateAdminRequest) (*UpdateAdminResponse, error)
	GenerateAuthToken(context.Context, *GenerateAuthTokenRequest) (*GenerateAuthTokenResponse, error)
	RefreshAuthToken(context.Context, *RefreshAuthTokenRequest) (*RefreshAuthTokenResponse, error)
	mustEmbedUnimplementedAdminsServiceServer()
}

// UnimplementedAdminsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdminsServiceServer struct{}

func (UnimplementedAdminsServiceServer) GetAdmin(context.Context, *GetAdminRequest) (*GetAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdmin not implemented")
}
func (UnimplementedAdminsServiceServer) UpdateAdmin(context.Context, *UpdateAdminRequest) (*UpdateAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdmin not implemented")
}
func (UnimplementedAdminsServiceServer) GenerateAuthToken(context.Context, *GenerateAuthTokenRequest) (*GenerateAuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAuthToken not implemented")
}
func (UnimplementedAdminsServiceServer) RefreshAuthToken(context.Context, *RefreshAuthTokenRequest) (*RefreshAuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshAuthToken not implemented")
}
func (UnimplementedAdminsServiceServer) mustEmbedUnimplementedAdminsServiceServer() {}
func (UnimplementedAdminsServiceServer) testEmbeddedByValue()                       {}

// UnsafeAdminsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminsServiceServer will
// result in compilation errors.
type UnsafeAdminsServiceServer interface {
	mustEmbedUnimplementedAdminsServiceServer()
}

func RegisterAdminsServiceServer(s grpc.ServiceRegistrar, srv AdminsServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdminsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdminsService_ServiceDesc, srv)
}

func _AdminsService_GetAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminsServiceServer).GetAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminsService_GetAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminsServiceServer).GetAdmin(ctx, req.(*GetAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminsService_UpdateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminsServiceServer).UpdateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminsService_UpdateAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminsServiceServer).UpdateAdmin(ctx, req.(*UpdateAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminsService_GenerateAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminsServiceServer).GenerateAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminsService_GenerateAuthToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminsServiceServer).GenerateAuthToken(ctx, req.(*GenerateAuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminsService_RefreshAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshAuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminsServiceServer).RefreshAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminsService_RefreshAuthToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminsServiceServer).RefreshAuthToken(ctx, req.(*RefreshAuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminsService_ServiceDesc is the grpc.ServiceDesc for AdminsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_admin_v1.AdminsService",
	HandlerType: (*AdminsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdmin",
			Handler:    _AdminsService_GetAdmin_Handler,
		},
		{
			MethodName: "UpdateAdmin",
			Handler:    _AdminsService_UpdateAdmin_Handler,
		},
		{
			MethodName: "GenerateAuthToken",
			Handler:    _AdminsService_GenerateAuthToken_Handler,
		},
		{
			MethodName: "RefreshAuthToken",
			Handler:    _AdminsService_RefreshAuthToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/v1/admin.proto",
}
