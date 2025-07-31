package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// MessageRepository 消息数据访问接口
type MessageRepository interface {
	Create(ctx context.Context, message *model.Message) error
	GetMessages(ctx context.Context, fromUserID, toUserID string, page, pageSize int) ([]*model.MessageInfo, int64, error)
	GetByID(ctx context.Context, messageID string) (*model.Message, error)
	MarkAsRead(ctx context.Context, messageID string) error
}

// messageRepository 消息数据访问实现
type messageRepository struct {
	db *gorm.DB
}

// NewMessageRepository 创建消息数据访问实例
func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

// Create 创建消息
func (r *messageRepository) Create(ctx context.Context, message *model.Message) error {
	return r.db.WithContext(ctx).Create(message).Error
}

// GetMessages 获取消息列表
func (r *messageRepository) GetMessages(ctx context.Context, fromUserID, toUserID string, page, pageSize int) ([]*model.MessageInfo, int64, error) {
	var messages []*model.MessageInfo
	var total int64

	query := r.db.WithContext(ctx).
		Table("messages m").
		Select(`
			m.id, m.content, m.type, m.created_at,
			fu.id as from_user_id, fu.email as from_user_email,
			tu.id as to_user_id, tu.email as to_user_email
		`).
		Joins("LEFT JOIN users fu ON m.from_user_id = fu.id").
		Joins("LEFT JOIN users tu ON m.to_user_id = tu.id").
		Where("(m.from_user_id = ? AND m.to_user_id = ?) OR (m.from_user_id = ? AND m.to_user_id = ?)",
			fromUserID, toUserID, toUserID, fromUserID)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("m.created_at DESC").Offset(offset).Limit(pageSize).Find(&messages).Error

	return messages, total, err
}

// GetByID 根据ID获取消息
func (r *messageRepository) GetByID(ctx context.Context, messageID string) (*model.Message, error) {
	var message model.Message
	err := r.db.WithContext(ctx).
		Preload("FromUser").
		Preload("ToUser").
		Where("id = ?", messageID).
		First(&message).Error
	return &message, err
}

// MarkAsRead 标记消息为已读
func (r *messageRepository) MarkAsRead(ctx context.Context, messageID string) error {
	return r.db.WithContext(ctx).
		Model(&model.Message{}).
		Where("id = ?", messageID).
		Update("is_read", true).Error
}
