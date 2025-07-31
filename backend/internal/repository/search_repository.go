package repository

import (
	"context"
	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// SearchRepository 搜索数据访问接口
type SearchRepository interface {
	SearchMentors(ctx context.Context, query, domain string, page, pageSize int) ([]*model.MentorSearchItem, int64, error)
	SearchCourses(ctx context.Context, query, domain string, page, pageSize int) ([]*model.CourseSearchItem, int64, error)
	SearchPosts(ctx context.Context, query, domain string, page, pageSize int) ([]*model.PostSearchItem, int64, error)
}

type searchRepository struct {
	db *gorm.DB
}

func NewSearchRepository(db *gorm.DB) SearchRepository {
	return &searchRepository{db: db}
}

func (r *searchRepository) SearchMentors(ctx context.Context, query, domain string, page, pageSize int) ([]*model.MentorSearchItem, int64, error) {
	var mentors []*model.MentorSearchItem
	var total int64

	dbQuery := r.db.WithContext(ctx).
		Table("mentors m").
		Select(`m.id, up.name, ui.domain, m.rating, m.student_count, m.hourly_rate, m.is_online, up.avatar`).
		Joins("LEFT JOIN user_identities ui ON m.identity_id = ui.id").
		Joins("LEFT JOIN user_profiles up ON ui.id = up.identity_id").
		Where("ui.identity_type = ? AND ui.status = ?", "master", "active")

	// 添加搜索条件
	if query != "" {
		dbQuery = dbQuery.Where("up.name ILIKE ? OR ui.domain ILIKE ?", "%"+query+"%", "%"+query+"%")
	}
	if domain != "" {
		dbQuery = dbQuery.Where("ui.domain = ?", domain)
	}

	// 获取总数
	dbQuery.Count(&total)

	// 分页查询
	offset := (page - 1) * pageSize
	err := dbQuery.Order("m.rating DESC").Offset(offset).Limit(pageSize).Find(&mentors).Error

	return mentors, total, err
}

func (r *searchRepository) SearchCourses(ctx context.Context, query, domain string, page, pageSize int) ([]*model.CourseSearchItem, int64, error) {
	var courses []*model.CourseSearchItem
	var total int64

	dbQuery := r.db.WithContext(ctx).
		Table("courses c").
		Select(`c.id, c.title, c.description, up.name as mentor_name, c.price, c.rating, c.difficulty, c.cover_image`).
		Joins("LEFT JOIN mentors m ON c.mentor_id = m.id").
		Joins("LEFT JOIN user_identities ui ON m.identity_id = ui.id").
		Joins("LEFT JOIN user_profiles up ON ui.id = up.identity_id").
		Where("c.status = ?", "published")

	// 添加搜索条件
	if query != "" {
		dbQuery = dbQuery.Where("c.title ILIKE ? OR c.description ILIKE ?", "%"+query+"%", "%"+query+"%")
	}
	if domain != "" {
		dbQuery = dbQuery.Where("ui.domain = ?", domain)
	}

	// 获取总数
	dbQuery.Count(&total)

	// 分页查询
	offset := (page - 1) * pageSize
	err := dbQuery.Order("c.rating DESC").Offset(offset).Limit(pageSize).Find(&courses).Error

	return courses, total, err
}

func (r *searchRepository) SearchPosts(ctx context.Context, query, domain string, page, pageSize int) ([]*model.PostSearchItem, int64, error) {
	var posts []*model.PostSearchItem
	var total int64

	dbQuery := r.db.WithContext(ctx).
		Table("posts p").
		Select(`p.id, p.content, up.name as author, c.name as circle, p.like_count, p.created_at`).
		Joins("LEFT JOIN user_profiles up ON p.user_id = up.user_id").
		Joins("LEFT JOIN circles c ON p.circle_id = c.id").
		Where("p.status = ?", "active")

	// 添加搜索条件
	if query != "" {
		dbQuery = dbQuery.Where("p.content ILIKE ?", "%"+query+"%")
	}
	if domain != "" {
		dbQuery = dbQuery.Where("c.domain = ?", domain)
	}

	// 获取总数
	dbQuery.Count(&total)

	// 分页查询
	offset := (page - 1) * pageSize
	err := dbQuery.Order("p.created_at DESC").Offset(offset).Limit(pageSize).Find(&posts).Error

	return posts, total, err
}
