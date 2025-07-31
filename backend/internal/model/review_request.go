package model

// CreateReviewRequest 创建评价请求
type CreateReviewRequest struct {
	ReviewedID    string `json:"reviewed_id" binding:"required"`
	CourseID      string `json:"course_id"`
	AppointmentID string `json:"appointment_id"`
	Rating        int    `json:"rating" binding:"required,min=1,max=5"`
	Content       string `json:"content" binding:"required"`
	ReviewType    string `json:"review_type" binding:"required,oneof=course mentor appointment"`
}

// ReviewListRequest 获取评价列表请求
type ReviewListRequest struct {
	PaginationRequest
	ReviewedID string `json:"reviewed_id" form:"reviewed_id"`
	ReviewType string `json:"review_type" form:"review_type"`
	Rating     int    `json:"rating" form:"rating"`
}

// UpdateReviewRequest 更新评价请求
type UpdateReviewRequest struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Content string `json:"content" binding:"required"`
}

// ReviewStatsRequest 获取评价统计请求
type ReviewStatsRequest struct {
	ReviewedID string `json:"reviewed_id" form:"reviewed_id" binding:"required"`
	ReviewType string `json:"review_type" form:"review_type" binding:"required"`
}
