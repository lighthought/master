package model

// Message 消息模型
type Message struct {
	BaseModel
	FromUserID string `json:"from_user_id" gorm:"not null"`
	ToUserID   string `json:"to_user_id"`
	CircleID   string `json:"circle_id"`
	Content    string `json:"content" gorm:"not null"`
	Type       string `json:"type" gorm:"default:'text'"`
	IsRead     bool   `json:"is_read" gorm:"default:false"`

	// 关联关系
	FromUser *User   `json:"from_user,omitempty" gorm:"foreignKey:FromUserID"`
	ToUser   *User   `json:"to_user,omitempty" gorm:"foreignKey:ToUserID"`
	Circle   *Circle `json:"circle,omitempty" gorm:"foreignKey:CircleID"`
}

// TableName 指定表名
func (Message) TableName() string {
	return "messages"
}
