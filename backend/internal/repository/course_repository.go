package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// CourseRepository 课程数据访问接口
type CourseRepository interface {
	GetCourses(ctx context.Context, domain, difficulty string, minPrice, maxPrice float64, sortBy string, page, pageSize int) ([]*model.Course, int64, error)
	GetCourseByID(ctx context.Context, courseID string) (*model.Course, error)
	CreateCourse(ctx context.Context, course *model.Course) error
	SearchCourses(ctx context.Context, query, domain, difficulty string, minPrice, maxPrice float64, sortBy string, page, pageSize int) ([]*model.Course, int64, error)
	GetRecommendedCourses(ctx context.Context, userID string, limit int) ([]*model.Course, error)
	GetEnrolledCourses(ctx context.Context, userID, status string, page, pageSize int) ([]*model.Course, int64, error)
	EnrollCourse(ctx context.Context, userID, courseID string) error
	GetCourseProgress(ctx context.Context, userID, courseID string) (*model.LearningRecordModel, error)
	GetCompletedContents(ctx context.Context, userID, courseID string) ([]string, error)
}

// courseRepository 课程数据访问实现
type courseRepository struct {
	db *gorm.DB
}

// NewCourseRepository 创建课程数据访问实例
func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

// GetCourses 获取课程列表
func (r *courseRepository) GetCourses(ctx context.Context, domain, difficulty string, minPrice, maxPrice float64, sortBy string, page, pageSize int) ([]*model.Course, int64, error) {
	var courses []*model.Course
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Course{}).
		Preload("Mentor").
		Preload("Mentor.Profile").
		Where("courses.status = ?", "published")

	if domain != "" {
		query = query.Joins("JOIN mentors m ON courses.mentor_id = m.id").
			Joins("JOIN user_identities ui ON m.identity_id = ui.id").
			Where("ui.domain = ?", domain)
	}

	if difficulty != "" {
		query = query.Where("courses.difficulty = ?", difficulty)
	}

	if minPrice > 0 {
		query = query.Where("courses.price >= ?", minPrice)
	}

	if maxPrice > 0 {
		query = query.Where("courses.price <= ?", maxPrice)
	}

	// 排序
	switch sortBy {
	case "rating":
		query = query.Order("courses.rating DESC")
	case "price":
		query = query.Order("courses.price ASC")
	case "created_at":
		query = query.Order("courses.created_at DESC")
	default:
		query = query.Order("courses.created_at DESC")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&courses).Error

	return courses, total, err
}

// GetCourseByID 根据ID获取课程详情
func (r *courseRepository) GetCourseByID(ctx context.Context, courseID string) (*model.Course, error) {
	var course model.Course
	err := r.db.WithContext(ctx).
		Preload("Mentor").
		Preload("Mentor.Profile").
		Where("id = ?", courseID).
		First(&course).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

// CreateCourse 创建课程
func (r *courseRepository) CreateCourse(ctx context.Context, course *model.Course) error {
	return r.db.WithContext(ctx).Create(course).Error
}

// SearchCourses 搜索课程
func (r *courseRepository) SearchCourses(ctx context.Context, query, domain, difficulty string, minPrice, maxPrice float64, sortBy string, page, pageSize int) ([]*model.Course, int64, error) {
	var courses []*model.Course
	var total int64

	dbQuery := r.db.WithContext(ctx).Model(&model.Course{}).
		Preload("Mentor").
		Preload("Mentor.Profile").
		Joins("JOIN mentors m ON courses.mentor_id = m.id").
		Joins("JOIN user_identities ui ON m.identity_id = ui.id").
		Joins("JOIN user_profiles up ON ui.id = up.identity_id").
		Where("courses.status = ?", "published")

	if query != "" {
		dbQuery = dbQuery.Where("courses.title ILIKE ? OR courses.description ILIKE ? OR up.name ILIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%")
	}

	if domain != "" {
		dbQuery = dbQuery.Where("ui.domain = ?", domain)
	}

	if difficulty != "" {
		dbQuery = dbQuery.Where("courses.difficulty = ?", difficulty)
	}

	if minPrice > 0 {
		dbQuery = dbQuery.Where("courses.price >= ?", minPrice)
	}

	if maxPrice > 0 {
		dbQuery = dbQuery.Where("courses.price <= ?", maxPrice)
	}

	// 排序
	switch sortBy {
	case "rating":
		dbQuery = dbQuery.Order("courses.rating DESC")
	case "price":
		dbQuery = dbQuery.Order("courses.price ASC")
	case "created_at":
		dbQuery = dbQuery.Order("courses.created_at DESC")
	default:
		dbQuery = dbQuery.Order("courses.created_at DESC")
	}

	// 获取总数
	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := dbQuery.Offset(offset).Limit(pageSize).Find(&courses).Error

	return courses, total, err
}

// GetRecommendedCourses 获取推荐课程
func (r *courseRepository) GetRecommendedCourses(ctx context.Context, userID string, limit int) ([]*model.Course, error) {
	var courses []*model.Course

	// 这里可以实现更复杂的推荐算法
	// 目前简单按评分和学生数量排序
	err := r.db.WithContext(ctx).
		Preload("Mentor").
		Preload("Mentor.Profile").
		Where("courses.status = ?", "published").
		Order("courses.rating DESC, courses.student_count DESC").
		Limit(limit).
		Find(&courses).Error

	return courses, err
}

// GetEnrolledCourses 获取已报名课程
func (r *courseRepository) GetEnrolledCourses(ctx context.Context, userID, status string, page, pageSize int) ([]*model.Course, int64, error) {
	var courses []*model.Course
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Course{}).
		Preload("Mentor").
		Preload("Mentor.Profile").
		Joins("JOIN learning_records lr ON courses.id = lr.course_id").
		Where("lr.user_id = ?", userID)

	if status != "" {
		query = query.Where("lr.status = ?", status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("lr.enrolled_at DESC").Offset(offset).Limit(pageSize).Find(&courses).Error

	return courses, total, err
}

// EnrollCourse 报名课程
func (r *courseRepository) EnrollCourse(ctx context.Context, userID, courseID string) error {
	// 检查是否已经报名
	var count int64
	err := r.db.WithContext(ctx).Model(&model.LearningRecordModel{}).
		Where("user_id = ? AND course_id = ?", userID, courseID).
		Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return gorm.ErrRecordNotFound // 使用这个错误表示已报名
	}

	// 创建学习记录
	learningRecord := &model.LearningRecordModel{
		UserID:   userID,
		CourseID: courseID,
		Status:   "enrolled",
	}

	return r.db.WithContext(ctx).Create(learningRecord).Error
}

// GetCourseProgress 获取课程进度
func (r *courseRepository) GetCourseProgress(ctx context.Context, userID, courseID string) (*model.LearningRecordModel, error) {
	var record model.LearningRecordModel
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND course_id = ?", userID, courseID).
		First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// GetCompletedContents 获取已完成的内容
func (r *courseRepository) GetCompletedContents(ctx context.Context, userID, courseID string) ([]string, error) {
	var contentIDs []string
	err := r.db.WithContext(ctx).Model(&model.ContentProgressModel{}).
		Joins("JOIN course_contents cc ON content_progress.content_id = cc.id").
		Where("content_progress.user_id = ? AND cc.course_id = ? AND content_progress.is_completed = ?",
			userID, courseID, true).
		Pluck("content_progress.content_id", &contentIDs).Error
	return contentIDs, err
}
