package handlers

import (
	"net/http"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// CourseHandler 课程处理器
type CourseHandler struct {
	courseService service.CourseService
}

// NewCourseHandler 创建课程处理器
func NewCourseHandler(courseService service.CourseService) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
	}
}

// GetCourses 获取课程列表
// @Summary 获取课程列表
// @Description 获取课程列表，支持分页和筛选
// @Tags 课程管理
// @Accept json
// @Produce json
// @Param domain query string false "专业领域"
// @Param difficulty query string false "难度级别"
// @Param min_price query number false "最低价格"
// @Param max_price query number false "最高价格"
// @Param sort_by query string false "排序方式" Enums(rating, price, created_at)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.CourseListResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /courses [get]
func (h *CourseHandler) GetCourses(c *gin.Context) {
	var req model.CourseListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	response, err := h.courseService.GetCourses(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:      500,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "success",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetCourseDetail 获取课程详情
// @Summary 获取课程详情
// @Description 获取指定课程的详细信息
// @Tags 课程管理
// @Accept json
// @Produce json
// @Param course_id path string true "课程ID"
// @Success 200 {object} model.Response{data=model.CourseDetailResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /courses/{course_id} [get]
func (h *CourseHandler) GetCourseDetail(c *gin.Context) {
	courseID := c.Param("course_id")
	if courseID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "课程ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.courseService.GetCourseDetail(c.Request.Context(), courseID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "课程不存在" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, model.Response{
			Code:      statusCode,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "success",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// CreateCourse 创建课程
// @Summary 创建课程
// @Description 大师创建新课程
// @Tags 课程管理
// @Accept json
// @Produce json
// @Param course body model.CreateCourseRequest true "课程信息"
// @Success 200 {object} model.Response{data=model.CreateCourseResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /courses [post]
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var req model.CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从JWT中获取大师ID，这里暂时使用模拟值
	mentorID := "MENTOR_00000000001" // TODO: 从JWT中获取

	response, err := h.courseService.CreateCourse(c.Request.Context(), mentorID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:      500,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "课程创建成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// EnrollCourse 报名课程
// @Summary 报名课程
// @Description 用户报名指定课程
// @Tags 课程管理
// @Accept json
// @Produce json
// @Param course_id path string true "课程ID"
// @Param enrollment body model.EnrollCourseRequest true "报名信息"
// @Success 200 {object} model.Response{data=model.EnrollCourseResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 409 {object} model.ErrorResponse
// @Router /courses/{course_id}/enroll [post]
func (h *CourseHandler) EnrollCourse(c *gin.Context) {
	courseID := c.Param("course_id")
	if courseID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "课程ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	var req model.EnrollCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从JWT中获取用户ID，这里暂时使用模拟值
	userID := "USER_00000000001" // TODO: 从JWT中获取

	response, err := h.courseService.EnrollCourse(c.Request.Context(), userID, courseID, &req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "已报名该课程" {
			statusCode = http.StatusConflict
		}
		c.JSON(statusCode, model.Response{
			Code:      statusCode,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "报名成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetCourseProgress 获取课程进度
// @Summary 获取课程进度
// @Description 获取用户在指定课程中的学习进度
// @Tags 课程管理
// @Accept json
// @Produce json
// @Param course_id path string true "课程ID"
// @Success 200 {object} model.Response{data=model.CourseProgressResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /courses/{course_id}/progress [get]
func (h *CourseHandler) GetCourseProgress(c *gin.Context) {
	courseID := c.Param("course_id")
	if courseID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "课程ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从JWT中获取用户ID，这里暂时使用模拟值
	userID := "USER_00000000001" // TODO: 从JWT中获取

	response, err := h.courseService.GetCourseProgress(c.Request.Context(), userID, courseID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "未找到学习记录" {
			statusCode = http.StatusNotFound
		}
		c.JSON(statusCode, model.Response{
			Code:      statusCode,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "success",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// SearchCourses 搜索课程
// @Summary 搜索课程
// @Description 根据关键词搜索课程
// @Tags 课程管理
// @Accept json
// @Produce json
// @Param q query string false "搜索关键词"
// @Param domain query string false "专业领域"
// @Param difficulty query string false "难度级别"
// @Param min_price query number false "最低价格"
// @Param max_price query number false "最高价格"
// @Param sort_by query string false "排序方式" Enums(rating, price, created_at)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.CourseSearchResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /courses/search [get]
func (h *CourseHandler) SearchCourses(c *gin.Context) {
	var req model.CourseSearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	response, err := h.courseService.SearchCourses(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:      500,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "success",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetRecommendedCourses 获取推荐课程
// @Summary 获取推荐课程
// @Description 获取推荐课程列表
// @Tags 课程管理
// @Accept json
// @Produce json
// @Param user_id query string false "用户ID（用于个性化推荐）"
// @Success 200 {object} model.Response{data=model.RecommendedCoursesResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /courses/recommended [get]
func (h *CourseHandler) GetRecommendedCourses(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		// 从JWT中获取用户ID，这里暂时使用模拟值
		userID = "USER_00000000001" // TODO: 从JWT中获取
	}

	response, err := h.courseService.GetRecommendedCourses(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:      500,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "success",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetEnrolledCourses 获取已报名课程
// @Summary 获取已报名课程
// @Description 获取用户已报名的课程列表
// @Tags 课程管理
// @Accept json
// @Produce json
// @Param status query string false "课程状态" Enums(learning, completed, paused)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.EnrolledCoursesResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /courses/enrolled [get]
func (h *CourseHandler) GetEnrolledCourses(c *gin.Context) {
	var req model.EnrolledCoursesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	// 从JWT中获取用户ID，这里暂时使用模拟值
	userID := "USER_00000000001" // TODO: 从JWT中获取

	response, err := h.courseService.GetEnrolledCourses(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:      500,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "success",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
