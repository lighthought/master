package model

import "time"

// PaymentRefund 退款模型
type PaymentRefund struct {
	ID                  string     `json:"id" gorm:"primaryKey;type:varchar(32)"`
	PaymentID           string     `json:"payment_id"`
	Amount              float64    `json:"amount"`
	Status              string     `json:"status"`
	Reason              string     `json:"reason"`
	Description         string     `json:"description"`
	CreatedAt           time.Time  `json:"created_at"`
	CompletedAt         *time.Time `json:"completed_at"`
	RefundTransactionID string     `json:"refund_transaction_id"`
}

func (PaymentRefund) TableName() string {
	return "payment_refunds"
}
