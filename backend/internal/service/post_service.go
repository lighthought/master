package service

import (
	"context"
	"errors"
	"math"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"

	"gorm.io/gorm"
)

// PostService 动态服务接口
type PostService interface {
	GetPosts(ctx context.Context, circleID string, req *model.PostListRequest) (*model.PostListResponse, error)
	CreatePost(ctx context.Context, userID, identityID, circleID string, req *model.CreatePostRequest) (*model.CreatePostResponse, error)
	LikePost(ctx context.Context, userID, postID string) (*model.LikePostResponse, error)
	UnlikePost(ctx context.Context, userID, postID string) (*model.LikePostResponse, error)
}

// postService 动态服务实现
type postService struct {
	postRepo repository.PostRepository
}

// NewPostService 创建动态服务实例
func NewPostService(postRepo repository.PostRepository) PostService {
	return &postService{
		postRepo: postRepo,
	}
}

// GetPosts 获取动态列表
func (s *postService) GetPosts(ctx context.Context, circleID string, req *model.PostListRequest) (*model.PostListResponse, error) {
	posts, total, err := s.postRepo.GetPosts(ctx, circleID, req.PostType, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	postInfos := make([]*model.PostInfo, len(posts))
	for i, post := range posts {
		postInfos[i] = s.convertToPostInfo(post)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.PostListResponse{
		Posts: postInfos,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// CreatePost 创建动态
func (s *postService) CreatePost(ctx context.Context, userID, identityID, circleID string, req *model.CreatePostRequest) (*model.CreatePostResponse, error) {
	post := &model.Post{
		UserID:     userID,
		IdentityID: identityID,
		CircleID:   circleID,
		Content:    req.Content,
		MediaURLs:  req.MediaURLs,
		PostType:   req.PostType,
		Status:     "active",
	}

	err := s.postRepo.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return &model.CreatePostResponse{
		PostID: post.ID,
	}, nil
}

// LikePost 点赞动态
func (s *postService) LikePost(ctx context.Context, userID, postID string) (*model.LikePostResponse, error) {
	err := s.postRepo.LikePost(ctx, userID, postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("已经点赞过该动态")
		}
		return nil, err
	}

	// 获取更新后的点赞数量
	post, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	return &model.LikePostResponse{
		PostID:    postID,
		LikeCount: post.LikeCount,
	}, nil
}

// UnlikePost 取消点赞动态
func (s *postService) UnlikePost(ctx context.Context, userID, postID string) (*model.LikePostResponse, error) {
	err := s.postRepo.UnlikePost(ctx, userID, postID)
	if err != nil {
		return nil, err
	}

	// 获取更新后的点赞数量
	post, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	return &model.LikePostResponse{
		PostID:    postID,
		LikeCount: post.LikeCount,
	}, nil
}

// convertToPostInfo 转换为动态信息
func (s *postService) convertToPostInfo(post *model.Post) *model.PostInfo {
	postInfo := &model.PostInfo{
		ID:           post.ID,
		Content:      post.Content,
		MediaURLs:    post.MediaURLs,
		PostType:     post.PostType,
		LikeCount:    post.LikeCount,
		CommentCount: post.CommentCount,
		CreatedAt:    post.CreatedAt,
	}

	// 转换用户信息
	if post.User != nil {
		postInfo.User = &model.UserInfo{
			ID: post.User.ID,
		}
		// 这里应该从用户档案中获取姓名和头像
		// 暂时使用邮箱作为姓名
		postInfo.User.Email = post.User.Email
	}

	// 转换身份信息
	if post.Identity != nil {
		postInfo.Identity = &model.CircleIdentityInfo{
			IdentityType: post.Identity.IdentityType,
			Domain:       post.Identity.Domain,
		}
	}

	return postInfo
}
