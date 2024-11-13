// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v5.28.3
// source: order-server/v1/order-server.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationOrderServiceCreateOrder = "/order_server.v1.OrderService/CreateOrder"
const OperationOrderServiceDetail = "/order_server.v1.OrderService/Detail"

type OrderServiceHTTPServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	Detail(context.Context, *DetailRequest) (*DetailResponse, error)
}

func RegisterOrderServiceHTTPServer(s *http.Server, srv OrderServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/order", _OrderService_CreateOrder0_HTTP_Handler(srv))
	r.GET("/v1/order/{order_id}", _OrderService_Detail0_HTTP_Handler(srv))
}

func _OrderService_CreateOrder0_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServiceCreateOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateOrder(ctx, req.(*CreateOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateOrderResponse)
		return ctx.Result(200, reply)
	}
}

func _OrderService_Detail0_HTTP_Handler(srv OrderServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DetailRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServiceDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Detail(ctx, req.(*DetailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DetailResponse)
		return ctx.Result(200, reply)
	}
}

type OrderServiceHTTPClient interface {
	CreateOrder(ctx context.Context, req *CreateOrderRequest, opts ...http.CallOption) (rsp *CreateOrderResponse, err error)
	Detail(ctx context.Context, req *DetailRequest, opts ...http.CallOption) (rsp *DetailResponse, err error)
}

type OrderServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewOrderServiceHTTPClient(client *http.Client) OrderServiceHTTPClient {
	return &OrderServiceHTTPClientImpl{client}
}

func (c *OrderServiceHTTPClientImpl) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...http.CallOption) (*CreateOrderResponse, error) {
	var out CreateOrderResponse
	pattern := "/v1/order"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrderServiceCreateOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OrderServiceHTTPClientImpl) Detail(ctx context.Context, in *DetailRequest, opts ...http.CallOption) (*DetailResponse, error) {
	var out DetailResponse
	pattern := "/v1/order/{order_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrderServiceDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
