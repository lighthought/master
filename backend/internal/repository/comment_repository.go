package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// CommentRepository 评论数据访问接口
type CommentRepository interface {
	GetComments(ctx context.Context, postID string, page, pageSize int) ([]*model.Comment, int64, error)
	GetCommentByID(ctx context.Context, commentID string) (*model.Comment, error)
	CreateComment(ctx context.Context, comment *model.Comment) error
	DeleteComment(ctx context.Context, commentID string) error
	LikeComment(ctx context.Context, userID, commentID string) error
	UnlikeComment(ctx context.Context, userID, commentID string) error
	IsCommentLiked(ctx context.Context, userID, commentID string) (bool, error)
}

// commentRepository 评论数据访问实现
type commentRepository struct {
	db *gorm.DB
}

// NewCommentRepository 创建评论数据访问实例
func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

// GetComments 获取评论列表
func (r *commentRepository) GetComments(ctx context.Context, postID string, page, pageSize int) ([]*model.Comment, int64, error) {
	var comments []*model.Comment
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Comment{}).
		Preload("User").
		Preload("User.Profile").
		Preload("Identity").
		Preload("Replies").
		Preload("Replies.User").
		Preload("Replies.User.Profile").
		Where("post_id = ? AND parent_id IS NULL", postID)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&comments).Error

	return comments, total, err
}

// GetCommentByID 根据ID获取评论
func (r *commentRepository) GetCommentByID(ctx context.Context, commentID string) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("User.Profile").
		Preload("Identity").
		Preload("Replies").
		Preload("Replies.User").
		Preload("Replies.User.Profile").
		Where("id = ?", commentID).
		First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// CreateComment 创建评论
func (r *commentRepository) CreateComment(ctx context.Context, comment *model.Comment) error {
	err := r.db.WithContext(ctx).Create(comment).Error
	if err != nil {
		return err
	}

	// 更新动态评论数量
	err = r.db.WithContext(ctx).Model(&model.Post{}).
		Where("id = ?", comment.PostID).
		UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error

	return err
}

// DeleteComment 删除评论
func (r *commentRepository) DeleteComment(ctx context.Context, commentID string) error {
	// 先获取评论信息
	comment, err := r.GetCommentByID(ctx, commentID)
	if err != nil {
		return err
	}

	// 删除评论
	err = r.db.WithContext(ctx).Delete(&model.Comment{}, commentID).Error
	if err != nil {
		return err
	}

	// 更新动态评论数量
	err = r.db.WithContext(ctx).Model(&model.Post{}).
		Where("id = ?", comment.PostID).
		UpdateColumn("comment_count", gorm.Expr("comment_count - 1")).Error

	return err
}

// LikeComment 点赞评论
func (r *commentRepository) LikeComment(ctx context.Context, userID, commentID string) error {
	// 检查是否已经点赞
	exists, err := r.IsCommentLiked(ctx, userID, commentID)
	if err != nil {
		return err
	}
	if exists {
		return gorm.ErrRecordNotFound // 使用这个错误表示已点赞
	}

	// 创建点赞记录
	like := &model.CommentLike{
		CommentID: commentID,
		UserID:    userID,
	}

	err = r.db.WithContext(ctx).Create(like).Error
	if err != nil {
		return err
	}

	// 更新评论点赞数量
	err = r.db.WithContext(ctx).Model(&model.Comment{}).
		Where("id = ?", commentID).
		UpdateColumn("like_count", gorm.Expr("like_count + 1")).Error

	return err
}

// UnlikeComment 取消点赞评论
func (r *commentRepository) UnlikeComment(ctx context.Context, userID, commentID string) error {
	// 删除点赞记录
	result := r.db.WithContext(ctx).Where("comment_id = ? AND user_id = ?", commentID, userID).
		Delete(&model.CommentLike{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	// 更新评论点赞数量
	err := r.db.WithContext(ctx).Model(&model.Comment{}).
		Where("id = ?", commentID).
		UpdateColumn("like_count", gorm.Expr("like_count - 1")).Error

	return err
}

// IsCommentLiked 检查用户是否已点赞评论
func (r *commentRepository) IsCommentLiked(ctx context.Context, userID, commentID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.CommentLike{}).
		Where("comment_id = ? AND user_id = ?", commentID, userID).
		Count(&count).Error
	return count > 0, err
}
