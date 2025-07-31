package handlers

import (
	"net/http"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// AppointmentHandler 预约处理器
type AppointmentHandler struct {
	appointmentService service.AppointmentService
}

// NewAppointmentHandler 创建预约处理器
func NewAppointmentHandler(appointmentService service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentService: appointmentService,
	}
}

// CreateAppointment 创建预约
// @Summary 创建预约
// @Description 学生创建预约大师
// @Tags 预约管理
// @Accept json
// @Produce json
// @Param appointment body model.CreateAppointmentRequest true "预约信息"
// @Success 200 {object} model.Response{data=model.CreateAppointmentResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Router /appointments [post]
func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var req model.CreateAppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从JWT中获取学生ID，这里暂时使用模拟值
	studentID := "USER_00000000001" // TODO: 从JWT中获取

	response, err := h.appointmentService.CreateAppointment(c.Request.Context(), studentID, &req)
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
		Message:   "预约创建成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetAppointments 获取预约列表
// @Summary 获取预约列表
// @Description 获取预约列表，支持分页和筛选
// @Tags 预约管理
// @Accept json
// @Produce json
// @Param status query string false "预约状态"
// @Param type query string false "预约类型" Enums(student, mentor)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.AppointmentListResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /appointments [get]
func (h *AppointmentHandler) GetAppointments(c *gin.Context) {
	var req model.AppointmentListRequest
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

	response, err := h.appointmentService.GetAppointments(c.Request.Context(), userID, &req)
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

// GetAppointmentDetail 获取预约详情
// @Summary 获取预约详情
// @Description 获取指定预约的详细信息
// @Tags 预约管理
// @Accept json
// @Produce json
// @Param appointment_id path string true "预约ID"
// @Success 200 {object} model.Response{data=model.AppointmentDetailResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /appointments/{appointment_id} [get]
func (h *AppointmentHandler) GetAppointmentDetail(c *gin.Context) {
	appointmentID := c.Param("appointment_id")
	if appointmentID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "预约ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.appointmentService.GetAppointmentDetail(c.Request.Context(), appointmentID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "预约不存在" {
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

// UpdateAppointmentStatus 更新预约状态
// @Summary 更新预约状态
// @Description 更新预约状态（确认、完成等）
// @Tags 预约管理
// @Accept json
// @Produce json
// @Param appointment_id path string true "预约ID"
// @Param status body model.UpdateAppointmentStatusRequest true "状态信息"
// @Success 200 {object} model.Response{data=model.UpdateAppointmentStatusResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /appointments/{appointment_id}/status [put]
func (h *AppointmentHandler) UpdateAppointmentStatus(c *gin.Context) {
	appointmentID := c.Param("appointment_id")
	if appointmentID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "预约ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	var req model.UpdateAppointmentStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.appointmentService.UpdateAppointmentStatus(c.Request.Context(), appointmentID, &req)
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
		Message:   "预约状态更新成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// CancelAppointment 取消预约
// @Summary 取消预约
// @Description 取消指定预约
// @Tags 预约管理
// @Accept json
// @Produce json
// @Param appointment_id path string true "预约ID"
// @Success 200 {object} model.Response{data=model.CancelAppointmentResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /appointments/{appointment_id} [delete]
func (h *AppointmentHandler) CancelAppointment(c *gin.Context) {
	appointmentID := c.Param("appointment_id")
	if appointmentID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "预约ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.appointmentService.CancelAppointment(c.Request.Context(), appointmentID)
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
		Message:   "预约取消成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetMentorAppointmentStats 获取大师预约统计
// @Summary 获取大师预约统计
// @Description 获取大师的预约统计信息
// @Tags 预约管理
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=model.MentorAppointmentStatsResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /appointments/mentor-stats [get]
func (h *AppointmentHandler) GetMentorAppointmentStats(c *gin.Context) {
	// 从JWT中获取大师ID，这里暂时使用模拟值
	mentorID := "MENTOR_00000000001" // TODO: 从JWT中获取

	response, err := h.appointmentService.GetMentorAppointmentStats(c.Request.Context(), mentorID)
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
