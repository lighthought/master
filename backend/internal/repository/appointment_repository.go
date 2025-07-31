package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// AppointmentRepository 预约数据访问接口
type AppointmentRepository interface {
	CreateAppointment(ctx context.Context, appointment *model.AppointmentModel) error
	GetAppointments(ctx context.Context, userID, status, appointmentType string, page, pageSize int) ([]*model.AppointmentModel, int64, error)
	GetAppointmentByID(ctx context.Context, appointmentID string) (*model.AppointmentModel, error)
	UpdateAppointmentStatus(ctx context.Context, appointmentID, status string) error
	CancelAppointment(ctx context.Context, appointmentID string) error
	GetMentorAppointmentStats(ctx context.Context, mentorID string) (*model.MentorAppointmentStats, error)
	GetMentorStudentCount(ctx context.Context, mentorID string) (int64, error)
	GetMentorTotalHours(ctx context.Context, mentorID string) (int64, error)
}

// appointmentRepository 预约数据访问实现
type appointmentRepository struct {
	db *gorm.DB
}

// NewAppointmentRepository 创建预约数据访问实例
func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db: db}
}

// CreateAppointment 创建预约
func (r *appointmentRepository) CreateAppointment(ctx context.Context, appointment *model.AppointmentModel) error {
	return r.db.WithContext(ctx).Create(appointment).Error
}

// GetAppointments 获取预约列表
func (r *appointmentRepository) GetAppointments(ctx context.Context, userID, status, appointmentType string, page, pageSize int) ([]*model.AppointmentModel, int64, error) {
	var appointments []*model.AppointmentModel
	var total int64

	query := r.db.WithContext(ctx).Model(&model.AppointmentModel{}).
		Preload("Student").
		Preload("Student.Profile").
		Preload("Mentor").
		Preload("Mentor.Profile")

	// 根据类型筛选
	switch appointmentType {
	case "student":
		query = query.Where("student_id = ?", userID)
	case "mentor":
		query = query.Where("mentor_id = ?", userID)
	default:
		// 如果没有指定类型，查询用户相关的所有预约
		query = query.Where("student_id = ? OR mentor_id = ?", userID, userID)
	}

	// 根据状态筛选
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("appointment_time DESC").Offset(offset).Limit(pageSize).Find(&appointments).Error

	return appointments, total, err
}

// GetAppointmentByID 根据ID获取预约详情
func (r *appointmentRepository) GetAppointmentByID(ctx context.Context, appointmentID string) (*model.AppointmentModel, error) {
	var appointment model.AppointmentModel
	err := r.db.WithContext(ctx).
		Preload("Student").
		Preload("Student.Profile").
		Preload("Mentor").
		Preload("Mentor.Profile").
		Where("id = ?", appointmentID).
		First(&appointment).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

// UpdateAppointmentStatus 更新预约状态
func (r *appointmentRepository) UpdateAppointmentStatus(ctx context.Context, appointmentID, status string) error {
	return r.db.WithContext(ctx).
		Model(&model.AppointmentModel{}).
		Where("id = ?", appointmentID).
		Update("status", status).Error
}

// CancelAppointment 取消预约
func (r *appointmentRepository) CancelAppointment(ctx context.Context, appointmentID string) error {
	return r.db.WithContext(ctx).
		Model(&model.AppointmentModel{}).
		Where("id = ?", appointmentID).
		Update("status", "cancelled").Error
}

// GetMentorAppointmentStats 获取大师预约统计
func (r *appointmentRepository) GetMentorAppointmentStats(ctx context.Context, mentorID string) (*model.MentorAppointmentStats, error) {
	var stats model.MentorAppointmentStats

	// 获取各种状态的预约数量
	err := r.db.WithContext(ctx).Model(&model.AppointmentModel{}).
		Where("mentor_id = ?", mentorID).
		Select("COUNT(*) as total_appointments").
		Scan(&stats.TotalAppointments).Error
	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).Model(&model.AppointmentModel{}).
		Where("mentor_id = ? AND status = ?", mentorID, "pending").
		Select("COUNT(*) as pending_appointments").
		Scan(&stats.PendingAppointments).Error
	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).Model(&model.AppointmentModel{}).
		Where("mentor_id = ? AND status = ?", mentorID, "confirmed").
		Select("COUNT(*) as confirmed_appointments").
		Scan(&stats.ConfirmedAppointments).Error
	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).Model(&model.AppointmentModel{}).
		Where("mentor_id = ? AND status = ?", mentorID, "completed").
		Select("COUNT(*) as completed_appointments").
		Scan(&stats.CompletedAppointments).Error
	if err != nil {
		return nil, err
	}

	err = r.db.WithContext(ctx).Model(&model.AppointmentModel{}).
		Where("mentor_id = ? AND status = ?", mentorID, "cancelled").
		Select("COUNT(*) as cancelled_appointments").
		Scan(&stats.CancelledAppointments).Error
	if err != nil {
		return nil, err
	}

	// 获取总收入
	err = r.db.WithContext(ctx).Model(&model.AppointmentModel{}).
		Where("mentor_id = ? AND status IN (?, ?)", mentorID, "confirmed", "completed").
		Select("COALESCE(SUM(price), 0) as total_earnings").
		Scan(&stats.TotalEarnings).Error
	if err != nil {
		return nil, err
	}

	// 获取平均评分
	err = r.db.WithContext(ctx).Model(&model.Review{}).
		Joins("JOIN appointments a ON reviews.appointment_id = a.id").
		Where("a.mentor_id = ? AND reviews.review_type = ?", mentorID, "appointment").
		Select("COALESCE(AVG(reviews.rating), 0) as average_rating").
		Scan(&stats.AverageRating).Error
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// GetMentorStudentCount 获取导师的学生数量
func (r *appointmentRepository) GetMentorStudentCount(ctx context.Context, mentorID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.AppointmentModel{}).
		Where("mentor_id = ?", mentorID).
		Distinct("student_id").
		Count(&count).Error
	return count, err
}

// GetMentorTotalHours 获取导师的总教学小时数
func (r *appointmentRepository) GetMentorTotalHours(ctx context.Context, mentorID string) (int64, error) {
	var totalMinutes int64
	err := r.db.WithContext(ctx).
		Model(&model.AppointmentModel{}).
		Where("mentor_id = ? AND status IN (?, ?)", mentorID, "confirmed", "completed").
		Select("COALESCE(SUM(duration_minutes), 0)").
		Scan(&totalMinutes).Error
	if err != nil {
		return 0, err
	}
	// 转换为小时
	return totalMinutes / 60, nil
}
