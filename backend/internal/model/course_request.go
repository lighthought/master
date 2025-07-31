package model

// CourseListRequest 获取课程列表请求
type CourseListRequest struct {
	PaginationRequest
	Domain     string  `json:"domain" form:"domain"`
	Difficulty string  `json:"difficulty" form:"difficulty"`
	MinPrice   float64 `json:"min_price" form:"min_price"`
	MaxPrice   float64 `json:"max_price" form:"max_price"`
	SortBy     string  `json:"sort_by" form:"sort_by"`
}

// CourseSearchRequest 搜索课程请求
type CourseSearchRequest struct {
	PaginationRequest
	Query      string  `json:"q" form:"q"`
	Domain     string  `json:"domain" form:"domain"`
	Difficulty string  `json:"difficulty" form:"difficulty"`
	MinPrice   float64 `json:"min_price" form:"min_price"`
	MaxPrice   float64 `json:"max_price" form:"max_price"`
	SortBy     string  `json:"sort_by" form:"sort_by"`
}

// CreateCourseRequest 创建课程请求
type CreateCourseRequest struct {
	Title         string                `json:"title" binding:"required"`
	Description   string                `json:"description"`
	CoverImage    string                `json:"cover_image"`
	Price         float64               `json:"price" binding:"required,min=0"`
	DurationHours int                   `json:"duration_hours" binding:"required,min=1"`
	Difficulty    string                `json:"difficulty"`
	MaxStudents   *int                  `json:"max_students"`
	Contents      []*CourseContentInput `json:"contents"`
}

// CourseContentInput 课程内容输入
type CourseContentInput struct {
	Title           string `json:"title" binding:"required"`
	ContentType     string `json:"content_type" binding:"required"`
	ContentURL      string `json:"content_url"`
	ContentText     string `json:"content_text"`
	DurationMinutes int    `json:"duration_minutes"`
	OrderIndex      int    `json:"order_index" binding:"required"`
}

// EnrollCourseRequest 报名课程请求
type EnrollCourseRequest struct {
	PaymentMethod string          `json:"payment_method" binding:"required"`
	UserInfo      *EnrollUserInfo `json:"user_info"`
}

// EnrollUserInfo 报名用户信息
type EnrollUserInfo struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

// EnrolledCoursesRequest 获取已报名课程请求
type EnrolledCoursesRequest struct {
	PaginationRequest
	Status string `json:"status" form:"status"`
}
