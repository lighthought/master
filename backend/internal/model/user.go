package model

import (
	"time"
)

// User 用户模型
type User struct {
	BaseModel
	Email         string     `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash  string     `json:"-" gorm:"not null"`
	Phone         string     `json:"phone"`
	Status        string     `json:"status" gorm:"default:'active'"`
	EmailVerified bool       `json:"email_verified" gorm:"default:false"`
	PhoneVerified bool       `json:"phone_verified" gorm:"default:false"`
	LastLoginAt   *time.Time `json:"last_login_at"`
}

// UserIdentity 用户身份模型
type UserIdentity struct {
	BaseModel
	UserID             string `json:"user_id" gorm:"not null"`
	IdentityType       string `json:"identity_type" gorm:"not null"` // master, apprentice
	Domain             string `json:"domain" gorm:"not null"`
	Status             string `json:"status" gorm:"default:'pending'"`
	VerificationStatus string `json:"verification_status" gorm:"default:'unverified'"`

	// 关联关系
	User    *User        `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Profile *UserProfile `json:"profile,omitempty" gorm:"foreignKey:IdentityID"`
}

// UserProfile 用户档案模型
type UserProfile struct {
	BaseModel
	UserID          string   `json:"user_id" gorm:"not null"`
	IdentityID      string   `json:"identity_id" gorm:"not null;uniqueIndex"`
	Name            string   `json:"name" gorm:"not null"`
	Avatar          string   `json:"avatar"`
	Bio             string   `json:"bio"`
	Skills          []string `json:"skills" gorm:"type:text[]"`
	ExperienceYears int      `json:"experience_years" gorm:"default:0"`
	HourlyRate      float64  `json:"hourly_rate"`

	// 关联关系
	User     *User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Identity *UserIdentity `json:"identity,omitempty" gorm:"foreignKey:IdentityID"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

func (UserIdentity) TableName() string {
	return "user_identities"
}

func (UserProfile) TableName() string {
	return "user_profiles"
}
