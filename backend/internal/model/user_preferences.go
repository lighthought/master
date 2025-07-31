package model

import "time"

// UserPreferences 用户偏好
type UserPreferences struct {
	BaseModel
	UserID           string    `json:"user_id" gorm:"not null"`
	LearningStyle    string    `json:"learning_style" gorm:"not null"`
	TimePreference   string    `json:"time_preference" gorm:"not null"`
	BudgetRange      string    `json:"budget_range" gorm:"not null"`
	LearningGoals    []string  `json:"learning_goals" gorm:"type:text[]"`
	PreferredDomains []string  `json:"preferred_domains" gorm:"type:text[]"`
	ExperienceLevel  string    `json:"experience_level" gorm:"not null"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// 关联关系
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName 指定表名
func (UserPreferences) TableName() string {
	return "user_preferences"
}
