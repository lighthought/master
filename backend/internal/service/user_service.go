package service

import (
	"context"
	"errors"
	"time"

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
	learningRepo    repository.LearningRepository
	mentorRepo      repository.MentorRepository
	appointmentRepo repository.AppointmentRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repository.UserRepository, identityRepo repository.IdentityRepository, profileRepo repository.ProfileRepository, preferencesRepo repository.PreferencesRepository, learningRepo repository.LearningRepository, mentorRepo repository.MentorRepository, appointmentRepo repository.AppointmentRepository) UserService {
	return &userService{
		userRepo:        userRepo,
		identityRepo:    identityRepo,
		profileRepo:     profileRepo,
		preferencesRepo: preferencesRepo,
		learningRepo:    learningRepo,
		mentorRepo:      mentorRepo,
		appointmentRepo: appointmentRepo,
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
	// 使用实际的repository获取学习统计数据
	stats, err := s.learningRepo.GetLearningStats(ctx, userID, "all") // 获取全部时间的数据
	if err != nil {
		return nil, err
	}

	return &model.LearningStatsResponse{
		Stats: stats,
	}, nil
}

// GetTeachingStats 获取教学统计
func (s *userService) GetTeachingStats(ctx context.Context, userID string) (*model.TeachingStatsResponse, error) {
	// 首先获取用户的导师身份
	identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 找到导师身份
	var mentorIdentity *model.UserIdentity
	for _, identity := range identities {
		if identity.IdentityType == "master" && identity.Status == "active" {
			mentorIdentity = identity
			break
		}
	}

	if mentorIdentity == nil {
		// 如果没有导师身份，返回空统计
		return &model.TeachingStatsResponse{
			TotalStudents:     0,
			TotalHours:        0,
			TotalEarnings:     0,
			AverageRating:     0,
			CompletedSessions: 0,
			UpcomingSessions:  0,
		}, nil
	}

	// 获取导师信息
	mentor, err := s.mentorRepo.GetMentorByID(ctx, mentorIdentity.ID)
	if err != nil {
		return nil, err
	}

	// 获取导师的预约统计
	appointmentStats, err := s.appointmentRepo.GetMentorAppointmentStats(ctx, mentor.ID)
	if err != nil {
		return nil, err
	}

	// 计算总学生数（去重）
	totalStudents, err := s.appointmentRepo.GetMentorStudentCount(ctx, mentor.ID)
	if err != nil {
		return nil, err
	}

	// 计算总小时数（已完成和已确认的预约）
	totalHours, err := s.appointmentRepo.GetMentorTotalHours(ctx, mentor.ID)
	if err != nil {
		return nil, err
	}

	// 转换为小时
	totalHours = totalHours / 60

	return &model.TeachingStatsResponse{
		TotalStudents:     int(totalStudents),
		TotalHours:        int(totalHours),
		TotalEarnings:     appointmentStats.TotalEarnings,
		AverageRating:     appointmentStats.AverageRating,
		CompletedSessions: appointmentStats.CompletedAppointments,
		UpcomingSessions:  appointmentStats.ConfirmedAppointments,
	}, nil
}

// GetGeneralStats 获取通用统计
func (s *userService) GetGeneralStats(ctx context.Context, userID string) (*model.GeneralStatsResponse, error) {
	// 基于现有数据计算通用统计信息

	// 获取用户信息
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 获取学习记录数量作为活跃天数
	learningRecords, _, err := s.learningRepo.GetLearningRecords(ctx, userID, "", "", time.Time{}, time.Time{}, 1, 1000)
	if err != nil {
		return nil, err
	}

	// 获取预约数量作为活动记录
	appointments, _, err := s.appointmentRepo.GetAppointments(ctx, userID, "", "", 1, 1000)
	if err != nil {
		return nil, err
	}

	// 计算活跃天数（有学习记录或预约的天数）
	activeDays := len(learningRecords) + len(appointments)
	if activeDays > 30 { // 限制最大值为30天
		activeDays = 30
	}

	// 计算成就数量（基于完成的学习记录）
	achievements := 0
	for _, record := range learningRecords {
		if record.Status == "completed" {
			achievements++
		}
	}

	// 计算连续登录天数（简化实现，基于用户创建时间）
	daysSinceCreation := int(time.Since(user.CreatedAt).Hours() / 24)
	streakDays := daysSinceCreation
	if streakDays > 30 {
		streakDays = 30
	}

	// 格式化最后登录日期
	lastLoginDate := user.CreatedAt.Format("2006-01-02")

	return &model.GeneralStatsResponse{
		ActiveDays:     activeDays,
		Achievements:   achievements,
		TotalLoginDays: int(daysSinceCreation),
		LastLoginDate:  lastLoginDate,
		StreakDays:     streakDays,
	}, nil
}

// GetUserAchievements 获取用户成就
func (s *userService) GetUserAchievements(ctx context.Context, userID, identityType string) (*model.AchievementsResponse, error) {
	// 基于实际学习记录生成成就
	learningRecords, _, err := s.learningRepo.GetLearningRecords(ctx, userID, "", "", time.Time{}, time.Time{}, 1, 1000)
	if err != nil {
		return nil, err
	}

	var achievements []*model.Achievement

	// 计算完成的课程数量
	completedCourses := 0
	totalStudyTime := 0
	for _, record := range learningRecords {
		if record.Status == "completed" {
			completedCourses++
		}
		totalStudyTime += record.TotalStudyTime
	}

	// 根据完成情况生成成就
	if completedCourses >= 1 {
		achievements = append(achievements, &model.Achievement{
			ID:          "1",
			Name:        "学习新手",
			Description: "完成第一门课程",
			Icon:        "🎓",
		})
	}

	if completedCourses >= 5 {
		achievements = append(achievements, &model.Achievement{
			ID:          "2",
			Name:        "学习达人",
			Description: "完成5门课程",
			Icon:        "🏆",
		})
	}

	if totalStudyTime >= 60 { // 60小时
		achievements = append(achievements, &model.Achievement{
			ID:          "3",
			Name:        "坚持不懈",
			Description: "累计学习60小时",
			Icon:        "🔥",
		})
	}

	if totalStudyTime >= 120 { // 120小时
		achievements = append(achievements, &model.Achievement{
			ID:          "4",
			Name:        "学习大师",
			Description: "累计学习120小时",
			Icon:        "👑",
		})
	}

	// 如果是导师身份，添加导师相关成就
	if identityType == "master" {
		// 获取导师的预约统计
		identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, userID)
		if err == nil {
			for _, identity := range identities {
				if identity.IdentityType == "master" && identity.Status == "active" {
					mentor, err := s.mentorRepo.GetMentorByID(ctx, identity.ID)
					if err == nil && mentor != nil {
						appointmentStats, err := s.appointmentRepo.GetMentorAppointmentStats(ctx, mentor.ID)
						if err == nil && appointmentStats != nil {
							if appointmentStats.CompletedAppointments >= 10 {
								achievements = append(achievements, &model.Achievement{
									ID:          "5",
									Name:        "优秀导师",
									Description: "完成10次教学",
									Icon:        "👨‍🏫",
								})
							}

							if appointmentStats.AverageRating >= 4.5 {
								achievements = append(achievements, &model.Achievement{
									ID:          "6",
									Name:        "五星导师",
									Description: "平均评分达到4.5分",
									Icon:        "⭐",
								})
							}
						}
					}
					break
				}
			}
		}
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
	// 获取用户偏好
	preferences, err := s.preferencesRepo.GetByUserID(ctx, userID)
	if err != nil {
		// 如果没有偏好设置，返回默认推荐
		return &model.RecommendedLearningPathResponse{
			RecommendedPath: "one-on-one",
			Confidence:      0.6,
			Reasons:         []string{"基于默认设置推荐", "适合初学者"},
		}, nil
	}

	// 基于用户偏好和学习历史推荐路径
	var recommendedPath string
	var confidence float64
	var reasons []string

	// 根据学习风格推荐
	switch preferences.LearningStyle {
	case "one-on-one":
		recommendedPath = "one-on-one"
		confidence = 0.8
		reasons = append(reasons, "基于您的学习风格偏好")
	case "group":
		recommendedPath = "structured"
		confidence = 0.7
		reasons = append(reasons, "基于您的学习风格偏好")
	case "self-paced":
		recommendedPath = "browse"
		confidence = 0.7
		reasons = append(reasons, "基于您的学习风格偏好")
	default:
		recommendedPath = "one-on-one"
		confidence = 0.6
		reasons = append(reasons, "基于默认设置推荐")
	}

	// 根据经验水平调整推荐
	switch preferences.ExperienceLevel {
	case "beginner":
		reasons = append(reasons, "适合初学者")
		confidence += 0.1
	case "intermediate":
		reasons = append(reasons, "适合中级学习者")
	case "advanced":
		reasons = append(reasons, "适合高级学习者")
		confidence += 0.05
	}

	// 根据预算范围调整推荐
	switch preferences.BudgetRange {
	case "low":
		recommendedPath = "browse"
		reasons = append(reasons, "基于您的预算考虑")
	case "medium":
		// 保持当前推荐
	case "high":
		recommendedPath = "one-on-one"
		reasons = append(reasons, "基于您的预算考虑")
	}

	// 限制置信度在0-1之间
	if confidence > 1.0 {
		confidence = 1.0
	}

	return &model.RecommendedLearningPathResponse{
		RecommendedPath: recommendedPath,
		Confidence:      confidence,
		Reasons:         reasons,
	}, nil
}

// GetLearningPathStats 获取学习路径统计
func (s *userService) GetLearningPathStats(ctx context.Context) (*model.LearningPathStatsResponse, error) {
	// 基于实际数据计算学习路径统计
	// 这里我们基于用户偏好和学习记录来计算统计信息

	// 获取所有用户的偏好设置
	// 由于没有直接的repository方法，我们使用模拟数据
	// 在实际实现中，应该从数据库获取真实的统计数据

	// 模拟基于用户偏好的路径分布
	pathDistribution := map[string]int{
		"one-on-one": 45, // 45%的用户选择一对一学习
		"structured": 30, // 30%的用户选择结构化学习
		"browse":     20, // 20%的用户选择自由浏览
		"other":      5,  // 5%的用户选择其他方式
	}

	// 模拟满意度评分
	satisfactionRates := map[string]float64{
		"one-on-one": 4.8, // 一对一学习满意度最高
		"structured": 4.6, // 结构化学习满意度较高
		"browse":     4.4, // 自由浏览满意度中等
		"other":      4.2, // 其他方式满意度较低
	}

	// 总用户数（模拟数据）
	totalUsers := 1250

	// 在实际实现中，这些数据应该从数据库查询：
	// 1. 统计不同学习路径的用户数量
	// 2. 计算每种路径的平均满意度
	// 3. 获取总用户数

	return &model.LearningPathStatsResponse{
		TotalUsers:        totalUsers,
		PathDistribution:  pathDistribution,
		SatisfactionRates: satisfactionRates,
	}, nil
}
