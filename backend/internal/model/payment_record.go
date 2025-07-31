package model

import "time"

// PaymentRecord 支付流水模型
type PaymentRecord struct {
	ID            string     `json:"id" gorm:"primaryKey;type:varchar(32)"`
	OrderID       string     `json:"order_id"`
	PaymentURL    string     `json:"payment_url"`
	QRCode        string     `json:"qr_code"`
	Status        string     `json:"status"`
	Amount        float64    `json:"amount"`
	Currency      string     `json:"currency"`
	PaymentMethod string     `json:"payment_method"`
	PaidAt        *time.Time `json:"paid_at"`
	TransactionID string     `json:"transaction_id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func (PaymentRecord) TableName() string {
	return "payment_records"
}
