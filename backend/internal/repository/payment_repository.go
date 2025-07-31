package repository

import (
	"context"
	"time"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// PaymentRepository 支付数据访问接口
type PaymentRepository interface {
	CreateOrder(ctx context.Context, order *model.PaymentOrder) error
	GetOrderByID(ctx context.Context, id string) (*model.PaymentOrder, error)
	GetOrderByRef(ctx context.Context, orderType, orderRefID string) (*model.PaymentOrder, error)
	UpdateOrderStatus(ctx context.Context, id, status string) error

	CreatePaymentRecord(ctx context.Context, record *model.PaymentRecord) error
	GetPaymentRecordByID(ctx context.Context, id string) (*model.PaymentRecord, error)
	GetPaymentRecordByOrderID(ctx context.Context, orderID string) (*model.PaymentRecord, error)
	UpdatePaymentRecordStatus(ctx context.Context, id, status string, paidAt *time.Time, transactionID string) error

	ListPaymentHistory(ctx context.Context, req *model.PaymentHistoryRequest) ([]*model.PaymentHistoryItem, int64, error)

	CreateRefund(ctx context.Context, refund *model.PaymentRefund) error
	GetRefundByID(ctx context.Context, id string) (*model.PaymentRefund, error)
	GetRefundByPaymentID(ctx context.Context, paymentID string) (*model.PaymentRefund, error)

	ListPaymentMethods(ctx context.Context) ([]*model.PaymentMethod, error)

	GetPaymentStats(ctx context.Context, req *model.PaymentStatsRequest) (*model.PaymentStats, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) CreateOrder(ctx context.Context, order *model.PaymentOrder) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *paymentRepository) GetOrderByID(ctx context.Context, id string) (*model.PaymentOrder, error) {
	var order model.PaymentOrder
	err := r.db.WithContext(ctx).First(&order, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *paymentRepository) GetOrderByRef(ctx context.Context, orderType, orderRefID string) (*model.PaymentOrder, error) {
	var order model.PaymentOrder
	err := r.db.WithContext(ctx).First(&order, "order_type = ? AND order_ref_id = ?", orderType, orderRefID).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *paymentRepository) UpdateOrderStatus(ctx context.Context, id, status string) error {
	return r.db.WithContext(ctx).Model(&model.PaymentOrder{}).Where("id = ?", id).Update("status", status).Error
}

func (r *paymentRepository) CreatePaymentRecord(ctx context.Context, record *model.PaymentRecord) error {
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *paymentRepository) GetPaymentRecordByID(ctx context.Context, id string) (*model.PaymentRecord, error) {
	var rec model.PaymentRecord
	err := r.db.WithContext(ctx).First(&rec, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &rec, nil
}

func (r *paymentRepository) GetPaymentRecordByOrderID(ctx context.Context, orderID string) (*model.PaymentRecord, error) {
	var rec model.PaymentRecord
	err := r.db.WithContext(ctx).First(&rec, "order_id = ?", orderID).Error
	if err != nil {
		return nil, err
	}
	return &rec, nil
}

func (r *paymentRepository) UpdatePaymentRecordStatus(ctx context.Context, id, status string, paidAt *time.Time, transactionID string) error {
	return r.db.WithContext(ctx).Model(&model.PaymentRecord{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":         status,
		"paid_at":        paidAt,
		"transaction_id": transactionID,
	}).Error
}

func (r *paymentRepository) ListPaymentHistory(ctx context.Context, req *model.PaymentHistoryRequest) ([]*model.PaymentHistoryItem, int64, error) {
	var items []*model.PaymentHistoryItem
	var total int64
	query := r.db.WithContext(ctx).Table("payment_records pr").
		Select(`pr.id, pr.order_id, po.order_type as type, pr.amount, pr.currency, pr.payment_method, pr.status, po.description, pr.created_at, pr.paid_at, pr.transaction_id, po.metadata`).
		Joins("LEFT JOIN payment_orders po ON pr.order_id = po.id")
	if req.Type != "" {
		query = query.Where("po.order_type = ?", req.Type)
	}
	if req.Status != "" {
		query = query.Where("pr.status = ?", req.Status)
	}
	if !req.StartDate.IsZero() && !req.EndDate.IsZero() {
		query = query.Where("pr.created_at BETWEEN ? AND ?", req.StartDate, req.EndDate)
	}
	query.Count(&total)
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	offset := (req.Page - 1) * req.PageSize
	err := query.Order("pr.created_at DESC").Offset(offset).Limit(req.PageSize).Find(&items).Error
	return items, total, err
}

func (r *paymentRepository) CreateRefund(ctx context.Context, refund *model.PaymentRefund) error {
	return r.db.WithContext(ctx).Create(refund).Error
}

func (r *paymentRepository) GetRefundByID(ctx context.Context, id string) (*model.PaymentRefund, error) {
	var refund model.PaymentRefund
	err := r.db.WithContext(ctx).First(&refund, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &refund, nil
}

func (r *paymentRepository) GetRefundByPaymentID(ctx context.Context, paymentID string) (*model.PaymentRefund, error) {
	var refund model.PaymentRefund
	err := r.db.WithContext(ctx).First(&refund, "payment_id = ?", paymentID).Error
	if err != nil {
		return nil, err
	}
	return &refund, nil
}

func (r *paymentRepository) ListPaymentMethods(ctx context.Context) ([]*model.PaymentMethod, error) {
	var methods []*model.PaymentMethod
	err := r.db.WithContext(ctx).Find(&methods).Error
	return methods, err
}

func (r *paymentRepository) GetPaymentStats(ctx context.Context, req *model.PaymentStatsRequest) (*model.PaymentStats, error) {
	// 统计实现略，返回空结构体，后续可补充
	return &model.PaymentStats{}, nil
}
