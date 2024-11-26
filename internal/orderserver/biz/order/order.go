package order

import (
	"context"

	"github.com/costa92/micros-service/internal/orderserver/model"
	"github.com/costa92/micros-service/internal/orderserver/store"
	v1 "github.com/costa92/micros-service/pkg/api/orderserver/v1"
	"github.com/costa92/micros-service/pkg/store/where"
)

type IOrderBiz interface {
	CreateOrder(ctx context.Context, obj *model.Order) error

	GetOrder(ctx context.Context, req *v1.DetailRequest) (*model.Order, error)
}

type orderBiz struct {
	ds store.IStore
}

var _ IOrderBiz = (*orderBiz)(nil)

func NewOrderBiz(store store.IStore) IOrderBiz {
	return &orderBiz{
		ds: store,
	}
}

func (b *orderBiz) CreateOrder(ctx context.Context, obj *model.Order) error {
	return nil
}

// GetOrder retrieves an order by its ID.
func (b *orderBiz) GetOrder(ctx context.Context, req *v1.DetailRequest) (*model.Order, error) {
	whr := where.T(ctx).F("order_id", req.OrderId)
	order, err := b.ds.Orders().Get(ctx, whr)
	if err != nil {
		return nil, err
	}
	return order, nil
}
