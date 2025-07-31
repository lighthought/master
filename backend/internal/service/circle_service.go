package service

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// CircleService 圈子服务接口
type CircleService interface {
	GetCircles(ctx context.Context, req *model.CircleListRequest) (*model.CircleListResponse, error)
	GetRecommendedCircles(ctx context.Context, userID string) (*model.RecommendedCirclesResponse, error)
	JoinCircle(ctx context.Context, userID, circleID string) (*model.JoinCircleResponse, error)
	LeaveCircle(ctx context.Context, userID, circleID string) (*model.LeaveCircleResponse, error)
}

// circleService 圈子服务实现
type circleService struct {
	circleRepo repository.CircleRepository
}

// NewCircleService 创建圈子服务实例
func NewCircleService(circleRepo repository.CircleRepository) CircleService {
	return &circleService{
		circleRepo: circleRepo,
	}
}

// GetCircles 获取圈子列表
func (s *circleService) GetCircles(ctx context.Context, req *model.CircleListRequest) (*model.CircleListResponse, error) {
	circles, _, err := s.circleRepo.GetCircles(ctx, req.Domain, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	circleInfos := make([]*model.CircleInfo, len(circles))
	for i, circle := range circles {
		circleInfos[i] = s.convertToCircleInfo(circle)
	}

	return &model.CircleListResponse{
		Circles: circleInfos,
	}, nil
}

// GetRecommendedCircles 获取推荐圈子
func (s *circleService) GetRecommendedCircles(ctx context.Context, userID string) (*model.RecommendedCirclesResponse, error) {
	circles, err := s.circleRepo.GetRecommendedCircles(ctx, userID, 10)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	recommendedCircles := make([]*model.RecommendedCircle, len(circles))
	for i, circle := range circles {
		recommendedCircles[i] = s.convertToRecommendedCircle(circle)
	}

	return &model.RecommendedCirclesResponse{
		Circles: recommendedCircles,
	}, nil
}

// JoinCircle 加入圈子
func (s *circleService) JoinCircle(ctx context.Context, userID, circleID string) (*model.JoinCircleResponse, error) {
	err := s.circleRepo.JoinCircle(ctx, userID, circleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("已加入该圈子")
		}
		return nil, err
	}

	return &model.JoinCircleResponse{
		CircleID: circleID,
	}, nil
}

// LeaveCircle 退出圈子
func (s *circleService) LeaveCircle(ctx context.Context, userID, circleID string) (*model.LeaveCircleResponse, error) {
	err := s.circleRepo.LeaveCircle(ctx, userID, circleID)
	if err != nil {
		return nil, err
	}

	return &model.LeaveCircleResponse{
		CircleID: circleID,
	}, nil
}

// convertToCircleInfo 转换为圈子信息
func (s *circleService) convertToCircleInfo(circle *model.Circle) *model.CircleInfo {
	return &model.CircleInfo{
		ID:          circle.ID,
		Name:        circle.Name,
		Description: circle.Description,
		Domain:      circle.Domain,
		MemberCount: circle.MemberCount,
		IsJoined:    false, // TODO: 根据用户是否加入设置
	}
}

// convertToRecommendedCircle 转换为推荐圈子
func (s *circleService) convertToRecommendedCircle(circle *model.Circle) *model.RecommendedCircle {
	return &model.RecommendedCircle{
		ID:                   circle.ID,
		Name:                 circle.Name,
		Description:          circle.Description,
		Domain:               circle.Domain,
		MemberCount:          circle.MemberCount,
		IsJoined:             false, // TODO: 根据用户是否加入设置
		RecommendationReason: "基于圈子活跃度和成员数量推荐",
	}
}
