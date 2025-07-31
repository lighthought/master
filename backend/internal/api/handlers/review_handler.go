package handlers

import (
	"net/http"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// ReviewHandler 评价处理器
type ReviewHandler struct {
	reviewService service.ReviewService
}

// NewReviewHandler 创建评价处理器
func NewReviewHandler(reviewService service.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		reviewService: reviewService,
	}
}

// GetReviews 获取评价列表
// @Summary 获取评价列表
// @Description 获取评价列表，支持分页和筛选
// @Tags 评价管理
// @Accept json
// @Produce json
// @Param reviewed_id query string false "被评价对象ID"
// @Param review_type query string false "评价类型"
// @Param rating query int false "评分筛选"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.ReviewListResponse}
// @Router /reviews [get]
func (h *ReviewHandler) GetReviews(c *gin.Context) {
	var req model.ReviewListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
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

	resp, err := h.reviewService.GetReviews(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取评价列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// GetReviewByID 获取评价详情
// @Summary 获取评价详情
// @Description 根据评价ID获取评价详情
// @Tags 评价管理
// @Accept json
// @Produce json
// @Param review_id path string true "评价ID"
// @Success 200 {object} model.Response{data=model.ReviewDetailResponse}
// @Router /reviews/{review_id} [get]
func (h *ReviewHandler) GetReviewByID(c *gin.Context) {
	reviewID := c.Param("review_id")

	resp, err := h.reviewService.GetReviewByID(c.Request.Context(), reviewID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Code:    404,
			Message: "评价不存在",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// CreateReview 创建评价
// @Summary 创建评价
// @Description 创建新的评价
// @Tags 评价管理
// @Accept json
// @Produce json
// @Param request body model.CreateReviewRequest true "评价信息"
// @Success 200 {object} model.Response{data=model.CreateReviewResponse}
// @Router /reviews [post]
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	userID := c.GetString("user_id") // 从JWT中获取

	var req model.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.reviewService.CreateReview(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "评价创建成功",
		Data:    resp,
	})
}

// UpdateReview 更新评价
// @Summary 更新评价
// @Description 更新指定评价
// @Tags 评价管理
// @Accept json
// @Produce json
// @Param review_id path string true "评价ID"
// @Param request body model.UpdateReviewRequest true "更新信息"
// @Success 200 {object} model.Response{data=model.UpdateReviewResponse}
// @Router /reviews/{review_id} [put]
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	reviewID := c.Param("review_id")
	userID := c.GetString("user_id") // 从JWT中获取

	var req model.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.reviewService.UpdateReview(c.Request.Context(), reviewID, userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "评价更新成功",
		Data:    resp,
	})
}

// DeleteReview 删除评价
// @Summary 删除评价
// @Description 删除指定评价
// @Tags 评价管理
// @Accept json
// @Produce json
// @Param review_id path string true "评价ID"
// @Success 200 {object} model.Response{data=model.DeleteReviewResponse}
// @Router /reviews/{review_id} [delete]
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	reviewID := c.Param("review_id")
	userID := c.GetString("user_id") // 从JWT中获取

	resp, err := h.reviewService.DeleteReview(c.Request.Context(), reviewID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "评价删除成功",
		Data:    resp,
	})
}

// GetReviewStats 获取评价统计
// @Summary 获取评价统计
// @Description 获取指定对象的评价统计信息
// @Tags 评价管理
// @Accept json
// @Produce json
// @Param reviewed_id query string true "被评价对象ID"
// @Param review_type query string true "评价类型"
// @Success 200 {object} model.Response{data=model.ReviewStatsResponse}
// @Router /reviews/stats [get]
func (h *ReviewHandler) GetReviewStats(c *gin.Context) {
	var req model.ReviewStatsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.reviewService.GetReviewStats(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取评价统计失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}
