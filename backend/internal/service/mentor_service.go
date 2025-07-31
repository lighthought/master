package service

import (
	"context"
	"errors"
	"math"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// MentorService 大师服务接口
type MentorService interface {
	GetMentors(ctx context.Context, req *model.MentorListRequest) (*model.MentorListResponse, error)
	GetMentorDetail(ctx context.Context, mentorID string) (*model.MentorDetailResponse, error)
	SearchMentors(ctx context.Context, req *model.MentorSearchRequest) (*model.MentorSearchResponse, error)
	GetRecommendedMentors(ctx context.Context, userID string) (*model.RecommendedMentorsResponse, error)
	GetMentorReviews(ctx context.Context, mentorID string, req *model.MentorReviewsRequest) (*model.MentorReviewsResponse, error)
}

// mentorService 大师服务实现
type mentorService struct {
	mentorRepo repository.MentorRepository
}

// NewMentorService 创建大师服务实例
func NewMentorService(mentorRepo repository.MentorRepository) MentorService {
	return &mentorService{
		mentorRepo: mentorRepo,
	}
}

// GetMentors 获取大师列表
func (s *mentorService) GetMentors(ctx context.Context, req *model.MentorListRequest) (*model.MentorListResponse, error) {
	mentors, total, err := s.mentorRepo.GetMentors(ctx, req.Domain, req.MinRating, req.MaxPrice, req.IsOnline, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	mentorInfos := make([]*model.MentorInfo, len(mentors))
	for i, mentor := range mentors {
		mentorInfos[i] = s.convertToMentorInfo(mentor)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.MentorListResponse{
		Mentors: mentorInfos,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// GetMentorDetail 获取大师详情
func (s *mentorService) GetMentorDetail(ctx context.Context, mentorID string) (*model.MentorDetailResponse, error) {
	mentor, err := s.mentorRepo.GetMentorByID(ctx, mentorID)
	if err != nil {
		return nil, errors.New("大师不存在")
	}

	// 获取大师课程
	courses, err := s.mentorRepo.GetMentorCourses(ctx, mentorID)
	if err != nil {
		return nil, err
	}

	// 获取大师评价（前5条）
	reviews, _, err := s.mentorRepo.GetMentorReviews(ctx, mentorID, 1, 5)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	mentorDetail := s.convertToMentorDetail(mentor, courses, reviews)

	return &model.MentorDetailResponse{
		Mentor: mentorDetail,
	}, nil
}

// SearchMentors 搜索大师
func (s *mentorService) SearchMentors(ctx context.Context, req *model.MentorSearchRequest) (*model.MentorSearchResponse, error) {
	mentors, total, err := s.mentorRepo.SearchMentors(ctx, req.Query, req.Domain, req.MinRating, req.MaxPrice, req.IsOnline, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	mentorInfos := make([]*model.MentorInfo, len(mentors))
	for i, mentor := range mentors {
		mentorInfos[i] = s.convertToMentorInfo(mentor)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.MentorSearchResponse{
		Mentors: mentorInfos,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// GetRecommendedMentors 获取推荐大师
func (s *mentorService) GetRecommendedMentors(ctx context.Context, userID string) (*model.RecommendedMentorsResponse, error) {
	mentors, err := s.mentorRepo.GetRecommendedMentors(ctx, userID, 10)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	recommendedMentors := make([]*model.RecommendedMentor, len(mentors))
	for i, mentor := range mentors {
		recommendedMentors[i] = s.convertToRecommendedMentor(mentor)
	}

	return &model.RecommendedMentorsResponse{
		Mentors: recommendedMentors,
	}, nil
}

// GetMentorReviews 获取大师评价
func (s *mentorService) GetMentorReviews(ctx context.Context, mentorID string, req *model.MentorReviewsRequest) (*model.MentorReviewsResponse, error) {
	reviews, total, err := s.mentorRepo.GetMentorReviews(ctx, mentorID, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	reviewInfos := make([]*model.MentorReview, len(reviews))
	for i, review := range reviews {
		reviewInfos[i] = s.convertToMentorReview(review)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.MentorReviewsResponse{
		Reviews: reviewInfos,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// convertToMentorInfo 转换为大师信息
func (s *mentorService) convertToMentorInfo(mentor *model.Mentor) *model.MentorInfo {
	mentorInfo := &model.MentorInfo{
		ID:           mentor.ID,
		IdentityID:   mentor.IdentityID,
		Rating:       mentor.Rating,
		StudentCount: mentor.StudentCount,
		HourlyRate:   mentor.HourlyRate,
		IsOnline:     mentor.IsOnline,
	}

	if mentor.Identity != nil {
		mentorInfo.Domain = mentor.Identity.Domain
	}

	if mentor.Profile != nil {
		mentorInfo.Name = mentor.Profile.Name
		mentorInfo.Avatar = mentor.Profile.Avatar
		mentorInfo.Bio = mentor.Profile.Bio
		mentorInfo.Skills = mentor.Profile.Skills
	}

	return mentorInfo
}

// convertToMentorDetail 转换为大师详情
func (s *mentorService) convertToMentorDetail(mentor *model.Mentor, courses []*model.Course, reviews []*model.MentorReviewModel) *model.MentorDetail {
	mentorDetail := &model.MentorDetail{
		ID:              mentor.ID,
		IdentityID:      mentor.IdentityID,
		Rating:          mentor.Rating,
		StudentCount:    mentor.StudentCount,
		HourlyRate:      mentor.HourlyRate,
		IsOnline:        mentor.IsOnline,
		ExperienceYears: mentor.ExperienceYears,
	}

	if mentor.Identity != nil {
		mentorDetail.Domain = mentor.Identity.Domain
	}

	if mentor.Profile != nil {
		mentorDetail.Name = mentor.Profile.Name
		mentorDetail.Avatar = mentor.Profile.Avatar
		mentorDetail.Bio = mentor.Profile.Bio
		mentorDetail.Skills = mentor.Profile.Skills
	}

	// 转换课程信息
	mentorDetail.Courses = make([]*model.MentorCourse, len(courses))
	for i, course := range courses {
		mentorDetail.Courses[i] = &model.MentorCourse{
			ID:           course.ID,
			Title:        course.Title,
			Price:        course.Price,
			StudentCount: course.ReviewCount, // 使用评价数量作为学生数量
		}
	}

	// 转换评价信息
	mentorDetail.Reviews = make([]*model.MentorReview, len(reviews))
	for i, review := range reviews {
		mentorDetail.Reviews[i] = s.convertToMentorReview(review)
	}

	return mentorDetail
}

// convertToRecommendedMentor 转换为推荐大师
func (s *mentorService) convertToRecommendedMentor(mentor *model.Mentor) *model.RecommendedMentor {
	recommendedMentor := &model.RecommendedMentor{
		ID:                   mentor.ID,
		IdentityID:           mentor.IdentityID,
		Rating:               mentor.Rating,
		StudentCount:         mentor.StudentCount,
		HourlyRate:           mentor.HourlyRate,
		IsOnline:             mentor.IsOnline,
		RecommendationReason: "基于评分和学生数量推荐",
	}

	if mentor.Identity != nil {
		recommendedMentor.Domain = mentor.Identity.Domain
	}

	if mentor.Profile != nil {
		recommendedMentor.Name = mentor.Profile.Name
		recommendedMentor.Avatar = mentor.Profile.Avatar
		recommendedMentor.Bio = mentor.Profile.Bio
		recommendedMentor.Skills = mentor.Profile.Skills
	}

	return recommendedMentor
}

// convertToMentorReview 转换为大师评价
func (s *mentorService) convertToMentorReview(review *model.MentorReviewModel) *model.MentorReview {
	mentorReview := &model.MentorReview{
		ID:        review.ID,
		Rating:    review.Rating,
		Content:   review.Content,
		CreatedAt: review.CreatedAt,
	}

	if review.Reviewer != nil && !review.IsAnonymous {
		mentorReview.ReviewerName = review.Reviewer.Email // 这里应该使用用户档案中的姓名
		// mentorReview.ReviewerAvatar = review.Reviewer.Avatar
	} else {
		mentorReview.ReviewerName = "匿名用户"
	}

	return mentorReview
}
