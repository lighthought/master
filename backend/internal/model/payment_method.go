package model

// PaymentMethod 支付方式模型
type PaymentMethod struct {
	ID        string  `json:"id" gorm:"primaryKey;type:varchar(32)"`
	Name      string  `json:"name"`
	Icon      string  `json:"icon"`
	Enabled   bool    `json:"enabled"`
	MinAmount float64 `json:"min_amount"`
	MaxAmount float64 `json:"max_amount"`
}

func (PaymentMethod) TableName() string {
	return "payment_methods"
}
