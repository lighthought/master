package model

import "time"

// WebSocketEvent WebSocket 事件基础结构
type WebSocketEvent struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

// AuthenticateRequest 连接认证请求
type AuthenticateRequest struct {
	Token string `json:"token" binding:"required"`
}

// AuthenticateResponse 连接认证响应
type AuthenticateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	UserID  string `json:"user_id,omitempty"`
}

// MessageEvent 消息事件
type MessageEvent struct {
	ID        string        `json:"id"`
	FromUser  *ChatUserInfo `json:"from_user"`
	Content   string        `json:"content"`
	Type      string        `json:"type"`
	CreatedAt time.Time     `json:"created_at"`
}

// ChatUserInfo 聊天用户信息
type ChatUserInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserStatusEvent 在线状态事件
type UserStatusEvent struct {
	UserID    string    `json:"user_id"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

// ChatMessage 聊天消息
type ChatMessage struct {
	ID        string        `json:"id"`
	FromUser  *ChatUserInfo `json:"from_user"`
	Content   string        `json:"content"`
	Type      string        `json:"type"`
	CreatedAt time.Time     `json:"created_at"`
	ToUserID  string        `json:"to_user_id,omitempty"`
	CircleID  string        `json:"circle_id,omitempty"`
}

// OnlineUser 在线用户
type OnlineUser struct {
	UserID   string    `json:"user_id"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
	IsOnline bool      `json:"is_online"`
	LastSeen time.Time `json:"last_seen"`
}

// ChatMessagesRequest 获取聊天记录请求
type ChatMessagesRequest struct {
	TargetID string `form:"target_id"`
	CircleID string `form:"circle_id"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

// ChatMessagesResponse 聊天记录响应
type ChatMessagesResponse struct {
	Messages   []*ChatMessage      `json:"messages"`
	Pagination *PaginationResponse `json:"pagination"`
}

// OnlineUsersResponse 在线用户响应
type OnlineUsersResponse struct {
	OnlineUsers []*OnlineUser `json:"online_users"`
}
