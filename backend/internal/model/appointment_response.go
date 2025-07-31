package model

import "time"

// AppointmentInfo 预约信息
type AppointmentInfo struct {
	ID              string      `json:"id"`
	Mentor          *MentorInfo `json:"mentor"`
	Student         *UserInfo   `json:"student"`
	AppointmentTime time.Time   `json:"appointment_time"`
	DurationMinutes int         `json:"duration_minutes"`
	MeetingType     string      `json:"meeting_type"`
	Status          string      `json:"status"`
	Price           float64     `json:"price"`
}

// AppointmentDetail 预约详情
type AppointmentDetail struct {
	ID              string      `json:"id"`
	Mentor          *MentorInfo `json:"mentor"`
	Student         *UserInfo   `json:"student"`
	AppointmentTime time.Time   `json:"appointment_time"`
	DurationMinutes int         `json:"duration_minutes"`
	MeetingType     string      `json:"meeting_type"`
	Status          string      `json:"status"`
	Price           float64     `json:"price"`
	Notes           string      `json:"notes"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

// CreateAppointmentResponse 创建预约响应
type CreateAppointmentResponse struct {
	AppointmentID string  `json:"appointment_id"`
	Status        string  `json:"status"`
	Price         float64 `json:"price"`
}

// UpdateAppointmentStatusResponse 更新预约状态响应
type UpdateAppointmentStatusResponse struct {
	AppointmentID string `json:"appointment_id"`
	Status        string `json:"status"`
}

// CancelAppointmentResponse 取消预约响应
type CancelAppointmentResponse struct {
	AppointmentID string `json:"appointment_id"`
}

// AppointmentListResponse 预约列表响应
type AppointmentListResponse struct {
	Appointments []*AppointmentInfo  `json:"appointments"`
	Pagination   *PaginationResponse `json:"pagination"`
}

// AppointmentDetailResponse 预约详情响应
type AppointmentDetailResponse struct {
	Appointment *AppointmentDetail `json:"appointment"`
}

// MentorAppointmentStats 大师预约统计
type MentorAppointmentStats struct {
	TotalAppointments     int     `json:"total_appointments"`
	PendingAppointments   int     `json:"pending_appointments"`
	ConfirmedAppointments int     `json:"confirmed_appointments"`
	CompletedAppointments int     `json:"completed_appointments"`
	CancelledAppointments int     `json:"cancelled_appointments"`
	TotalEarnings         float64 `json:"total_earnings"`
	AverageRating         float64 `json:"average_rating"`
}

// MentorAppointmentStatsResponse 大师预约统计响应
type MentorAppointmentStatsResponse struct {
	Stats *MentorAppointmentStats `json:"stats"`
}
