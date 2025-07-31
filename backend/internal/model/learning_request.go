package model

import "time"

// LearningRecordListRequest 获取学习记录列表请求
type LearningRecordListRequest struct {
	PaginationRequest
	CourseID  string    `json:"course_id" form:"course_id"`
	Status    string    `json:"status" form:"status"`
	StartDate time.Time `json:"start_date" form:"start_date" time:"2006-01-02"`
	EndDate   time.Time `json:"end_date" form:"end_date" time:"2006-01-02"`
}

// UpdateProgressRequest 更新学习进度请求
type UpdateProgressRequest struct {
	ProgressPercentage float64 `json:"progress_percentage" binding:"required,min=0,max=100"`
	CurrentChapter     string  `json:"current_chapter"`
	StudyTimeMinutes   int     `json:"study_time_minutes"`
}

// SubmitAssignmentRequest 提交作业请求
type SubmitAssignmentRequest struct {
	Title          string   `json:"title" binding:"required"`
	Content        string   `json:"content" binding:"required"`
	AttachmentURLs []string `json:"attachment_urls"`
}

// LearningStatsRequest 获取学习统计请求
type LearningStatsRequest struct {
	Period string `json:"period" form:"period" binding:"required,oneof=week month year all"`
}
