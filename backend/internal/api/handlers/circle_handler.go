package handlers

import (
	"net/http"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// CircleHandler 圈子处理器
type CircleHandler struct {
	circleService service.CircleService
}

// NewCircleHandler 创建圈子处理器
func NewCircleHandler(circleService service.CircleService) *CircleHandler {
	return &CircleHandler{
		circleService: circleService,
	}
}

// GetCircles 获取圈子列表
// @Summary 获取圈子列表
// @Description 获取圈子列表，支持分页和领域筛选
// @Tags 圈子管理
// @Accept json
// @Produce json
// @Param domain query string false "领域筛选"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.CircleListResponse}
// @Router /circles [get]
func (h *CircleHandler) GetCircles(c *gin.Context) {
	var req model.CircleListRequest
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

	resp, err := h.circleService.GetCircles(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取圈子列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// GetRecommendedCircles 获取推荐圈子
// @Summary 获取推荐圈子
// @Description 获取推荐圈子列表
// @Tags 圈子管理
// @Accept json
// @Produce json
// @Param user_id query string false "用户ID"
// @Success 200 {object} model.Response{data=model.RecommendedCirclesResponse}
// @Router /circles/recommended [get]
func (h *CircleHandler) GetRecommendedCircles(c *gin.Context) {
	userID := c.Query("user_id")

	resp, err := h.circleService.GetRecommendedCircles(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取推荐圈子失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// JoinCircle 加入圈子
// @Summary 加入圈子
// @Description 用户加入指定圈子
// @Tags 圈子管理
// @Accept json
// @Produce json
// @Param circle_id path string true "圈子ID"
// @Success 200 {object} model.Response{data=model.JoinCircleResponse}
// @Router /circles/{circle_id}/join [post]
func (h *CircleHandler) JoinCircle(c *gin.Context) {
	circleID := c.Param("circle_id")
	userID := c.GetString("user_id") // 从JWT中获取

	resp, err := h.circleService.JoinCircle(c.Request.Context(), userID, circleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "加入圈子成功",
		Data:    resp,
	})
}

// LeaveCircle 退出圈子
// @Summary 退出圈子
// @Description 用户退出指定圈子
// @Tags 圈子管理
// @Accept json
// @Produce json
// @Param circle_id path string true "圈子ID"
// @Success 200 {object} model.Response{data=model.LeaveCircleResponse}
// @Router /circles/{circle_id}/join [delete]
func (h *CircleHandler) LeaveCircle(c *gin.Context) {
	circleID := c.Param("circle_id")
	userID := c.GetString("user_id") // 从JWT中获取

	resp, err := h.circleService.LeaveCircle(c.Request.Context(), userID, circleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "退出圈子成功",
		Data:    resp,
	})
}
