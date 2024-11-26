package biz

import (
	"github.com/costa92/micros-service/internal/orderserver/biz/order"
	"github.com/costa92/micros-service/internal/orderserver/store"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewBiz)

type IBiz interface {
	Orders() order.IOrderBiz
}

type biz struct {
	ds store.IStore
}

func NewBiz(ds store.IStore) IBiz {
	return &biz{
		ds: ds,
	}
}

func (b *biz) Orders() order.IOrderBiz {
	return order.NewOrderBiz(b.ds)
}
