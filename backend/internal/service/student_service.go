package service

import (
	"context"
	"errors"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// StudentService 学生服务接口
type StudentService interface {
	GetStudents(ctx context.Context, mentorID string, req *model.StudentListRequest) (*model.StudentListResponse, error)
	GetStudentByID(ctx context.Context, studentID string) (*model.StudentDetailResponse, error)
	GetStudentStats(ctx context.Context, mentorID string) (*model.StudentStatsResponse, error)
	SendMessage(ctx context.Context, mentorID, studentID string, req *model.SendMessageRequest) (*model.SendMessageResponse, error)
	GetMessages(ctx context.Context, mentorID, studentID string, page, pageSize int) (*model.MessageListResponse, error)
	UpdateStudentProgress(ctx context.Context, mentorID, studentID, courseID string, req *model.StudentProgressRequest) (*model.StudentProgressResponse, error)
	GradeAssignment(ctx context.Context, mentorID, studentID, assignmentID string, req *model.GradeAssignmentRequest) (*model.GradeAssignmentResponse, error)
	GetStudentReport(ctx context.Context, mentorID, studentID string, req *model.StudentReportRequest) (*model.StudentReportResponse, error)
}

// studentService 学生服务实现
type studentService struct {
	studentRepo     repository.StudentRepository
	userRepo        repository.UserRepository
	identityRepo    repository.IdentityRepository
	appointmentRepo repository.AppointmentRepository
	messageRepo     repository.MessageRepository
}

// NewStudentService 创建学生服务实例
func NewStudentService(studentRepo repository.StudentRepository, userRepo repository.UserRepository, identityRepo repository.IdentityRepository, appointmentRepo repository.AppointmentRepository, messageRepo repository.MessageRepository) StudentService {
	return &studentService{
		studentRepo:     studentRepo,
		userRepo:        userRepo,
		identityRepo:    identityRepo,
		appointmentRepo: appointmentRepo,
		messageRepo:     messageRepo,
	}
}

// GetStudents 获取学生列表
func (s *studentService) GetStudents(ctx context.Context, mentorID string, req *model.StudentListRequest) (*model.StudentListResponse, error) {
	// 验证导师身份
	if mentorID != "" {
		identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, mentorID)
		if err != nil {
			return nil, err
		}

		isMentor := false
		for _, identity := range identities {
			if identity.IdentityType == "master" && identity.Status == "active" {
				isMentor = true
				break
			}
		}

		if !isMentor {
			return nil, errors.New("只有导师可以查看学生列表")
		}
	}

	// 获取学生列表
	students, total, err := s.studentRepo.GetStudents(ctx, mentorID, req.Status, req.CourseID, req.Search, req.SortBy, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 构建分页响应
	totalPages := (int(total) + req.PageSize - 1) / req.PageSize
	pagination := &model.PaginationResponse{
		Total:      total,
		Page:       req.Page,
		PageSize:   req.PageSize,
		TotalPages: totalPages,
	}

	return &model.StudentListResponse{
		Students:   students,
		Pagination: pagination,
	}, nil
}

// GetStudentByID 根据ID获取学生详情
func (s *studentService) GetStudentByID(ctx context.Context, studentID string) (*model.StudentDetailResponse, error) {
	// 验证学生是否存在
	_, err := s.userRepo.GetByID(ctx, studentID)
	if err != nil {
		return nil, errors.New("学生不存在")
	}

	// 验证学生身份
	identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, studentID)
	if err != nil {
		return nil, err
	}

	isStudent := false
	for _, identity := range identities {
		if identity.IdentityType == "apprentice" {
			isStudent = true
			break
		}
	}

	if !isStudent {
		return nil, errors.New("该用户不是学生")
	}

	// 获取学生详情
	student, err := s.studentRepo.GetStudentByID(ctx, studentID)
	if err != nil {
		return nil, err
	}

	return &model.StudentDetailResponse{
		Student: student,
	}, nil
}

// GetStudentStats 获取学生统计
func (s *studentService) GetStudentStats(ctx context.Context, mentorID string) (*model.StudentStatsResponse, error) {
	// 验证导师身份
	if mentorID != "" {
		identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, mentorID)
		if err != nil {
			return nil, err
		}

		isMentor := false
		for _, identity := range identities {
			if identity.IdentityType == "master" && identity.Status == "active" {
				isMentor = true
				break
			}
		}

		if !isMentor {
			return nil, errors.New("只有导师可以查看学生统计")
		}
	}

	// 获取学生统计
	stats, err := s.studentRepo.GetStudentStats(ctx, mentorID)
	if err != nil {
		return nil, err
	}

	return &model.StudentStatsResponse{
		Stats: stats,
	}, nil
}

// SendMessage 发送消息给学生
func (s *studentService) SendMessage(ctx context.Context, mentorID, studentID string, req *model.SendMessageRequest) (*model.SendMessageResponse, error) {
	// 验证导师身份
	identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, mentorID)
	if err != nil {
		return nil, err
	}

	isMentor := false
	for _, identity := range identities {
		if identity.IdentityType == "master" && identity.Status == "active" {
			isMentor = true
			break
		}
	}

	if !isMentor {
		return nil, errors.New("只有导师可以发送消息")
	}

	// 验证学生是否存在
	_, err = s.userRepo.GetByID(ctx, studentID)
	if err != nil {
		return nil, errors.New("学生不存在")
	}

	// 创建消息
	message := &model.Message{
		FromUserID: mentorID,
		ToUserID:   studentID,
		Content:    req.Content,
		Type:       req.Type,
	}

	err = s.messageRepo.Create(ctx, message)
	if err != nil {
		return nil, err
	}

	return &model.SendMessageResponse{
		MessageID: message.ID,
	}, nil
}

// GetMessages 获取与学生聊天记录
func (s *studentService) GetMessages(ctx context.Context, mentorID, studentID string, page, pageSize int) (*model.MessageListResponse, error) {
	// 验证导师身份
	identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, mentorID)
	if err != nil {
		return nil, err
	}

	isMentor := false
	for _, identity := range identities {
		if identity.IdentityType == "master" && identity.Status == "active" {
			isMentor = true
			break
		}
	}

	if !isMentor {
		return nil, errors.New("只有导师可以查看聊天记录")
	}

	// 验证学生是否存在
	_, err = s.userRepo.GetByID(ctx, studentID)
	if err != nil {
		return nil, errors.New("学生不存在")
	}

	// 获取消息列表
	messages, total, err := s.messageRepo.GetMessages(ctx, mentorID, studentID, page, pageSize)
	if err != nil {
		return nil, err
	}

	// 构建分页响应
	totalPages := (int(total) + pageSize - 1) / pageSize
	pagination := &model.PaginationResponse{
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}

	return &model.MessageListResponse{
		Messages:   messages,
		Pagination: pagination,
	}, nil
}

// UpdateStudentProgress 更新学生学习进度
func (s *studentService) UpdateStudentProgress(ctx context.Context, mentorID, studentID, courseID string, req *model.StudentProgressRequest) (*model.StudentProgressResponse, error) {
	// 验证导师身份
	identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, mentorID)
	if err != nil {
		return nil, err
	}

	isMentor := false
	for _, identity := range identities {
		if identity.IdentityType == "master" && identity.Status == "active" {
			isMentor = true
			break
		}
	}

	if !isMentor {
		return nil, errors.New("只有导师可以更新学生进度")
	}

	// 验证学生是否存在
	_, err = s.userRepo.GetByID(ctx, studentID)
	if err != nil {
		return nil, errors.New("学生不存在")
	}

	// 更新学习进度
	err = s.studentRepo.UpdateStudentProgress(ctx, studentID, courseID, req.ProgressPercentage, req.Notes)
	if err != nil {
		return nil, err
	}

	return &model.StudentProgressResponse{
		StudentID:          studentID,
		CourseID:           courseID,
		ProgressPercentage: req.ProgressPercentage,
	}, nil
}

// GradeAssignment 评价学生作业
func (s *studentService) GradeAssignment(ctx context.Context, mentorID, studentID, assignmentID string, req *model.GradeAssignmentRequest) (*model.GradeAssignmentResponse, error) {
	// 验证导师身份
	identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, mentorID)
	if err != nil {
		return nil, err
	}

	isMentor := false
	for _, identity := range identities {
		if identity.IdentityType == "master" && identity.Status == "active" {
			isMentor = true
			break
		}
	}

	if !isMentor {
		return nil, errors.New("只有导师可以评价作业")
	}

	// 验证学生是否存在
	_, err = s.userRepo.GetByID(ctx, studentID)
	if err != nil {
		return nil, errors.New("学生不存在")
	}

	// 评价作业
	err = s.studentRepo.GradeAssignment(ctx, assignmentID, req.Score, req.Feedback, req.Comments)
	if err != nil {
		return nil, err
	}

	return &model.GradeAssignmentResponse{
		AssignmentID: assignmentID,
		Score:        req.Score,
	}, nil
}

// GetStudentReport 获取学生学习报告
func (s *studentService) GetStudentReport(ctx context.Context, mentorID, studentID string, req *model.StudentReportRequest) (*model.StudentReportResponse, error) {
	// 验证导师身份
	identities, err := s.identityRepo.GetIdentitiesWithProfile(ctx, mentorID)
	if err != nil {
		return nil, err
	}

	isMentor := false
	for _, identity := range identities {
		if identity.IdentityType == "master" && identity.Status == "active" {
			isMentor = true
			break
		}
	}

	if !isMentor {
		return nil, errors.New("只有导师可以查看学习报告")
	}

	// 验证学生是否存在
	_, err = s.userRepo.GetByID(ctx, studentID)
	if err != nil {
		return nil, errors.New("学生不存在")
	}

	// 获取学习报告
	report, err := s.studentRepo.GetStudentReport(ctx, studentID, req.Period)
	if err != nil {
		return nil, err
	}

	return &model.StudentReportResponse{
		Report: report,
	}, nil
}
