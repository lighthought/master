package model

// UserStatsRequest 用户统计请求
type UserStatsRequest struct {
	UserID string `uri:"user_id" binding:"required"`
}

// UserStatsResponse 用户统计响应
type UserStatsResponse struct {
	LearningStats *UserLearningStats `json:"learning_stats"`
	TeachingStats *UserTeachingStats `json:"teaching_stats"`
}

// UserLearningStats 用户学习统计
type UserLearningStats struct {
	EnrolledCourses    int64   `json:"enrolled_courses"`
	CompletedCourses   int64   `json:"completed_courses"`
	TotalStudyHours    float64 `json:"total_study_hours"`
	CurrentCourses     int64   `json:"current_courses"`
	AverageProgress    float64 `json:"average_progress"`
	CertificatesEarned int64   `json:"certificates_earned"`
}

// UserTeachingStats 用户教学统计
type UserTeachingStats struct {
	TotalStudents    int64   `json:"total_students"`
	TotalIncome      float64 `json:"total_income"`
	AverageRating    float64 `json:"average_rating"`
	TotalCourses     int64   `json:"total_courses"`
	ActiveStudents   int64   `json:"active_students"`
	CompletedLessons int64   `json:"completed_lessons"`
}
