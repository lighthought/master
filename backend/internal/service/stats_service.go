package service

import (
	"context"
	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// StatsService 统计服务接口
type StatsService interface {
	GetUserStats(ctx context.Context, userID string) (*model.UserStatsResponse, error)
}

type statsService struct {
	statsRepo repository.StatsRepository
}

func NewStatsService(statsRepo repository.StatsRepository) StatsService {
	return &statsService{
		statsRepo: statsRepo,
	}
}

func (s *statsService) GetUserStats(ctx context.Context, userID string) (*model.UserStatsResponse, error) {
	// 获取学习统计
	learningStats, err := s.statsRepo.GetUserLearningStats(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 获取教学统计
	teachingStats, err := s.statsRepo.GetUserTeachingStats(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.UserStatsResponse{
		LearningStats: learningStats,
		TeachingStats: teachingStats,
	}, nil
}
