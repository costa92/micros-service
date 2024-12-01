package model


// OrderStatus 定义订单状态常量
type OrderStatus string

const (
	OrderStatusUnpaid    OrderStatus = "1" // 未支付
	OrderStatusPaid      OrderStatus = "2" // 已支付
	OrderStatusShipped   OrderStatus = "3" // 已发货
	OrderStatusReceived  OrderStatus = "4" // 已收货
	OrderStatusComplete  OrderStatus = "5" // 已完成
	OrderStatusCancelled OrderStatus = "6" // 已取消
)

// Order represents an order in the system
type Order struct {
	ID            int64       `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	OrderID       string      `json:"order_id" gorm:"column:order_id;uniqueIndex;type:varchar(64)"`
	UserID        int64       `json:"user_id" gorm:"column:user_id;index"`
	PaymentAmount float64     `json:"payment_amount" gorm:"column:payment_amount;type:decimal(10,2)"`
	OrderStatus   OrderStatus `json:"order_status" gorm:"column:order_status;type:varchar(32)"` 
	CreatedAt     int64      `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt     int64      `json:"updated_at" gorm:"column:updated_at;autoUpdateTime:milli"`
}

// TableName specifies the table name for Order model
func (Order) TableName() string {
	return "orders"
}

// IsValidStatus checks if the order status is valid
func (o *Order) IsValidStatus() bool {
	switch o.OrderStatus {
	case OrderStatusUnpaid, OrderStatusPaid, OrderStatusShipped,
		OrderStatusReceived, OrderStatusComplete, OrderStatusCancelled:
		return true
	}
	return false
}
