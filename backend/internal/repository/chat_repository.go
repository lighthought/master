package repository

import (
	"context"
	"master-guide-backend/internal/model"
	"time"

	"gorm.io/gorm"
)

// ChatRepository 聊天数据访问接口
type ChatRepository interface {
	GetChatMessages(ctx context.Context, targetID, circleID string, page, pageSize int) ([]*model.ChatMessage, int64, error)
	SaveMessage(ctx context.Context, message *model.ChatMessage) error
	GetUserProfile(ctx context.Context, userID string) (*model.ChatUserInfo, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) GetChatMessages(ctx context.Context, targetID, circleID string, page, pageSize int) ([]*model.ChatMessage, int64, error) {
	var messages []*Message
	var total int64

	query := r.db.WithContext(ctx).Table("messages m").
		Select(`m.id, m.content, m.message_type as message_type, m.created_at, 
		        up.name as from_user_name, m.from_user_id as from_user_id`).
		Joins("LEFT JOIN user_profiles up ON m.from_user_id = up.user_id")

	// 根据目标类型构建查询条件
	if targetID != "" {
		// 私聊消息
		query = query.Where("(m.from_user_id = ? AND m.to_user_id = ?) OR (m.from_user_id = ? AND m.to_user_id = ?)",
			targetID, targetID, targetID, targetID)
	} else if circleID != "" {
		// 圈子消息
		query = query.Where("m.circle_id = ?", circleID)
	}

	// 获取总数
	query.Count(&total)

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("m.created_at DESC").Offset(offset).Limit(pageSize).Find(&messages).Error
	if err != nil {
		return nil, 0, err
	}

	// 转换为 ChatMessage 格式
	var chatMessages []*model.ChatMessage
	for _, msg := range messages {
		createdAt, _ := time.Parse("2006-01-02 15:04:05", msg.CreatedAt)
		chatMessage := &model.ChatMessage{
			ID:        msg.ID,
			Content:   msg.Content,
			Type:      msg.MessageType,
			CreatedAt: createdAt,
			FromUser: &model.ChatUserInfo{
				ID:   msg.FromUserID,
				Name: msg.FromUserName,
			},
		}
		chatMessages = append(chatMessages, chatMessage)
	}

	return chatMessages, total, nil
}

func (r *chatRepository) SaveMessage(ctx context.Context, message *model.ChatMessage) error {
	// 保存到 messages 表
	dbMessage := &Message{
		ID:          message.ID,
		FromUserID:  message.FromUser.ID,
		Content:     message.Content,
		MessageType: message.Type,
		CreatedAt:   message.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	// 如果有目标用户，设置 to_user_id
	if message.ToUserID != "" {
		dbMessage.ToUserID = &message.ToUserID
	}

	// 如果有圈子，设置 circle_id
	if message.CircleID != "" {
		dbMessage.CircleID = &message.CircleID
	}

	return r.db.WithContext(ctx).Create(dbMessage).Error
}

func (r *chatRepository) GetUserProfile(ctx context.Context, userID string) (*model.ChatUserInfo, error) {
	var userInfo model.ChatUserInfo

	err := r.db.WithContext(ctx).
		Table("user_profiles").
		Select("user_id as id, name").
		Where("user_id = ?", userID).
		First(&userInfo).Error

	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}

// Message 数据库消息模型（用于内部使用）
type Message struct {
	ID           string  `json:"id" gorm:"primaryKey;type:varchar(32)"`
	FromUserID   string  `json:"from_user_id"`
	FromUserName string  `json:"from_user_name"`
	ToUserID     *string `json:"to_user_id"`
	CircleID     *string `json:"circle_id"`
	Content      string  `json:"content"`
	MessageType  string  `json:"message_type"`
	IsRead       bool    `json:"is_read"`
	CreatedAt    string  `json:"created_at"`
}
