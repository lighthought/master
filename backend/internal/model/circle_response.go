package model

import "time"

// CircleInfo 圈子信息
type CircleInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	MemberCount int    `json:"member_count"`
	IsJoined    bool   `json:"is_joined"`
}

// RecommendedCircle 推荐圈子
type RecommendedCircle struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Domain               string `json:"domain"`
	MemberCount          int    `json:"member_count"`
	IsJoined             bool   `json:"is_joined"`
	RecommendationReason string `json:"recommendation_reason"`
}

// PostInfo 动态信息
type PostInfo struct {
	ID           string              `json:"id"`
	User         *UserInfo           `json:"user"`
	Identity     *CircleIdentityInfo `json:"identity"`
	Content      string              `json:"content"`
	MediaURLs    []string            `json:"media_urls"`
	PostType     string              `json:"post_type"`
	LikeCount    int                 `json:"like_count"`
	CommentCount int                 `json:"comment_count"`
	CreatedAt    time.Time           `json:"created_at"`
}

// CommentInfo 评论信息
type CommentInfo struct {
	ID        string              `json:"id"`
	User      *UserInfo           `json:"user"`
	Identity  *CircleIdentityInfo `json:"identity"`
	Content   string              `json:"content"`
	LikeCount int                 `json:"like_count"`
	IsLiked   bool                `json:"is_liked"`
	CreatedAt time.Time           `json:"created_at"`
	Replies   []*ReplyInfo        `json:"replies"`
}

// ReplyInfo 回复信息
type ReplyInfo struct {
	ID        string    `json:"id"`
	User      *UserInfo `json:"user"`
	Content   string    `json:"content"`
	LikeCount int       `json:"like_count"`
	IsLiked   bool      `json:"is_liked"`
	CreatedAt time.Time `json:"created_at"`
}

// CircleIdentityInfo 圈子身份信息
type CircleIdentityInfo struct {
	IdentityType string `json:"identity_type"`
	Domain       string `json:"domain"`
}

// CircleListResponse 圈子列表响应
type CircleListResponse struct {
	Circles []*CircleInfo `json:"circles"`
}

// RecommendedCirclesResponse 推荐圈子响应
type RecommendedCirclesResponse struct {
	Circles []*RecommendedCircle `json:"circles"`
}

// PostListResponse 动态列表响应
type PostListResponse struct {
	Posts      []*PostInfo         `json:"posts"`
	Pagination *PaginationResponse `json:"pagination"`
}

// CreatePostResponse 创建动态响应
type CreatePostResponse struct {
	PostID string `json:"post_id"`
}

// LikePostResponse 点赞动态响应
type LikePostResponse struct {
	PostID    string `json:"post_id"`
	LikeCount int    `json:"like_count"`
}

// CommentListResponse 评论列表响应
type CommentListResponse struct {
	Comments   []*CommentInfo      `json:"comments"`
	Pagination *PaginationResponse `json:"pagination"`
}

// CreateCommentResponse 创建评论响应
type CreateCommentResponse struct {
	CommentID string `json:"comment_id"`
}

// CreateReplyResponse 创建回复响应
type CreateReplyResponse struct {
	ReplyID string `json:"reply_id"`
}

// LikeCommentResponse 点赞评论响应
type LikeCommentResponse struct {
	CommentID string `json:"comment_id"`
	LikeCount int    `json:"like_count"`
}

// LikeReplyResponse 点赞回复响应
type LikeReplyResponse struct {
	ReplyID   string `json:"reply_id"`
	LikeCount int    `json:"like_count"`
}

// DeleteCommentResponse 删除评论响应
type DeleteCommentResponse struct {
	CommentID string `json:"comment_id"`
}

// DeleteReplyResponse 删除回复响应
type DeleteReplyResponse struct {
	ReplyID string `json:"reply_id"`
}

// JoinCircleResponse 加入圈子响应
type JoinCircleResponse struct {
	CircleID string `json:"circle_id"`
}

// LeaveCircleResponse 退出圈子响应
type LeaveCircleResponse struct {
	CircleID string `json:"circle_id"`
}
