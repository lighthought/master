package model

import "time"

// CreateAppointmentRequest 创建预约请求
type CreateAppointmentRequest struct {
	MentorID        string    `json:"mentor_id" binding:"required"`
	AppointmentTime time.Time `json:"appointment_time" binding:"required"`
	DurationMinutes int       `json:"duration_minutes" binding:"required,min=15,max=480"`
	MeetingType     string    `json:"meeting_type" binding:"required,oneof=video voice text"`
	Notes           string    `json:"notes"`
}

// AppointmentListRequest 获取预约列表请求
type AppointmentListRequest struct {
	PaginationRequest
	Status string `json:"status" form:"status"`
	Type   string `json:"type" form:"type"` // student/mentor
}

// UpdateAppointmentStatusRequest 更新预约状态请求
type UpdateAppointmentStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending confirmed completed cancelled"`
}
