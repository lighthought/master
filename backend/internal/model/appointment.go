package model

import "time"

// AppointmentModel 预约模型
type AppointmentModel struct {
	BaseModel
	StudentID       string    `json:"student_id" gorm:"not null"`
	MentorID        string    `json:"mentor_id" gorm:"not null"`
	AppointmentTime time.Time `json:"appointment_time" gorm:"not null"`
	DurationMinutes int       `json:"duration_minutes" gorm:"not null"`
	MeetingType     string    `json:"meeting_type" gorm:"default:'video'"`
	Status          string    `json:"status" gorm:"default:'pending'"`
	Price           float64   `json:"price" gorm:"not null"`
	Notes           string    `json:"notes" gorm:"type:text"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// 关联关系
	Student *User   `json:"student,omitempty" gorm:"foreignKey:StudentID"`
	Mentor  *Mentor `json:"mentor,omitempty" gorm:"foreignKey:MentorID"`
}

// TableName 指定表名
func (AppointmentModel) TableName() string {
	return "appointments"
}
