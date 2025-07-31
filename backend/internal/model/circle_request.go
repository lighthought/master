package model

// CircleListRequest 获取圈子列表请求
type CircleListRequest struct {
	PaginationRequest
	Domain string `json:"domain" form:"domain"`
}

// PostListRequest 获取动态列表请求
type PostListRequest struct {
	PaginationRequest
	PostType string `json:"post_type" form:"post_type"`
}

// CreatePostRequest 创建动态请求
type CreatePostRequest struct {
	Content   string   `json:"content" binding:"required"`
	MediaURLs []string `json:"media_urls"`
	PostType  string   `json:"post_type" binding:"required,oneof=text image video link"`
}

// CreateCommentRequest 创建评论请求
type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// CreateReplyRequest 创建回复请求
type CreateReplyRequest struct {
	Content string `json:"content" binding:"required"`
}

// CommentListRequest 获取评论列表请求
type CommentListRequest struct {
	PaginationRequest
}
