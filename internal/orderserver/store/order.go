package store

import (
	"context"

	"github.com/costa92/micros-service/internal/orderserver/model"
	genericstore "github.com/costa92/micros-service/pkg/store"
	"github.com/costa92/micros-service/pkg/store/where"
)

// OrderStore defines the interface for order storage operations
type OrderStore interface {
	Create(ctx context.Context, order *model.Order) error
	Get(ctx context.Context, opts *where.WhereOptions) (*model.Order, error)
	Update(ctx context.Context, order *model.Order, opts *where.WhereOptions) error
	Delete(ctx context.Context, opts *where.WhereOptions) error
	List(ctx context.Context, opts *where.WhereOptions) (count int64, ret []*model.Order, err error)
}

type orderStore struct {
	*genericstore.Store[model.Order]
}

// NewOrderStore creates a new order store instance
func NewOrderStore(ds *datastore) OrderStore {
	return &orderStore{
		Store: genericstore.NewStore[model.Order](ds, nil),
	}
}

// Create creates a new order
func (s *orderStore) Create(ctx context.Context, order *model.Order) error {
	return s.Store.Create(ctx, order)
}

// Get retrieves an order by conditions
func (s *orderStore) Get(ctx context.Context, opts *where.WhereOptions) (*model.Order, error) {
	return s.Store.Get(ctx, opts)
}

// Update updates an order
func (s *orderStore) Update(ctx context.Context, order *model.Order, opts *where.WhereOptions) error {
	return s.Store.Update(ctx, order, opts)
}

// Delete deletes an order
func (s *orderStore) Delete(ctx context.Context, opts *where.WhereOptions) error {
	return s.Store.Delete(ctx, opts)
}

// List retrieves orders by conditions
func (s *orderStore) List(ctx context.Context, opts *where.WhereOptions) (count int64, ret []*model.Order, err error) {
	return s.Store.List(ctx, opts)
}
