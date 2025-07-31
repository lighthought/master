package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// ProfileRepository 用户档案数据访问接口
type ProfileRepository interface {
	Create(ctx context.Context, profile *model.UserProfile) error
	GetByIdentityID(ctx context.Context, identityID string) (*model.UserProfile, error)
	Update(ctx context.Context, profile *model.UserProfile) error
	GetByUserID(ctx context.Context, userID string) ([]*model.UserProfile, error)
}

// profileRepository 用户档案数据访问实现
type profileRepository struct {
	db *gorm.DB
}

// NewProfileRepository 创建用户档案数据访问实例
func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

// Create 创建用户档案
func (r *profileRepository) Create(ctx context.Context, profile *model.UserProfile) error {
	return r.db.WithContext(ctx).Create(profile).Error
}

// GetByIdentityID 根据身份ID获取用户档案
func (r *profileRepository) GetByIdentityID(ctx context.Context, identityID string) (*model.UserProfile, error) {
	var profile model.UserProfile
	err := r.db.WithContext(ctx).Where("identity_id = ?", identityID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// Update 更新用户档案
func (r *profileRepository) Update(ctx context.Context, profile *model.UserProfile) error {
	return r.db.WithContext(ctx).Save(profile).Error
}

// GetByUserID 根据用户ID获取用户档案列表
func (r *profileRepository) GetByUserID(ctx context.Context, userID string) ([]*model.UserProfile, error) {
	var profiles []*model.UserProfile
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&profiles).Error
	return profiles, err
}
