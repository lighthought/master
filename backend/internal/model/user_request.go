package model

// UpdateProfileRequest 更新用户档案请求
type UpdateProfileRequest struct {
	Name            string   `json:"name" binding:"required"`
	Avatar          string   `json:"avatar"`
	Bio             string   `json:"bio"`
	Skills          []string `json:"skills"`
	ExperienceYears int      `json:"experience_years" binding:"min=0"`
	HourlyRate      float64  `json:"hourly_rate" binding:"min=0"`
}

// CreateIdentityRequest 创建身份请求
type CreateIdentityRequest struct {
	IdentityType    string   `json:"identity_type" binding:"required,oneof=master apprentice"`
	Domain          string   `json:"domain" binding:"required"`
	Name            string   `json:"name" binding:"required"`
	Bio             string   `json:"bio"`
	Skills          []string `json:"skills"`
	ExperienceYears int      `json:"experience_years" binding:"min=0"`
	HourlyRate      float64  `json:"hourly_rate" binding:"min=0"`
}

// UpdateIdentityRequest 更新身份请求
type UpdateIdentityRequest struct {
	Name            string   `json:"name" binding:"required"`
	Bio             string   `json:"bio"`
	Skills          []string `json:"skills"`
	ExperienceYears int      `json:"experience_years" binding:"min=0"`
	HourlyRate      float64  `json:"hourly_rate" binding:"min=0"`
}

// UserPreferencesRequest 用户偏好请求
type UserPreferencesRequest struct {
	LearningStyle    string   `json:"learning_style" binding:"required,oneof=one-on-one structured browse"`
	TimePreference   string   `json:"time_preference" binding:"required,oneof=flexible fixed"`
	BudgetRange      string   `json:"budget_range" binding:"required,oneof=low medium high"`
	LearningGoals    []string `json:"learning_goals"`
	PreferredDomains []string `json:"preferred_domains"`
	ExperienceLevel  string   `json:"experience_level" binding:"required,oneof=beginner intermediate advanced"`
}
