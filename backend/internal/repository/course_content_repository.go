package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// CourseContentRepository 课程内容数据访问接口
type CourseContentRepository interface {
	GetCourseContents(ctx context.Context, courseID string) ([]*model.CourseContentModel, error)
	CreateCourseContent(ctx context.Context, content *model.CourseContentModel) error
	CreateBatchCourseContents(ctx context.Context, contents []*model.CourseContentModel) error
}

// courseContentRepository 课程内容数据访问实现
type courseContentRepository struct {
	db *gorm.DB
}

// NewCourseContentRepository 创建课程内容数据访问实例
func NewCourseContentRepository(db *gorm.DB) CourseContentRepository {
	return &courseContentRepository{db: db}
}

// GetCourseContents 获取课程内容
func (r *courseContentRepository) GetCourseContents(ctx context.Context, courseID string) ([]*model.CourseContentModel, error) {
	var contents []*model.CourseContentModel
	err := r.db.WithContext(ctx).
		Where("course_id = ?", courseID).
		Order("order_index ASC").
		Find(&contents).Error
	return contents, err
}

// CreateCourseContent 创建课程内容
func (r *courseContentRepository) CreateCourseContent(ctx context.Context, content *model.CourseContentModel) error {
	return r.db.WithContext(ctx).Create(content).Error
}

// CreateBatchCourseContents 批量创建课程内容
func (r *courseContentRepository) CreateBatchCourseContents(ctx context.Context, contents []*model.CourseContentModel) error {
	return r.db.WithContext(ctx).Create(contents).Error
}
