package service

import (
	"context"
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
	"master-guide-backend/internal/utils"
)

// UploadService 文件上传服务接口
type UploadService interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, fileType string, userID string) (*model.UploadFileResponse, error)
	GetUploadFile(ctx context.Context, fileID string) (*model.UploadFile, error)
	DeleteUploadFile(ctx context.Context, fileID string, userID string) error
}

type uploadService struct {
	uploadRepo repository.UploadRepository
	fileUtils  *utils.FileUtils
}

func NewUploadService(uploadRepo repository.UploadRepository) UploadService {
	return &uploadService{
		uploadRepo: uploadRepo,
		fileUtils:  utils.NewFileUtils(),
	}
}

func (s *uploadService) UploadFile(ctx context.Context, file *multipart.FileHeader, fileType string, userID string) (*model.UploadFileResponse, error) {
	// 验证文件类型
	if !s.fileUtils.IsValidFileType(fileType) {
		return nil, errors.New("不支持的文件类型")
	}

	// 验证文件大小
	if file.Size > s.fileUtils.GetFileSize() {
		return nil, errors.New("文件大小超过限制")
	}

	// 验证文件扩展名
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !s.fileUtils.IsValidFileExtension(ext) {
		return nil, errors.New("不支持的文件格式")
	}

	// 保存文件
	filePath, fileURL, err := s.fileUtils.SaveUploadedFile(file, fileType)
	if err != nil {
		return nil, err
	}

	// 获取文件格式
	fileFormat := strings.ToLower(filepath.Ext(file.Filename))
	if fileFormat == "" {
		fileFormat = "unknown"
	} else {
		fileFormat = fileFormat[1:] // 去掉点号
	}

	// 创建上传文件记录
	uploadFile := &model.UploadFile{
		OriginalName: file.Filename,
		FilePath:     filePath,
		FileURL:      fileURL,
		FileType:     fileFormat,
		FileSize:     file.Size,
		MimeType:     file.Header.Get("Content-Type"),
		UserID:       userID,
		UploadType:   fileType,
	}

	err = s.uploadRepo.CreateUploadFile(ctx, uploadFile)
	if err != nil {
		return nil, err
	}

	return &model.UploadFileResponse{
		FileURL: fileURL,
		FileID:  uploadFile.ID,
	}, nil
}

func (s *uploadService) GetUploadFile(ctx context.Context, fileID string) (*model.UploadFile, error) {
	return s.uploadRepo.GetUploadFileByID(ctx, fileID)
}

func (s *uploadService) DeleteUploadFile(ctx context.Context, fileID string, userID string) error {
	file, err := s.uploadRepo.GetUploadFileByID(ctx, fileID)
	if err != nil {
		return err
	}

	// 验证文件所有者
	if file.UserID != userID {
		return errors.New("无权限删除此文件")
	}

	return s.uploadRepo.DeleteUploadFile(ctx, fileID)
}
