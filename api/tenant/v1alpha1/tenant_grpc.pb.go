// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: tenant/v1alpha1/tenant.proto

package valpha1_tenant

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantServiceClient interface {
	CreateTenant(ctx context.Context, in *Tenant, opts ...grpc.CallOption) (*Tenant, error)
	GetTenant(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	DeleteTenant(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListTenants(ctx context.Context, in *TenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
}

type tenantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantServiceClient(cc grpc.ClientConnInterface) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) CreateTenant(ctx context.Context, in *Tenant, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/org.template.tenant.v1.TenantService/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) GetTenant(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/org.template.tenant.v1.TenantService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) DeleteTenant(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/org.template.tenant.v1.TenantService/DeleteTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) ListTenants(ctx context.Context, in *TenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, "/org.template.tenant.v1.TenantService/ListTenants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
// All implementations must embed UnimplementedTenantServiceServer
// for forward compatibility
type TenantServiceServer interface {
	CreateTenant(context.Context, *Tenant) (*Tenant, error)
	GetTenant(context.Context, *TenantRequest) (*Tenant, error)
	DeleteTenant(context.Context, *TenantRequest) (*emptypb.Empty, error)
	ListTenants(context.Context, *TenantsRequest) (*ListTenantsResponse, error)
	mustEmbedUnimplementedTenantServiceServer()
}

// UnimplementedTenantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTenantServiceServer struct {
}

func (UnimplementedTenantServiceServer) CreateTenant(context.Context, *Tenant) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedTenantServiceServer) GetTenant(context.Context, *TenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (UnimplementedTenantServiceServer) DeleteTenant(context.Context, *TenantRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (UnimplementedTenantServiceServer) ListTenants(context.Context, *TenantsRequest) (*ListTenantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenants not implemented")
}
func (UnimplementedTenantServiceServer) mustEmbedUnimplementedTenantServiceServer() {}

// UnsafeTenantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServiceServer will
// result in compilation errors.
type UnsafeTenantServiceServer interface {
	mustEmbedUnimplementedTenantServiceServer()
}

func RegisterTenantServiceServer(s grpc.ServiceRegistrar, srv TenantServiceServer) {
	s.RegisterService(&TenantService_ServiceDesc, srv)
}

func _TenantService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tenant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.template.tenant.v1.TenantService/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).CreateTenant(ctx, req.(*Tenant))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.template.tenant.v1.TenantService/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetTenant(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.template.tenant.v1.TenantService/DeleteTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).DeleteTenant(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_ListTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).ListTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.template.tenant.v1.TenantService/ListTenants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).ListTenants(ctx, req.(*TenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantService_ServiceDesc is the grpc.ServiceDesc for TenantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.template.tenant.v1.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTenant",
			Handler:    _TenantService_CreateTenant_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _TenantService_GetTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _TenantService_DeleteTenant_Handler,
		},
		{
			MethodName: "ListTenants",
			Handler:    _TenantService_ListTenants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenant/v1alpha1/tenant.proto",
}
