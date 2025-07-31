package model

import "time"

// NotificationInfo 通知信息
type NotificationInfo struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Title       string                 `json:"title"`
	Content     string                 `json:"content"`
	Status      string                 `json:"status"`
	Sender      *UserInfo              `json:"sender,omitempty"`
	RelatedData map[string]interface{} `json:"related_data,omitempty"`
	CreatedAt   time.Time              `json:"created_at"`
	ReadAt      *time.Time             `json:"read_at,omitempty"`
}

// NotificationListResponse 通知列表响应
type NotificationListResponse struct {
	Notifications []*NotificationInfo `json:"notifications"`
	Pagination    *PaginationResponse `json:"pagination"`
	UnreadCount   int                 `json:"unread_count"`
}

// MarkNotificationReadResponse 标记通知为已读响应
type MarkNotificationReadResponse struct {
	NotificationID string `json:"notification_id"`
}

// BatchMarkReadResponse 批量标记为已读响应
type BatchMarkReadResponse struct {
	MarkedCount int `json:"marked_count"`
}

// DeleteNotificationResponse 删除通知响应
type DeleteNotificationResponse struct {
	NotificationID string `json:"notification_id"`
}

// BatchDeleteResponse 批量删除响应
type BatchDeleteResponse struct {
	DeletedCount int `json:"deleted_count"`
}

// UnreadCountResponse 未读通知数量响应
type UnreadCountResponse struct {
	UnreadCount int            `json:"unread_count"`
	CountByType map[string]int `json:"count_by_type"`
}

// NotificationSettingsResponse 通知设置响应
type NotificationSettingsResponse struct {
	Settings *NotificationSettings `json:"settings"`
}

// NotificationSettings 通知设置
type NotificationSettings struct {
	EmailNotifications *NotificationTypeSettings `json:"email_notifications"`
	PushNotifications  *NotificationTypeSettings `json:"push_notifications"`
	InAppNotifications *NotificationTypeSettings `json:"in_app_notifications"`
	QuietHours         *QuietHoursSettings       `json:"quiet_hours"`
}

// UpdateSettingsResponse 更新设置响应
type UpdateSettingsResponse struct {
	SettingsUpdated bool `json:"settings_updated"`
}

// SendNotificationResponse 发送通知响应
type SendNotificationResponse struct {
	SentCount       int      `json:"sent_count"`
	NotificationIDs []string `json:"notification_ids"`
}
