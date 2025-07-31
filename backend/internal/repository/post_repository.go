package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// PostRepository 动态数据访问接口
type PostRepository interface {
	GetPosts(ctx context.Context, circleID, postType string, page, pageSize int) ([]*model.Post, int64, error)
	GetPostByID(ctx context.Context, postID string) (*model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) error
	DeletePost(ctx context.Context, postID string) error
	LikePost(ctx context.Context, userID, postID string) error
	UnlikePost(ctx context.Context, userID, postID string) error
	IsPostLiked(ctx context.Context, userID, postID string) (bool, error)
}

// postRepository 动态数据访问实现
type postRepository struct {
	db *gorm.DB
}

// NewPostRepository 创建动态数据访问实例
func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

// GetPosts 获取动态列表
func (r *postRepository) GetPosts(ctx context.Context, circleID, postType string, page, pageSize int) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Post{}).
		Preload("User").
		Preload("User.Profile").
		Preload("Identity").
		Where("circle_id = ? AND status = ?", circleID, "active")

	if postType != "" {
		query = query.Where("post_type = ?", postType)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&posts).Error

	return posts, total, err
}

// GetPostByID 根据ID获取动态
func (r *postRepository) GetPostByID(ctx context.Context, postID string) (*model.Post, error) {
	var post model.Post
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("User.Profile").
		Preload("Identity").
		Where("id = ?", postID).
		First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// CreatePost 创建动态
func (r *postRepository) CreatePost(ctx context.Context, post *model.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

// DeletePost 删除动态
func (r *postRepository) DeletePost(ctx context.Context, postID string) error {
	return r.db.WithContext(ctx).Model(&model.Post{}).
		Where("id = ?", postID).
		Update("status", "deleted").Error
}

// LikePost 点赞动态
func (r *postRepository) LikePost(ctx context.Context, userID, postID string) error {
	// 检查是否已经点赞
	exists, err := r.IsPostLiked(ctx, userID, postID)
	if err != nil {
		return err
	}
	if exists {
		return gorm.ErrRecordNotFound // 使用这个错误表示已点赞
	}

	// 创建点赞记录
	like := &model.PostLike{
		PostID: postID,
		UserID: userID,
	}

	err = r.db.WithContext(ctx).Create(like).Error
	if err != nil {
		return err
	}

	// 更新动态点赞数量
	err = r.db.WithContext(ctx).Model(&model.Post{}).
		Where("id = ?", postID).
		UpdateColumn("like_count", gorm.Expr("like_count + 1")).Error

	return err
}

// UnlikePost 取消点赞动态
func (r *postRepository) UnlikePost(ctx context.Context, userID, postID string) error {
	// 删除点赞记录
	result := r.db.WithContext(ctx).Where("post_id = ? AND user_id = ?", postID, userID).
		Delete(&model.PostLike{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	// 更新动态点赞数量
	err := r.db.WithContext(ctx).Model(&model.Post{}).
		Where("id = ?", postID).
		UpdateColumn("like_count", gorm.Expr("like_count - 1")).Error

	return err
}

// IsPostLiked 检查用户是否已点赞动态
func (r *postRepository) IsPostLiked(ctx context.Context, userID, postID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.PostLike{}).
		Where("post_id = ? AND user_id = ?", postID, userID).
		Count(&count).Error
	return count > 0, err
}
