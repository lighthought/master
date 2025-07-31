package model

import "time"

// LearningCourseInfo 学习记录中的课程信息
type LearningCourseInfo struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description,omitempty"`
	CoverImage    string `json:"cover_image,omitempty"`
	DurationHours int    `json:"duration_hours,omitempty"`
}

// LearningMentorInfo 学习记录中的导师信息
type LearningMentorInfo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar,omitempty"`
}

// StudySessionInfo 学习会话信息
type StudySessionInfo struct {
	ID              string     `json:"id"`
	StartTime       time.Time  `json:"start_time"`
	EndTime         *time.Time `json:"end_time,omitempty"`
	DurationMinutes int        `json:"duration_minutes"`
	Chapter         string     `json:"chapter,omitempty"`
	Notes           string     `json:"notes,omitempty"`
}

// AssignmentInfo 作业信息
type AssignmentInfo struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Status      string     `json:"status"`
	SubmittedAt time.Time  `json:"submitted_at"`
	ReviewedAt  *time.Time `json:"reviewed_at,omitempty"`
	Score       *float64   `json:"score,omitempty"`
}

// LearningRecordInfo 学习记录信息
type LearningRecordInfo struct {
	ID                 string              `json:"id"`
	Course             *LearningCourseInfo `json:"course"`
	Mentor             *LearningMentorInfo `json:"mentor"`
	EnrollmentDate     time.Time           `json:"enrollment_date"`
	LastStudyDate      time.Time           `json:"last_study_date"`
	TotalStudyTime     int                 `json:"total_study_time"`
	ProgressPercentage float64             `json:"progress_percentage"`
	Status             string              `json:"status"`
	CurrentChapter     string              `json:"current_chapter,omitempty"`
	CompletedChapters  []string            `json:"completed_chapters,omitempty"`
	CertificateIssued  bool                `json:"certificate_issued"`
}

// LearningRecordDetail 学习记录详情
type LearningRecordDetail struct {
	ID                 string              `json:"id"`
	Course             *LearningCourseInfo `json:"course"`
	Mentor             *LearningMentorInfo `json:"mentor"`
	EnrollmentDate     time.Time           `json:"enrollment_date"`
	LastStudyDate      time.Time           `json:"last_study_date"`
	TotalStudyTime     int                 `json:"total_study_time"`
	ProgressPercentage float64             `json:"progress_percentage"`
	Status             string              `json:"status"`
	CurrentChapter     string              `json:"current_chapter,omitempty"`
	CompletedChapters  []string            `json:"completed_chapters,omitempty"`
	StudySessions      []*StudySessionInfo `json:"study_sessions,omitempty"`
	Assignments        []*AssignmentInfo   `json:"assignments,omitempty"`
	CertificateIssued  bool                `json:"certificate_issued"`
	CertificateURL     *string             `json:"certificate_url,omitempty"`
}

// LearningRecordListResponse 学习记录列表响应
type LearningRecordListResponse struct {
	Records    []*LearningRecordInfo `json:"records"`
	Pagination *PaginationResponse   `json:"pagination"`
}

// LearningRecordDetailResponse 学习记录详情响应
type LearningRecordDetailResponse struct {
	Record *LearningRecordDetail `json:"record"`
}

// UpdateProgressResponse 更新进度响应
type UpdateProgressResponse struct {
	RecordID           string  `json:"record_id"`
	ProgressPercentage float64 `json:"progress_percentage"`
}

// SubmitAssignmentResponse 提交作业响应
type SubmitAssignmentResponse struct {
	AssignmentID string `json:"assignment_id"`
}

// LearningStats 学习统计
type LearningStats struct {
	TotalCourses         int64   `json:"total_courses"`
	CompletedCourses     int64   `json:"completed_courses"`
	TotalStudyHours      float64 `json:"total_study_hours"`
	AverageProgress      float64 `json:"average_progress"`
	CurrentStreakDays    int     `json:"current_streak_days"`
	TotalAssignments     int64   `json:"total_assignments"`
	CompletedAssignments int64   `json:"completed_assignments"`
	AverageScore         float64 `json:"average_score"`
	CertificatesEarned   int64   `json:"certificates_earned"`
}

// LearningStatsResponse 学习统计响应
type LearningStatsResponse struct {
	Stats *LearningStats `json:"stats"`
}

// LearningRecommendedCourse 学习推荐课程
type LearningRecommendedCourse struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	Reason            string `json:"reason"`
	EstimatedDuration int    `json:"estimated_duration"`
}

// LearningPath 学习路径
type LearningPath struct {
	CurrentLevel            string                       `json:"current_level"`
	NextCourses             []*LearningRecommendedCourse `json:"next_courses"`
	SkillsToDevelop         []string                     `json:"skills_to_develop"`
	EstimatedCompletionTime string                       `json:"estimated_completion_time"`
}

// LearningPathResponse 学习路径响应
type LearningPathResponse struct {
	Path *LearningPath `json:"path"`
}
