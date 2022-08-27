// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.7.1
// source: clientbase.proto

package api

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

// ClientbaseClient is the client API for Clientbase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientbaseClient interface {
	Add(ctx context.Context, in *AddClientRequest, opts ...grpc.CallOption) (*AddClientResponse, error)
	Delete(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientResponse, error)
	GetClients(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (*GetClientsResponse, error)
	GetClientsFromRedis(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (*GetClientsResponse, error)
}

type clientbaseClient struct {
	cc grpc.ClientConnInterface
}

func NewClientbaseClient(cc grpc.ClientConnInterface) ClientbaseClient {
	return &clientbaseClient{cc}
}

func (c *clientbaseClient) Add(ctx context.Context, in *AddClientRequest, opts ...grpc.CallOption) (*AddClientResponse, error) {
	out := new(AddClientResponse)
	err := c.cc.Invoke(ctx, "/clientbase.clientbase/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientbaseClient) Delete(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientResponse, error) {
	out := new(DeleteClientResponse)
	err := c.cc.Invoke(ctx, "/clientbase.clientbase/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientbaseClient) GetClients(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (*GetClientsResponse, error) {
	out := new(GetClientsResponse)
	err := c.cc.Invoke(ctx, "/clientbase.clientbase/GetClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientbaseClient) GetClientsFromRedis(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (*GetClientsResponse, error) {
	out := new(GetClientsResponse)
	err := c.cc.Invoke(ctx, "/clientbase.clientbase/GetClientsFromRedis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientbaseServer is the server API for Clientbase service.
// All implementations must embed UnimplementedClientbaseServer
// for forward compatibility
type ClientbaseServer interface {
	Add(context.Context, *AddClientRequest) (*AddClientResponse, error)
	Delete(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error)
	GetClients(context.Context, *GetClientsRequest) (*GetClientsResponse, error)
	GetClientsFromRedis(context.Context, *GetClientsRequest) (*GetClientsResponse, error)
	mustEmbedUnimplementedClientbaseServer()
}

// UnimplementedClientbaseServer must be embedded to have forward compatible implementations.
type UnimplementedClientbaseServer struct {
}

func (UnimplementedClientbaseServer) Add(context.Context, *AddClientRequest) (*AddClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedClientbaseServer) Delete(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClientbaseServer) GetClients(context.Context, *GetClientsRequest) (*GetClientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClients not implemented")
}
func (UnimplementedClientbaseServer) GetClientsFromRedis(context.Context, *GetClientsRequest) (*GetClientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientsFromRedis not implemented")
}
func (UnimplementedClientbaseServer) mustEmbedUnimplementedClientbaseServer() {}

// UnsafeClientbaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientbaseServer will
// result in compilation errors.
type UnsafeClientbaseServer interface {
	mustEmbedUnimplementedClientbaseServer()
}

func RegisterClientbaseServer(s grpc.ServiceRegistrar, srv ClientbaseServer) {
	s.RegisterService(&Clientbase_ServiceDesc, srv)
}

func _Clientbase_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientbaseServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientbase.clientbase/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientbaseServer).Add(ctx, req.(*AddClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clientbase_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientbaseServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientbase.clientbase/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientbaseServer).Delete(ctx, req.(*DeleteClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clientbase_GetClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientbaseServer).GetClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientbase.clientbase/GetClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientbaseServer).GetClients(ctx, req.(*GetClientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clientbase_GetClientsFromRedis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientbaseServer).GetClientsFromRedis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientbase.clientbase/GetClientsFromRedis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientbaseServer).GetClientsFromRedis(ctx, req.(*GetClientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Clientbase_ServiceDesc is the grpc.ServiceDesc for Clientbase service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Clientbase_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clientbase.clientbase",
	HandlerType: (*ClientbaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Clientbase_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Clientbase_Delete_Handler,
		},
		{
			MethodName: "GetClients",
			Handler:    _Clientbase_GetClients_Handler,
		},
		{
			MethodName: "GetClientsFromRedis",
			Handler:    _Clientbase_GetClientsFromRedis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clientbase.proto",
}
