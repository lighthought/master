package service

import (
	"context"
	"errors"
	"math"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// LearningService 学习记录服务接口
type LearningService interface {
	GetLearningRecords(ctx context.Context, userID string, req *model.LearningRecordListRequest) (*model.LearningRecordListResponse, error)
	GetLearningRecordByID(ctx context.Context, recordID string) (*model.LearningRecordDetailResponse, error)
	UpdateLearningProgress(ctx context.Context, recordID, userID string, req *model.UpdateProgressRequest) (*model.UpdateProgressResponse, error)
	SubmitAssignment(ctx context.Context, recordID, userID string, req *model.SubmitAssignmentRequest) (*model.SubmitAssignmentResponse, error)
	GetLearningStats(ctx context.Context, userID string, req *model.LearningStatsRequest) (*model.LearningStatsResponse, error)
	GetRecommendedPath(ctx context.Context, userID string) (*model.LearningPathResponse, error)
}

// learningService 学习记录服务实现
type learningService struct {
	learningRepo repository.LearningRepository
}

// NewLearningService 创建学习记录服务实例
func NewLearningService(learningRepo repository.LearningRepository) LearningService {
	return &learningService{
		learningRepo: learningRepo,
	}
}

// GetLearningRecords 获取学习记录列表
func (s *learningService) GetLearningRecords(ctx context.Context, userID string, req *model.LearningRecordListRequest) (*model.LearningRecordListResponse, error) {
	records, total, err := s.learningRepo.GetLearningRecords(ctx, userID, req.CourseID, req.Status, req.StartDate, req.EndDate, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	recordInfos := make([]*model.LearningRecordInfo, len(records))
	for i, record := range records {
		recordInfos[i] = s.convertToLearningRecordInfo(record)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.LearningRecordListResponse{
		Records: recordInfos,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// GetLearningRecordByID 根据ID获取学习记录详情
func (s *learningService) GetLearningRecordByID(ctx context.Context, recordID string) (*model.LearningRecordDetailResponse, error) {
	record, err := s.learningRepo.GetLearningRecordByID(ctx, recordID)
	if err != nil {
		return nil, err
	}

	// 获取学习会话
	sessions, err := s.learningRepo.GetStudySessions(ctx, recordID)
	if err != nil {
		return nil, err
	}

	// 获取作业
	assignments, err := s.learningRepo.GetAssignments(ctx, recordID)
	if err != nil {
		return nil, err
	}

	detail := s.convertToLearningRecordDetail(record, sessions, assignments)

	return &model.LearningRecordDetailResponse{
		Record: detail,
	}, nil
}

// UpdateLearningProgress 更新学习进度
func (s *learningService) UpdateLearningProgress(ctx context.Context, recordID, userID string, req *model.UpdateProgressRequest) (*model.UpdateProgressResponse, error) {
	// 检查学习记录是否存在且属于当前用户
	record, err := s.learningRepo.GetLearningRecordByID(ctx, recordID)
	if err != nil {
		return nil, err
	}

	if record.UserID != userID {
		return nil, errors.New("只能更新自己的学习记录")
	}

	// 创建学习会话
	if req.StudyTimeMinutes > 0 {
		session := &model.StudySession{
			LearningRecordID: recordID,
			StartTime:        time.Now().Add(-time.Duration(req.StudyTimeMinutes) * time.Minute),
			EndTime:          &[]time.Time{time.Now()}[0],
			DurationMinutes:  req.StudyTimeMinutes,
			Chapter:          req.CurrentChapter,
		}
		err = s.learningRepo.CreateStudySession(ctx, session)
		if err != nil {
			return nil, err
		}
	}

	// 更新学习进度
	err = s.learningRepo.UpdateLearningProgress(ctx, recordID, req.ProgressPercentage, req.CurrentChapter, req.StudyTimeMinutes)
	if err != nil {
		return nil, err
	}

	return &model.UpdateProgressResponse{
		RecordID:           recordID,
		ProgressPercentage: req.ProgressPercentage,
	}, nil
}

// SubmitAssignment 提交作业
func (s *learningService) SubmitAssignment(ctx context.Context, recordID, userID string, req *model.SubmitAssignmentRequest) (*model.SubmitAssignmentResponse, error) {
	// 检查学习记录是否存在且属于当前用户
	record, err := s.learningRepo.GetLearningRecordByID(ctx, recordID)
	if err != nil {
		return nil, err
	}

	if record.UserID != userID {
		return nil, errors.New("只能为自己的学习记录提交作业")
	}

	assignment := &model.Assignment{
		LearningRecordID: recordID,
		Title:            req.Title,
		Content:          req.Content,
		AttachmentURLs:   req.AttachmentURLs,
		Status:           "submitted",
	}

	err = s.learningRepo.CreateAssignment(ctx, assignment)
	if err != nil {
		return nil, err
	}

	return &model.SubmitAssignmentResponse{
		AssignmentID: assignment.ID,
	}, nil
}

// GetLearningStats 获取学习统计
func (s *learningService) GetLearningStats(ctx context.Context, userID string, req *model.LearningStatsRequest) (*model.LearningStatsResponse, error) {
	stats, err := s.learningRepo.GetLearningStats(ctx, userID, req.Period)
	if err != nil {
		return nil, err
	}

	return &model.LearningStatsResponse{
		Stats: stats,
	}, nil
}

// GetRecommendedPath 获取推荐学习路径
func (s *learningService) GetRecommendedPath(ctx context.Context, userID string) (*model.LearningPathResponse, error) {
	path, err := s.learningRepo.GetRecommendedPath(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.LearningPathResponse{
		Path: path,
	}, nil
}

// convertToLearningRecordInfo 转换为学习记录信息
func (s *learningService) convertToLearningRecordInfo(record *model.LearningRecord) *model.LearningRecordInfo {
	recordInfo := &model.LearningRecordInfo{
		ID:                 record.ID,
		EnrollmentDate:     record.EnrolledAt,
		LastStudyDate:      record.LastAccessedAt,
		TotalStudyTime:     record.TotalStudyTime,
		ProgressPercentage: record.ProgressPercentage,
		Status:             record.Status,
		CurrentChapter:     record.CurrentChapter,
		CompletedChapters:  record.CompletedChapters,
		CertificateIssued:  record.CertificateIssued,
	}

	// 转换课程信息
	if record.Course != nil {
		recordInfo.Course = &model.LearningCourseInfo{
			ID:            record.Course.ID,
			Title:         record.Course.Title,
			Description:   record.Course.Description,
			CoverImage:    record.Course.CoverImage,
			DurationHours: record.Course.DurationHours,
		}
	}

	// 转换导师信息
	if record.Course != nil && record.Course.Mentor != nil {
		recordInfo.Mentor = &model.LearningMentorInfo{
			ID:   record.Course.Mentor.ID,
			Name: record.Course.Mentor.Profile.Name,
		}
		if record.Course.Mentor.Profile != nil {
			recordInfo.Mentor.Avatar = record.Course.Mentor.Profile.Avatar
		}
	}

	return recordInfo
}

// convertToLearningRecordDetail 转换为学习记录详情
func (s *learningService) convertToLearningRecordDetail(record *model.LearningRecord, sessions []*model.StudySession, assignments []*model.Assignment) *model.LearningRecordDetail {
	detail := &model.LearningRecordDetail{
		ID:                 record.ID,
		EnrollmentDate:     record.EnrolledAt,
		LastStudyDate:      record.LastAccessedAt,
		TotalStudyTime:     record.TotalStudyTime,
		ProgressPercentage: record.ProgressPercentage,
		Status:             record.Status,
		CurrentChapter:     record.CurrentChapter,
		CompletedChapters:  record.CompletedChapters,
		CertificateIssued:  record.CertificateIssued,
		CertificateURL:     record.CertificateURL,
	}

	// 转换课程信息
	if record.Course != nil {
		detail.Course = &model.LearningCourseInfo{
			ID:            record.Course.ID,
			Title:         record.Course.Title,
			Description:   record.Course.Description,
			CoverImage:    record.Course.CoverImage,
			DurationHours: record.Course.DurationHours,
		}
	}

	// 转换导师信息
	if record.Course != nil && record.Course.Mentor != nil {
		detail.Mentor = &model.LearningMentorInfo{
			ID:   record.Course.Mentor.ID,
			Name: record.Course.Mentor.Profile.Name,
		}
		if record.Course.Mentor.Profile != nil {
			detail.Mentor.Avatar = record.Course.Mentor.Profile.Avatar
		}
	}

	// 转换学习会话
	detail.StudySessions = make([]*model.StudySessionInfo, len(sessions))
	for i, session := range sessions {
		detail.StudySessions[i] = &model.StudySessionInfo{
			ID:              session.ID,
			StartTime:       session.StartTime,
			EndTime:         session.EndTime,
			DurationMinutes: session.DurationMinutes,
			Chapter:         session.Chapter,
			Notes:           session.Notes,
		}
	}

	// 转换作业
	detail.Assignments = make([]*model.AssignmentInfo, len(assignments))
	for i, assignment := range assignments {
		detail.Assignments[i] = &model.AssignmentInfo{
			ID:          assignment.ID,
			Title:       assignment.Title,
			Status:      assignment.Status,
			SubmittedAt: assignment.SubmittedAt,
			ReviewedAt:  assignment.ReviewedAt,
			Score:       assignment.Score,
		}
	}

	return detail
}
