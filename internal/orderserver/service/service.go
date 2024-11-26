// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

package service

import (
	"context"
	"errors"

	"github.com/costa92/micros-service/internal/orderserver/biz"
	servermetrics "github.com/costa92/micros-service/internal/pkg/metrics"
	v1 "github.com/costa92/micros-service/pkg/api/orderserver/v1"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is a set of service providers, used for dependency injection.
var ProviderSet = wire.NewSet(NewOrderService)

type OrderService struct {
	v1.UnimplementedOrderServerServer
	OptMetrics *servermetrics.Metrics
	biz        biz.IBiz // A factory for creating business logic components.
}

func NewOrderService(optMetrics *servermetrics.Metrics, biz biz.IBiz) *OrderService {
	return &OrderService{
		OptMetrics: optMetrics,
		biz:        biz,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *v1.CreateOrderRequest) (*v1.CreateOrderResponse, error) {
	return &v1.CreateOrderResponse{}, nil
}

func (s *OrderService) Detail(ctx context.Context, req *v1.DetailRequest) (*v1.DetailResponse, error) {
	s.OptMetrics.IncrementLabelOrderCount(ctx, "order_id", req.OrderId)
	s.OptMetrics.IncrementOrderCount(ctx)
	order, err := s.biz.Orders().GetOrder(ctx, req)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, v1.ErrorDatabaseError("database error")
	}

	if order == nil {
		return nil, v1.ErrorOrderNotFound("order not found")
	}

	return &v1.DetailResponse{
		OrderId: req.OrderId,
	}, nil
}
