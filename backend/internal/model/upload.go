package model

import "time"

// UploadFile 文件上传模型
type UploadFile struct {
	ID           string    `json:"id" gorm:"primaryKey;type:varchar(32)"`
	OriginalName string    `json:"original_name"`
	FilePath     string    `json:"file_path"`
	FileURL      string    `json:"file_url"`
	FileType     string    `json:"file_type"`
	FileSize     int64     `json:"file_size"`
	MimeType     string    `json:"mime_type"`
	UserID       string    `json:"user_id"`
	UploadType   string    `json:"upload_type"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (UploadFile) TableName() string {
	return "upload_files"
}

// UploadFileRequest 文件上传请求
type UploadFileRequest struct {
	Type string `form:"type" binding:"required,oneof=avatar course_cover post_image"`
}

// UploadFileResponse 文件上传响应
type UploadFileResponse struct {
	FileURL string `json:"file_url"`
	FileID  string `json:"file_id"`
}
