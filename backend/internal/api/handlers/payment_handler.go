package handlers

import (
	"net/http"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// PaymentHandler 支付处理器
type PaymentHandler struct {
	paymentService service.PaymentService
}

// NewPaymentHandler 创建支付处理器
func NewPaymentHandler(paymentService service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

// CreatePaymentOrder 创建支付订单
// @Summary 创建支付订单
// @Description 创建支付订单，返回支付URL和二维码
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param request body model.CreatePaymentOrderRequest true "支付订单信息"
// @Success 200 {object} model.Response{data=model.CreatePaymentOrderResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /payments/orders [post]
func (h *PaymentHandler) CreatePaymentOrder(c *gin.Context) {
	var req model.CreatePaymentOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.paymentService.CreatePaymentOrder(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "支付订单创建成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// QueryPaymentStatus 查询支付状态
// @Summary 查询支付状态
// @Description 查询指定订单的支付状态
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param order_id path string true "订单ID"
// @Success 200 {object} model.Response{data=model.QueryPaymentStatusResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /payments/orders/{order_id}/status [get]
func (h *PaymentHandler) QueryPaymentStatus(c *gin.Context) {
	orderID := c.Param("order_id")
	if orderID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "订单ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.paymentService.QueryPaymentStatus(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Code:      404,
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

// ListPaymentHistory 获取支付历史
// @Summary 获取支付历史
// @Description 获取支付历史记录，支持分页和筛选
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param type query string false "支付类型" Enums(course_enrollment, appointment, refund)
// @Param status query string false "支付状态" Enums(pending, completed, failed, cancelled)
// @Param start_date query string false "开始日期" format(date)
// @Param end_date query string false "结束日期" format(date)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.PaymentHistoryResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /payments/history [get]
func (h *PaymentHandler) ListPaymentHistory(c *gin.Context) {
	var req model.PaymentHistoryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.paymentService.ListPaymentHistory(c.Request.Context(), &req)
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

// CreateRefund 申请退款
// @Summary 申请退款
// @Description 申请支付退款
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param request body model.CreateRefundRequest true "退款申请信息"
// @Success 200 {object} model.Response{data=model.CreateRefundResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /payments/refunds [post]
func (h *PaymentHandler) CreateRefund(c *gin.Context) {
	var req model.CreateRefundRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.paymentService.CreateRefund(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "退款申请提交成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// QueryRefundStatus 查询退款状态
// @Summary 查询退款状态
// @Description 查询指定退款的处理状态
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param refund_id path string true "退款ID"
// @Success 200 {object} model.Response{data=model.QueryRefundStatusResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /payments/refunds/{refund_id}/status [get]
func (h *PaymentHandler) QueryRefundStatus(c *gin.Context) {
	refundID := c.Param("refund_id")
	if refundID == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "退款ID不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.paymentService.QueryRefundStatus(c.Request.Context(), refundID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.Response{
			Code:      404,
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

// ListPaymentMethods 获取支付方式列表
// @Summary 获取支付方式列表
// @Description 获取可用的支付方式列表
// @Tags 支付管理
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=model.PaymentMethodListResponse}
// @Router /payments/methods [get]
func (h *PaymentHandler) ListPaymentMethods(c *gin.Context) {
	response, err := h.paymentService.ListPaymentMethods(c.Request.Context())
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

// GetPaymentStats 获取支付统计
// @Summary 获取支付统计
// @Description 获取支付统计数据
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param period query string false "统计周期" Enums(day, week, month, quarter, year)
// @Param start_date query string false "开始日期" format(date)
// @Param end_date query string false "结束日期" format(date)
// @Success 200 {object} model.Response{data=model.PaymentStatsResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /payments/stats [get]
func (h *PaymentHandler) GetPaymentStats(c *gin.Context) {
	var req model.PaymentStatsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.paymentService.GetPaymentStats(c.Request.Context(), &req)
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

// ProcessPaymentWebhook 支付回调处理
// @Summary 支付回调处理
// @Description 处理支付网关回调
// @Tags 支付管理
// @Accept json
// @Produce json
// @Param gateway path string true "支付网关"
// @Param request body model.PaymentWebhookRequest true "回调数据"
// @Success 200 {object} model.Response{data=model.PaymentWebhookResponse}
// @Failure 400 {object} model.ErrorResponse
// @Router /payments/webhook/{gateway} [post]
func (h *PaymentHandler) ProcessPaymentWebhook(c *gin.Context) {
	gateway := c.Param("gateway")
	if gateway == "" {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "支付网关不能为空",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	var req model.PaymentWebhookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.paymentService.ProcessPaymentWebhook(c.Request.Context(), gateway, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:      0,
		Message:   "回调处理成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
