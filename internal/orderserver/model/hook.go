package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order is a model for an order.

// 统一使用一套生成id的方案，这里使用int类型的id

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.OrderID == "" {
		// Generate a new UUID for SecretKey.
		o.OrderID = uuid.New().String()
	}
	return nil
}
