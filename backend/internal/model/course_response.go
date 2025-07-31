package model

import "time"

// CourseInfo 课程信息
type CourseInfo struct {
	ID            string      `json:"id"`
	Title         string      `json:"title"`
	Description   string      `json:"description"`
	CoverImage    string      `json:"cover_image"`
	Price         float64     `json:"price"`
	DurationHours int         `json:"duration_hours"`
	Difficulty    string      `json:"difficulty"`
	StudentCount  int         `json:"student_count"`
	Rating        float64     `json:"rating"`
	Mentor        *MentorInfo `json:"mentor"`
}

// CourseDetail 课程详情
type CourseDetail struct {
	ID            string           `json:"id"`
	Title         string           `json:"title"`
	Description   string           `json:"description"`
	CoverImage    string           `json:"cover_image"`
	Price         float64          `json:"price"`
	DurationHours int              `json:"duration_hours"`
	Difficulty    string           `json:"difficulty"`
	StudentCount  int              `json:"student_count"`
	Rating        float64          `json:"rating"`
	Mentor        *MentorInfo      `json:"mentor"`
	Contents      []*CourseContent `json:"contents"`
	Reviews       []*CourseReview  `json:"reviews"`
}

// CourseContent 课程内容
type CourseContent struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	ContentType     string `json:"content_type"`
	DurationMinutes int    `json:"duration_minutes"`
	OrderIndex      int    `json:"order_index"`
}

// CourseReview 课程评价
type CourseReview struct {
	ID           string    `json:"id"`
	Rating       int       `json:"rating"`
	Content      string    `json:"content"`
	ReviewerName string    `json:"reviewer_name"`
	CreatedAt    time.Time `json:"created_at"`
}

// RecommendedCourse 推荐课程
type RecommendedCourse struct {
	ID                   string      `json:"id"`
	Title                string      `json:"title"`
	Description          string      `json:"description"`
	CoverImage           string      `json:"cover_image"`
	Price                float64     `json:"price"`
	DurationHours        int         `json:"duration_hours"`
	Difficulty           string      `json:"difficulty"`
	StudentCount         int         `json:"student_count"`
	Rating               float64     `json:"rating"`
	Mentor               *MentorInfo `json:"mentor"`
	RecommendationReason string      `json:"recommendation_reason"`
}

// EnrolledCourse 已报名课程
type EnrolledCourse struct {
	ID                 string      `json:"id"`
	Title              string      `json:"title"`
	Description        string      `json:"description"`
	CoverImage         string      `json:"cover_image"`
	Price              float64     `json:"price"`
	DurationHours      int         `json:"duration_hours"`
	Difficulty         string      `json:"difficulty"`
	StudentCount       int         `json:"student_count"`
	Rating             float64     `json:"rating"`
	Mentor             *MentorInfo `json:"mentor"`
	EnrollmentStatus   string      `json:"enrollment_status"`
	ProgressPercentage float64     `json:"progress_percentage"`
	EnrolledAt         time.Time   `json:"enrolled_at"`
	LastAccessedAt     time.Time   `json:"last_accessed_at"`
}

// CourseProgress 课程进度
type CourseProgress struct {
	CourseID           string    `json:"course_id"`
	ProgressPercentage float64   `json:"progress_percentage"`
	Status             string    `json:"status"`
	EnrolledAt         time.Time `json:"enrolled_at"`
	LastAccessedAt     time.Time `json:"last_accessed_at"`
	CompletedContents  []string  `json:"completed_contents"`
}

// CourseListResponse 课程列表响应
type CourseListResponse struct {
	Courses    []*CourseInfo       `json:"courses"`
	Pagination *PaginationResponse `json:"pagination"`
}

// CourseDetailResponse 课程详情响应
type CourseDetailResponse struct {
	Course *CourseDetail `json:"course"`
}

// CreateCourseResponse 创建课程响应
type CreateCourseResponse struct {
	CourseID string `json:"course_id"`
}

// EnrollCourseResponse 报名课程响应
type EnrollCourseResponse struct {
	EnrollmentID string `json:"enrollment_id"`
	CourseID     string `json:"course_id"`
	Status       string `json:"status"`
	PaymentURL   string `json:"payment_url"`
}

// CourseProgressResponse 课程进度响应
type CourseProgressResponse struct {
	Progress *CourseProgress `json:"progress"`
}

// CourseSearchResponse 课程搜索响应
type CourseSearchResponse struct {
	Courses    []*CourseInfo       `json:"courses"`
	Pagination *PaginationResponse `json:"pagination"`
}

// RecommendedCoursesResponse 推荐课程响应
type RecommendedCoursesResponse struct {
	Courses []*RecommendedCourse `json:"courses"`
}

// EnrolledCoursesResponse 已报名课程响应
type EnrolledCoursesResponse struct {
	Courses    []*EnrolledCourse   `json:"courses"`
	Pagination *PaginationResponse `json:"pagination"`
}
