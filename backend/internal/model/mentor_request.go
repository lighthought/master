package model

// MentorListRequest 获取大师列表请求
type MentorListRequest struct {
	Domain    string  `form:"domain"`
	MinRating float64 `form:"min_rating" binding:"min=0,max=5"`
	MaxPrice  float64 `form:"max_price" binding:"min=0"`
	IsOnline  *bool   `form:"is_online"`
	Page      int     `form:"page" binding:"min=1"`
	PageSize  int     `form:"page_size" binding:"min=1,max=100"`
}

// MentorSearchRequest 搜索大师请求
type MentorSearchRequest struct {
	Query     string  `form:"q"`
	Domain    string  `form:"domain"`
	MinRating float64 `form:"min_rating" binding:"min=0,max=5"`
	MaxPrice  float64 `form:"max_price" binding:"min=0"`
	IsOnline  *bool   `form:"is_online"`
	Page      int     `form:"page" binding:"min=1"`
	PageSize  int     `form:"page_size" binding:"min=1,max=100"`
}

// MentorReviewsRequest 获取大师评价请求
type MentorReviewsRequest struct {
	Page     int `form:"page" binding:"min=1"`
	PageSize int `form:"page_size" binding:"min=1,max=100"`
}
