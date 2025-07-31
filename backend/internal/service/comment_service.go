package service

import (
	"context"
	"errors"
	"math"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"

	"gorm.io/gorm"
)

// CommentService 评论服务接口
type CommentService interface {
	GetComments(ctx context.Context, postID string, req *model.CommentListRequest) (*model.CommentListResponse, error)
	CreateComment(ctx context.Context, userID, identityID, postID string, req *model.CreateCommentRequest) (*model.CreateCommentResponse, error)
	CreateReply(ctx context.Context, userID, identityID, commentID string, req *model.CreateReplyRequest) (*model.CreateReplyResponse, error)
	LikeComment(ctx context.Context, userID, commentID string) (*model.LikeCommentResponse, error)
	UnlikeComment(ctx context.Context, userID, commentID string) (*model.LikeCommentResponse, error)
	DeleteComment(ctx context.Context, commentID string) (*model.DeleteCommentResponse, error)
}

// commentService 评论服务实现
type commentService struct {
	commentRepo repository.CommentRepository
}

// NewCommentService 创建评论服务实例
func NewCommentService(commentRepo repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

// GetComments 获取评论列表
func (s *commentService) GetComments(ctx context.Context, postID string, req *model.CommentListRequest) (*model.CommentListResponse, error) {
	comments, total, err := s.commentRepo.GetComments(ctx, postID, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	commentInfos := make([]*model.CommentInfo, len(comments))
	for i, comment := range comments {
		commentInfos[i] = s.convertToCommentInfo(comment)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.CommentListResponse{
		Comments: commentInfos,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// CreateComment 创建评论
func (s *commentService) CreateComment(ctx context.Context, userID, identityID, postID string, req *model.CreateCommentRequest) (*model.CreateCommentResponse, error) {
	comment := &model.Comment{
		PostID:     postID,
		UserID:     userID,
		IdentityID: identityID,
		Content:    req.Content,
	}

	err := s.commentRepo.CreateComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &model.CreateCommentResponse{
		CommentID: comment.ID,
	}, nil
}

// CreateReply 创建回复
func (s *commentService) CreateReply(ctx context.Context, userID, identityID, commentID string, req *model.CreateReplyRequest) (*model.CreateReplyResponse, error) {
	// 先获取父评论信息
	parentComment, err := s.commentRepo.GetCommentByID(ctx, commentID)
	if err != nil {
		return nil, err
	}

	reply := &model.Comment{
		PostID:     parentComment.PostID,
		UserID:     userID,
		IdentityID: identityID,
		Content:    req.Content,
		ParentID:   &commentID,
	}

	err = s.commentRepo.CreateComment(ctx, reply)
	if err != nil {
		return nil, err
	}

	return &model.CreateReplyResponse{
		ReplyID: reply.ID,
	}, nil
}

// LikeComment 点赞评论
func (s *commentService) LikeComment(ctx context.Context, userID, commentID string) (*model.LikeCommentResponse, error) {
	err := s.commentRepo.LikeComment(ctx, userID, commentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("已经点赞过该评论")
		}
		return nil, err
	}

	// 获取更新后的点赞数量
	comment, err := s.commentRepo.GetCommentByID(ctx, commentID)
	if err != nil {
		return nil, err
	}

	return &model.LikeCommentResponse{
		CommentID: commentID,
		LikeCount: comment.LikeCount,
	}, nil
}

// UnlikeComment 取消点赞评论
func (s *commentService) UnlikeComment(ctx context.Context, userID, commentID string) (*model.LikeCommentResponse, error) {
	err := s.commentRepo.UnlikeComment(ctx, userID, commentID)
	if err != nil {
		return nil, err
	}

	// 获取更新后的点赞数量
	comment, err := s.commentRepo.GetCommentByID(ctx, commentID)
	if err != nil {
		return nil, err
	}

	return &model.LikeCommentResponse{
		CommentID: commentID,
		LikeCount: comment.LikeCount,
	}, nil
}

// DeleteComment 删除评论
func (s *commentService) DeleteComment(ctx context.Context, commentID string) (*model.DeleteCommentResponse, error) {
	err := s.commentRepo.DeleteComment(ctx, commentID)
	if err != nil {
		return nil, err
	}

	return &model.DeleteCommentResponse{
		CommentID: commentID,
	}, nil
}

// convertToCommentInfo 转换为评论信息
func (s *commentService) convertToCommentInfo(comment *model.Comment) *model.CommentInfo {
	commentInfo := &model.CommentInfo{
		ID:        comment.ID,
		Content:   comment.Content,
		LikeCount: comment.LikeCount,
		IsLiked:   false, // TODO: 根据用户是否点赞设置
		CreatedAt: comment.CreatedAt,
	}

	// 转换用户信息
	if comment.User != nil {
		commentInfo.User = &model.UserInfo{
			ID: comment.User.ID,
		}
		// 这里应该从用户档案中获取姓名和头像
		// 暂时使用邮箱作为姓名
		commentInfo.User.Email = comment.User.Email
	}

	// 转换身份信息
	if comment.Identity != nil {
		commentInfo.Identity = &model.CircleIdentityInfo{
			IdentityType: comment.Identity.IdentityType,
			Domain:       comment.Identity.Domain,
		}
	}

	// 转换回复信息
	if len(comment.Replies) > 0 {
		replyInfos := make([]*model.ReplyInfo, len(comment.Replies))
		for i, reply := range comment.Replies {
			replyInfos[i] = s.convertToReplyInfo(reply)
		}
		commentInfo.Replies = replyInfos
	}

	return commentInfo
}

// convertToReplyInfo 转换为回复信息
func (s *commentService) convertToReplyInfo(reply *model.Comment) *model.ReplyInfo {
	replyInfo := &model.ReplyInfo{
		ID:        reply.ID,
		Content:   reply.Content,
		LikeCount: reply.LikeCount,
		IsLiked:   false, // TODO: 根据用户是否点赞设置
		CreatedAt: reply.CreatedAt,
	}

	// 转换用户信息
	if reply.User != nil {
		replyInfo.User = &model.UserInfo{
			ID: reply.User.ID,
		}
		// 这里应该从用户档案中获取姓名和头像
		// 暂时使用邮箱作为姓名
		replyInfo.User.Email = reply.User.Email
	}

	return replyInfo
}
