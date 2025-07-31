package model

// StudentListRequest 获取学生列表请求
type StudentListRequest struct {
	PaginationRequest
	Status   string `json:"status" form:"status"`
	CourseID string `json:"course_id" form:"course_id"`
	Search   string `json:"search" form:"search"`
	SortBy   string `json:"sort_by" form:"sort_by" binding:"oneof=name enrollment_date progress"`
}

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	Content string `json:"content" binding:"required"`
	Type    string `json:"type" binding:"required,oneof=text image file"`
}

// StudentProgressRequest 更新学习进度请求
type StudentProgressRequest struct {
	ProgressPercentage float64 `json:"progress_percentage" binding:"required,min=0,max=100"`
	Notes              string  `json:"notes"`
}

// GradeAssignmentRequest 评价作业请求
type GradeAssignmentRequest struct {
	Score    float64 `json:"score" binding:"required,min=0,max=100"`
	Feedback string  `json:"feedback" binding:"required"`
	Comments string  `json:"comments"`
}

// StudentReportRequest 获取学习报告请求
type StudentReportRequest struct {
	Period string `json:"period" form:"period" binding:"required,oneof=week month quarter year"`
}
