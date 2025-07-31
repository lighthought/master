package model

import "time"

// ContentProgressModel 内容进度模型
type ContentProgressModel struct {
	BaseModel
	UserID             string    `json:"user_id" gorm:"not null"`
	ContentID          string    `json:"content_id" gorm:"not null"`
	IsCompleted        bool      `json:"is_completed" gorm:"default:false"`
	ProgressPercentage float64   `json:"progress_percentage" gorm:"default:0"`
	StudyTimeMinutes   int       `json:"study_time_minutes" gorm:"default:0"`
	LastAccessedAt     time.Time `json:"last_accessed_at" gorm:"autoCreateTime"`

	// 关联关系
	User    *User               `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Content *CourseContentModel `json:"content,omitempty" gorm:"foreignKey:ContentID"`
}

// TableName 指定表名
func (ContentProgressModel) TableName() string {
	return "content_progress"
}
