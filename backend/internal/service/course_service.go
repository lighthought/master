package service

import (
	"context"
	"errors"
	"math"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"

	"gorm.io/gorm"
)

// CourseService 课程服务接口
type CourseService interface {
	GetCourses(ctx context.Context, req *model.CourseListRequest) (*model.CourseListResponse, error)
	GetCourseDetail(ctx context.Context, courseID string) (*model.CourseDetailResponse, error)
	CreateCourse(ctx context.Context, mentorID string, req *model.CreateCourseRequest) (*model.CreateCourseResponse, error)
	EnrollCourse(ctx context.Context, userID, courseID string, req *model.EnrollCourseRequest) (*model.EnrollCourseResponse, error)
	GetCourseProgress(ctx context.Context, userID, courseID string) (*model.CourseProgressResponse, error)
	SearchCourses(ctx context.Context, req *model.CourseSearchRequest) (*model.CourseSearchResponse, error)
	GetRecommendedCourses(ctx context.Context, userID string) (*model.RecommendedCoursesResponse, error)
	GetEnrolledCourses(ctx context.Context, userID string, req *model.EnrolledCoursesRequest) (*model.EnrolledCoursesResponse, error)
}

// courseService 课程服务实现
type courseService struct {
	courseRepo        repository.CourseRepository
	courseContentRepo repository.CourseContentRepository
}

// NewCourseService 创建课程服务实例
func NewCourseService(courseRepo repository.CourseRepository, courseContentRepo repository.CourseContentRepository) CourseService {
	return &courseService{
		courseRepo:        courseRepo,
		courseContentRepo: courseContentRepo,
	}
}

// GetCourses 获取课程列表
func (s *courseService) GetCourses(ctx context.Context, req *model.CourseListRequest) (*model.CourseListResponse, error) {
	courses, total, err := s.courseRepo.GetCourses(ctx, req.Domain, req.Difficulty, req.MinPrice, req.MaxPrice, req.SortBy, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	courseInfos := make([]*model.CourseInfo, len(courses))
	for i, course := range courses {
		courseInfos[i] = s.convertToCourseInfo(course)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.CourseListResponse{
		Courses: courseInfos,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// GetCourseDetail 获取课程详情
func (s *courseService) GetCourseDetail(ctx context.Context, courseID string) (*model.CourseDetailResponse, error) {
	course, err := s.courseRepo.GetCourseByID(ctx, courseID)
	if err != nil {
		return nil, errors.New("课程不存在")
	}

	// 获取课程内容
	contents, err := s.courseContentRepo.GetCourseContents(ctx, courseID)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	courseDetail := s.convertToCourseDetail(course, contents)

	return &model.CourseDetailResponse{
		Course: courseDetail,
	}, nil
}

// CreateCourse 创建课程
func (s *courseService) CreateCourse(ctx context.Context, mentorID string, req *model.CreateCourseRequest) (*model.CreateCourseResponse, error) {
	// 创建课程
	course := &model.Course{
		MentorID:      mentorID,
		Title:         req.Title,
		Description:   req.Description,
		CoverImage:    req.CoverImage,
		Price:         req.Price,
		DurationHours: req.DurationHours,
		Difficulty:    req.Difficulty,
		MaxStudents:   req.MaxStudents,
		Status:        "draft", // 默认为草稿状态
	}

	err := s.courseRepo.CreateCourse(ctx, course)
	if err != nil {
		return nil, err
	}

	// 创建课程内容
	if len(req.Contents) > 0 {
		contents := make([]*model.CourseContentModel, len(req.Contents))
		for i, contentInput := range req.Contents {
			contents[i] = &model.CourseContentModel{
				CourseID:        course.ID,
				Title:           contentInput.Title,
				ContentType:     contentInput.ContentType,
				ContentURL:      contentInput.ContentURL,
				ContentText:     contentInput.ContentText,
				DurationMinutes: contentInput.DurationMinutes,
				OrderIndex:      contentInput.OrderIndex,
			}
		}

		err = s.courseContentRepo.CreateBatchCourseContents(ctx, contents)
		if err != nil {
			return nil, err
		}
	}

	return &model.CreateCourseResponse{
		CourseID: course.ID,
	}, nil
}

// EnrollCourse 报名课程
func (s *courseService) EnrollCourse(ctx context.Context, userID, courseID string, req *model.EnrollCourseRequest) (*model.EnrollCourseResponse, error) {
	err := s.courseRepo.EnrollCourse(ctx, userID, courseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("已报名该课程")
		}
		return nil, err
	}

	// 这里应该调用支付接口，目前返回模拟的支付URL
	paymentURL := "https://example.com/payment?course_id=" + courseID

	return &model.EnrollCourseResponse{
		EnrollmentID: "ENROLL_" + courseID + "_" + userID,
		CourseID:     courseID,
		Status:       "enrolled",
		PaymentURL:   paymentURL,
	}, nil
}

// GetCourseProgress 获取课程进度
func (s *courseService) GetCourseProgress(ctx context.Context, userID, courseID string) (*model.CourseProgressResponse, error) {
	progress, err := s.courseRepo.GetCourseProgress(ctx, userID, courseID)
	if err != nil {
		return nil, errors.New("未找到学习记录")
	}

	// 获取已完成的内容
	completedContents, err := s.courseRepo.GetCompletedContents(ctx, userID, courseID)
	if err != nil {
		return nil, err
	}

	courseProgress := &model.CourseProgress{
		CourseID:           progress.CourseID,
		ProgressPercentage: progress.ProgressPercentage,
		Status:             progress.Status,
		EnrolledAt:         progress.EnrolledAt,
		LastAccessedAt:     progress.LastAccessedAt,
		CompletedContents:  completedContents,
	}

	return &model.CourseProgressResponse{
		Progress: courseProgress,
	}, nil
}

// SearchCourses 搜索课程
func (s *courseService) SearchCourses(ctx context.Context, req *model.CourseSearchRequest) (*model.CourseSearchResponse, error) {
	courses, total, err := s.courseRepo.SearchCourses(ctx, req.Query, req.Domain, req.Difficulty, req.MinPrice, req.MaxPrice, req.SortBy, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	courseInfos := make([]*model.CourseInfo, len(courses))
	for i, course := range courses {
		courseInfos[i] = s.convertToCourseInfo(course)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.CourseSearchResponse{
		Courses: courseInfos,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// GetRecommendedCourses 获取推荐课程
func (s *courseService) GetRecommendedCourses(ctx context.Context, userID string) (*model.RecommendedCoursesResponse, error) {
	courses, err := s.courseRepo.GetRecommendedCourses(ctx, userID, 10)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	recommendedCourses := make([]*model.RecommendedCourse, len(courses))
	for i, course := range courses {
		recommendedCourses[i] = s.convertToRecommendedCourse(course)
	}

	return &model.RecommendedCoursesResponse{
		Courses: recommendedCourses,
	}, nil
}

// GetEnrolledCourses 获取已报名课程
func (s *courseService) GetEnrolledCourses(ctx context.Context, userID string, req *model.EnrolledCoursesRequest) (*model.EnrolledCoursesResponse, error) {
	courses, total, err := s.courseRepo.GetEnrolledCourses(ctx, userID, req.Status, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	enrolledCourses := make([]*model.EnrolledCourse, len(courses))
	for i, course := range courses {
		enrolledCourses[i] = s.convertToEnrolledCourse(course, userID)
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.EnrolledCoursesResponse{
		Courses: enrolledCourses,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// convertToCourseInfo 转换为课程信息
func (s *courseService) convertToCourseInfo(course *model.Course) *model.CourseInfo {
	courseInfo := &model.CourseInfo{
		ID:            course.ID,
		Title:         course.Title,
		Description:   course.Description,
		CoverImage:    course.CoverImage,
		Price:         course.Price,
		DurationHours: course.DurationHours,
		Difficulty:    course.Difficulty,
		StudentCount:  course.ReviewCount, // 使用评价数量作为学生数量
		Rating:        course.Rating,
	}

	if course.Mentor != nil {
		courseInfo.Mentor = &model.MentorInfo{
			ID: course.Mentor.ID,
		}
		if course.Mentor.Profile != nil {
			courseInfo.Mentor.Name = course.Mentor.Profile.Name
			courseInfo.Mentor.Avatar = course.Mentor.Profile.Avatar
		}
	}

	return courseInfo
}

// convertToCourseDetail 转换为课程详情
func (s *courseService) convertToCourseDetail(course *model.Course, contents []*model.CourseContentModel) *model.CourseDetail {
	courseDetail := &model.CourseDetail{
		ID:            course.ID,
		Title:         course.Title,
		Description:   course.Description,
		CoverImage:    course.CoverImage,
		Price:         course.Price,
		DurationHours: course.DurationHours,
		Difficulty:    course.Difficulty,
		StudentCount:  course.ReviewCount,
		Rating:        course.Rating,
	}

	if course.Mentor != nil {
		courseDetail.Mentor = &model.MentorInfo{
			ID: course.Mentor.ID,
		}
		if course.Mentor.Profile != nil {
			courseDetail.Mentor.Name = course.Mentor.Profile.Name
			courseDetail.Mentor.Avatar = course.Mentor.Profile.Avatar
		}
	}

	// 转换课程内容
	courseDetail.Contents = make([]*model.CourseContent, len(contents))
	for i, content := range contents {
		courseDetail.Contents[i] = &model.CourseContent{
			ID:              content.ID,
			Title:           content.Title,
			ContentType:     content.ContentType,
			DurationMinutes: content.DurationMinutes,
			OrderIndex:      content.OrderIndex,
		}
	}

	return courseDetail
}

// convertToRecommendedCourse 转换为推荐课程
func (s *courseService) convertToRecommendedCourse(course *model.Course) *model.RecommendedCourse {
	recommendedCourse := &model.RecommendedCourse{
		ID:                   course.ID,
		Title:                course.Title,
		Description:          course.Description,
		CoverImage:           course.CoverImage,
		Price:                course.Price,
		DurationHours:        course.DurationHours,
		Difficulty:           course.Difficulty,
		StudentCount:         course.ReviewCount,
		Rating:               course.Rating,
		RecommendationReason: "基于评分和学生数量推荐",
	}

	if course.Mentor != nil {
		recommendedCourse.Mentor = &model.MentorInfo{
			ID: course.Mentor.ID,
		}
		if course.Mentor.Profile != nil {
			recommendedCourse.Mentor.Name = course.Mentor.Profile.Name
			recommendedCourse.Mentor.Avatar = course.Mentor.Profile.Avatar
		}
	}

	return recommendedCourse
}

// convertToEnrolledCourse 转换为已报名课程
func (s *courseService) convertToEnrolledCourse(course *model.Course, userID string) *model.EnrolledCourse {
	enrolledCourse := &model.EnrolledCourse{
		ID:            course.ID,
		Title:         course.Title,
		Description:   course.Description,
		CoverImage:    course.CoverImage,
		Price:         course.Price,
		DurationHours: course.DurationHours,
		Difficulty:    course.Difficulty,
		StudentCount:  course.ReviewCount,
		Rating:        course.Rating,
		// 这里应该从学习记录中获取实际状态和进度
		EnrollmentStatus:   "learning",
		ProgressPercentage: 0.0,
		EnrolledAt:         course.CreatedAt,
		LastAccessedAt:     course.UpdatedAt,
	}

	if course.Mentor != nil {
		enrolledCourse.Mentor = &model.MentorInfo{
			ID: course.Mentor.ID,
		}
		if course.Mentor.Profile != nil {
			enrolledCourse.Mentor.Name = course.Mentor.Profile.Name
			enrolledCourse.Mentor.Avatar = course.Mentor.Profile.Avatar
		}
	}

	return enrolledCourse
}
