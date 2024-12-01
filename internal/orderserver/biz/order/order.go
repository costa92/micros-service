package order

import (
	"context"
	"fmt"

	"github.com/costa92/micros-service/internal/orderserver/model"
	"github.com/costa92/micros-service/internal/orderserver/store"
	v1 "github.com/costa92/micros-service/pkg/api/orderserver/v1"
	"github.com/costa92/micros-service/pkg/store/where"
)

// IOrderBiz defines the order business interface
type IOrderBiz interface {
	CreateOrder(ctx context.Context, order *model.Order) error
	GetOrder(ctx context.Context, req *v1.DetailRequest) (*model.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID string, status model.OrderStatus) error
}

type orderBiz struct {
	ds store.IStore
}

// NewOrderBiz creates a new order business instance
func NewOrderBiz(store store.IStore) IOrderBiz {
	return &orderBiz{
		ds: store,
	}
}

// CreateOrder creates a new order
func (b *orderBiz) CreateOrder(ctx context.Context, order *model.Order) error {
	if !order.IsValidStatus() {
		return fmt.Errorf("invalid order status: %s", order.OrderStatus)
	}
	
	return b.ds.Orders().Create(ctx, order)
}

// GetOrder retrieves an order by its ID
func (b *orderBiz) GetOrder(ctx context.Context, req *v1.DetailRequest) (*model.Order, error) {
	whr := where.T(ctx).F("order_id", req.OrderId)
	return b.ds.Orders().Get(ctx, whr)
}

// UpdateOrderStatus updates the order status
func (b *orderBiz) UpdateOrderStatus(ctx context.Context, orderID string, status model.OrderStatus) error {
	order := &model.Order{
		OrderStatus: status,
	}
	
	if !order.IsValidStatus() {
		return fmt.Errorf("invalid order status: %s", status)
	}

	whr := where.T(ctx).F("order_id", orderID)
	return b.ds.Orders().Update(ctx, order, whr)
}
