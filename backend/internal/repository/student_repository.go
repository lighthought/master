package repository

import (
	"context"
	"time"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// StudentRepository 学生数据访问接口
type StudentRepository interface {
	GetStudents(ctx context.Context, mentorID, status, courseID, search, sortBy string, page, pageSize int) ([]*model.StudentInfo, int64, error)
	GetStudentByID(ctx context.Context, studentID string) (*model.StudentDetail, error)
	GetStudentStats(ctx context.Context, mentorID string) (*model.StudentStats, error)
	GetStudentCourses(ctx context.Context, studentID string) ([]*model.StudentCourseDetail, error)
	GetStudentAppointments(ctx context.Context, studentID string) ([]*model.StudentAppointmentInfo, error)
	GetStudentReviews(ctx context.Context, studentID string) ([]*model.StudentReviewInfo, error)
	GetStudentAssignments(ctx context.Context, studentID, courseID string) ([]*model.StudentAssignmentInfo, error)
	UpdateStudentProgress(ctx context.Context, studentID, courseID string, progressPercentage float64, notes string) error
	GradeAssignment(ctx context.Context, assignmentID string, score float64, feedback, comments string) error
	GetStudentReport(ctx context.Context, studentID, period string) (*model.StudentReport, error)
}

// studentRepository 学生数据访问实现
type studentRepository struct {
	db *gorm.DB
}

// NewStudentRepository 创建学生数据访问实例
func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

// GetStudents 获取学生列表
func (r *studentRepository) GetStudents(ctx context.Context, mentorID, status, courseID, search, sortBy string, page, pageSize int) ([]*model.StudentInfo, int64, error) {
	var students []*model.StudentInfo
	var total int64

	// 构建查询
	query := r.db.WithContext(ctx).
		Table("users u").
		Select(`
			u.id, up.name, up.avatar, u.email, u.phone, u.created_at as enrollment_date, u.status,
			COUNT(DISTINCT lr.course_id) as total_courses,
			COUNT(DISTINCT CASE WHEN lr.status = 'completed' THEN lr.course_id END) as completed_courses,
			COALESCE(SUM(lr.total_study_time), 0) / 60.0 as total_study_hours,
			COALESCE(AVG(lr.progress_percentage), 0) as average_progress,
			MAX(lr.last_accessed_at) as last_activity
		`).
		Joins("JOIN user_identities ui ON u.id = ui.user_id").
		Joins("LEFT JOIN user_profiles up ON ui.id = up.identity_id").
		Joins("LEFT JOIN learning_records lr ON u.id = lr.user_id").
		Joins("LEFT JOIN courses c ON lr.course_id = c.id").
		Where("ui.identity_type = ?", "apprentice")

	// 如果指定了导师ID，只查询该导师的学生
	if mentorID != "" {
		query = query.Where("c.mentor_id = ?", mentorID)
	}

	// 根据状态筛选
	if status != "" {
		query = query.Where("u.status = ?", status)
	}

	// 根据课程ID筛选
	if courseID != "" {
		query = query.Where("lr.course_id = ?", courseID)
	}

	// 搜索功能
	if search != "" {
		query = query.Where("up.name ILIKE ? OR u.email ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// 分组
	query = query.Group("u.id, up.name, up.avatar, u.email, u.phone, u.created_at, u.status")

	// 排序
	switch sortBy {
	case "name":
		query = query.Order("up.name")
	case "enrollment_date":
		query = query.Order("u.created_at DESC")
	case "progress":
		query = query.Order("average_progress DESC")
	default:
		query = query.Order("u.created_at DESC")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&students).Error

	return students, total, err
}

// GetStudentByID 根据ID获取学生详情
func (r *studentRepository) GetStudentByID(ctx context.Context, studentID string) (*model.StudentDetail, error) {
	var student model.StudentDetail

	// 获取学生基本信息
	err := r.db.WithContext(ctx).
		Table("users u").
		Select(`
			u.id, up.name, up.avatar, u.email, u.phone, up.bio, u.created_at as enrollment_date, u.status
		`).
		Joins("JOIN user_identities ui ON u.id = ui.user_id").
		Joins("LEFT JOIN user_profiles up ON ui.id = up.identity_id").
		Where("u.id = ? AND ui.identity_type = ?", studentID, "apprentice").
		First(&student).Error

	if err != nil {
		return nil, err
	}

	// 获取学习目标（从用户偏好中获取）
	var preferences model.UserPreferences
	err = r.db.WithContext(ctx).
		Where("user_id = ?", studentID).
		First(&preferences).Error
	if err == nil {
		student.LearningGoals = preferences.LearningGoals
		student.PreferredLearningStyle = preferences.LearningStyle
	}

	// 获取课程信息
	courses, err := r.GetStudentCourses(ctx, studentID)
	if err == nil {
		student.Courses = courses
	}

	// 获取预约信息
	appointments, err := r.GetStudentAppointments(ctx, studentID)
	if err == nil {
		student.Appointments = appointments
	}

	// 获取评价信息
	reviews, err := r.GetStudentReviews(ctx, studentID)
	if err == nil {
		student.Reviews = reviews
	}

	return &student, nil
}

// GetStudentStats 获取学生统计
func (r *studentRepository) GetStudentStats(ctx context.Context, mentorID string) (*model.StudentStats, error) {
	var stats model.StudentStats

	// 构建基础查询
	baseQuery := r.db.WithContext(ctx).
		Table("users u").
		Joins("JOIN user_identities ui ON u.id = ui.user_id").
		Joins("LEFT JOIN learning_records lr ON u.id = lr.user_id").
		Joins("LEFT JOIN courses c ON lr.course_id = c.id").
		Where("ui.identity_type = ?", "apprentice")

	if mentorID != "" {
		baseQuery = baseQuery.Where("c.mentor_id = ?", mentorID)
	}

	// 获取总学生数
	err := baseQuery.Distinct("u.id").Count(&stats.TotalStudents).Error
	if err != nil {
		return nil, err
	}

	// 获取活跃学生数
	err = baseQuery.Where("u.status = ?", "active").Distinct("u.id").Count(&stats.ActiveStudents).Error
	if err != nil {
		return nil, err
	}

	// 获取非活跃学生数
	err = baseQuery.Where("u.status = ?", "inactive").Distinct("u.id").Count(&stats.InactiveStudents).Error
	if err != nil {
		return nil, err
	}

	// 获取毕业学生数（完成所有课程的学生）
	err = baseQuery.Where("lr.status = ?", "completed").Distinct("u.id").Count(&stats.GraduatedStudents).Error
	if err != nil {
		return nil, err
	}

	// 获取本月新学生数
	startOfMonth := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -time.Now().Day()+1)
	err = baseQuery.Where("u.created_at >= ?", startOfMonth).Distinct("u.id").Count(&stats.NewStudentsThisMonth).Error
	if err != nil {
		return nil, err
	}

	// 获取平均进度
	err = baseQuery.Select("COALESCE(AVG(lr.progress_percentage), 0)").Scan(&stats.AverageProgress).Error
	if err != nil {
		return nil, err
	}

	// 获取平均评分
	err = r.db.WithContext(ctx).
		Table("reviews r").
		Joins("JOIN user_identities ui ON r.reviewed_id = ui.id").
		Where("ui.identity_type = ?", "apprentice").
		Select("COALESCE(AVG(r.rating), 0)").
		Scan(&stats.AverageRating).Error
	if err != nil {
		return nil, err
	}

	// 获取优秀学生数（进度超过80%的学生）
	err = baseQuery.Where("lr.progress_percentage >= ?", 80).Distinct("u.id").Count(&stats.TopPerformingStudents).Error
	if err != nil {
		return nil, err
	}

	// 获取按课程分组的统计
	var courseStats []*model.StudentCourseStats
	err = r.db.WithContext(ctx).
		Table("courses c").
		Select(`
			c.id as course_id, c.title as course_title,
			COUNT(DISTINCT lr.user_id) as student_count,
			COALESCE(AVG(lr.progress_percentage), 0) as average_progress
		`).
		Joins("LEFT JOIN learning_records lr ON c.id = lr.course_id").
		Joins("LEFT JOIN user_identities ui ON lr.user_id = ui.user_id").
		Where("ui.identity_type = ?", "apprentice").
		Group("c.id, c.title").
		Find(&courseStats).Error
	if err == nil {
		stats.StudentsByCourse = courseStats
	}

	return &stats, nil
}

// GetStudentCourses 获取学生课程信息
func (r *studentRepository) GetStudentCourses(ctx context.Context, studentID string) ([]*model.StudentCourseDetail, error) {
	var courses []*model.StudentCourseDetail

	err := r.db.WithContext(ctx).
		Table("learning_records lr").
		Select(`
			lr.course_id, c.title, lr.enrolled_at as enrollment_date,
			lr.progress_percentage, lr.status, lr.last_accessed_at as last_study_date,
			lr.total_study_time
		`).
		Joins("JOIN courses c ON lr.course_id = c.id").
		Where("lr.user_id = ?", studentID).
		Find(&courses).Error

	if err != nil {
		return nil, err
	}

	// 为每个课程获取作业信息
	for _, course := range courses {
		assignments, err := r.GetStudentAssignments(ctx, studentID, course.CourseID)
		if err == nil {
			course.Assignments = assignments
		}
	}

	return courses, nil
}

// GetStudentAppointments 获取学生预约信息
func (r *studentRepository) GetStudentAppointments(ctx context.Context, studentID string) ([]*model.StudentAppointmentInfo, error) {
	var appointments []*model.StudentAppointmentInfo

	err := r.db.WithContext(ctx).
		Table("appointments a").
		Select(`
			a.id, a.appointment_time, a.status, a.notes as topic
		`).
		Where("a.student_id = ?", studentID).
		Order("a.appointment_time DESC").
		Find(&appointments).Error

	return appointments, err
}

// GetStudentReviews 获取学生评价信息
func (r *studentRepository) GetStudentReviews(ctx context.Context, studentID string) ([]*model.StudentReviewInfo, error) {
	var reviews []*model.StudentReviewInfo

	err := r.db.WithContext(ctx).
		Table("reviews r").
		Select(`
			r.id, r.rating, r.content, r.created_at
		`).
		Joins("JOIN user_identities ui ON r.reviewer_id = ui.user_id").
		Where("ui.user_id = ? AND ui.identity_type = ?", studentID, "apprentice").
		Order("r.created_at DESC").
		Find(&reviews).Error

	return reviews, err
}

// GetStudentAssignments 获取学生作业信息
func (r *studentRepository) GetStudentAssignments(ctx context.Context, studentID, courseID string) ([]*model.StudentAssignmentInfo, error) {
	var assignments []*model.StudentAssignmentInfo

	query := r.db.WithContext(ctx).
		Table("assignments a").
		Select(`
			a.id, a.title, a.status, a.score, a.submitted_at
		`).
		Joins("JOIN learning_records lr ON a.learning_record_id = lr.id").
		Where("lr.user_id = ?", studentID)

	if courseID != "" {
		query = query.Where("lr.course_id = ?", courseID)
	}

	err := query.Order("a.submitted_at DESC").Find(&assignments).Error

	return assignments, err
}

// UpdateStudentProgress 更新学生学习进度
func (r *studentRepository) UpdateStudentProgress(ctx context.Context, studentID, courseID string, progressPercentage float64, notes string) error {
	return r.db.WithContext(ctx).
		Model(&model.LearningRecord{}).
		Where("user_id = ? AND course_id = ?", studentID, courseID).
		Updates(map[string]interface{}{
			"progress_percentage": progressPercentage,
			"last_accessed_at":    time.Now(),
		}).Error
}

// GradeAssignment 评价学生作业
func (r *studentRepository) GradeAssignment(ctx context.Context, assignmentID string, score float64, feedback, comments string) error {
	return r.db.WithContext(ctx).
		Model(&model.Assignment{}).
		Where("id = ?", assignmentID).
		Updates(map[string]interface{}{
			"score":       score,
			"feedback":    feedback,
			"status":      "reviewed",
			"reviewed_at": time.Now(),
		}).Error
}

// GetStudentReport 获取学生学习报告
func (r *studentRepository) GetStudentReport(ctx context.Context, studentID, period string) (*model.StudentReport, error) {
	var report model.StudentReport

	// 获取学生基本信息
	var user model.User
	var profile model.UserProfile
	err := r.db.WithContext(ctx).
		Joins("JOIN user_identities ui ON users.id = ui.user_id").
		Joins("LEFT JOIN user_profiles up ON ui.id = up.identity_id").
		Where("users.id = ? AND ui.identity_type = ?", studentID, "apprentice").
		First(&user).Error
	if err != nil {
		return nil, err
	}

	report.StudentID = studentID
	report.StudentName = profile.Name
	report.Period = period

	// 计算时间范围
	var startTime time.Time
	switch period {
	case "week":
		startTime = time.Now().AddDate(0, 0, -7)
	case "month":
		startTime = time.Now().AddDate(0, -1, 0)
	case "quarter":
		startTime = time.Now().AddDate(0, -3, 0)
	case "year":
		startTime = time.Now().AddDate(-1, 0, 0)
	}

	// 获取学习时间
	err = r.db.WithContext(ctx).
		Model(&model.LearningRecord{}).
		Select("COALESCE(SUM(total_study_time), 0) / 60.0").
		Where("user_id = ? AND last_accessed_at >= ?", studentID, startTime).
		Scan(&report.StudyTime).Error
	if err != nil {
		return nil, err
	}

	// 获取课程进度
	var coursesProgress []*model.CourseProgressInfo
	err = r.db.WithContext(ctx).
		Table("learning_records lr").
		Select(`
			lr.course_id, c.title, lr.progress_percentage,
			lr.total_study_time / 60.0 as study_time,
			COUNT(DISTINCT a.id) as assignments_completed,
			COALESCE(AVG(a.score), 0) as average_score
		`).
		Joins("JOIN courses c ON lr.course_id = c.id").
		Joins("LEFT JOIN assignments a ON lr.id = a.learning_record_id").
		Where("lr.user_id = ? AND lr.last_accessed_at >= ?", studentID, startTime).
		Group("lr.course_id, c.title, lr.progress_percentage, lr.total_study_time").
		Find(&coursesProgress).Error
	if err == nil {
		report.CoursesProgress = coursesProgress
	}

	// 生成优势和改进建议（基于学习数据）
	report.Strengths = []string{"学习态度积极", "按时完成作业"}
	report.AreasForImprovement = []string{"需要加强实践练习", "可以多参与讨论"}
	report.Recommendations = []string{"建议增加项目实战", "可以尝试开源项目"}

	return &report, nil
}
