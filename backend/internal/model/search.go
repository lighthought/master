package model

// SearchRequest 搜索请求
type SearchRequest struct {
	Query    string `form:"q" binding:"required"`
	Type     string `form:"type" binding:"omitempty,oneof=mentors courses posts"`
	Domain   string `form:"domain"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

// SearchResult 搜索结果
type SearchResult struct {
	Mentors      []*MentorSearchItem `json:"mentors"`
	Courses      []*CourseSearchItem `json:"courses"`
	Posts        []*PostSearchItem   `json:"posts"`
	TotalResults int64               `json:"total_results"`
}

// MentorSearchItem 导师搜索结果
type MentorSearchItem struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Domain       string  `json:"domain"`
	Rating       float64 `json:"rating"`
	StudentCount int     `json:"student_count"`
	HourlyRate   float64 `json:"hourly_rate"`
	IsOnline     bool    `json:"is_online"`
	Avatar       string  `json:"avatar"`
}

// CourseSearchItem 课程搜索结果
type CourseSearchItem struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	MentorName  string  `json:"mentor_name"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
	Difficulty  string  `json:"difficulty"`
	CoverImage  string  `json:"cover_image"`
}

// PostSearchItem 帖子搜索结果
type PostSearchItem struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	Circle    string `json:"circle"`
	LikeCount int    `json:"like_count"`
	CreatedAt string `json:"created_at"`
}
