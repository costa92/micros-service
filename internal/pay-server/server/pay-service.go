package service

import (
	context "context"

	payserver "github.com/costa92/micros-service/pkg/api/pay-server/v1"
	v1 "github.com/costa92/micros-service/pkg/api/pay-server/v1"
)

type PayService struct {
	v1.UnimplementedPayServiceServer
}

func NewPayService() *PayService {
	return &PayService{}
}

func (s *PayService) Pay(ctx context.Context, req *payserver.PayRequest) (*payserver.PayResponse, error) {
	return &payserver.PayResponse{}, nil
}

func (s *PayService) Detail(ctx context.Context, req *payserver.DetailRequest) (*payserver.DetailResponse, error) {
	return &payserver.DetailResponse{
		OrderId: req.OrderId,
	}, nil
}
