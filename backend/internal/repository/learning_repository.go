package repository

import (
	"context"
	"time"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// LearningRepository 学习记录数据访问接口
type LearningRepository interface {
	GetLearningRecords(ctx context.Context, userID, courseID, status string, startDate, endDate time.Time, page, pageSize int) ([]*model.LearningRecord, int64, error)
	GetLearningRecordByID(ctx context.Context, recordID string) (*model.LearningRecord, error)
	UpdateLearningProgress(ctx context.Context, recordID string, progressPercentage float64, currentChapter string, studyTimeMinutes int) error
	CreateStudySession(ctx context.Context, session *model.StudySession) error
	GetStudySessions(ctx context.Context, learningRecordID string) ([]*model.StudySession, error)
	CreateAssignment(ctx context.Context, assignment *model.Assignment) error
	GetAssignments(ctx context.Context, learningRecordID string) ([]*model.Assignment, error)
	GetLearningStats(ctx context.Context, userID, period string) (*model.LearningStats, error)
	GetRecommendedPath(ctx context.Context, userID string) (*model.LearningPath, error)
}

// learningRepository 学习记录数据访问实现
type learningRepository struct {
	db *gorm.DB
}

// NewLearningRepository 创建学习记录数据访问实例
func NewLearningRepository(db *gorm.DB) LearningRepository {
	return &learningRepository{db: db}
}

// GetLearningRecords 获取学习记录列表
func (r *learningRepository) GetLearningRecords(ctx context.Context, userID, courseID, status string, startDate, endDate time.Time, page, pageSize int) ([]*model.LearningRecord, int64, error) {
	var records []*model.LearningRecord
	var total int64

	query := r.db.WithContext(ctx).Model(&model.LearningRecord{}).
		Preload("Course").
		Preload("Course.Mentor").
		Preload("Course.Mentor.Profile").
		Preload("User").
		Preload("User.Profile").
		Where("user_id = ?", userID)

	if courseID != "" {
		query = query.Where("course_id = ?", courseID)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if !startDate.IsZero() {
		query = query.Where("enrolled_at >= ?", startDate)
	}

	if !endDate.IsZero() {
		query = query.Where("enrolled_at <= ?", endDate)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("enrolled_at DESC").Offset(offset).Limit(pageSize).Find(&records).Error

	return records, total, err
}

// GetLearningRecordByID 根据ID获取学习记录
func (r *learningRepository) GetLearningRecordByID(ctx context.Context, recordID string) (*model.LearningRecord, error) {
	var record model.LearningRecord
	err := r.db.WithContext(ctx).
		Preload("Course").
		Preload("Course.Mentor").
		Preload("Course.Mentor.Profile").
		Preload("User").
		Preload("User.Profile").
		Where("id = ?", recordID).
		First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// UpdateLearningProgress 更新学习进度
func (r *learningRepository) UpdateLearningProgress(ctx context.Context, recordID string, progressPercentage float64, currentChapter string, studyTimeMinutes int) error {
	updates := map[string]interface{}{
		"progress_percentage": progressPercentage,
		"last_accessed_at":    time.Now(),
		"total_study_time":    gorm.Expr("total_study_time + ?", studyTimeMinutes),
	}

	if currentChapter != "" {
		updates["current_chapter"] = currentChapter
	}

	// 如果进度达到100%，标记为完成
	if progressPercentage >= 100 {
		now := time.Now()
		updates["status"] = "completed"
		updates["completed_at"] = &now
		updates["certificate_issued"] = true
	}

	return r.db.WithContext(ctx).Model(&model.LearningRecord{}).
		Where("id = ?", recordID).
		Updates(updates).Error
}

// CreateStudySession 创建学习会话
func (r *learningRepository) CreateStudySession(ctx context.Context, session *model.StudySession) error {
	return r.db.WithContext(ctx).Create(session).Error
}

// GetStudySessions 获取学习会话列表
func (r *learningRepository) GetStudySessions(ctx context.Context, learningRecordID string) ([]*model.StudySession, error) {
	var sessions []*model.StudySession
	err := r.db.WithContext(ctx).
		Where("learning_record_id = ?", learningRecordID).
		Order("start_time DESC").
		Find(&sessions).Error
	return sessions, err
}

// CreateAssignment 创建作业
func (r *learningRepository) CreateAssignment(ctx context.Context, assignment *model.Assignment) error {
	return r.db.WithContext(ctx).Create(assignment).Error
}

// GetAssignments 获取作业列表
func (r *learningRepository) GetAssignments(ctx context.Context, learningRecordID string) ([]*model.Assignment, error) {
	var assignments []*model.Assignment
	err := r.db.WithContext(ctx).
		Where("learning_record_id = ?", learningRecordID).
		Order("submitted_at DESC").
		Find(&assignments).Error
	return assignments, err
}

// GetLearningStats 获取学习统计
func (r *learningRepository) GetLearningStats(ctx context.Context, userID, period string) (*model.LearningStats, error) {
	var stats model.LearningStats

	// 构建时间范围查询
	var timeFilter string
	var args []interface{}

	switch period {
	case "week":
		timeFilter = "enrolled_at >= NOW() - INTERVAL '7 days'"
	case "month":
		timeFilter = "enrolled_at >= NOW() - INTERVAL '1 month'"
	case "year":
		timeFilter = "enrolled_at >= NOW() - INTERVAL '1 year'"
	case "all":
		timeFilter = "1=1"
	}

	// 获取总课程数
	err := r.db.WithContext(ctx).Model(&model.LearningRecord{}).
		Where("user_id = ? AND "+timeFilter, append([]interface{}{userID}, args...)...).
		Count(&stats.TotalCourses).Error
	if err != nil {
		return nil, err
	}

	// 获取已完成课程数
	err = r.db.WithContext(ctx).Model(&model.LearningRecord{}).
		Where("user_id = ? AND status = 'completed' AND "+timeFilter, append([]interface{}{userID}, args...)...).
		Count(&stats.CompletedCourses).Error
	if err != nil {
		return nil, err
	}

	// 获取总学习时间
	err = r.db.WithContext(ctx).Model(&model.LearningRecord{}).
		Select("COALESCE(SUM(total_study_time), 0)").
		Where("user_id = ? AND "+timeFilter, append([]interface{}{userID}, args...)...).
		Scan(&stats.TotalStudyHours).Error
	if err != nil {
		return nil, err
	}
	stats.TotalStudyHours = stats.TotalStudyHours / 60.0 // 转换为小时

	// 获取平均进度
	err = r.db.WithContext(ctx).Model(&model.LearningRecord{}).
		Select("COALESCE(AVG(progress_percentage), 0)").
		Where("user_id = ? AND "+timeFilter, append([]interface{}{userID}, args...)...).
		Scan(&stats.AverageProgress).Error
	if err != nil {
		return nil, err
	}

	// 获取作业统计
	err = r.db.WithContext(ctx).Model(&model.Assignment{}).
		Joins("JOIN learning_records ON assignments.learning_record_id = learning_records.id").
		Where("learning_records.user_id = ? AND learning_records."+timeFilter, append([]interface{}{userID}, args...)...).
		Count(&stats.TotalAssignments).Error
	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).Model(&model.Assignment{}).
		Joins("JOIN learning_records ON assignments.learning_record_id = learning_records.id").
		Where("learning_records.user_id = ? AND assignments.status = 'approved' AND learning_records."+timeFilter, append([]interface{}{userID}, args...)...).
		Count(&stats.CompletedAssignments).Error
	if err != nil {
		return nil, err
	}

	// 获取平均分数
	err = r.db.WithContext(ctx).Model(&model.Assignment{}).
		Joins("JOIN learning_records ON assignments.learning_record_id = learning_records.id").
		Select("COALESCE(AVG(assignments.score), 0)").
		Where("learning_records.user_id = ? AND assignments.score IS NOT NULL AND learning_records."+timeFilter, append([]interface{}{userID}, args...)...).
		Scan(&stats.AverageScore).Error
	if err != nil {
		return nil, err
	}

	// 获取证书数量
	err = r.db.WithContext(ctx).Model(&model.LearningRecord{}).
		Where("user_id = ? AND certificate_issued = true AND "+timeFilter, append([]interface{}{userID}, args...)...).
		Count(&stats.CertificatesEarned).Error
	if err != nil {
		return nil, err
	}

	// 计算连续学习天数（简化实现）
	stats.CurrentStreakDays = 7 // 暂时返回固定值

	return &stats, nil
}

// GetRecommendedPath 获取推荐学习路径
func (r *learningRepository) GetRecommendedPath(ctx context.Context, userID string) (*model.LearningPath, error) {
	// 这里实现推荐算法，暂时返回示例数据
	path := &model.LearningPath{
		CurrentLevel: "intermediate",
		NextCourses: []*model.LearningRecommendedCourse{
			{
				ID:                "COURSE_00000000001",
				Title:             "Go微服务架构",
				Reason:            "基于您当前的学习进度推荐",
				EstimatedDuration: 25,
			},
		},
		SkillsToDevelop:         []string{"微服务设计", "服务治理", "容器化部署"},
		EstimatedCompletionTime: "3个月",
	}

	return path, nil
}
