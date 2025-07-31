package service

import (
	"context"
	"errors"
	"math"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// NotificationService 通知服务接口
type NotificationService interface {
	GetNotifications(ctx context.Context, userID string, req *model.NotificationListRequest) (*model.NotificationListResponse, error)
	MarkNotificationRead(ctx context.Context, userID, notificationID string) (*model.MarkNotificationReadResponse, error)
	BatchMarkNotificationsRead(ctx context.Context, userID string, req *model.MarkNotificationReadRequest) (*model.BatchMarkReadResponse, error)
	DeleteNotification(ctx context.Context, userID, notificationID string) (*model.DeleteNotificationResponse, error)
	BatchDeleteNotifications(ctx context.Context, userID string, req *model.DeleteNotificationRequest) (*model.BatchDeleteResponse, error)
	GetUnreadCount(ctx context.Context, userID string) (*model.UnreadCountResponse, error)
	GetNotificationSettings(ctx context.Context, userID string) (*model.NotificationSettingsResponse, error)
	UpdateNotificationSettings(ctx context.Context, userID string, req *model.NotificationSettingsRequest) (*model.UpdateSettingsResponse, error)
	SendNotification(ctx context.Context, req *model.SendNotificationRequest) (*model.SendNotificationResponse, error)
}

// notificationService 通知服务实现
type notificationService struct {
	notificationRepo repository.NotificationRepository
}

// NewNotificationService 创建通知服务实例
func NewNotificationService(notificationRepo repository.NotificationRepository) NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
	}
}

// GetNotifications 获取通知列表
func (s *notificationService) GetNotifications(ctx context.Context, userID string, req *model.NotificationListRequest) (*model.NotificationListResponse, error) {
	notifications, total, err := s.notificationRepo.GetNotifications(ctx, userID, req.Type, req.Status, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	notificationInfos := make([]*model.NotificationInfo, len(notifications))
	for i, notification := range notifications {
		notificationInfos[i] = s.convertToNotificationInfo(notification)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	// 获取未读数量
	unreadCount, _, err := s.notificationRepo.GetUnreadCount(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.NotificationListResponse{
		Notifications: notificationInfos,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
		UnreadCount: unreadCount,
	}, nil
}

// MarkNotificationRead 标记通知为已读
func (s *notificationService) MarkNotificationRead(ctx context.Context, userID, notificationID string) (*model.MarkNotificationReadResponse, error) {
	// 检查通知是否存在且属于当前用户
	notification, err := s.notificationRepo.GetNotificationByID(ctx, notificationID)
	if err != nil {
		return nil, err
	}

	if notification.UserID != userID {
		return nil, errors.New("只能操作自己的通知")
	}

	err = s.notificationRepo.MarkNotificationRead(ctx, notificationID)
	if err != nil {
		return nil, err
	}

	return &model.MarkNotificationReadResponse{
		NotificationID: notificationID,
	}, nil
}

// BatchMarkNotificationsRead 批量标记通知为已读
func (s *notificationService) BatchMarkNotificationsRead(ctx context.Context, userID string, req *model.MarkNotificationReadRequest) (*model.BatchMarkReadResponse, error) {
	markedCount, err := s.notificationRepo.BatchMarkNotificationsRead(ctx, userID, req.NotificationIDs, req.MarkAll)
	if err != nil {
		return nil, err
	}

	return &model.BatchMarkReadResponse{
		MarkedCount: markedCount,
	}, nil
}

// DeleteNotification 删除通知
func (s *notificationService) DeleteNotification(ctx context.Context, userID, notificationID string) (*model.DeleteNotificationResponse, error) {
	// 检查通知是否存在且属于当前用户
	notification, err := s.notificationRepo.GetNotificationByID(ctx, notificationID)
	if err != nil {
		return nil, err
	}

	if notification.UserID != userID {
		return nil, errors.New("只能删除自己的通知")
	}

	err = s.notificationRepo.DeleteNotification(ctx, notificationID)
	if err != nil {
		return nil, err
	}

	return &model.DeleteNotificationResponse{
		NotificationID: notificationID,
	}, nil
}

// BatchDeleteNotifications 批量删除通知
func (s *notificationService) BatchDeleteNotifications(ctx context.Context, userID string, req *model.DeleteNotificationRequest) (*model.BatchDeleteResponse, error) {
	deletedCount, err := s.notificationRepo.BatchDeleteNotifications(ctx, userID, req.NotificationIDs, req.DeleteAll)
	if err != nil {
		return nil, err
	}

	return &model.BatchDeleteResponse{
		DeletedCount: deletedCount,
	}, nil
}

// GetUnreadCount 获取未读通知数量
func (s *notificationService) GetUnreadCount(ctx context.Context, userID string) (*model.UnreadCountResponse, error) {
	unreadCount, countByType, err := s.notificationRepo.GetUnreadCount(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.UnreadCountResponse{
		UnreadCount: unreadCount,
		CountByType: countByType,
	}, nil
}

// GetNotificationSettings 获取通知设置
func (s *notificationService) GetNotificationSettings(ctx context.Context, userID string) (*model.NotificationSettingsResponse, error) {
	settings, err := s.notificationRepo.GetNotificationSettings(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.NotificationSettingsResponse{
		Settings: settings,
	}, nil
}

// UpdateNotificationSettings 更新通知设置
func (s *notificationService) UpdateNotificationSettings(ctx context.Context, userID string, req *model.NotificationSettingsRequest) (*model.UpdateSettingsResponse, error) {
	settings := &model.NotificationSettings{
		EmailNotifications: req.EmailNotifications,
		PushNotifications:  req.PushNotifications,
		InAppNotifications: req.InAppNotifications,
		QuietHours:         req.QuietHours,
	}

	err := s.notificationRepo.UpdateNotificationSettings(ctx, userID, settings)
	if err != nil {
		return nil, err
	}

	return &model.UpdateSettingsResponse{
		SettingsUpdated: true,
	}, nil
}

// SendNotification 发送通知
func (s *notificationService) SendNotification(ctx context.Context, req *model.SendNotificationRequest) (*model.SendNotificationResponse, error) {
	notifications := make([]*model.Notification, len(req.UserIDs))
	notificationIDs := make([]string, 0, len(req.UserIDs))

	for i, userID := range req.UserIDs {
		notification := &model.Notification{
			UserID:           userID,
			Title:            req.Title,
			Content:          req.Content,
			NotificationType: req.Type,
			IsRead:           false,
			Metadata:         model.JSONMap(req.RelatedData),
		}

		notifications[i] = notification
	}

	err := s.notificationRepo.BatchCreateNotifications(ctx, notifications)
	if err != nil {
		return nil, err
	}

	// 收集创建的通知ID
	for _, notification := range notifications {
		notificationIDs = append(notificationIDs, notification.ID)
	}

	return &model.SendNotificationResponse{
		SentCount:       len(req.UserIDs),
		NotificationIDs: notificationIDs,
	}, nil
}

// convertToNotificationInfo 转换为通知信息
func (s *notificationService) convertToNotificationInfo(notification *model.Notification) *model.NotificationInfo {
	notificationInfo := &model.NotificationInfo{
		ID:          notification.ID,
		Type:        notification.NotificationType,
		Title:       notification.Title,
		Content:     notification.Content,
		Status:      s.getNotificationStatus(notification.IsRead),
		RelatedData: map[string]interface{}(notification.Metadata),
		CreatedAt:   notification.CreatedAt,
		ReadAt:      notification.ReadAt,
	}

	// 转换发送者信息
	if notification.User != nil {
		notificationInfo.Sender = &model.UserInfo{
			ID: notification.User.ID,
		}
		// 这里应该从用户档案中获取姓名和头像
		// 暂时使用邮箱作为姓名
		notificationInfo.Sender.Email = notification.User.Email
	}

	return notificationInfo
}

// getNotificationStatus 获取通知状态
func (s *notificationService) getNotificationStatus(isRead bool) string {
	if isRead {
		return "read"
	}
	return "unread"
}
