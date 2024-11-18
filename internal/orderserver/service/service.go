// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

package service

import (
	"context"

	servermetrics "github.com/costa92/micros-service/internal/pkg/metrics"
	v1 "github.com/costa92/micros-service/pkg/api/orderserver/v1"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// ProviderSet is a set of service providers, used for dependency injection.
var ProviderSet = wire.NewSet(NewOrderService)

type OrderService struct {
	v1.UnimplementedOrderServerServer
	OptMetrics *servermetrics.Metrics
}

func NewOrderService(optMetrics *servermetrics.Metrics) *OrderService {
	return &OrderService{
		OptMetrics: optMetrics,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *v1.CreateOrderRequest) (*v1.CreateOrderResponse, error) {
	return &v1.CreateOrderResponse{}, nil
}

func (s *OrderService) Detail(ctx context.Context, req *v1.DetailRequest) (*v1.DetailResponse, error) {
	if req.OrderId == "11" {
		return nil, v1.ErrorOrderNotFound("order not found")
	}

	counter, _ := s.OptMetrics.Meter.Int64Counter("order_id", metric.WithUnit("1"), metric.WithDescription("order id"))

	// Increment the counter.
	counter.Add(ctx, 1, metric.WithAttributes(attribute.String("foo", "bar")))
	counter.Add(ctx, 10, metric.WithAttributes(attribute.String("hello", "world")))

	if req.OrderId == "22" {
		return nil, v1.ErrorOrderAlreadyExists("order already exists")
	}
	return &v1.DetailResponse{
		OrderId: req.OrderId,
	}, nil
}
