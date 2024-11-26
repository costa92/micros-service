package store

import (
	"context"

	"github.com/costa92/micros-service/internal/orderserver/model"
	genericstore "github.com/costa92/micros-service/pkg/store"
	"github.com/costa92/micros-service/pkg/store/where"
)

type OrderStore interface {
	Create(ctx context.Context, obj *model.Order) error

	Get(ctx context.Context, opts *where.WhereOptions) (*model.Order, error)
}

type orderStore struct {
	*genericstore.Store[model.Order]
}

var _ OrderStore = (*orderStore)(nil)

// NewOrderStore creates a new order store.
func NewOrderStore(ds *datastore) OrderStore {
	return &orderStore{
		Store: genericstore.NewStore[model.Order](ds, nil),
	}
}
