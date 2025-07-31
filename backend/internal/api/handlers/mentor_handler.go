package handlers

import (
	"net/http"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// MentorHandler 大师处理器
type MentorHandler struct {
	mentorService service.MentorService
}

// NewMentorHandler 创建大师处理器
func NewMentorHandler(mentorService service.MentorService) *MentorHandler {
	return &MentorHandler{
		mentorService: mentorService,
	}
}

// GetMentors 获取大师列表
// @Summary 获取大师列表
// @Description 获取大师列表，支持分页和筛选
// @Tags 大师管理
// @Accept json
// @Produce json
// @Param domain query string false "专业领域"
// @Param min_rating query number false "最低评分"
// @Param max_price query number false "最高价格"
// @Param is_online query boolean false "是否在线"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.MentorListResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /mentors [get]
func (h *MentorHandler) GetMentors(c *gin.Context) {
	var req model.MentorListRequest
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

	response, err := h.mentorService.GetMentors(c.Request.Context(), &req)
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

// GetMentorDetail 获取大师详情
// @Summary 获取大师详情
// @Description 获取指定大师的详细信息
// @Tags 大师管理
// @Accept json
// @Produce json
// @Param mentor_id path string true "大师ID"
// @Success 200 {object} model.Response{data=model.MentorDetailResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /mentors/{mentor_id} [get]
func (h *MentorHandler) GetMentorDetail(c *gin.Context) {
	mentorID := c.Param("mentor_id")
	if mentorID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "大师ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.mentorService.GetMentorDetail(c.Request.Context(), mentorID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "大师不存在" {
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

// SearchMentors 搜索大师
// @Summary 搜索大师
// @Description 根据关键词搜索大师
// @Tags 大师管理
// @Accept json
// @Produce json
// @Param q query string false "搜索关键词"
// @Param domain query string false "专业领域"
// @Param min_rating query number false "最低评分"
// @Param max_price query number false "最高价格"
// @Param is_online query boolean false "是否在线"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.MentorSearchResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /mentors/search [get]
func (h *MentorHandler) SearchMentors(c *gin.Context) {
	var req model.MentorSearchRequest
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

	response, err := h.mentorService.SearchMentors(c.Request.Context(), &req)
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

// GetRecommendedMentors 获取推荐大师
// @Summary 获取推荐大师
// @Description 获取推荐的大师列表
// @Tags 大师管理
// @Accept json
// @Produce json
// @Param user_id query string false "用户ID（用于个性化推荐）"
// @Success 200 {object} model.Response{data=model.RecommendedMentorsResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /mentors/recommended [get]
func (h *MentorHandler) GetRecommendedMentors(c *gin.Context) {
	userID := c.Query("user_id")

	response, err := h.mentorService.GetRecommendedMentors(c.Request.Context(), userID)
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

// GetMentorReviews 获取大师评价
// @Summary 获取大师评价
// @Description 获取指定大师的评价列表
// @Tags 大师管理
// @Accept json
// @Produce json
// @Param mentor_id path string true "大师ID"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} model.Response{data=model.MentorReviewsResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /mentors/{mentor_id}/reviews [get]
func (h *MentorHandler) GetMentorReviews(c *gin.Context) {
	mentorID := c.Param("mentor_id")
	if mentorID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "大师ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	var req model.MentorReviewsRequest
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
		req.PageSize = 10
	}

	response, err := h.mentorService.GetMentorReviews(c.Request.Context(), mentorID, &req)
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
