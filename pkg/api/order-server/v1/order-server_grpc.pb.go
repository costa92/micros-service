// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.3
// source: order-server/v1/order-server.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	OrderServer_CreateOrder_FullMethodName = "/order_server.v1.OrderServer/CreateOrder"
	OrderServer_Detail_FullMethodName      = "/order_server.v1.OrderServer/Detail"
)

// OrderServerClient is the client API for OrderServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServerClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
}

type orderServerClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServerClient(cc grpc.ClientConnInterface) OrderServerClient {
	return &orderServerClient{cc}
}

func (c *orderServerClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, OrderServer_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServerClient) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetailResponse)
	err := c.cc.Invoke(ctx, OrderServer_Detail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServerServer is the server API for OrderServer service.
// All implementations must embed UnimplementedOrderServerServer
// for forward compatibility
type OrderServerServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	Detail(context.Context, *DetailRequest) (*DetailResponse, error)
	mustEmbedUnimplementedOrderServerServer()
}

// UnimplementedOrderServerServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServerServer struct {
}

func (UnimplementedOrderServerServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServerServer) Detail(context.Context, *DetailRequest) (*DetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedOrderServerServer) mustEmbedUnimplementedOrderServerServer() {}

// UnsafeOrderServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServerServer will
// result in compilation errors.
type UnsafeOrderServerServer interface {
	mustEmbedUnimplementedOrderServerServer()
}

func RegisterOrderServerServer(s grpc.ServiceRegistrar, srv OrderServerServer) {
	s.RegisterService(&OrderServer_ServiceDesc, srv)
}

func _OrderServer_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServerServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServer_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServerServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderServer_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServerServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderServer_Detail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServerServer).Detail(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderServer_ServiceDesc is the grpc.ServiceDesc for OrderServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order_server.v1.OrderServer",
	HandlerType: (*OrderServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderServer_CreateOrder_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _OrderServer_Detail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order-server/v1/order-server.proto",
}
