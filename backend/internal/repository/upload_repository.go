package repository

import (
	"context"
	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// UploadRepository 文件上传数据访问接口
type UploadRepository interface {
	CreateUploadFile(ctx context.Context, file *model.UploadFile) error
	GetUploadFileByID(ctx context.Context, id string) (*model.UploadFile, error)
	GetUploadFilesByUserID(ctx context.Context, userID string) ([]*model.UploadFile, error)
	DeleteUploadFile(ctx context.Context, id string) error
}

type uploadRepository struct {
	db *gorm.DB
}

func NewUploadRepository(db *gorm.DB) UploadRepository {
	return &uploadRepository{db: db}
}

func (r *uploadRepository) CreateUploadFile(ctx context.Context, file *model.UploadFile) error {
	return r.db.WithContext(ctx).Create(file).Error
}

func (r *uploadRepository) GetUploadFileByID(ctx context.Context, id string) (*model.UploadFile, error) {
	var file model.UploadFile
	err := r.db.WithContext(ctx).First(&file, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *uploadRepository) GetUploadFilesByUserID(ctx context.Context, userID string) ([]*model.UploadFile, error) {
	var files []*model.UploadFile
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&files).Error
	return files, err
}

func (r *uploadRepository) DeleteUploadFile(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.UploadFile{}, "id = ?", id).Error
}
