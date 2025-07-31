package model

import "time"

// Circle 圈子模型
type Circle struct {
	BaseModel
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"type:text"`
	Domain      string    `json:"domain" gorm:"not null"`
	CreatedBy   string    `json:"created_by" gorm:"not null"`
	Status      string    `json:"status" gorm:"default:'active'"`
	MemberCount int       `json:"member_count" gorm:"default:0"`
	Tags        []string  `json:"tags" gorm:"type:text[]"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// 关联关系
	Creator *UserIdentity `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
}

// TableName 指定表名
func (Circle) TableName() string {
	return "circles"
}

// CircleMember 圈子成员模型
type CircleMember struct {
	BaseModel
	CircleID string    `json:"circle_id" gorm:"not null"`
	UserID   string    `json:"user_id" gorm:"not null"`
	Role     string    `json:"role" gorm:"default:'member'"`
	JoinedAt time.Time `json:"joined_at" gorm:"autoCreateTime"`

	// 关联关系
	Circle *Circle `json:"circle,omitempty" gorm:"foreignKey:CircleID"`
	User   *User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName 指定表名
func (CircleMember) TableName() string {
	return "circle_members"
}

// Post 动态模型
type Post struct {
	BaseModel
	UserID       string    `json:"user_id" gorm:"not null"`
	IdentityID   string    `json:"identity_id" gorm:"not null"`
	CircleID     string    `json:"circle_id" gorm:"not null"`
	Content      string    `json:"content" gorm:"not null;type:text"`
	MediaURLs    []string  `json:"media_urls" gorm:"type:text[]"`
	PostType     string    `json:"post_type" gorm:"default:'text'"`
	Status       string    `json:"status" gorm:"default:'active'"`
	LikeCount    int       `json:"like_count" gorm:"default:0"`
	CommentCount int       `json:"comment_count" gorm:"default:0"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// 关联关系
	User     *User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Identity *UserIdentity `json:"identity,omitempty" gorm:"foreignKey:IdentityID"`
	Circle   *Circle       `json:"circle,omitempty" gorm:"foreignKey:CircleID"`
}

// TableName 指定表名
func (Post) TableName() string {
	return "posts"
}

// Comment 评论模型
type Comment struct {
	BaseModel
	PostID     string    `json:"post_id" gorm:"not null"`
	UserID     string    `json:"user_id" gorm:"not null"`
	IdentityID string    `json:"identity_id" gorm:"not null"`
	Content    string    `json:"content" gorm:"not null;type:text"`
	ParentID   *string   `json:"parent_id"`
	LikeCount  int       `json:"like_count" gorm:"default:0"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`

	// 关联关系
	Post     *Post         `json:"post,omitempty" gorm:"foreignKey:PostID"`
	User     *User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Identity *UserIdentity `json:"identity,omitempty" gorm:"foreignKey:IdentityID"`
	Parent   *Comment      `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Replies  []*Comment    `json:"replies,omitempty" gorm:"foreignKey:ParentID"`
}

// TableName 指定表名
func (Comment) TableName() string {
	return "comments"
}

// PostLike 动态点赞模型
type PostLike struct {
	BaseModel
	PostID    string    `json:"post_id" gorm:"not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	// 关联关系
	Post *Post `json:"post,omitempty" gorm:"foreignKey:PostID"`
	User *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName 指定表名
func (PostLike) TableName() string {
	return "post_likes"
}

// CommentLike 评论点赞模型
type CommentLike struct {
	BaseModel
	CommentID string    `json:"comment_id" gorm:"not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	// 关联关系
	Comment *Comment `json:"comment,omitempty" gorm:"foreignKey:CommentID"`
	User    *User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName 指定表名
func (CommentLike) TableName() string {
	return "comment_likes"
}
