package model

import "time"

// LearningRecordModel 学习记录模型
type LearningRecordModel struct {
	BaseModel
	UserID             string     `json:"user_id" gorm:"not null"`
	CourseID           string     `json:"course_id" gorm:"not null"`
	ProgressPercentage float64    `json:"progress_percentage" gorm:"default:0"`
	Status             string     `json:"status" gorm:"default:'enrolled'"`
	EnrolledAt         time.Time  `json:"enrolled_at" gorm:"autoCreateTime"`
	CompletedAt        *time.Time `json:"completed_at"`
	LastAccessedAt     time.Time  `json:"last_accessed_at" gorm:"autoCreateTime"`

	// 关联关系
	User   *User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Course *Course `json:"course,omitempty" gorm:"foreignKey:CourseID"`
}

// TableName 指定表名
func (LearningRecordModel) TableName() string {
	return "learning_records"
}
