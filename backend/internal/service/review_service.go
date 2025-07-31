package service

import (
	"context"
	"errors"
	"math"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// ReviewService 评价服务接口
type ReviewService interface {
	GetReviews(ctx context.Context, req *model.ReviewListRequest) (*model.ReviewListResponse, error)
	GetReviewByID(ctx context.Context, reviewID string) (*model.ReviewDetailResponse, error)
	CreateReview(ctx context.Context, reviewerID string, req *model.CreateReviewRequest) (*model.CreateReviewResponse, error)
	UpdateReview(ctx context.Context, reviewID, reviewerID string, req *model.UpdateReviewRequest) (*model.UpdateReviewResponse, error)
	DeleteReview(ctx context.Context, reviewID, reviewerID string) (*model.DeleteReviewResponse, error)
	GetReviewStats(ctx context.Context, req *model.ReviewStatsRequest) (*model.ReviewStatsResponse, error)
}

// reviewService 评价服务实现
type reviewService struct {
	reviewRepo repository.ReviewRepository
}

// NewReviewService 创建评价服务实例
func NewReviewService(reviewRepo repository.ReviewRepository) ReviewService {
	return &reviewService{
		reviewRepo: reviewRepo,
	}
}

// GetReviews 获取评价列表
func (s *reviewService) GetReviews(ctx context.Context, req *model.ReviewListRequest) (*model.ReviewListResponse, error) {
	reviews, total, err := s.reviewRepo.GetReviews(ctx, req.ReviewedID, req.ReviewType, req.Rating, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	reviewInfos := make([]*model.ReviewInfo, len(reviews))
	for i, review := range reviews {
		reviewInfos[i] = s.convertToReviewInfo(review)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.ReviewListResponse{
		Reviews: reviewInfos,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// GetReviewByID 根据ID获取评价详情
func (s *reviewService) GetReviewByID(ctx context.Context, reviewID string) (*model.ReviewDetailResponse, error) {
	review, err := s.reviewRepo.GetReviewByID(ctx, reviewID)
	if err != nil {
		return nil, err
	}

	return &model.ReviewDetailResponse{
		Review: s.convertToReviewDetail(review),
	}, nil
}

// CreateReview 创建评价
func (s *reviewService) CreateReview(ctx context.Context, reviewerID string, req *model.CreateReviewRequest) (*model.CreateReviewResponse, error) {
	// 检查用户是否已经评价过
	canReview, err := s.reviewRepo.CheckUserCanReview(ctx, reviewerID, req.ReviewedID, req.ReviewType)
	if err != nil {
		return nil, err
	}
	if !canReview {
		return nil, errors.New("您已经评价过该对象")
	}

	review := &model.Review{
		ReviewerID:    reviewerID,
		ReviewedID:    req.ReviewedID,
		CourseID:      &req.CourseID,
		AppointmentID: &req.AppointmentID,
		Rating:        req.Rating,
		Content:       req.Content,
		ReviewType:    req.ReviewType,
	}

	// 处理可选字段
	if req.CourseID == "" {
		review.CourseID = nil
	}
	if req.AppointmentID == "" {
		review.AppointmentID = nil
	}

	err = s.reviewRepo.CreateReview(ctx, review)
	if err != nil {
		return nil, err
	}

	return &model.CreateReviewResponse{
		ReviewID: review.ID,
	}, nil
}

// UpdateReview 更新评价
func (s *reviewService) UpdateReview(ctx context.Context, reviewID, reviewerID string, req *model.UpdateReviewRequest) (*model.UpdateReviewResponse, error) {
	// 检查评价是否存在且属于当前用户
	review, err := s.reviewRepo.GetReviewByID(ctx, reviewID)
	if err != nil {
		return nil, err
	}

	if review.ReviewerID != reviewerID {
		return nil, errors.New("只能修改自己的评价")
	}

	err = s.reviewRepo.UpdateReview(ctx, reviewID, req.Rating, req.Content)
	if err != nil {
		return nil, err
	}

	return &model.UpdateReviewResponse{
		ReviewID: reviewID,
	}, nil
}

// DeleteReview 删除评价
func (s *reviewService) DeleteReview(ctx context.Context, reviewID, reviewerID string) (*model.DeleteReviewResponse, error) {
	// 检查评价是否存在且属于当前用户
	review, err := s.reviewRepo.GetReviewByID(ctx, reviewID)
	if err != nil {
		return nil, err
	}

	if review.ReviewerID != reviewerID {
		return nil, errors.New("只能删除自己的评价")
	}

	err = s.reviewRepo.DeleteReview(ctx, reviewID)
	if err != nil {
		return nil, err
	}

	return &model.DeleteReviewResponse{
		ReviewID: reviewID,
	}, nil
}

// GetReviewStats 获取评价统计
func (s *reviewService) GetReviewStats(ctx context.Context, req *model.ReviewStatsRequest) (*model.ReviewStatsResponse, error) {
	stats, err := s.reviewRepo.GetReviewStats(ctx, req.ReviewedID, req.ReviewType)
	if err != nil {
		return nil, err
	}

	return &model.ReviewStatsResponse{
		Stats: stats,
	}, nil
}

// convertToReviewInfo 转换为评价信息
func (s *reviewService) convertToReviewInfo(review *model.Review) *model.ReviewInfo {
	reviewInfo := &model.ReviewInfo{
		ID:         review.ID,
		ReviewedID: review.ReviewedID,
		Rating:     review.Rating,
		Content:    review.Content,
		ReviewType: review.ReviewType,
		CreatedAt:  review.CreatedAt,
	}

	// 转换评价者信息
	if review.Reviewer != nil {
		reviewInfo.Reviewer = &model.UserInfo{
			ID: review.Reviewer.ID,
		}
		// 这里应该从用户档案中获取姓名和头像
		// 暂时使用邮箱作为姓名
		reviewInfo.Reviewer.Email = review.Reviewer.Email
	}

	// 转换关联信息
	if review.CourseID != nil {
		reviewInfo.CourseID = *review.CourseID
	}
	if review.AppointmentID != nil {
		reviewInfo.AppointmentID = *review.AppointmentID
	}

	return reviewInfo
}

// convertToReviewDetail 转换为评价详情
func (s *reviewService) convertToReviewDetail(review *model.Review) *model.ReviewDetail {
	reviewDetail := &model.ReviewDetail{
		ID:         review.ID,
		ReviewedID: review.ReviewedID,
		Rating:     review.Rating,
		Content:    review.Content,
		ReviewType: review.ReviewType,
		CreatedAt:  review.CreatedAt,
		UpdatedAt:  review.UpdatedAt,
	}

	// 转换评价者信息
	if review.Reviewer != nil {
		reviewDetail.Reviewer = &model.UserInfo{
			ID: review.Reviewer.ID,
		}
		// 这里应该从用户档案中获取姓名和头像
		// 暂时使用邮箱作为姓名
		reviewDetail.Reviewer.Email = review.Reviewer.Email
	}

	// 转换关联信息
	if review.CourseID != nil {
		reviewDetail.CourseID = *review.CourseID
	}
	if review.AppointmentID != nil {
		reviewDetail.AppointmentID = *review.AppointmentID
	}

	return reviewDetail
}
