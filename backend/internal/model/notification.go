package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Notification 通知模型
type Notification struct {
	BaseModel
	UserID           string     `json:"user_id" gorm:"not null"`
	Title            string     `json:"title" gorm:"not null"`
	Content          string     `json:"content" gorm:"type:text;not null"`
	NotificationType string     `json:"notification_type" gorm:"not null"`
	IsRead           bool       `json:"is_read" gorm:"default:false"`
	Metadata         JSONMap    `json:"metadata" gorm:"type:jsonb"`
	ReadAt           *time.Time `json:"read_at"`
	CreatedAt        time.Time  `json:"created_at" gorm:"autoCreateTime"`

	// 关联关系
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName 指定表名
func (Notification) TableName() string {
	return "notifications"
}

// JSONMap 用于处理JSONB字段
type JSONMap map[string]interface{}

// Value 实现driver.Valuer接口
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan 实现sql.Scanner接口
func (j *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return nil
	}

	return json.Unmarshal(bytes, j)
}
