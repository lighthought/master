package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// MentorRepository 大师数据访问接口
type MentorRepository interface {
	GetMentors(ctx context.Context, domain string, minRating, maxPrice float64, isOnline *bool, page, pageSize int) ([]*model.Mentor, int64, error)
	GetMentorByID(ctx context.Context, mentorID string) (*model.Mentor, error)
	SearchMentors(ctx context.Context, query, domain string, minRating, maxPrice float64, isOnline *bool, page, pageSize int) ([]*model.Mentor, int64, error)
	GetRecommendedMentors(ctx context.Context, userID string, limit int) ([]*model.Mentor, error)
	GetMentorReviews(ctx context.Context, mentorID string, page, pageSize int) ([]*model.MentorReviewModel, int64, error)
	GetMentorCourses(ctx context.Context, mentorID string) ([]*model.Course, error)
}

// mentorRepository 大师数据访问实现
type mentorRepository struct {
	db *gorm.DB
}

// NewMentorRepository 创建大师数据访问实例
func NewMentorRepository(db *gorm.DB) MentorRepository {
	return &mentorRepository{db: db}
}

// GetMentors 获取大师列表
func (r *mentorRepository) GetMentors(ctx context.Context, domain string, minRating, maxPrice float64, isOnline *bool, page, pageSize int) ([]*model.Mentor, int64, error) {
	var mentors []*model.Mentor
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Mentor{}).
		Preload("Identity").
		Preload("User").
		Preload("Profile").
		Where("mentors.status = ?", "active")

	if domain != "" {
		query = query.Joins("JOIN user_identities ui ON mentors.identity_id = ui.id").
			Where("ui.domain = ?", domain)
	}

	if minRating > 0 {
		query = query.Where("mentors.rating >= ?", minRating)
	}

	if maxPrice > 0 {
		query = query.Where("mentors.hourly_rate <= ?", maxPrice)
	}

	if isOnline != nil {
		query = query.Where("mentors.is_online = ?", *isOnline)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&mentors).Error

	return mentors, total, err
}

// GetMentorByID 根据ID获取大师详情
func (r *mentorRepository) GetMentorByID(ctx context.Context, mentorID string) (*model.Mentor, error) {
	var mentor model.Mentor
	err := r.db.WithContext(ctx).
		Preload("Identity").
		Preload("User").
		Preload("Profile").
		Where("id = ?", mentorID).
		First(&mentor).Error
	if err != nil {
		return nil, err
	}
	return &mentor, nil
}

// SearchMentors 搜索大师
func (r *mentorRepository) SearchMentors(ctx context.Context, query, domain string, minRating, maxPrice float64, isOnline *bool, page, pageSize int) ([]*model.Mentor, int64, error) {
	var mentors []*model.Mentor
	var total int64

	dbQuery := r.db.WithContext(ctx).Model(&model.Mentor{}).
		Preload("Identity").
		Preload("User").
		Preload("Profile").
		Joins("JOIN user_identities ui ON mentors.identity_id = ui.id").
		Joins("JOIN user_profiles up ON ui.id = up.identity_id").
		Where("mentors.status = ?", "active")

	if query != "" {
		dbQuery = dbQuery.Where("up.name ILIKE ? OR up.bio ILIKE ? OR up.skills::text ILIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%")
	}

	if domain != "" {
		dbQuery = dbQuery.Where("ui.domain = ?", domain)
	}

	if minRating > 0 {
		dbQuery = dbQuery.Where("mentors.rating >= ?", minRating)
	}

	if maxPrice > 0 {
		dbQuery = dbQuery.Where("mentors.hourly_rate <= ?", maxPrice)
	}

	if isOnline != nil {
		dbQuery = dbQuery.Where("mentors.is_online = ?", *isOnline)
	}

	// 获取总数
	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := dbQuery.Offset(offset).Limit(pageSize).Find(&mentors).Error

	return mentors, total, err
}

// GetRecommendedMentors 获取推荐大师
func (r *mentorRepository) GetRecommendedMentors(ctx context.Context, userID string, limit int) ([]*model.Mentor, error) {
	var mentors []*model.Mentor

	// 这里可以实现更复杂的推荐算法
	// 目前简单按评分和学生数量排序
	err := r.db.WithContext(ctx).
		Preload("Identity").
		Preload("User").
		Preload("Profile").
		Where("mentors.status = ?", "active").
		Order("mentors.rating DESC, mentors.student_count DESC").
		Limit(limit).
		Find(&mentors).Error

	return mentors, err
}

// GetMentorReviews 获取大师评价
func (r *mentorRepository) GetMentorReviews(ctx context.Context, mentorID string, page, pageSize int) ([]*model.MentorReviewModel, int64, error) {
	var reviews []*model.MentorReviewModel
	var total int64

	query := r.db.WithContext(ctx).Model(&model.MentorReviewModel{}).
		Preload("Reviewer").
		Where("mentor_id = ?", mentorID)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&reviews).Error

	return reviews, total, err
}

// GetMentorCourses 获取大师课程
func (r *mentorRepository) GetMentorCourses(ctx context.Context, mentorID string) ([]*model.Course, error) {
	var courses []*model.Course
	err := r.db.WithContext(ctx).
		Where("mentor_id = ? AND status = ?", mentorID, "published").
		Order("created_at DESC").
		Find(&courses).Error
	return courses, err
}
