package model

type Order struct {
	ID            int     `json:"id" gorm:"column:id;autoIncrement;primaryKey"`
	OrderID       string  `json:"order_id" gorm:"column:order_id;uniqueIndex"`
	UserId        int     `json:"user_id"`
	PaymentAmount float64 `json:"payment_amount"`
	OrderStatus   string  `json:"order_status"` //订单状态： 1 - 未支付 2 - 已支付 3 - 已发货 4 - 已收货 5 - 已完成 6 - 已取消
	CreatedAt     int64   `json:"created_at"`
	UpdatedAt     int64   `json:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}
