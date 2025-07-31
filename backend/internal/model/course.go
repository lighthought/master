package model

// Course 课程模型
type Course struct {
	BaseModel
	MentorID      string   `json:"mentor_id" gorm:"not null"`
	Title         string   `json:"title" gorm:"not null"`
	Description   string   `json:"description" gorm:"type:text"`
	CoverImage    string   `json:"cover_image"`
	Price         float64  `json:"price" gorm:"not null"`
	DurationHours int      `json:"duration_hours" gorm:"not null"`
	Difficulty    string   `json:"difficulty"`
	Status        string   `json:"status" gorm:"default:'draft'"`
	MaxStudents   *int     `json:"max_students"`
	Rating        float64  `json:"rating" gorm:"default:0"`
	ReviewCount   int      `json:"review_count" gorm:"default:0"`
	Tags          []string `json:"tags" gorm:"type:text[]"`

	// 关联关系
	Mentor *Mentor `json:"mentor,omitempty" gorm:"foreignKey:MentorID"`
}

// TableName 指定表名
func (Course) TableName() string {
	return "courses"
}
