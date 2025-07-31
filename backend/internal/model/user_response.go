package model

import "time"

// UserProfileResponse 用户档案响应
type UserProfileResponse struct {
	User            *UserInfo            `json:"user"`
	CurrentIdentity *IdentityWithProfile `json:"current_identity"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// IdentityWithProfile 带档案的身份信息
type IdentityWithProfile struct {
	ID           string           `json:"id"`
	IdentityType string           `json:"identity_type"`
	Domain       string           `json:"domain"`
	Status       string           `json:"status"`
	Profile      *UserProfileInfo `json:"profile"`
}

// UserProfileInfo 用户档案信息
type UserProfileInfo struct {
	Name            string   `json:"name"`
	Avatar          string   `json:"avatar"`
	Bio             string   `json:"bio"`
	Skills          []string `json:"skills"`
	ExperienceYears int      `json:"experience_years"`
	HourlyRate      float64  `json:"hourly_rate"`
}

// IdentityListResponse 身份列表响应
type IdentityListResponse struct {
	Identities []*IdentityWithProfile `json:"identities"`
}

// CreateIdentityResponse 创建身份响应
type CreateIdentityResponse struct {
	IdentityID string `json:"identity_id"`
	Status     string `json:"status"`
}

// LearningStatsResponse 学习统计响应
type LearningStatsResponse struct {
	TotalCourses     int    `json:"total_courses"`
	Progress         int    `json:"progress"`
	CompletedLessons int    `json:"completed_lessons"`
	TotalLessons     int    `json:"total_lessons"`
	CurrentCourse    string `json:"current_course"`
	NextLesson       string `json:"next_lesson"`
}

// TeachingStatsResponse 教学统计响应
type TeachingStatsResponse struct {
	TotalStudents     int     `json:"total_students"`
	TotalHours        int     `json:"total_hours"`
	TotalEarnings     float64 `json:"total_earnings"`
	AverageRating     float64 `json:"average_rating"`
	CompletedSessions int     `json:"completed_sessions"`
	UpcomingSessions  int     `json:"upcoming_sessions"`
}

// GeneralStatsResponse 通用统计响应
type GeneralStatsResponse struct {
	ActiveDays     int    `json:"active_days"`
	Achievements   int    `json:"achievements"`
	TotalLoginDays int    `json:"total_login_days"`
	LastLoginDate  string `json:"last_login_date"`
	StreakDays     int    `json:"streak_days"`
}

// Achievement 成就
type Achievement struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// AchievementsResponse 成就列表响应
type AchievementsResponse struct {
	Achievements []*Achievement `json:"achievements"`
}

// UserPreferencesResponse 用户偏好响应
type UserPreferencesResponse struct {
	LearningStyle    string    `json:"learning_style"`
	TimePreference   string    `json:"time_preference"`
	BudgetRange      string    `json:"budget_range"`
	LearningGoals    []string  `json:"learning_goals"`
	PreferredDomains []string  `json:"preferred_domains"`
	ExperienceLevel  string    `json:"experience_level"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// RecommendedLearningPathResponse 推荐学习路径响应
type RecommendedLearningPathResponse struct {
	RecommendedPath string   `json:"recommended_path"`
	Confidence      float64  `json:"confidence"`
	Reasons         []string `json:"reasons"`
}

// LearningPathStatsResponse 学习路径统计响应
type LearningPathStatsResponse struct {
	TotalUsers        int                `json:"total_users"`
	PathDistribution  map[string]int     `json:"path_distribution"`
	SatisfactionRates map[string]float64 `json:"satisfaction_rates"`
}
