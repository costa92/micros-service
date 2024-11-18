// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.0--rc1
// source: orderserver/v1/order-server.proto

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

const OperationOrderServerCreateOrder = "/order_server.v1.OrderServer/CreateOrder"
const OperationOrderServerDetail = "/order_server.v1.OrderServer/Detail"

type OrderServerHTTPServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	Detail(context.Context, *DetailRequest) (*DetailResponse, error)
}

func RegisterOrderServerHTTPServer(s *http.Server, srv OrderServerHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/order", _OrderServer_CreateOrder0_HTTP_Handler(srv))
	r.GET("/v1/order/{order_id}", _OrderServer_Detail0_HTTP_Handler(srv))
}

func _OrderServer_CreateOrder0_HTTP_Handler(srv OrderServerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServerCreateOrder)
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

func _OrderServer_Detail0_HTTP_Handler(srv OrderServerHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DetailRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrderServerDetail)
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

type OrderServerHTTPClient interface {
	CreateOrder(ctx context.Context, req *CreateOrderRequest, opts ...http.CallOption) (rsp *CreateOrderResponse, err error)
	Detail(ctx context.Context, req *DetailRequest, opts ...http.CallOption) (rsp *DetailResponse, err error)
}

type OrderServerHTTPClientImpl struct {
	cc *http.Client
}

func NewOrderServerHTTPClient(client *http.Client) OrderServerHTTPClient {
	return &OrderServerHTTPClientImpl{client}
}

func (c *OrderServerHTTPClientImpl) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...http.CallOption) (*CreateOrderResponse, error) {
	var out CreateOrderResponse
	pattern := "/v1/order"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrderServerCreateOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OrderServerHTTPClientImpl) Detail(ctx context.Context, in *DetailRequest, opts ...http.CallOption) (*DetailResponse, error) {
	var out DetailResponse
	pattern := "/v1/order/{order_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrderServerDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
