package model

import "time"

// MentorInfo 大师信息
type MentorInfo struct {
	ID           string   `json:"id"`
	IdentityID   string   `json:"identity_id"`
	Name         string   `json:"name"`
	Avatar       string   `json:"avatar"`
	Domain       string   `json:"domain"`
	Rating       float64  `json:"rating"`
	StudentCount int      `json:"student_count"`
	HourlyRate   float64  `json:"hourly_rate"`
	IsOnline     bool     `json:"is_online"`
	Skills       []string `json:"skills"`
	Bio          string   `json:"bio"`
}

// MentorDetail 大师详情
type MentorDetail struct {
	ID              string          `json:"id"`
	IdentityID      string          `json:"identity_id"`
	Name            string          `json:"name"`
	Avatar          string          `json:"avatar"`
	Domain          string          `json:"domain"`
	Rating          float64         `json:"rating"`
	StudentCount    int             `json:"student_count"`
	HourlyRate      float64         `json:"hourly_rate"`
	IsOnline        bool            `json:"is_online"`
	Skills          []string        `json:"skills"`
	Bio             string          `json:"bio"`
	ExperienceYears int             `json:"experience_years"`
	Courses         []*MentorCourse `json:"courses"`
	Reviews         []*MentorReview `json:"reviews"`
}

// MentorCourse 大师课程
type MentorCourse struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	Price        float64 `json:"price"`
	StudentCount int     `json:"student_count"`
}

// MentorReview 大师评价
type MentorReview struct {
	ID             string    `json:"id"`
	Rating         int       `json:"rating"`
	Content        string    `json:"content"`
	ReviewerName   string    `json:"reviewer_name"`
	ReviewerAvatar string    `json:"reviewer_avatar"`
	CreatedAt      time.Time `json:"created_at"`
}

// RecommendedMentor 推荐大师
type RecommendedMentor struct {
	ID                   string   `json:"id"`
	IdentityID           string   `json:"identity_id"`
	Name                 string   `json:"name"`
	Avatar               string   `json:"avatar"`
	Domain               string   `json:"domain"`
	Rating               float64  `json:"rating"`
	StudentCount         int      `json:"student_count"`
	HourlyRate           float64  `json:"hourly_rate"`
	IsOnline             bool     `json:"is_online"`
	Skills               []string `json:"skills"`
	Bio                  string   `json:"bio"`
	RecommendationReason string   `json:"recommendation_reason"`
}

// MentorListResponse 大师列表响应
type MentorListResponse struct {
	Mentors    []*MentorInfo       `json:"mentors"`
	Pagination *PaginationResponse `json:"pagination"`
}

// MentorDetailResponse 大师详情响应
type MentorDetailResponse struct {
	Mentor *MentorDetail `json:"mentor"`
}

// MentorSearchResponse 大师搜索响应
type MentorSearchResponse struct {
	Mentors    []*MentorInfo       `json:"mentors"`
	Pagination *PaginationResponse `json:"pagination"`
}

// RecommendedMentorsResponse 推荐大师响应
type RecommendedMentorsResponse struct {
	Mentors []*RecommendedMentor `json:"mentors"`
}

// MentorReviewsResponse 大师评价响应
type MentorReviewsResponse struct {
	Reviews    []*MentorReview     `json:"reviews"`
	Pagination *PaginationResponse `json:"pagination"`
}
