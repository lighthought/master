package service

import (
	"context"
	"errors"
	"math"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// PaymentService 支付服务接口
type PaymentService interface {
	CreatePaymentOrder(ctx context.Context, req *model.CreatePaymentOrderRequest) (*model.CreatePaymentOrderResponse, error)
	QueryPaymentStatus(ctx context.Context, orderID string) (*model.QueryPaymentStatusResponse, error)
	ListPaymentHistory(ctx context.Context, req *model.PaymentHistoryRequest) (*model.PaymentHistoryResponse, error)
	CreateRefund(ctx context.Context, req *model.CreateRefundRequest) (*model.CreateRefundResponse, error)
	QueryRefundStatus(ctx context.Context, refundID string) (*model.QueryRefundStatusResponse, error)
	ListPaymentMethods(ctx context.Context) (*model.PaymentMethodListResponse, error)
	GetPaymentStats(ctx context.Context, req *model.PaymentStatsRequest) (*model.PaymentStatsResponse, error)
	ProcessPaymentWebhook(ctx context.Context, gateway string, req *model.PaymentWebhookRequest) (*model.PaymentWebhookResponse, error)
}

type paymentService struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository) PaymentService {
	return &paymentService{
		paymentRepo: paymentRepo,
	}
}

func (s *paymentService) CreatePaymentOrder(ctx context.Context, req *model.CreatePaymentOrderRequest) (*model.CreatePaymentOrderResponse, error) {
	// 检查是否已存在订单
	existingOrder, err := s.paymentRepo.GetOrderByRef(ctx, req.OrderType, req.OrderID)
	if err == nil && existingOrder != nil {
		return nil, errors.New("订单已存在")
	}

	// 创建支付订单
	order := &model.PaymentOrder{
		OrderType:     req.OrderType,
		OrderRefID:    req.OrderID,
		Amount:        req.Amount,
		Currency:      req.Currency,
		PaymentMethod: req.PaymentMethod,
		Description:   req.Description,
		Status:        "pending",
		ExpiresAt:     &[]time.Time{time.Now().Add(30 * time.Minute)}[0], // 30分钟过期
	}

	err = s.paymentRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	// 创建支付记录
	paymentRecord := &model.PaymentRecord{
		OrderID:       order.ID,
		PaymentURL:    s.generatePaymentURL(order.ID, req.PaymentMethod),
		QRCode:        s.generateQRCode(order.ID),
		Status:        "pending",
		Amount:        req.Amount,
		Currency:      req.Currency,
		PaymentMethod: req.PaymentMethod,
	}

	err = s.paymentRepo.CreatePaymentRecord(ctx, paymentRecord)
	if err != nil {
		return nil, err
	}

	return &model.CreatePaymentOrderResponse{
		OrderID:    order.ID,
		PaymentID:  paymentRecord.ID,
		PaymentURL: paymentRecord.PaymentURL,
		QRCode:     paymentRecord.QRCode,
		ExpiresAt:  *order.ExpiresAt,
	}, nil
}

func (s *paymentService) QueryPaymentStatus(ctx context.Context, orderID string) (*model.QueryPaymentStatusResponse, error) {
	order, err := s.paymentRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, errors.New("订单不存在")
	}

	paymentRecord, err := s.paymentRepo.GetPaymentRecordByOrderID(ctx, orderID)
	if err != nil {
		return nil, errors.New("支付记录不存在")
	}

	return &model.QueryPaymentStatusResponse{
		OrderID:       order.ID,
		PaymentID:     paymentRecord.ID,
		Status:        paymentRecord.Status,
		Amount:        paymentRecord.Amount,
		Currency:      paymentRecord.Currency,
		PaymentMethod: paymentRecord.PaymentMethod,
		PaidAt:        paymentRecord.PaidAt,
		TransactionID: paymentRecord.TransactionID,
	}, nil
}

func (s *paymentService) ListPaymentHistory(ctx context.Context, req *model.PaymentHistoryRequest) (*model.PaymentHistoryResponse, error) {
	items, total, err := s.paymentRepo.ListPaymentHistory(ctx, req)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.PaymentHistoryResponse{
		Payments: items,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *paymentService) CreateRefund(ctx context.Context, req *model.CreateRefundRequest) (*model.CreateRefundResponse, error) {
	// 检查支付记录是否存在
	paymentRecord, err := s.paymentRepo.GetPaymentRecordByID(ctx, req.PaymentID)
	if err != nil {
		return nil, errors.New("支付记录不存在")
	}

	// 检查支付状态
	if paymentRecord.Status != "completed" {
		return nil, errors.New("支付未完成，无法退款")
	}

	// 检查退款金额
	if req.Amount > paymentRecord.Amount {
		return nil, errors.New("退款金额不能超过支付金额")
	}

	// 检查是否已有退款
	existingRefund, err := s.paymentRepo.GetRefundByPaymentID(ctx, req.PaymentID)
	if err == nil && existingRefund != nil {
		return nil, errors.New("该支付已有退款记录")
	}

	// 创建退款记录
	refund := &model.PaymentRefund{
		PaymentID:   req.PaymentID,
		Amount:      req.Amount,
		Status:      "pending",
		Reason:      req.Reason,
		Description: req.Description,
	}

	err = s.paymentRepo.CreateRefund(ctx, refund)
	if err != nil {
		return nil, err
	}

	estimatedCompletionTime := time.Now().AddDate(0, 0, 3) // 3天后

	return &model.CreateRefundResponse{
		RefundID:                refund.ID,
		Status:                  refund.Status,
		EstimatedCompletionTime: estimatedCompletionTime,
	}, nil
}

func (s *paymentService) QueryRefundStatus(ctx context.Context, refundID string) (*model.QueryRefundStatusResponse, error) {
	refund, err := s.paymentRepo.GetRefundByID(ctx, refundID)
	if err != nil {
		return nil, errors.New("退款记录不存在")
	}

	return &model.QueryRefundStatusResponse{
		RefundID:            refund.ID,
		PaymentID:           refund.PaymentID,
		Amount:              refund.Amount,
		Status:              refund.Status,
		Reason:              refund.Reason,
		CreatedAt:           refund.CreatedAt,
		CompletedAt:         refund.CompletedAt,
		RefundTransactionID: refund.RefundTransactionID,
	}, nil
}

func (s *paymentService) ListPaymentMethods(ctx context.Context) (*model.PaymentMethodListResponse, error) {
	methods, err := s.paymentRepo.ListPaymentMethods(ctx)
	if err != nil {
		return nil, err
	}

	return &model.PaymentMethodListResponse{
		PaymentMethods: methods,
	}, nil
}

func (s *paymentService) GetPaymentStats(ctx context.Context, req *model.PaymentStatsRequest) (*model.PaymentStatsResponse, error) {
	stats, err := s.paymentRepo.GetPaymentStats(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.PaymentStatsResponse{
		Stats: stats,
	}, nil
}

func (s *paymentService) ProcessPaymentWebhook(ctx context.Context, gateway string, req *model.PaymentWebhookRequest) (*model.PaymentWebhookResponse, error) {
	// 验证签名（简化处理）
	if !s.verifySignature(req) {
		return nil, errors.New("签名验证失败")
	}

	// 更新支付状态
	now := time.Now()
	err := s.paymentRepo.UpdatePaymentRecordStatus(ctx, req.OrderID, req.Status, &now, req.TransactionID)
	if err != nil {
		return nil, err
	}

	// 更新订单状态
	orderStatus := "pending"
	if req.Status == "success" {
		orderStatus = "completed"
	} else if req.Status == "failed" {
		orderStatus = "failed"
	}

	err = s.paymentRepo.UpdateOrderStatus(ctx, req.OrderID, orderStatus)
	if err != nil {
		return nil, err
	}

	return &model.PaymentWebhookResponse{
		Processed: true,
	}, nil
}

// 辅助方法
func (s *paymentService) generatePaymentURL(orderID, paymentMethod string) string {
	return "https://example.com/payment/gateway?order_id=" + orderID + "&method=" + paymentMethod
}

func (s *paymentService) generateQRCode(orderID string) string {
	return "https://example.com/qr-code/" + orderID + ".png"
}

func (s *paymentService) verifySignature(req *model.PaymentWebhookRequest) bool {
	// 简化签名验证，实际应该使用加密算法
	return req.Signature != ""
}
