package handlers

import (
	"net/http"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// StatsHandler 统计处理器
type StatsHandler struct {
	statsService service.StatsService
}

// NewStatsHandler 创建统计处理器
func NewStatsHandler(statsService service.StatsService) *StatsHandler {
	return &StatsHandler{
		statsService: statsService,
	}
}

// GetUserStats 获取用户统计
// @Summary 获取用户统计
// @Description 获取用户的学习和教学统计数据
// @Tags 统计
// @Accept json
// @Produce json
// @Param user_id path string true "用户ID"
// @Success 200 {object} model.Response{data=model.UserStatsResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /stats/user/{user_id} [get]
func (h *StatsHandler) GetUserStats(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "用户ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 获取用户统计
	result, err := h.statsService.GetUserStats(c.Request.Context(), userID)
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
		Data:      result,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
