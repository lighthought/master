package model

// NotificationListRequest 获取通知列表请求
type NotificationListRequest struct {
	PaginationRequest
	Type   string `json:"type" form:"type"`
	Status string `json:"status" form:"status"` // unread, read, all
}

// MarkNotificationReadRequest 标记通知为已读请求
type MarkNotificationReadRequest struct {
	NotificationIDs []string `json:"notification_ids"`
	MarkAll         bool     `json:"mark_all"`
}

// DeleteNotificationRequest 删除通知请求
type DeleteNotificationRequest struct {
	NotificationIDs []string `json:"notification_ids"`
	DeleteAll       bool     `json:"delete_all"`
}

// NotificationSettingsRequest 通知设置请求
type NotificationSettingsRequest struct {
	EmailNotifications *NotificationTypeSettings `json:"email_notifications"`
	PushNotifications  *NotificationTypeSettings `json:"push_notifications"`
	InAppNotifications *NotificationTypeSettings `json:"in_app_notifications"`
	QuietHours         *QuietHoursSettings       `json:"quiet_hours"`
}

// NotificationTypeSettings 通知类型设置
type NotificationTypeSettings struct {
	Enabled bool     `json:"enabled"`
	Types   []string `json:"types"`
}

// QuietHoursSettings 静默时间设置
type QuietHoursSettings struct {
	Enabled   bool   `json:"enabled"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// SendNotificationRequest 发送通知请求
type SendNotificationRequest struct {
	UserIDs     []string               `json:"user_ids" binding:"required"`
	Type        string                 `json:"type" binding:"required"`
	Title       string                 `json:"title" binding:"required"`
	Content     string                 `json:"content" binding:"required"`
	RelatedData map[string]interface{} `json:"related_data"`
	SenderID    string                 `json:"sender_id"`
}
