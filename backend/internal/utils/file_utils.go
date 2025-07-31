package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileUtils 文件工具类
type FileUtils struct{}

// NewFileUtils 创建文件工具实例
func NewFileUtils() *FileUtils {
	return &FileUtils{}
}

// SaveUploadedFile 保存上传的文件
func (f *FileUtils) SaveUploadedFile(file *multipart.FileHeader, uploadType string) (string, string, error) {
	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return "", "", fmt.Errorf("打开文件失败: %w", err)
	}
	defer src.Close()

	// 获取文件扩展名
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext == "" {
		ext = ".unknown"
	}

	// 去掉点号作为格式
	format := ext[1:]

	// 获取当前日期
	now := time.Now()
	month := fmt.Sprintf("%02d", now.Month())
	day := fmt.Sprintf("%02d", now.Day())

	// 生成文件名
	fileName := f.generateGUID() + ext

	// 构建目标目录路径：data/format/MM/DD
	targetDir := filepath.Join("data", format, month, day)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return "", "", fmt.Errorf("创建目标目录失败: %w", err)
	}

	targetPath := filepath.Join(targetDir, fileName)

	// 创建目标文件
	dst, err := os.Create(targetPath)
	if err != nil {
		return "", "", fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer dst.Close()

	// 复制文件内容
	if _, err = io.Copy(dst, src); err != nil {
		return "", "", fmt.Errorf("复制文件失败: %w", err)
	}

	// 构建文件URL
	fileURL := fmt.Sprintf("/data/%s/%s/%s/%s", format, month, day, fileName)

	return targetPath, fileURL, nil
}

// generateGUID 生成GUID
func (f *FileUtils) generateGUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// IsValidFileType 验证文件类型
func (f *FileUtils) IsValidFileType(fileType string) bool {
	validTypes := []string{"avatar", "course_cover", "post_image"}
	for _, t := range validTypes {
		if t == fileType {
			return true
		}
	}
	return false
}

// IsValidFileExtension 验证文件扩展名
func (f *FileUtils) IsValidFileExtension(ext string) bool {
	validExts := []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".pdf", ".doc", ".docx"}
	for _, e := range validExts {
		if e == ext {
			return true
		}
	}
	return false
}

// GetFileSize 获取文件大小限制
func (f *FileUtils) GetFileSize() int64 {
	return 10 * 1024 * 1024 // 10MB
}
