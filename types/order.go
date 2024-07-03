package types

import "time"

// Order 订单
type Order struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `sql:"index" json:"deletedAt"`
	UserID       uint       `json:"userId"`
	TotalPrice   float64    `json:"totalPrice"`
	Payment      float64    `json:"payment"`
	Freight      float64    `json:"freight"`
	Remark       string     `json:"remark"`
	Discount     int        `json:"discount"`
	DeliverStart time.Time  `json:"deliverStart"`
	DeliverEnd   time.Time  `json:"deliverEnd"`
	Status       int        `json:"status"`
	PayAt        time.Time  `json:"payAt"`
}

const (
	// OrderStatusPending 未支付
	OrderStatusPending = 0

	// OrderStatusPaid 已支付
	OrderStatusPaid = 1
)

// OrderPerDay 每天的订单数
type OrderPerDay []struct {
	Count     int    `json:"count"`
	CreatedAt string `gorm:"column:createdAt" json:"createdAt"`
}

// AmountPerDay 每天的销售额
type AmountPerDay []struct {
	Amount float64 `json:"amount"`
	PayAt  string  `gorm:"column:payAt" json:"payAt"`
}
