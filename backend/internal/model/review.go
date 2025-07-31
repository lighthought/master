package model

import "time"

// Review 评价模型
type Review struct {
	BaseModel
	ReviewerID    string    `json:"reviewer_id" gorm:"not null"`
	ReviewedID    string    `json:"reviewed_id" gorm:"not null"`
	CourseID      *string   `json:"course_id"`
	AppointmentID *string   `json:"appointment_id"`
	Rating        int       `json:"rating" gorm:"not null"`
	Content       string    `json:"content" gorm:"type:text"`
	ReviewType    string    `json:"review_type" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// 关联关系
	Reviewer    *User             `json:"reviewer,omitempty" gorm:"foreignKey:ReviewerID"`
	Reviewed    *UserIdentity     `json:"reviewed,omitempty" gorm:"foreignKey:ReviewedID"`
	Course      *Course           `json:"course,omitempty" gorm:"foreignKey:CourseID"`
	Appointment *AppointmentModel `json:"appointment,omitempty" gorm:"foreignKey:AppointmentID"`
}

// TableName 指定表名
func (Review) TableName() string {
	return "reviews"
}
