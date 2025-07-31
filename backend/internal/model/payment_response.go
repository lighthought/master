package model

import "time"

// CreatePaymentOrderResponse 创建支付订单响应
type CreatePaymentOrderResponse struct {
	OrderID    string    `json:"order_id"`
	PaymentID  string    `json:"payment_id"`
	PaymentURL string    `json:"payment_url"`
	QRCode     string    `json:"qr_code"`
	ExpiresAt  time.Time `json:"expires_at"`
}

// QueryPaymentStatusResponse 查询支付状态响应
type QueryPaymentStatusResponse struct {
	OrderID       string     `json:"order_id"`
	PaymentID     string     `json:"payment_id"`
	Status        string     `json:"status"`
	Amount        float64    `json:"amount"`
	Currency      string     `json:"currency"`
	PaymentMethod string     `json:"payment_method"`
	PaidAt        *time.Time `json:"paid_at"`
	TransactionID string     `json:"transaction_id"`
}

// PaymentHistoryItem 支付历史单条
type PaymentHistoryItem struct {
	ID            string                 `json:"id"`
	OrderID       string                 `json:"order_id"`
	Type          string                 `json:"type"`
	Amount        float64                `json:"amount"`
	Currency      string                 `json:"currency"`
	PaymentMethod string                 `json:"payment_method"`
	Status        string                 `json:"status"`
	Description   string                 `json:"description"`
	CreatedAt     time.Time              `json:"created_at"`
	PaidAt        *time.Time             `json:"paid_at"`
	TransactionID string                 `json:"transaction_id"`
	Metadata      map[string]interface{} `json:"metadata"`
}

// PaymentHistoryResponse 支付历史响应
type PaymentHistoryResponse struct {
	Payments   []*PaymentHistoryItem `json:"payments"`
	Pagination *PaginationResponse   `json:"pagination"`
}

// CreateRefundResponse 申请退款响应
type CreateRefundResponse struct {
	RefundID                string    `json:"refund_id"`
	Status                  string    `json:"status"`
	EstimatedCompletionTime time.Time `json:"estimated_completion_time"`
}

// QueryRefundStatusResponse 查询退款状态响应
type QueryRefundStatusResponse struct {
	RefundID            string     `json:"refund_id"`
	PaymentID           string     `json:"payment_id"`
	Amount              float64    `json:"amount"`
	Status              string     `json:"status"`
	Reason              string     `json:"reason"`
	CreatedAt           time.Time  `json:"created_at"`
	CompletedAt         *time.Time `json:"completed_at"`
	RefundTransactionID string     `json:"refund_transaction_id"`
}

// PaymentMethodListResponse 支付方式列表响应
type PaymentMethodListResponse struct {
	PaymentMethods []*PaymentMethod `json:"payment_methods"`
}

// PaymentStatsResponse 支付统计响应
type PaymentStatsResponse struct {
	Stats *PaymentStats `json:"stats"`
}

type PaymentStats struct {
	TotalAmount                float64             `json:"total_amount"`
	TotalTransactions          int64               `json:"total_transactions"`
	SuccessfulTransactions     int64               `json:"successful_transactions"`
	FailedTransactions         int64               `json:"failed_transactions"`
	RefundAmount               float64             `json:"refund_amount"`
	RefundCount                int64               `json:"refund_count"`
	PaymentMethodsDistribution map[string]int64    `json:"payment_methods_distribution"`
	DailyStats                 []*PaymentDailyStat `json:"daily_stats"`
}

type PaymentDailyStat struct {
	Date         string  `json:"date"`
	Amount       float64 `json:"amount"`
	Transactions int     `json:"transactions"`
}
