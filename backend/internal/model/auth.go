package model

// RegisterRequest 注册请求
type RegisterRequest struct {
	Email           string          `json:"email" binding:"required,email"`
	Password        string          `json:"password" binding:"required,min=8"`
	Phone           string          `json:"phone"`
	PrimaryIdentity PrimaryIdentity `json:"primary_identity" binding:"required"`
}

// PrimaryIdentity 主要身份信息
type PrimaryIdentity struct {
	IdentityType string `json:"identity_type" binding:"required,oneof=master apprentice"`
	Domain       string `json:"domain" binding:"required"`
	Name         string `json:"name" binding:"required"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// SwitchIdentityRequest 身份切换请求
type SwitchIdentityRequest struct {
	IdentityID string `json:"identity_id" binding:"required"`
}

// ChangePasswordRequest 修改密码请求
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
}

// AuthResponse 认证响应
type AuthResponse struct {
	UserID          string         `json:"user_id"`
	Token           string         `json:"token"`
	CurrentIdentity *IdentityInfo  `json:"current_identity,omitempty"`
	Identities      []IdentityInfo `json:"identities,omitempty"`
}

// IdentityInfo 身份信息
type IdentityInfo struct {
	ID           string `json:"id"`
	IdentityType string `json:"identity_type"`
	Domain       string `json:"domain"`
	Status       string `json:"status"`
}

// TokenResponse Token响应
type TokenResponse struct {
	Token string `json:"token"`
}
