package repository

import (
	"context"
	"time"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// NotificationRepository 通知数据访问接口
type NotificationRepository interface {
	GetNotifications(ctx context.Context, userID, notificationType, status string, page, pageSize int) ([]*model.Notification, int64, error)
	GetNotificationByID(ctx context.Context, notificationID string) (*model.Notification, error)
	CreateNotification(ctx context.Context, notification *model.Notification) error
	BatchCreateNotifications(ctx context.Context, notifications []*model.Notification) error
	MarkNotificationRead(ctx context.Context, notificationID string) error
	BatchMarkNotificationsRead(ctx context.Context, userID string, notificationIDs []string, markAll bool) (int, error)
	DeleteNotification(ctx context.Context, notificationID string) error
	BatchDeleteNotifications(ctx context.Context, userID string, notificationIDs []string, deleteAll bool) (int, error)
	GetUnreadCount(ctx context.Context, userID string) (int, map[string]int, error)
	GetNotificationSettings(ctx context.Context, userID string) (*model.NotificationSettings, error)
	UpdateNotificationSettings(ctx context.Context, userID string, settings *model.NotificationSettings) error
}

// notificationRepository 通知数据访问实现
type notificationRepository struct {
	db *gorm.DB
}

// NewNotificationRepository 创建通知数据访问实例
func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

// GetNotifications 获取通知列表
func (r *notificationRepository) GetNotifications(ctx context.Context, userID, notificationType, status string, page, pageSize int) ([]*model.Notification, int64, error) {
	var notifications []*model.Notification
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Notification{}).
		Preload("User").
		Preload("User.Profile").
		Where("user_id = ?", userID)

	if notificationType != "" {
		query = query.Where("notification_type = ?", notificationType)
	}

	if status != "" && status != "all" {
		if status == "unread" {
			query = query.Where("is_read = ?", false)
		} else if status == "read" {
			query = query.Where("is_read = ?", true)
		}
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&notifications).Error

	return notifications, total, err
}

// GetNotificationByID 根据ID获取通知
func (r *notificationRepository) GetNotificationByID(ctx context.Context, notificationID string) (*model.Notification, error) {
	var notification model.Notification
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("User.Profile").
		Where("id = ?", notificationID).
		First(&notification).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

// CreateNotification 创建通知
func (r *notificationRepository) CreateNotification(ctx context.Context, notification *model.Notification) error {
	return r.db.WithContext(ctx).Create(notification).Error
}

// BatchCreateNotifications 批量创建通知
func (r *notificationRepository) BatchCreateNotifications(ctx context.Context, notifications []*model.Notification) error {
	if len(notifications) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).CreateInBatches(notifications, 100).Error
}

// MarkNotificationRead 标记通知为已读
func (r *notificationRepository) MarkNotificationRead(ctx context.Context, notificationID string) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&model.Notification{}).
		Where("id = ?", notificationID).
		Updates(map[string]interface{}{
			"is_read": true,
			"read_at": &now,
		}).Error
}

// BatchMarkNotificationsRead 批量标记通知为已读
func (r *notificationRepository) BatchMarkNotificationsRead(ctx context.Context, userID string, notificationIDs []string, markAll bool) (int, error) {
	now := time.Now()
	query := r.db.WithContext(ctx).Model(&model.Notification{}).Where("user_id = ? AND is_read = ?", userID, false)

	if !markAll && len(notificationIDs) > 0 {
		query = query.Where("id IN ?", notificationIDs)
	}

	result := query.Updates(map[string]interface{}{
		"is_read": true,
		"read_at": &now,
	})

	return int(result.RowsAffected), result.Error
}

// DeleteNotification 删除通知
func (r *notificationRepository) DeleteNotification(ctx context.Context, notificationID string) error {
	return r.db.WithContext(ctx).Delete(&model.Notification{}, notificationID).Error
}

// BatchDeleteNotifications 批量删除通知
func (r *notificationRepository) BatchDeleteNotifications(ctx context.Context, userID string, notificationIDs []string, deleteAll bool) (int, error) {
	query := r.db.WithContext(ctx).Model(&model.Notification{}).Where("user_id = ?", userID)

	if !deleteAll && len(notificationIDs) > 0 {
		query = query.Where("id IN ?", notificationIDs)
	}

	result := query.Delete(&model.Notification{})
	return int(result.RowsAffected), result.Error
}

// GetUnreadCount 获取未读通知数量
func (r *notificationRepository) GetUnreadCount(ctx context.Context, userID string) (int, map[string]int, error) {
	var totalUnread int64
	var countByType []struct {
		NotificationType string `gorm:"column:notification_type"`
		Count            int64  `gorm:"column:count"`
	}

	// 获取总未读数量
	err := r.db.WithContext(ctx).Model(&model.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Count(&totalUnread).Error
	if err != nil {
		return 0, nil, err
	}

	// 获取各类型未读数量
	err = r.db.WithContext(ctx).Model(&model.Notification{}).
		Select("notification_type, COUNT(*) as count").
		Where("user_id = ? AND is_read = ?", userID, false).
		Group("notification_type").
		Scan(&countByType).Error
	if err != nil {
		return 0, nil, err
	}

	// 转换为map
	countMap := make(map[string]int)
	for _, item := range countByType {
		countMap[item.NotificationType] = int(item.Count)
	}

	return int(totalUnread), countMap, nil
}

// GetNotificationSettings 获取通知设置
func (r *notificationRepository) GetNotificationSettings(ctx context.Context, userID string) (*model.NotificationSettings, error) {
	// 这里可以从用户偏好设置表或系统配置表获取
	// 暂时返回默认设置
	return &model.NotificationSettings{
		EmailNotifications: &model.NotificationTypeSettings{
			Enabled: true,
			Types:   []string{"message", "system", "activity"},
		},
		PushNotifications: &model.NotificationTypeSettings{
			Enabled: true,
			Types:   []string{"message", "system"},
		},
		InAppNotifications: &model.NotificationTypeSettings{
			Enabled: true,
			Types:   []string{"message", "system", "activity", "reminder"},
		},
		QuietHours: &model.QuietHoursSettings{
			Enabled:   false,
			StartTime: "22:00",
			EndTime:   "08:00",
		},
	}, nil
}

// UpdateNotificationSettings 更新通知设置
func (r *notificationRepository) UpdateNotificationSettings(ctx context.Context, userID string, settings *model.NotificationSettings) error {
	// 这里可以更新用户偏好设置表或系统配置表
	// 暂时返回成功
	return nil
}
