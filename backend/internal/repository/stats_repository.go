package repository

import (
	"context"
	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// StatsRepository 统计数据访问接口
type StatsRepository interface {
	GetUserLearningStats(ctx context.Context, userID string) (*model.UserLearningStats, error)
	GetUserTeachingStats(ctx context.Context, userID string) (*model.UserTeachingStats, error)
}

type statsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) StatsRepository {
	return &statsRepository{db: db}
}

func (r *statsRepository) GetUserLearningStats(ctx context.Context, userID string) (*model.UserLearningStats, error) {
	var stats model.UserLearningStats

	// 获取已报名课程数
	err := r.db.WithContext(ctx).
		Table("learning_records").
		Where("user_id = ?", userID).
		Count(&stats.EnrolledCourses).Error
	if err != nil {
		return nil, err
	}

	// 获取已完成课程数
	err = r.db.WithContext(ctx).
		Table("learning_records").
		Where("user_id = ? AND status = ?", userID, "completed").
		Count(&stats.CompletedCourses).Error
	if err != nil {
		return nil, err
	}

	// 获取当前学习中的课程数
	err = r.db.WithContext(ctx).
		Table("learning_records").
		Where("user_id = ? AND status IN (?)", userID, []string{"enrolled", "learning"}).
		Count(&stats.CurrentCourses).Error
	if err != nil {
		return nil, err
	}

	// 获取总学习时长（小时）
	err = r.db.WithContext(ctx).
		Table("learning_records").
		Select("COALESCE(SUM(total_study_time), 0)").
		Where("user_id = ?", userID).
		Scan(&stats.TotalStudyHours).Error
	if err != nil {
		return nil, err
	}
	stats.TotalStudyHours = stats.TotalStudyHours / 60.0 // 转换为小时

	// 获取平均进度
	err = r.db.WithContext(ctx).
		Table("learning_records").
		Select("COALESCE(AVG(progress_percentage), 0)").
		Where("user_id = ?", userID).
		Scan(&stats.AverageProgress).Error
	if err != nil {
		return nil, err
	}

	// 获取获得的证书数
	err = r.db.WithContext(ctx).
		Table("learning_records").
		Where("user_id = ? AND certificate_issued = ?", userID, true).
		Count(&stats.CertificatesEarned).Error
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func (r *statsRepository) GetUserTeachingStats(ctx context.Context, userID string) (*model.UserTeachingStats, error) {
	var stats model.UserTeachingStats

	// 获取导师身份ID
	var mentorID string
	err := r.db.WithContext(ctx).
		Table("mentors m").
		Select("m.id").
		Joins("LEFT JOIN user_identities ui ON m.identity_id = ui.id").
		Where("ui.user_id = ?", userID).
		Scan(&mentorID).Error
	if err != nil {
		return nil, err
	}

	if mentorID == "" {
		// 用户不是导师，返回空统计
		return &stats, nil
	}

	// 获取总课程数
	err = r.db.WithContext(ctx).
		Table("courses").
		Where("mentor_id = ?", mentorID).
		Count(&stats.TotalCourses).Error
	if err != nil {
		return nil, err
	}

	// 获取总学生数
	err = r.db.WithContext(ctx).
		Table("learning_records lr").
		Joins("LEFT JOIN courses c ON lr.course_id = c.id").
		Where("c.mentor_id = ?", mentorID).
		Distinct("lr.user_id").
		Count(&stats.TotalStudents).Error
	if err != nil {
		return nil, err
	}

	// 获取活跃学生数（最近30天有学习记录）
	err = r.db.WithContext(ctx).
		Table("learning_records lr").
		Joins("LEFT JOIN courses c ON lr.course_id = c.id").
		Where("c.mentor_id = ? AND lr.last_accessed_at >= NOW() - INTERVAL '30 days'", mentorID).
		Distinct("lr.user_id").
		Count(&stats.ActiveStudents).Error
	if err != nil {
		return nil, err
	}

	// 获取总收入
	err = r.db.WithContext(ctx).
		Table("income_transactions").
		Select("COALESCE(SUM(net_income), 0)").
		Where("mentor_id = ? AND status = ?", mentorID, "completed").
		Scan(&stats.TotalIncome).Error
	if err != nil {
		return nil, err
	}

	// 获取平均评分
	err = r.db.WithContext(ctx).
		Table("reviews r").
		Joins("LEFT JOIN user_identities ui ON r.reviewed_id = ui.id").
		Select("COALESCE(AVG(r.rating), 0)").
		Where("ui.user_id = ? AND r.review_type = ?", userID, "mentor").
		Scan(&stats.AverageRating).Error
	if err != nil {
		return nil, err
	}

	// 获取已完成课程数（作为已完成课程数）
	err = r.db.WithContext(ctx).
		Table("learning_records lr").
		Joins("LEFT JOIN courses c ON lr.course_id = c.id").
		Where("c.mentor_id = ? AND lr.status = ?", mentorID, "completed").
		Count(&stats.CompletedLessons).Error
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
