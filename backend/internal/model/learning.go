package model

import "time"

// LearningRecord 学习记录模型
type LearningRecord struct {
	BaseModel
	UserID             string     `json:"user_id" gorm:"not null"`
	CourseID           string     `json:"course_id" gorm:"not null"`
	ProgressPercentage float64    `json:"progress_percentage" gorm:"type:decimal(5,2);default:0"`
	Status             string     `json:"status" gorm:"default:'enrolled'"`
	EnrolledAt         time.Time  `json:"enrolled_at" gorm:"autoCreateTime"`
	CompletedAt        *time.Time `json:"completed_at"`
	LastAccessedAt     time.Time  `json:"last_accessed_at" gorm:"autoCreateTime"`
	CurrentChapter     string     `json:"current_chapter"`
	CompletedChapters  []string   `json:"completed_chapters" gorm:"type:text[]"`
	TotalStudyTime     int        `json:"total_study_time" gorm:"default:0"`
	CertificateIssued  bool       `json:"certificate_issued" gorm:"default:false"`
	CertificateURL     *string    `json:"certificate_url"`

	// 关联关系
	User   *User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Course *Course `json:"course,omitempty" gorm:"foreignKey:CourseID"`
}

// TableName 指定表名
func (LearningRecord) TableName() string {
	return "learning_records"
}

// StudySession 学习会话模型
type StudySession struct {
	BaseModel
	LearningRecordID string     `json:"learning_record_id" gorm:"not null"`
	StartTime        time.Time  `json:"start_time" gorm:"not null"`
	EndTime          *time.Time `json:"end_time"`
	DurationMinutes  int        `json:"duration_minutes" gorm:"default:0"`
	Chapter          string     `json:"chapter"`
	Notes            string     `json:"notes" gorm:"type:text"`
	CreatedAt        time.Time  `json:"created_at" gorm:"autoCreateTime"`

	// 关联关系
	LearningRecord *LearningRecord `json:"learning_record,omitempty" gorm:"foreignKey:LearningRecordID"`
}

// TableName 指定表名
func (StudySession) TableName() string {
	return "study_sessions"
}

// Assignment 作业模型
type Assignment struct {
	BaseModel
	LearningRecordID string     `json:"learning_record_id" gorm:"not null"`
	Title            string     `json:"title" gorm:"not null"`
	Content          string     `json:"content" gorm:"type:text"`
	AttachmentURLs   []string   `json:"attachment_urls" gorm:"type:text[]"`
	Status           string     `json:"status" gorm:"default:'submitted'"`
	SubmittedAt      time.Time  `json:"submitted_at" gorm:"autoCreateTime"`
	ReviewedAt       *time.Time `json:"reviewed_at"`
	Score            *float64   `json:"score" gorm:"type:decimal(5,2)"`
	Feedback         string     `json:"feedback" gorm:"type:text"`
	CreatedAt        time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time  `json:"updated_at" gorm:"autoUpdateTime"`

	// 关联关系
	LearningRecord *LearningRecord `json:"learning_record,omitempty" gorm:"foreignKey:LearningRecordID"`
}

// TableName 指定表名
func (Assignment) TableName() string {
	return "assignments"
}
