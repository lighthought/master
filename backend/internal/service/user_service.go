package service

import (
	"context"
	"errors"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// UserService 用户服务接口
type UserService interface {
	GetUserProfile(ctx context.Context, userID string) (*model.UserProfileResponse, error)
	UpdateUserProfile(ctx context.Context, userID, identityID string, req *model.UpdateProfileRequest) error
	GetUserIdentities(ctx context.Context, userID string) (*model.IdentityListResponse, error)
	CreateUserIdentity(ctx context.Context, userID string, req *model.CreateIdentityRequest) (*model.CreateIdentityResponse, error)
	UpdateUserIdentity(ctx context.Context, userID, identityID string, req *model.UpdateIdentityRequest) error
	GetLearningStats(ctx context.Context, userID string) (*model.LearningStatsResponse, error)
	GetTeachingStats(ctx context.Context, userID string) (*model.TeachingStatsResponse, error)
	GetGeneralStats(ctx context.Context, userID string) (*model.GeneralStatsResponse, error)
	GetUserAchievements(ctx context.Context, userID, identityType string) (*model.AchievementsResponse, error)
	GetUserPreferences(ctx context.Context, userID string) (*model.UserPreferencesResponse, error)
	SaveUserPreferences(ctx context.Context, userID string, req *model.UserPreferencesRequest) error
	GetRecommendedLearningPath(ctx context.Context, userID string) (*model.RecommendedLearningPathResponse, error)
	GetLearningPathStats(ctx context.Context) (*model.LearningPathStatsResponse, error)
}

// userService 用户服务实现
type userService struct {
	userRepo        repository.UserRepository
	identityRepo    repository.IdentityRepository
	profileRepo     repository.ProfileRepository
	preferencesRepo repository.PreferencesRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repository.UserRepository, identityRepo repository.IdentityRepository, profileRepo repository.ProfileRepository, preferencesRepo repository.PreferencesRepository) UserService {
	return &userService{
		userRepo:        userRepo,
		identityRepo:    identityRepo,
		profileRepo:     profileRepo,
		preferencesRepo: preferencesRepo,
	}
}

// GetUserProfile 获取用户档案
func (s *userService) GetUserProfile(ctx context.Context, userID string) (*model.UserProfileResponse, error) {
	// 获取用户信息
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 获取用户身份列表
	identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 找到当前活跃身份
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

	// 构建响应
	response := &model.UserProfileResponse{
		User: &model.UserInfo{
			ID:        user.ID,
			Email:     user.Email,
			Phone:     user.Phone,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
		},
		CurrentIdentity: &model.IdentityWithProfile{
			ID:           currentIdentity.ID,
			IdentityType: currentIdentity.IdentityType,
			Domain:       currentIdentity.Domain,
			Status:       currentIdentity.Status,
		},
	}

	// 如果有档案信息，添加到响应中
	if currentIdentity.Profile != nil {
		response.CurrentIdentity.Profile = &model.UserProfileInfo{
			Name:            currentIdentity.Profile.Name,
			Avatar:          currentIdentity.Profile.Avatar,
			Bio:             currentIdentity.Profile.Bio,
			Skills:          currentIdentity.Profile.Skills,
			ExperienceYears: currentIdentity.Profile.ExperienceYears,
			HourlyRate:      currentIdentity.Profile.HourlyRate,
		}
	}

	return response, nil
}

// UpdateUserProfile 更新用户档案
func (s *userService) UpdateUserProfile(ctx context.Context, userID, identityID string, req *model.UpdateProfileRequest) error {
	// 验证身份是否属于用户
	identity, err := s.identityRepo.GetByID(ctx, identityID)
	if err != nil {
		return errors.New("身份不存在")
	}

	if identity.UserID != userID {
		return errors.New("无权访问此身份")
	}

	// 获取或创建档案
	profile, err := s.profileRepo.GetByIdentityID(ctx, identityID)
	if err != nil {
		// 创建新档案
		profile = &model.UserProfile{
			UserID:          userID,
			IdentityID:      identityID,
			Name:            req.Name,
			Avatar:          req.Avatar,
			Bio:             req.Bio,
			Skills:          req.Skills,
			ExperienceYears: req.ExperienceYears,
			HourlyRate:      req.HourlyRate,
		}
		return s.profileRepo.Create(ctx, profile)
	}

	// 更新档案
	profile.Name = req.Name
	profile.Avatar = req.Avatar
	profile.Bio = req.Bio
	profile.Skills = req.Skills
	profile.ExperienceYears = req.ExperienceYears
	profile.HourlyRate = req.HourlyRate

	return s.profileRepo.Update(ctx, profile)
}

// GetUserIdentities 获取用户身份列表
func (s *userService) GetUserIdentities(ctx context.Context, userID string) (*model.IdentityListResponse, error) {
	identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	identityList := make([]*model.IdentityWithProfile, len(identities))
	for i, identity := range identities {
		identityList[i] = &model.IdentityWithProfile{
			ID:           identity.ID,
			IdentityType: identity.IdentityType,
			Domain:       identity.Domain,
			Status:       identity.Status,
		}

		if identity.Profile != nil {
			identityList[i].Profile = &model.UserProfileInfo{
				Name:            identity.Profile.Name,
				Avatar:          identity.Profile.Avatar,
				Bio:             identity.Profile.Bio,
				Skills:          identity.Profile.Skills,
				ExperienceYears: identity.Profile.ExperienceYears,
				HourlyRate:      identity.Profile.HourlyRate,
			}
		}
	}

	return &model.IdentityListResponse{Identities: identityList}, nil
}

// CreateUserIdentity 创建用户身份
func (s *userService) CreateUserIdentity(ctx context.Context, userID string, req *model.CreateIdentityRequest) (*model.CreateIdentityResponse, error) {
	// 检查是否已存在相同类型和领域的身份
	existingIdentity, err := s.identityRepo.GetByUserIDAndType(ctx, userID, req.IdentityType, req.Domain)
	if err == nil && existingIdentity != nil {
		return nil, errors.New("已存在相同类型和领域的身份")
	}

	// 创建身份
	identity := &model.UserIdentity{
		UserID:       userID,
		IdentityType: req.IdentityType,
		Domain:       req.Domain,
		Status:       "pending", // 新身份默认为待审核状态
	}

	if err := s.identityRepo.Create(ctx, identity); err != nil {
		return nil, err
	}

	// 创建档案
	profile := &model.UserProfile{
		UserID:          userID,
		IdentityID:      identity.ID,
		Name:            req.Name,
		Bio:             req.Bio,
		Skills:          req.Skills,
		ExperienceYears: req.ExperienceYears,
		HourlyRate:      req.HourlyRate,
	}

	if err := s.profileRepo.Create(ctx, profile); err != nil {
		return nil, err
	}

	return &model.CreateIdentityResponse{
		IdentityID: identity.ID,
		Status:     identity.Status,
	}, nil
}

// UpdateUserIdentity 更新用户身份
func (s *userService) UpdateUserIdentity(ctx context.Context, userID, identityID string, req *model.UpdateIdentityRequest) error {
	// 验证身份是否属于用户
	identity, err := s.identityRepo.GetByID(ctx, identityID)
	if err != nil {
		return errors.New("身份不存在")
	}

	if identity.UserID != userID {
		return errors.New("无权访问此身份")
	}

	// 获取档案
	profile, err := s.profileRepo.GetByIdentityID(ctx, identityID)
	if err != nil {
		return errors.New("档案不存在")
	}

	// 更新档案
	profile.Name = req.Name
	profile.Bio = req.Bio
	profile.Skills = req.Skills
	profile.ExperienceYears = req.ExperienceYears
	profile.HourlyRate = req.HourlyRate

	return s.profileRepo.Update(ctx, profile)
}

// GetLearningStats 获取学习统计
func (s *userService) GetLearningStats(ctx context.Context, userID string) (*model.LearningStatsResponse, error) {
	// 这里应该从学习记录表中获取实际数据
	// 暂时返回模拟数据
	return &model.LearningStatsResponse{
		TotalCourses:     12,
		Progress:         65,
		CompletedLessons: 8,
		TotalLessons:     15,
		CurrentCourse:    "Vue.js 进阶开发",
		NextLesson:       "组件通信与状态管理",
	}, nil
}

// GetTeachingStats 获取教学统计
func (s *userService) GetTeachingStats(ctx context.Context, userID string) (*model.TeachingStatsResponse, error) {
	// 这里应该从教学记录表中获取实际数据
	// 暂时返回模拟数据
	return &model.TeachingStatsResponse{
		TotalStudents:     8,
		TotalHours:        24,
		TotalEarnings:     2400,
		AverageRating:     4.8,
		CompletedSessions: 12,
		UpcomingSessions:  3,
	}, nil
}

// GetGeneralStats 获取通用统计
func (s *userService) GetGeneralStats(ctx context.Context, userID string) (*model.GeneralStatsResponse, error) {
	// 这里应该从用户活动记录表中获取实际数据
	// 暂时返回模拟数据
	return &model.GeneralStatsResponse{
		ActiveDays:     7,
		Achievements:   3,
		TotalLoginDays: 15,
		LastLoginDate:  "2024-01-15",
		StreakDays:     5,
	}, nil
}

// GetUserAchievements 获取用户成就
func (s *userService) GetUserAchievements(ctx context.Context, userID, identityType string) (*model.AchievementsResponse, error) {
	// 这里应该从成就表中获取实际数据
	// 暂时返回模拟数据
	achievements := []*model.Achievement{
		{
			ID:          "1",
			Name:        "学习新手",
			Description: "完成第一门课程",
			Icon:        "🎓",
		},
		{
			ID:          "2",
			Name:        "坚持不懈",
			Description: "连续学习7天",
			Icon:        "🔥",
		},
	}

	return &model.AchievementsResponse{Achievements: achievements}, nil
}

// GetUserPreferences 获取用户偏好
func (s *userService) GetUserPreferences(ctx context.Context, userID string) (*model.UserPreferencesResponse, error) {
	preferences, err := s.preferencesRepo.GetByUserID(ctx, userID)
	if err != nil {
		// 如果不存在，返回默认偏好
		return &model.UserPreferencesResponse{
			LearningStyle:    "one-on-one",
			TimePreference:   "flexible",
			BudgetRange:      "medium",
			LearningGoals:    []string{},
			PreferredDomains: []string{},
			ExperienceLevel:  "beginner",
		}, nil
	}

	return &model.UserPreferencesResponse{
		LearningStyle:    preferences.LearningStyle,
		TimePreference:   preferences.TimePreference,
		BudgetRange:      preferences.BudgetRange,
		LearningGoals:    preferences.LearningGoals,
		PreferredDomains: preferences.PreferredDomains,
		ExperienceLevel:  preferences.ExperienceLevel,
		UpdatedAt:        preferences.UpdatedAt,
	}, nil
}

// SaveUserPreferences 保存用户偏好
func (s *userService) SaveUserPreferences(ctx context.Context, userID string, req *model.UserPreferencesRequest) error {
	preferences := &model.UserPreferences{
		UserID:           userID,
		LearningStyle:    req.LearningStyle,
		TimePreference:   req.TimePreference,
		BudgetRange:      req.BudgetRange,
		LearningGoals:    req.LearningGoals,
		PreferredDomains: req.PreferredDomains,
		ExperienceLevel:  req.ExperienceLevel,
	}

	return s.preferencesRepo.Upsert(ctx, preferences)
}

// GetRecommendedLearningPath 获取推荐学习路径
func (s *userService) GetRecommendedLearningPath(ctx context.Context, userID string) (*model.RecommendedLearningPathResponse, error) {
	// 这里应该基于用户偏好和历史数据计算推荐路径
	// 暂时返回模拟数据
	return &model.RecommendedLearningPathResponse{
		RecommendedPath: "one-on-one",
		Confidence:      0.7,
		Reasons:         []string{"基于用户偏好推荐", "适合初学者"},
	}, nil
}

// GetLearningPathStats 获取学习路径统计
func (s *userService) GetLearningPathStats(ctx context.Context) (*model.LearningPathStatsResponse, error) {
	// 这里应该从统计数据中获取实际数据
	// 暂时返回模拟数据
	return &model.LearningPathStatsResponse{
		TotalUsers: 1250,
		PathDistribution: map[string]int{
			"one-on-one": 45,
			"structured": 30,
			"browse":     20,
			"other":      5,
		},
		SatisfactionRates: map[string]float64{
			"one-on-one": 4.8,
			"structured": 4.6,
			"browse":     4.4,
			"other":      4.2,
		},
	}, nil
}
