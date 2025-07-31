package model

// CourseContentModel 课程内容模型
type CourseContentModel struct {
	BaseModel
	CourseID        string `json:"course_id" gorm:"not null"`
	Title           string `json:"title" gorm:"not null"`
	ContentType     string `json:"content_type" gorm:"not null"`
	ContentURL      string `json:"content_url"`
	ContentText     string `json:"content_text" gorm:"type:text"`
	OrderIndex      int    `json:"order_index" gorm:"not null"`
	DurationMinutes int    `json:"duration_minutes"`

	// 关联关系
	Course *Course `json:"course,omitempty" gorm:"foreignKey:CourseID"`
}

// TableName 指定表名
func (CourseContentModel) TableName() string {
	return "course_contents"
}
