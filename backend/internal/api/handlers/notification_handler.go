package handlers

import (
	"net/http"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// NotificationHandler 通知处理器
type NotificationHandler struct {
	notificationService service.NotificationService
}

// NewNotificationHandler 创建通知处理器
func NewNotificationHandler(notificationService service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		notificationService: notificationService,
	}
}

// GetNotifications 获取通知列表
// @Summary 获取通知列表
// @Description 获取用户的通知列表，支持分页和筛选
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param type query string false "通知类型"
// @Param status query string false "通知状态" Enums(unread, read, all)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.NotificationListResponse}
// @Router /notifications [get]
func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	userID := c.GetString("user_id") // 从JWT中获取

	var req model.NotificationListRequest
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

	resp, err := h.notificationService.GetNotifications(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取通知列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// MarkNotificationRead 标记通知为已读
// @Summary 标记通知为已读
// @Description 标记指定通知为已读状态
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param notification_id path string true "通知ID"
// @Success 200 {object} model.Response{data=model.MarkNotificationReadResponse}
// @Router /notifications/{notification_id}/read [put]
func (h *NotificationHandler) MarkNotificationRead(c *gin.Context) {
	userID := c.GetString("user_id")
	notificationID := c.Param("notification_id")

	resp, err := h.notificationService.MarkNotificationRead(c.Request.Context(), userID, notificationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "通知已标记为已读",
		Data:    resp,
	})
}

// BatchMarkNotificationsRead 批量标记通知为已读
// @Summary 批量标记通知为已读
// @Description 批量标记通知为已读状态
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param request body model.MarkNotificationReadRequest true "批量标记请求"
// @Success 200 {object} model.Response{data=model.BatchMarkReadResponse}
// @Router /notifications/read [put]
func (h *NotificationHandler) BatchMarkNotificationsRead(c *gin.Context) {
	userID := c.GetString("user_id")

	var req model.MarkNotificationReadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.notificationService.BatchMarkNotificationsRead(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "通知已批量标记为已读",
		Data:    resp,
	})
}

// DeleteNotification 删除通知
// @Summary 删除通知
// @Description 删除指定通知
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param notification_id path string true "通知ID"
// @Success 200 {object} model.Response{data=model.DeleteNotificationResponse}
// @Router /notifications/{notification_id} [delete]
func (h *NotificationHandler) DeleteNotification(c *gin.Context) {
	userID := c.GetString("user_id")
	notificationID := c.Param("notification_id")

	resp, err := h.notificationService.DeleteNotification(c.Request.Context(), userID, notificationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "通知删除成功",
		Data:    resp,
	})
}

// BatchDeleteNotifications 批量删除通知
// @Summary 批量删除通知
// @Description 批量删除通知
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param request body model.DeleteNotificationRequest true "批量删除请求"
// @Success 200 {object} model.Response{data=model.BatchDeleteResponse}
// @Router /notifications [delete]
func (h *NotificationHandler) BatchDeleteNotifications(c *gin.Context) {
	userID := c.GetString("user_id")

	var req model.DeleteNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.notificationService.BatchDeleteNotifications(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "通知批量删除成功",
		Data:    resp,
	})
}

// GetUnreadCount 获取未读通知数量
// @Summary 获取未读通知数量
// @Description 获取用户的未读通知数量和分类统计
// @Tags 通知管理
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=model.UnreadCountResponse}
// @Router /notifications/unread-count [get]
func (h *NotificationHandler) GetUnreadCount(c *gin.Context) {
	userID := c.GetString("user_id")

	resp, err := h.notificationService.GetUnreadCount(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取未读通知数量失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// GetNotificationSettings 获取通知设置
// @Summary 获取通知设置
// @Description 获取用户的通知设置
// @Tags 通知管理
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=model.NotificationSettingsResponse}
// @Router /notifications/settings [get]
func (h *NotificationHandler) GetNotificationSettings(c *gin.Context) {
	userID := c.GetString("user_id")

	resp, err := h.notificationService.GetNotificationSettings(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    500,
			Message: "获取通知设置失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "success",
		Data:    resp,
	})
}

// UpdateNotificationSettings 更新通知设置
// @Summary 更新通知设置
// @Description 更新用户的通知设置
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param request body model.NotificationSettingsRequest true "通知设置"
// @Success 200 {object} model.Response{data=model.UpdateSettingsResponse}
// @Router /notifications/settings [put]
func (h *NotificationHandler) UpdateNotificationSettings(c *gin.Context) {
	userID := c.GetString("user_id")

	var req model.NotificationSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.notificationService.UpdateNotificationSettings(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "通知设置更新成功",
		Data:    resp,
	})
}

// SendNotification 发送通知
// @Summary 发送通知
// @Description 系统发送通知给指定用户
// @Tags 通知管理
// @Accept json
// @Produce json
// @Param request body model.SendNotificationRequest true "发送通知请求"
// @Success 200 {object} model.Response{data=model.SendNotificationResponse}
// @Router /notifications/send [post]
func (h *NotificationHandler) SendNotification(c *gin.Context) {
	var req model.SendNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: "请求参数错误",
		})
		return
	}

	resp, err := h.notificationService.SendNotification(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "通知发送成功",
		Data:    resp,
	})
}
