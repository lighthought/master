package model

import "time"

// ReviewInfo 评价信息
type ReviewInfo struct {
	ID            string    `json:"id"`
	Reviewer      *UserInfo `json:"reviewer"`
	ReviewedID    string    `json:"reviewed_id"`
	CourseID      string    `json:"course_id,omitempty"`
	AppointmentID string    `json:"appointment_id,omitempty"`
	Rating        int       `json:"rating"`
	Content       string    `json:"content"`
	ReviewType    string    `json:"review_type"`
	CreatedAt     time.Time `json:"created_at"`
}

// ReviewDetail 评价详情
type ReviewDetail struct {
	ID            string    `json:"id"`
	Reviewer      *UserInfo `json:"reviewer"`
	ReviewedID    string    `json:"reviewed_id"`
	CourseID      string    `json:"course_id,omitempty"`
	AppointmentID string    `json:"appointment_id,omitempty"`
	Rating        int       `json:"rating"`
	Content       string    `json:"content"`
	ReviewType    string    `json:"review_type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// CreateReviewResponse 创建评价响应
type CreateReviewResponse struct {
	ReviewID string `json:"review_id"`
}

// UpdateReviewResponse 更新评价响应
type UpdateReviewResponse struct {
	ReviewID string `json:"review_id"`
}

// DeleteReviewResponse 删除评价响应
type DeleteReviewResponse struct {
	ReviewID string `json:"review_id"`
}

// ReviewListResponse 评价列表响应
type ReviewListResponse struct {
	Reviews    []*ReviewInfo       `json:"reviews"`
	Pagination *PaginationResponse `json:"pagination"`
}

// ReviewDetailResponse 评价详情响应
type ReviewDetailResponse struct {
	Review *ReviewDetail `json:"review"`
}

// RatingDistribution 评分分布
type RatingDistribution struct {
	Five  int `json:"5"`
	Four  int `json:"4"`
	Three int `json:"3"`
	Two   int `json:"2"`
	One   int `json:"1"`
}

// ReviewStats 评价统计
type ReviewStats struct {
	TotalReviews       int                 `json:"total_reviews"`
	AverageRating      float64             `json:"average_rating"`
	RatingDistribution *RatingDistribution `json:"rating_distribution"`
}

// ReviewStatsResponse 评价统计响应
type ReviewStatsResponse struct {
	Stats *ReviewStats `json:"stats"`
}
