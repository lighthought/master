package service

import (
	"context"
	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// SearchService 搜索服务接口
type SearchService interface {
	GlobalSearch(ctx context.Context, req *model.SearchRequest) (*model.SearchResult, error)
}

type searchService struct {
	searchRepo repository.SearchRepository
}

func NewSearchService(searchRepo repository.SearchRepository) SearchService {
	return &searchService{
		searchRepo: searchRepo,
	}
}

func (s *searchService) GlobalSearch(ctx context.Context, req *model.SearchRequest) (*model.SearchResult, error) {
	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	result := &model.SearchResult{}

	// 根据搜索类型执行不同的搜索
	switch req.Type {
	case "mentors":
		mentors, total, err := s.searchRepo.SearchMentors(ctx, req.Query, req.Domain, req.Page, req.PageSize)
		if err != nil {
			return nil, err
		}
		result.Mentors = mentors
		result.TotalResults = total

	case "courses":
		courses, total, err := s.searchRepo.SearchCourses(ctx, req.Query, req.Domain, req.Page, req.PageSize)
		if err != nil {
			return nil, err
		}
		result.Courses = courses
		result.TotalResults = total

	case "posts":
		posts, total, err := s.searchRepo.SearchPosts(ctx, req.Query, req.Domain, req.Page, req.PageSize)
		if err != nil {
			return nil, err
		}
		result.Posts = posts
		result.TotalResults = total

	default:
		// 全局搜索，搜索所有类型
		var totalMentors, totalCourses, totalPosts int64

		// 搜索导师
		mentors, mentorTotal, err := s.searchRepo.SearchMentors(ctx, req.Query, req.Domain, req.Page, req.PageSize)
		if err != nil {
			return nil, err
		}
		result.Mentors = mentors
		totalMentors = mentorTotal

		// 搜索课程
		courses, courseTotal, err := s.searchRepo.SearchCourses(ctx, req.Query, req.Domain, req.Page, req.PageSize)
		if err != nil {
			return nil, err
		}
		result.Courses = courses
		totalCourses = courseTotal

		// 搜索帖子
		posts, postTotal, err := s.searchRepo.SearchPosts(ctx, req.Query, req.Domain, req.Page, req.PageSize)
		if err != nil {
			return nil, err
		}
		result.Posts = posts
		totalPosts = postTotal

		// 计算总结果数
		result.TotalResults = totalMentors + totalCourses + totalPosts
	}

	return result, nil
}
