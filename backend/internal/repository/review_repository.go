package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// ReviewRepository 评价数据访问接口
type ReviewRepository interface {
	GetReviews(ctx context.Context, reviewedID, reviewType string, rating int, page, pageSize int) ([]*model.Review, int64, error)
	GetReviewByID(ctx context.Context, reviewID string) (*model.Review, error)
	CreateReview(ctx context.Context, review *model.Review) error
	UpdateReview(ctx context.Context, reviewID string, rating int, content string) error
	DeleteReview(ctx context.Context, reviewID string) error
	GetReviewStats(ctx context.Context, reviewedID, reviewType string) (*model.ReviewStats, error)
	CheckUserCanReview(ctx context.Context, reviewerID, reviewedID, reviewType string) (bool, error)
}

// reviewRepository 评价数据访问实现
type reviewRepository struct {
	db *gorm.DB
}

// NewReviewRepository 创建评价数据访问实例
func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

// GetReviews 获取评价列表
func (r *reviewRepository) GetReviews(ctx context.Context, reviewedID, reviewType string, rating int, page, pageSize int) ([]*model.Review, int64, error) {
	var reviews []*model.Review
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Review{}).
		Preload("Reviewer").
		Preload("Reviewer.Profile").
		Preload("Reviewed").
		Preload("Course").
		Preload("Appointment")

	if reviewedID != "" {
		query = query.Where("reviewed_id = ?", reviewedID)
	}

	if reviewType != "" {
		query = query.Where("review_type = ?", reviewType)
	}

	if rating > 0 {
		query = query.Where("rating = ?", rating)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&reviews).Error

	return reviews, total, err
}

// GetReviewByID 根据ID获取评价
func (r *reviewRepository) GetReviewByID(ctx context.Context, reviewID string) (*model.Review, error) {
	var review model.Review
	err := r.db.WithContext(ctx).
		Preload("Reviewer").
		Preload("Reviewer.Profile").
		Preload("Reviewed").
		Preload("Course").
		Preload("Appointment").
		Where("id = ?", reviewID).
		First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// CreateReview 创建评价
func (r *reviewRepository) CreateReview(ctx context.Context, review *model.Review) error {
	return r.db.WithContext(ctx).Create(review).Error
}

// UpdateReview 更新评价
func (r *reviewRepository) UpdateReview(ctx context.Context, reviewID string, rating int, content string) error {
	return r.db.WithContext(ctx).Model(&model.Review{}).
		Where("id = ?", reviewID).
		Updates(map[string]interface{}{
			"rating":  rating,
			"content": content,
		}).Error
}

// DeleteReview 删除评价
func (r *reviewRepository) DeleteReview(ctx context.Context, reviewID string) error {
	return r.db.WithContext(ctx).Delete(&model.Review{}, reviewID).Error
}

// GetReviewStats 获取评价统计
func (r *reviewRepository) GetReviewStats(ctx context.Context, reviewedID, reviewType string) (*model.ReviewStats, error) {
	var totalReviews int64
	var averageRating float64
	var ratingDistribution struct {
		Five  int64 `gorm:"column:five"`
		Four  int64 `gorm:"column:four"`
		Three int64 `gorm:"column:three"`
		Two   int64 `gorm:"column:two"`
		One   int64 `gorm:"column:one"`
	}

	// 获取总评价数和平均评分
	err := r.db.WithContext(ctx).Model(&model.Review{}).
		Where("reviewed_id = ? AND review_type = ?", reviewedID, reviewType).
		Select("COUNT(*) as total_reviews, AVG(rating) as average_rating").
		Scan(&struct {
			TotalReviews  int64   `gorm:"column:total_reviews"`
			AverageRating float64 `gorm:"column:average_rating"`
		}{TotalReviews: totalReviews, AverageRating: averageRating}).Error
	if err != nil {
		return nil, err
	}

	// 获取评分分布
	err = r.db.WithContext(ctx).Model(&model.Review{}).
		Where("reviewed_id = ? AND review_type = ?", reviewedID, reviewType).
		Select(`
			SUM(CASE WHEN rating = 5 THEN 1 ELSE 0 END) as five,
			SUM(CASE WHEN rating = 4 THEN 1 ELSE 0 END) as four,
			SUM(CASE WHEN rating = 3 THEN 1 ELSE 0 END) as three,
			SUM(CASE WHEN rating = 2 THEN 1 ELSE 0 END) as two,
			SUM(CASE WHEN rating = 1 THEN 1 ELSE 0 END) as one
		`).
		Scan(&ratingDistribution).Error
	if err != nil {
		return nil, err
	}

	return &model.ReviewStats{
		TotalReviews:  int(totalReviews),
		AverageRating: averageRating,
		RatingDistribution: &model.RatingDistribution{
			Five:  int(ratingDistribution.Five),
			Four:  int(ratingDistribution.Four),
			Three: int(ratingDistribution.Three),
			Two:   int(ratingDistribution.Two),
			One:   int(ratingDistribution.One),
		},
	}, nil
}

// CheckUserCanReview 检查用户是否可以评价
func (r *reviewRepository) CheckUserCanReview(ctx context.Context, reviewerID, reviewedID, reviewType string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Review{}).
		Where("reviewer_id = ? AND reviewed_id = ? AND review_type = ?", reviewerID, reviewedID, reviewType).
		Count(&count).Error
	return count == 0, err // 如果没有评价过，则可以评价
}
