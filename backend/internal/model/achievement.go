package model

// AchievementModel 成就模型
type AchievementModel struct {
	BaseModel
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Category    string `json:"category" gorm:"not null"`
	Points      int    `json:"points" gorm:"default:0"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}

// UserAchievement 用户成就关联
type UserAchievement struct {
	BaseModel
	UserID        string `json:"user_id" gorm:"not null"`
	AchievementID string `json:"achievement_id" gorm:"not null"`
	EarnedAt      string `json:"earned_at" gorm:"not null"`

	// 关联关系
	User        *User             `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Achievement *AchievementModel `json:"achievement,omitempty" gorm:"foreignKey:AchievementID"`
}

// TableName 指定表名
func (AchievementModel) TableName() string {
	return "achievements"
}

func (UserAchievement) TableName() string {
	return "user_achievements"
}
