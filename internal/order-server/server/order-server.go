// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

package server

import (
	"context"

	v1 "github.com/costa92/micros-service/pkg/api/order-server/v1"
)

type OrderService struct {
	v1.UnimplementedOrderServiceServer
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *v1.CreateOrderRequest) (*v1.CreateOrderResponse, error) {
	return &v1.CreateOrderResponse{}, nil
}

func (s *OrderService) Detail(ctx context.Context, req *v1.DetailRequest) (*v1.DetailResponse, error) {
	return &v1.DetailResponse{
		OrderId: req.OrderId,
	}, nil
}