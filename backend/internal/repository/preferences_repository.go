package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// PreferencesRepository 用户偏好数据访问接口
type PreferencesRepository interface {
	Create(ctx context.Context, preferences *model.UserPreferences) error
	GetByUserID(ctx context.Context, userID string) (*model.UserPreferences, error)
	Update(ctx context.Context, preferences *model.UserPreferences) error
	Upsert(ctx context.Context, preferences *model.UserPreferences) error
}

// preferencesRepository 用户偏好数据访问实现
type preferencesRepository struct {
	db *gorm.DB
}

// NewPreferencesRepository 创建用户偏好数据访问实例
func NewPreferencesRepository(db *gorm.DB) PreferencesRepository {
	return &preferencesRepository{db: db}
}

// Create 创建用户偏好
func (r *preferencesRepository) Create(ctx context.Context, preferences *model.UserPreferences) error {
	return r.db.WithContext(ctx).Create(preferences).Error
}

// GetByUserID 根据用户ID获取用户偏好
func (r *preferencesRepository) GetByUserID(ctx context.Context, userID string) (*model.UserPreferences, error) {
	var preferences model.UserPreferences
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&preferences).Error
	if err != nil {
		return nil, err
	}
	return &preferences, nil
}

// Update 更新用户偏好
func (r *preferencesRepository) Update(ctx context.Context, preferences *model.UserPreferences) error {
	return r.db.WithContext(ctx).Save(preferences).Error
}

// Upsert 插入或更新用户偏好
func (r *preferencesRepository) Upsert(ctx context.Context, preferences *model.UserPreferences) error {
	return r.db.WithContext(ctx).Save(preferences).Error
}
