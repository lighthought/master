package handlers

import (
	"net/http"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// LearningHandler 学习记录处理器
type LearningHandler struct {
	learningService service.LearningService
}

// NewLearningHandler 创建学习记录处理器
func NewLearningHandler(learningService service.LearningService) *LearningHandler {
	return &LearningHandler{
		learningService: learningService,
	}
}

// GetLearningRecords 获取学习记录列表
// @Summary 获取学习记录列表
// @Description 获取用户的学习记录列表，支持分页和筛选
// @Tags 学习记录管理
// @Accept json
// @Produce json
// @Param course_id query string false "课程ID"
// @Param status query string false "学习状态" Enums(learning, completed, paused)
// @Param start_date query string false "开始日期" format(date)
// @Param end_date query string false "结束日期" format(date)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.LearningRecordListResponse}
// @Router /learning-records [get]
func (h *LearningHandler) GetLearningRecords(c *gin.Context) {
	userID := c.GetString("user_id") // 从JWT中获取

	var req model.LearningRecordListRequest
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

	resp, err := h.learningService.GetLearningRecords(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取学习记录列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// GetLearningRecordByID 获取学习记录详情
// @Summary 获取学习记录详情
// @Description 根据学习记录ID获取详细信息
// @Tags 学习记录管理
// @Accept json
// @Produce json
// @Param record_id path string true "学习记录ID"
// @Success 200 {object} model.Response{data=model.LearningRecordDetailResponse}
// @Router /learning-records/{record_id} [get]
func (h *LearningHandler) GetLearningRecordByID(c *gin.Context) {
	recordID := c.Param("record_id")

	resp, err := h.learningService.GetLearningRecordByID(c.Request.Context(), recordID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Code:    404,
			Message: "学习记录不存在",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// UpdateLearningProgress 更新学习进度
// @Summary 更新学习进度
// @Description 更新指定学习记录的进度
// @Tags 学习记录管理
// @Accept json
// @Produce json
// @Param record_id path string true "学习记录ID"
// @Param request body model.UpdateProgressRequest true "进度更新信息"
// @Success 200 {object} model.Response{data=model.UpdateProgressResponse}
// @Router /learning-records/{record_id}/progress [put]
func (h *LearningHandler) UpdateLearningProgress(c *gin.Context) {
	userID := c.GetString("user_id")
	recordID := c.Param("record_id")

	var req model.UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.learningService.UpdateLearningProgress(c.Request.Context(), recordID, userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "学习进度更新成功",
		Data:    resp,
	})
}

// SubmitAssignment 提交作业
// @Summary 提交作业
// @Description 为指定学习记录提交作业
// @Tags 学习记录管理
// @Accept json
// @Produce json
// @Param record_id path string true "学习记录ID"
// @Param request body model.SubmitAssignmentRequest true "作业信息"
// @Success 200 {object} model.Response{data=model.SubmitAssignmentResponse}
// @Router /learning-records/{record_id}/assignments [post]
func (h *LearningHandler) SubmitAssignment(c *gin.Context) {
	userID := c.GetString("user_id")
	recordID := c.Param("record_id")

	var req model.SubmitAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.learningService.SubmitAssignment(c.Request.Context(), recordID, userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "作业提交成功",
		Data:    resp,
	})
}

// GetLearningStats 获取学习统计
// @Summary 获取学习统计
// @Description 获取用户的学习统计数据
// @Tags 学习记录管理
// @Accept json
// @Produce json
// @Param period query string true "统计周期" Enums(week, month, year, all)
// @Success 200 {object} model.Response{data=model.LearningStatsResponse}
// @Router /learning-records/stats [get]
func (h *LearningHandler) GetLearningStats(c *gin.Context) {
	userID := c.GetString("user_id")

	var req model.LearningStatsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.learningService.GetLearningStats(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取学习统计失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// GetRecommendedPath 获取学习路径推荐
// @Summary 获取学习路径推荐
// @Description 获取基于用户学习进度的推荐学习路径
// @Tags 学习记录管理
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=model.LearningPathResponse}
// @Router /learning-records/recommended-path [get]
func (h *LearningHandler) GetRecommendedPath(c *gin.Context) {
	userID := c.GetString("user_id")

	resp, err := h.learningService.GetRecommendedPath(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取推荐学习路径失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}
