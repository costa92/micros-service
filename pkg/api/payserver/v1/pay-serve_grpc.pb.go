// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0--rc1
// source: payserver/v1/pay-serve.proto

package v1

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
	PayServer_Pay_FullMethodName    = "/pay_server.v1.PayServer/Pay"
	PayServer_Detail_FullMethodName = "/pay_server.v1.PayServer/Detail"
)

// PayServerClient is the client API for PayServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayServerClient interface {
	Pay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error)
	Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
}

type payServerClient struct {
	cc grpc.ClientConnInterface
}

func NewPayServerClient(cc grpc.ClientConnInterface) PayServerClient {
	return &payServerClient{cc}
}

func (c *payServerClient) Pay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayResponse)
	err := c.cc.Invoke(ctx, PayServer_Pay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServerClient) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetailResponse)
	err := c.cc.Invoke(ctx, PayServer_Detail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayServerServer is the server API for PayServer service.
// All implementations must embed UnimplementedPayServerServer
// for forward compatibility.
type PayServerServer interface {
	Pay(context.Context, *PayRequest) (*PayResponse, error)
	Detail(context.Context, *DetailRequest) (*DetailResponse, error)
	mustEmbedUnimplementedPayServerServer()
}

// UnimplementedPayServerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPayServerServer struct{}

func (UnimplementedPayServerServer) Pay(context.Context, *PayRequest) (*PayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}
func (UnimplementedPayServerServer) Detail(context.Context, *DetailRequest) (*DetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedPayServerServer) mustEmbedUnimplementedPayServerServer() {}
func (UnimplementedPayServerServer) testEmbeddedByValue()                   {}

// UnsafePayServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayServerServer will
// result in compilation errors.
type UnsafePayServerServer interface {
	mustEmbedUnimplementedPayServerServer()
}

func RegisterPayServerServer(s grpc.ServiceRegistrar, srv PayServerServer) {
	// If the following call pancis, it indicates UnimplementedPayServerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PayServer_ServiceDesc, srv)
}

func _PayServer_Pay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServerServer).Pay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayServer_Pay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServerServer).Pay(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayServer_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServerServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayServer_Detail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServerServer).Detail(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayServer_ServiceDesc is the grpc.ServiceDesc for PayServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pay_server.v1.PayServer",
	HandlerType: (*PayServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pay",
			Handler:    _PayServer_Pay_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _PayServer_Detail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payserver/v1/pay-serve.proto",
}
