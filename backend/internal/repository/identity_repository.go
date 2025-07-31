package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// IdentityRepository 身份数据访问接口
type IdentityRepository interface {
	Create(ctx context.Context, identity *model.UserIdentity) error
	GetByUserID(ctx context.Context, userID string) ([]*model.UserIdentity, error)
	GetByID(ctx context.Context, id string) (*model.UserIdentity, error)
	Update(ctx context.Context, identity *model.UserIdentity) error
	GetByUserIDAndType(ctx context.Context, userID, identityType, domain string) (*model.UserIdentity, error)
	GetIdentitiesWithProfile(ctx context.Context, userID string) ([]*model.UserIdentity, error)
	GetIdentityWithProfile(ctx context.Context, identityID string) (*model.UserIdentity, error)
}

// identityRepository 身份数据访问实现
type identityRepository struct {
	db *gorm.DB
}

// NewIdentityRepository 创建身份数据访问实例
func NewIdentityRepository(db *gorm.DB) IdentityRepository {
	return &identityRepository{db: db}
}

// Create 创建身份
func (r *identityRepository) Create(ctx context.Context, identity *model.UserIdentity) error {
	return r.db.WithContext(ctx).Create(identity).Error
}

// GetByUserID 根据用户ID获取身份列表
func (r *identityRepository) GetByUserID(ctx context.Context, userID string) ([]*model.UserIdentity, error) {
	var identities []*model.UserIdentity
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&identities).Error
	return identities, err
}

// GetByID 根据ID获取身份
func (r *identityRepository) GetByID(ctx context.Context, id string) (*model.UserIdentity, error) {
	var identity model.UserIdentity
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&identity).Error
	if err != nil {
		return nil, err
	}
	return &identity, nil
}

// Update 更新身份
func (r *identityRepository) Update(ctx context.Context, identity *model.UserIdentity) error {
	return r.db.WithContext(ctx).Save(identity).Error
}

// GetByUserIDAndType 根据用户ID、身份类型和领域获取身份
func (r *identityRepository) GetByUserIDAndType(ctx context.Context, userID, identityType, domain string) (*model.UserIdentity, error) {
	var identity model.UserIdentity
	err := r.db.WithContext(ctx).Where("user_id = ? AND identity_type = ? AND domain = ?", userID, identityType, domain).First(&identity).Error
	if err != nil {
		return nil, err
	}
	return &identity, nil
}

// GetIdentitiesWithProfile 获取用户身份列表及其档案信息
func (r *identityRepository) GetIdentitiesWithProfile(ctx context.Context, userID string) ([]*model.UserIdentity, error) {
	var identities []*model.UserIdentity
	err := r.db.WithContext(ctx).Preload("Profile").Where("user_id = ?", userID).Find(&identities).Error
	return identities, err
}

// GetIdentityWithProfile 获取身份及其档案信息
func (r *identityRepository) GetIdentityWithProfile(ctx context.Context, identityID string) (*model.UserIdentity, error) {
	var identity model.UserIdentity
	err := r.db.WithContext(ctx).Preload("Profile").Where("id = ?", identityID).First(&identity).Error
	if err != nil {
		return nil, err
	}
	return &identity, nil
}
