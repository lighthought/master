package service

import (
	"context"
	"errors"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
	"master-guide-backend/internal/utils"
)

// AuthService 认证服务接口
type AuthService interface {
	Register(ctx context.Context, req *model.RegisterRequest) (*model.AuthResponse, error)
	Login(ctx context.Context, req *model.LoginRequest) (*model.AuthResponse, error)
	SwitchIdentity(ctx context.Context, userID, identityID string) (*model.AuthResponse, error)
	RefreshToken(ctx context.Context, userID string) (*model.TokenResponse, error)
	ChangePassword(ctx context.Context, userID, currentPassword, newPassword string) error
}

// authService 认证服务实现
type authService struct {
	userRepo     repository.UserRepository
	identityRepo repository.IdentityRepository
	jwtSecret    string
	jwtExpire    int
}

// NewAuthService 创建认证服务实例
func NewAuthService(userRepo repository.UserRepository, identityRepo repository.IdentityRepository, jwtSecret string, jwtExpire int) AuthService {
	return &authService{
		userRepo:     userRepo,
		identityRepo: identityRepo,
		jwtSecret:    jwtSecret,
		jwtExpire:    jwtExpire,
	}
}

// Register 用户注册
func (s *authService) Register(ctx context.Context, req *model.RegisterRequest) (*model.AuthResponse, error) {
	// 检查邮箱是否已存在
	exists, err := s.userRepo.ExistsByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("邮箱已存在")
	}

	// 加密密码
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &model.User{
		Email:        req.Email,
		PasswordHash: passwordHash,
		Phone:        req.Phone,
		Status:       "active",
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	// 创建主要身份
	identity := &model.UserIdentity{
		UserID:       user.ID,
		IdentityType: req.PrimaryIdentity.IdentityType,
		Domain:       req.PrimaryIdentity.Domain,
		Status:       "active",
	}
	if err := s.identityRepo.Create(ctx, identity); err != nil {
		return nil, err
	}

	// 生成Token
	token, err := utils.GenerateToken(user.ID, identity.ID, s.jwtSecret, s.jwtExpire)
	if err != nil {
		return nil, err
	}

	// 构建响应
	response := &model.AuthResponse{
		UserID: user.ID,
		Token:  token,
		CurrentIdentity: &model.IdentityInfo{
			ID:           identity.ID,
			IdentityType: identity.IdentityType,
			Domain:       identity.Domain,
			Status:       identity.Status,
		},
		Identities: []model.IdentityInfo{
			{
				ID:           identity.ID,
				IdentityType: identity.IdentityType,
				Domain:       identity.Domain,
				Status:       identity.Status,
			},
		},
	}

	return response, nil
}

// Login 用户登录
func (s *authService) Login(ctx context.Context, req *model.LoginRequest) (*model.AuthResponse, error) {
	// 获取用户
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("用户名或密码错误")
	}

	// 检查用户状态
	if user.Status != "active" {
		return nil, errors.New("账户已被禁用")
	}

	// 获取用户身份列表
	identities, err := s.identityRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	if len(identities) == 0 {
		return nil, errors.New("用户没有可用身份")
	}

	// 使用第一个活跃身份作为当前身份
	var currentIdentity *model.UserIdentity
	for _, identity := range identities {
		if identity.Status == "active" {
			currentIdentity = identity
			break
		}
	}

	if currentIdentity == nil {
		return nil, errors.New("用户没有活跃身份")
	}

	// 更新最后登录时间
	now := time.Now()
	user.LastLoginAt = &now
	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	// 生成Token
	token, err := utils.GenerateToken(user.ID, currentIdentity.ID, s.jwtSecret, s.jwtExpire)
	if err != nil {
		return nil, err
	}

	// 构建身份信息列表
	identityInfos := make([]model.IdentityInfo, len(identities))
	for i, identity := range identities {
		identityInfos[i] = model.IdentityInfo{
			ID:           identity.ID,
			IdentityType: identity.IdentityType,
			Domain:       identity.Domain,
			Status:       identity.Status,
		}
	}

	// 构建响应
	response := &model.AuthResponse{
		UserID: user.ID,
		Token:  token,
		CurrentIdentity: &model.IdentityInfo{
			ID:           currentIdentity.ID,
			IdentityType: currentIdentity.IdentityType,
			Domain:       currentIdentity.Domain,
			Status:       currentIdentity.Status,
		},
		Identities: identityInfos,
	}

	return response, nil
}

// SwitchIdentity 身份切换
func (s *authService) SwitchIdentity(ctx context.Context, userID, identityID string) (*model.AuthResponse, error) {
	// 获取身份
	identity, err := s.identityRepo.GetByID(ctx, identityID)
	if err != nil {
		return nil, errors.New("身份不存在")
	}

	// 检查身份是否属于当前用户
	if identity.UserID != userID {
		return nil, errors.New("无权访问此身份")
	}

	// 检查身份状态
	if identity.Status != "active" {
		return nil, errors.New("身份未激活")
	}

	// 生成新Token
	token, err := utils.GenerateToken(userID, identity.ID, s.jwtSecret, s.jwtExpire)
	if err != nil {
		return nil, err
	}

	// 构建响应
	response := &model.AuthResponse{
		UserID: userID,
		Token:  token,
		CurrentIdentity: &model.IdentityInfo{
			ID:           identity.ID,
			IdentityType: identity.IdentityType,
			Domain:       identity.Domain,
			Status:       identity.Status,
		},
	}

	return response, nil
}

// RefreshToken 刷新Token
func (s *authService) RefreshToken(ctx context.Context, userID string) (*model.TokenResponse, error) {
	// 获取用户
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 检查用户状态
	if user.Status != "active" {
		return nil, errors.New("账户已被禁用")
	}

	// 获取用户第一个活跃身份
	identities, err := s.identityRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var currentIdentity *model.UserIdentity
	for _, identity := range identities {
		if identity.Status == "active" {
			currentIdentity = identity
			break
		}
	}

	if currentIdentity == nil {
		return nil, errors.New("用户没有活跃身份")
	}

	// 生成新Token
	token, err := utils.GenerateToken(userID, currentIdentity.ID, s.jwtSecret, s.jwtExpire)
	if err != nil {
		return nil, err
	}

	return &model.TokenResponse{Token: token}, nil
}

// ChangePassword 修改密码
func (s *authService) ChangePassword(ctx context.Context, userID, currentPassword, newPassword string) error {
	// 获取用户
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 验证当前密码
	if !utils.CheckPassword(currentPassword, user.PasswordHash) {
		return errors.New("当前密码错误")
	}

	// 加密新密码
	newPasswordHash, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// 更新密码
	user.PasswordHash = newPasswordHash
	return s.userRepo.Update(ctx, user)
}
