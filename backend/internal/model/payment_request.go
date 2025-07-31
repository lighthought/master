package model

import "time"

// CreatePaymentOrderRequest 创建支付订单请求
// POST /payments/orders
// order_type, order_id, amount, currency, payment_method, description, metadata
// metadata 用 map[string]interface{} 反序列化
// swagger:parameters CreatePaymentOrderRequest
//
type CreatePaymentOrderRequest struct {
	OrderType     string                 `json:"order_type" binding:"required,oneof=course_enrollment appointment refund"`
	OrderID       string                 `json:"order_id" binding:"required"`
	Amount        float64                `json:"amount" binding:"required"`
	Currency      string                 `json:"currency" binding:"required"`
	PaymentMethod string                 `json:"payment_method" binding:"required"`
	Description   string                 `json:"description"`
	Metadata      map[string]interface{} `json:"metadata"`
}

// QueryPaymentStatusRequest 查询支付状态请求
// GET /payments/orders/{order_id}/status
// swagger:parameters QueryPaymentStatusRequest
//
type QueryPaymentStatusRequest struct {
	OrderID string `uri:"order_id" binding:"required"`
}

// PaymentHistoryRequest 获取支付历史请求
// GET /payments/history
// swagger:parameters PaymentHistoryRequest
//
type PaymentHistoryRequest struct {
	Type      string    `form:"type" binding:"omitempty,oneof=course_enrollment appointment refund"`
	Status    string    `form:"status" binding:"omitempty,oneof=pending completed failed cancelled"`
	StartDate time.Time `form:"start_date" time_format:"2006-01-02"`
	EndDate   time.Time `form:"end_date" time_format:"2006-01-02"`
	Page      int       `form:"page"`
	PageSize  int       `form:"page_size"`
}

// CreateRefundRequest 申请退款请求
// POST /payments/refunds
// swagger:parameters CreateRefundRequest
//
type CreateRefundRequest struct {
	PaymentID   string  `json:"payment_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Reason      string  `json:"reason"`
	Description string  `json:"description"`
}

// QueryRefundStatusRequest 查询退款状态请求
// GET /payments/refunds/{refund_id}/status
// swagger:parameters QueryRefundStatusRequest
//
type QueryRefundStatusRequest struct {
	RefundID string `uri:"refund_id" binding:"required"`
}

// PaymentStatsRequest 获取支付统计请求
// GET /payments/stats
// swagger:parameters PaymentStatsRequest
//
type PaymentStatsRequest struct {
	Period    string    `form:"period" binding:"omitempty,oneof=day week month quarter year"`
	StartDate time.Time `form:"start_date" time_format:"2006-01-02"`
	EndDate   time.Time `form:"end_date" time_format:"2006-01-02"`
}
