package model

// PaymentWebhookRequest 支付回调请求
type PaymentWebhookRequest struct {
	OrderID       string  `json:"order_id" binding:"required"`
	TransactionID string  `json:"transaction_id" binding:"required"`
	Status        string  `json:"status" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
	Signature     string  `json:"signature" binding:"required"`
}

// PaymentWebhookResponse 支付回调响应
type PaymentWebhookResponse struct {
	Processed bool `json:"processed"`
}
