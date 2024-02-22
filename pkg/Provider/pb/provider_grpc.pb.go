// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/Provider/pb/provider.proto

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

// ProviderServicesClient is the client API for ProviderServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderServicesClient interface {
	EditProvider(ctx context.Context, in *ProviderDataReq, opts ...grpc.CallOption) (*Res, error)
	DeleteProviderById(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Res, error)
	FindProviderById(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Res, error)
	CreateProvider(ctx context.Context, in *ProviderDataReq, opts ...grpc.CallOption) (*Res, error)
}

type providerServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderServicesClient(cc grpc.ClientConnInterface) ProviderServicesClient {
	return &providerServicesClient{cc}
}

func (c *providerServicesClient) EditProvider(ctx context.Context, in *ProviderDataReq, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/pb.ProviderServices/EditProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServicesClient) DeleteProviderById(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/pb.ProviderServices/DeleteProviderById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServicesClient) FindProviderById(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/pb.ProviderServices/FindProviderById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServicesClient) CreateProvider(ctx context.Context, in *ProviderDataReq, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/pb.ProviderServices/CreateProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServicesServer is the server API for ProviderServices service.
// All implementations must embed UnimplementedProviderServicesServer
// for forward compatibility
type ProviderServicesServer interface {
	EditProvider(context.Context, *ProviderDataReq) (*Res, error)
	DeleteProviderById(context.Context, *IdReq) (*Res, error)
	FindProviderById(context.Context, *IdReq) (*Res, error)
	CreateProvider(context.Context, *ProviderDataReq) (*Res, error)
	mustEmbedUnimplementedProviderServicesServer()
}

// UnimplementedProviderServicesServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServicesServer struct {
}

func (UnimplementedProviderServicesServer) EditProvider(context.Context, *ProviderDataReq) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProvider not implemented")
}
func (UnimplementedProviderServicesServer) DeleteProviderById(context.Context, *IdReq) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProviderById not implemented")
}
func (UnimplementedProviderServicesServer) FindProviderById(context.Context, *IdReq) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindProviderById not implemented")
}
func (UnimplementedProviderServicesServer) CreateProvider(context.Context, *ProviderDataReq) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProvider not implemented")
}
func (UnimplementedProviderServicesServer) mustEmbedUnimplementedProviderServicesServer() {}

// UnsafeProviderServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServicesServer will
// result in compilation errors.
type UnsafeProviderServicesServer interface {
	mustEmbedUnimplementedProviderServicesServer()
}

func RegisterProviderServicesServer(s grpc.ServiceRegistrar, srv ProviderServicesServer) {
	s.RegisterService(&ProviderServices_ServiceDesc, srv)
}

func _ProviderServices_EditProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServicesServer).EditProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProviderServices/EditProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServicesServer).EditProvider(ctx, req.(*ProviderDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderServices_DeleteProviderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServicesServer).DeleteProviderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProviderServices/DeleteProviderById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServicesServer).DeleteProviderById(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderServices_FindProviderById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServicesServer).FindProviderById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProviderServices/FindProviderById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServicesServer).FindProviderById(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderServices_CreateProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServicesServer).CreateProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProviderServices/CreateProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServicesServer).CreateProvider(ctx, req.(*ProviderDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProviderServices_ServiceDesc is the grpc.ServiceDesc for ProviderServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProviderServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProviderServices",
	HandlerType: (*ProviderServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EditProvider",
			Handler:    _ProviderServices_EditProvider_Handler,
		},
		{
			MethodName: "DeleteProviderById",
			Handler:    _ProviderServices_DeleteProviderById_Handler,
		},
		{
			MethodName: "FindProviderById",
			Handler:    _ProviderServices_FindProviderById_Handler,
		},
		{
			MethodName: "CreateProvider",
			Handler:    _ProviderServices_CreateProvider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/Provider/pb/provider.proto",
}