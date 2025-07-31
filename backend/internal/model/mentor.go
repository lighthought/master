package model

import "time"

// Mentor 大师模型
type Mentor struct {
	BaseModel
	IdentityID      string  `json:"identity_id" gorm:"not null"`
	UserID          string  `json:"user_id" gorm:"not null"`
	Rating          float64 `json:"rating" gorm:"default:0"`
	StudentCount    int     `json:"student_count" gorm:"default:0"`
	HourlyRate      float64 `json:"hourly_rate" gorm:"not null"`
	IsOnline        bool    `json:"is_online" gorm:"default:false"`
	ExperienceYears int     `json:"experience_years" gorm:"default:0"`
	Status          string  `json:"status" gorm:"default:'active'"`

	// 关联关系
	Identity *UserIdentity `json:"identity,omitempty" gorm:"foreignKey:IdentityID"`
	User     *User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Profile  *UserProfile  `json:"profile,omitempty" gorm:"foreignKey:IdentityID;references:IdentityID"`
}

// MentorReviewModel 大师评价模型
type MentorReviewModel struct {
	BaseModel
	MentorID    string    `json:"mentor_id" gorm:"not null"`
	ReviewerID  string    `json:"reviewer_id" gorm:"not null"`
	Rating      int       `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Content     string    `json:"content" gorm:"type:text"`
	IsAnonymous bool      `json:"is_anonymous" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`

	// 关联关系
	Mentor   *Mentor `json:"mentor,omitempty" gorm:"foreignKey:MentorID"`
	Reviewer *User   `json:"reviewer,omitempty" gorm:"foreignKey:ReviewerID"`
}

// TableName 指定表名
func (Mentor) TableName() string {
	return "mentors"
}

func (MentorReviewModel) TableName() string {
	return "mentor_reviews"
}
