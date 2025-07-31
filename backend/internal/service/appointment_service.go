package service

import (
	"context"
	"errors"
	"math"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// AppointmentService 预约服务接口
type AppointmentService interface {
	CreateAppointment(ctx context.Context, studentID string, req *model.CreateAppointmentRequest) (*model.CreateAppointmentResponse, error)
	GetAppointments(ctx context.Context, userID string, req *model.AppointmentListRequest) (*model.AppointmentListResponse, error)
	GetAppointmentDetail(ctx context.Context, appointmentID string) (*model.AppointmentDetailResponse, error)
	UpdateAppointmentStatus(ctx context.Context, appointmentID string, req *model.UpdateAppointmentStatusRequest) (*model.UpdateAppointmentStatusResponse, error)
	CancelAppointment(ctx context.Context, appointmentID string) (*model.CancelAppointmentResponse, error)
	GetMentorAppointmentStats(ctx context.Context, mentorID string) (*model.MentorAppointmentStatsResponse, error)
}

// appointmentService 预约服务实现
type appointmentService struct {
	appointmentRepo repository.AppointmentRepository
	mentorRepo      repository.MentorRepository
}

// NewAppointmentService 创建预约服务实例
func NewAppointmentService(appointmentRepo repository.AppointmentRepository, mentorRepo repository.MentorRepository) AppointmentService {
	return &appointmentService{
		appointmentRepo: appointmentRepo,
		mentorRepo:      mentorRepo,
	}
}

// CreateAppointment 创建预约
func (s *appointmentService) CreateAppointment(ctx context.Context, studentID string, req *model.CreateAppointmentRequest) (*model.CreateAppointmentResponse, error) {
	// 获取大师信息以计算价格
	mentor, err := s.mentorRepo.GetMentorByID(ctx, req.MentorID)
	if err != nil {
		return nil, errors.New("大师不存在")
	}

	// 计算价格（基于时长和大师时薪）
	price := float64(req.DurationMinutes) / 60.0 * mentor.HourlyRate

	// 创建预约
	appointment := &model.AppointmentModel{
		StudentID:       studentID,
		MentorID:        req.MentorID,
		AppointmentTime: req.AppointmentTime,
		DurationMinutes: req.DurationMinutes,
		MeetingType:     req.MeetingType,
		Status:          "pending",
		Price:           price,
		Notes:           req.Notes,
	}

	err = s.appointmentRepo.CreateAppointment(ctx, appointment)
	if err != nil {
		return nil, err
	}

	return &model.CreateAppointmentResponse{
		AppointmentID: appointment.ID,
		Status:        appointment.Status,
		Price:         appointment.Price,
	}, nil
}

// GetAppointments 获取预约列表
func (s *appointmentService) GetAppointments(ctx context.Context, userID string, req *model.AppointmentListRequest) (*model.AppointmentListResponse, error) {
	appointments, total, err := s.appointmentRepo.GetAppointments(ctx, userID, req.Status, req.Type, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	appointmentInfos := make([]*model.AppointmentInfo, len(appointments))
	for i, appointment := range appointments {
		appointmentInfos[i] = s.convertToAppointmentInfo(appointment)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.AppointmentListResponse{
		Appointments: appointmentInfos,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// GetAppointmentDetail 获取预约详情
func (s *appointmentService) GetAppointmentDetail(ctx context.Context, appointmentID string) (*model.AppointmentDetailResponse, error) {
	appointment, err := s.appointmentRepo.GetAppointmentByID(ctx, appointmentID)
	if err != nil {
		return nil, errors.New("预约不存在")
	}

	// 转换为响应格式
	appointmentDetail := s.convertToAppointmentDetail(appointment)

	return &model.AppointmentDetailResponse{
		Appointment: appointmentDetail,
	}, nil
}

// UpdateAppointmentStatus 更新预约状态
func (s *appointmentService) UpdateAppointmentStatus(ctx context.Context, appointmentID string, req *model.UpdateAppointmentStatusRequest) (*model.UpdateAppointmentStatusResponse, error) {
	err := s.appointmentRepo.UpdateAppointmentStatus(ctx, appointmentID, req.Status)
	if err != nil {
		return nil, err
	}

	return &model.UpdateAppointmentStatusResponse{
		AppointmentID: appointmentID,
		Status:        req.Status,
	}, nil
}

// CancelAppointment 取消预约
func (s *appointmentService) CancelAppointment(ctx context.Context, appointmentID string) (*model.CancelAppointmentResponse, error) {
	err := s.appointmentRepo.CancelAppointment(ctx, appointmentID)
	if err != nil {
		return nil, err
	}

	return &model.CancelAppointmentResponse{
		AppointmentID: appointmentID,
	}, nil
}

// GetMentorAppointmentStats 获取大师预约统计
func (s *appointmentService) GetMentorAppointmentStats(ctx context.Context, mentorID string) (*model.MentorAppointmentStatsResponse, error) {
	stats, err := s.appointmentRepo.GetMentorAppointmentStats(ctx, mentorID)
	if err != nil {
		return nil, err
	}

	return &model.MentorAppointmentStatsResponse{
		Stats: stats,
	}, nil
}

// convertToAppointmentInfo 转换为预约信息
func (s *appointmentService) convertToAppointmentInfo(appointment *model.AppointmentModel) *model.AppointmentInfo {
	appointmentInfo := &model.AppointmentInfo{
		ID:              appointment.ID,
		AppointmentTime: appointment.AppointmentTime,
		DurationMinutes: appointment.DurationMinutes,
		MeetingType:     appointment.MeetingType,
		Status:          appointment.Status,
		Price:           appointment.Price,
	}

	// 转换大师信息
	if appointment.Mentor != nil {
		appointmentInfo.Mentor = &model.MentorInfo{
			ID: appointment.Mentor.ID,
		}
		if appointment.Mentor.Profile != nil {
			appointmentInfo.Mentor.Name = appointment.Mentor.Profile.Name
			appointmentInfo.Mentor.Avatar = appointment.Mentor.Profile.Avatar
		}
	}

	// 转换学生信息
	if appointment.Student != nil {
		appointmentInfo.Student = &model.UserInfo{
			ID: appointment.Student.ID,
		}
		// 这里应该从用户档案中获取姓名和头像
		// 暂时使用邮箱作为姓名
		appointmentInfo.Student.Email = appointment.Student.Email
	}

	return appointmentInfo
}

// convertToAppointmentDetail 转换为预约详情
func (s *appointmentService) convertToAppointmentDetail(appointment *model.AppointmentModel) *model.AppointmentDetail {
	appointmentDetail := &model.AppointmentDetail{
		ID:              appointment.ID,
		AppointmentTime: appointment.AppointmentTime,
		DurationMinutes: appointment.DurationMinutes,
		MeetingType:     appointment.MeetingType,
		Status:          appointment.Status,
		Price:           appointment.Price,
		Notes:           appointment.Notes,
		CreatedAt:       appointment.CreatedAt,
		UpdatedAt:       appointment.UpdatedAt,
	}

	// 转换大师信息
	if appointment.Mentor != nil {
		appointmentDetail.Mentor = &model.MentorInfo{
			ID: appointment.Mentor.ID,
		}
		if appointment.Mentor.Profile != nil {
			appointmentDetail.Mentor.Name = appointment.Mentor.Profile.Name
			appointmentDetail.Mentor.Avatar = appointment.Mentor.Profile.Avatar
		}
	}

	// 转换学生信息
	if appointment.Student != nil {
		appointmentDetail.Student = &model.UserInfo{
			ID: appointment.Student.ID,
		}
		// 这里应该从用户档案中获取姓名和头像
		// 暂时使用邮箱作为姓名
		appointmentDetail.Student.Email = appointment.Student.Email
	}

	return appointmentDetail
}
