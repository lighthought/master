package model

import "time"

// StudentInfo 学生信息
type StudentInfo struct {
	ID               string               `json:"id"`
	Name             string               `json:"name"`
	Avatar           string               `json:"avatar"`
	Email            string               `json:"email"`
	Phone            string               `json:"phone"`
	EnrollmentDate   time.Time            `json:"enrollment_date"`
	Status           string               `json:"status"`
	TotalCourses     int                  `json:"total_courses"`
	CompletedCourses int                  `json:"completed_courses"`
	TotalStudyHours  float64              `json:"total_study_hours"`
	AverageProgress  float64              `json:"average_progress"`
	LastActivity     time.Time            `json:"last_activity"`
	CurrentCourses   []*StudentCourseInfo `json:"current_courses"`
}

// StudentCourseInfo 学生课程信息
type StudentCourseInfo struct {
	CourseID           string  `json:"course_id"`
	Title              string  `json:"title"`
	ProgressPercentage float64 `json:"progress_percentage"`
	Status             string  `json:"status"`
}

// StudentDetail 学生详情
type StudentDetail struct {
	ID                     string                    `json:"id"`
	Name                   string                    `json:"name"`
	Avatar                 string                    `json:"avatar"`
	Email                  string                    `json:"email"`
	Phone                  string                    `json:"phone"`
	Bio                    string                    `json:"bio"`
	EnrollmentDate         time.Time                 `json:"enrollment_date"`
	Status                 string                    `json:"status"`
	LearningGoals          []string                  `json:"learning_goals"`
	PreferredLearningStyle string                    `json:"preferred_learning_style"`
	Courses                []*StudentCourseDetail    `json:"courses"`
	Appointments           []*StudentAppointmentInfo `json:"appointments"`
	Reviews                []*StudentReviewInfo      `json:"reviews"`
}

// StudentCourseDetail 学生课程详情
type StudentCourseDetail struct {
	CourseID           string                   `json:"course_id"`
	Title              string                   `json:"title"`
	EnrollmentDate     time.Time                `json:"enrollment_date"`
	ProgressPercentage float64                  `json:"progress_percentage"`
	Status             string                   `json:"status"`
	LastStudyDate      time.Time                `json:"last_study_date"`
	TotalStudyTime     int                      `json:"total_study_time"`
	Assignments        []*StudentAssignmentInfo `json:"assignments"`
}

// StudentAssignmentInfo 学生作业信息
type StudentAssignmentInfo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	Score       *float64  `json:"score"`
	SubmittedAt time.Time `json:"submitted_at"`
}

// StudentAppointmentInfo 学生预约信息
type StudentAppointmentInfo struct {
	ID              string    `json:"id"`
	AppointmentTime time.Time `json:"appointment_time"`
	Status          string    `json:"status"`
	Topic           string    `json:"topic"`
}

// StudentReviewInfo 学生评价信息
type StudentReviewInfo struct {
	ID        string    `json:"id"`
	Rating    int       `json:"rating"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// StudentListResponse 学生列表响应
type StudentListResponse struct {
	Students   []*StudentInfo      `json:"students"`
	Pagination *PaginationResponse `json:"pagination"`
}

// StudentDetailResponse 学生详情响应
type StudentDetailResponse struct {
	Student *StudentDetail `json:"student"`
}

// StudentStats 学生统计
type StudentStats struct {
	TotalStudents         int64                 `json:"total_students"`
	ActiveStudents        int64                 `json:"active_students"`
	InactiveStudents      int64                 `json:"inactive_students"`
	GraduatedStudents     int64                 `json:"graduated_students"`
	NewStudentsThisMonth  int64                 `json:"new_students_this_month"`
	AverageProgress       float64               `json:"average_progress"`
	AverageRating         float64               `json:"average_rating"`
	TopPerformingStudents int64                 `json:"top_performing_students"`
	StudentsByCourse      []*StudentCourseStats `json:"students_by_course"`
}

// StudentCourseStats 学生课程统计
type StudentCourseStats struct {
	CourseID        string  `json:"course_id"`
	CourseTitle     string  `json:"course_title"`
	StudentCount    int     `json:"student_count"`
	AverageProgress float64 `json:"average_progress"`
}

// StudentStatsResponse 学生统计响应
type StudentStatsResponse struct {
	Stats *StudentStats `json:"stats"`
}

// SendMessageResponse 发送消息响应
type SendMessageResponse struct {
	MessageID string `json:"message_id"`
}

// MessageInfo 消息信息
type MessageInfo struct {
	ID            string    `json:"id"`
	Content       string    `json:"content"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	FromUserID    string    `json:"from_user_id"`
	FromUserEmail string    `json:"from_user_email"`
	ToUserID      string    `json:"to_user_id"`
	ToUserEmail   string    `json:"to_user_email"`
}

// MessageListResponse 消息列表响应
type MessageListResponse struct {
	Messages   []*MessageInfo      `json:"messages"`
	Pagination *PaginationResponse `json:"pagination"`
}

// StudentProgressResponse 更新进度响应
type StudentProgressResponse struct {
	StudentID          string  `json:"student_id"`
	CourseID           string  `json:"course_id"`
	ProgressPercentage float64 `json:"progress_percentage"`
}

// GradeAssignmentResponse 评价作业响应
type GradeAssignmentResponse struct {
	AssignmentID string  `json:"assignment_id"`
	Score        float64 `json:"score"`
}

// StudentReport 学生学习报告
type StudentReport struct {
	StudentID           string                `json:"student_id"`
	StudentName         string                `json:"student_name"`
	Period              string                `json:"period"`
	StudyTime           float64               `json:"study_time"`
	CoursesProgress     []*CourseProgressInfo `json:"courses_progress"`
	Strengths           []string              `json:"strengths"`
	AreasForImprovement []string              `json:"areas_for_improvement"`
	Recommendations     []string              `json:"recommendations"`
}

// CourseProgressInfo 课程进度信息
type CourseProgressInfo struct {
	CourseID             string  `json:"course_id"`
	Title                string  `json:"title"`
	ProgressPercentage   float64 `json:"progress_percentage"`
	StudyTime            float64 `json:"study_time"`
	AssignmentsCompleted int     `json:"assignments_completed"`
	AverageScore         float64 `json:"average_score"`
}

// StudentReportResponse 学习报告响应
type StudentReportResponse struct {
	Report *StudentReport `json:"report"`
}
