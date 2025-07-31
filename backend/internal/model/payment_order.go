package model

import "time"

// PaymentOrder 支付订单模型
type PaymentOrder struct {
	ID            string     `json:"id" gorm:"primaryKey;type:varchar(32)"`
	OrderType     string     `json:"order_type"`
	OrderRefID    string     `json:"order_id"`
	Amount        float64    `json:"amount"`
	Currency      string     `json:"currency"`
	PaymentMethod string     `json:"payment_method"`
	Description   string     `json:"description"`
	Metadata      string     `json:"metadata" gorm:"type:jsonb"`
	Status        string     `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	ExpiresAt     *time.Time `json:"expires_at"`
}

func (PaymentOrder) TableName() string {
	return "payment_orders"
}
