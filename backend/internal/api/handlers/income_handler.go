package handlers

import (
	"net/http"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// IncomeHandler 收入处理器
type IncomeHandler struct {
	incomeService service.IncomeService
}

// NewIncomeHandler 创建收入处理器
func NewIncomeHandler(incomeService service.IncomeService) *IncomeHandler {
	return &IncomeHandler{
		incomeService: incomeService,
	}
}

// GetIncomeStats 获取收入统计
// @Summary 获取收入统计
// @Description 获取大师的收入统计信息，包括总收入、交易数、平均收入等
// @Tags 收入管理
// @Accept json
// @Produce json
// @Param period query string false "统计周期" Enums(week, month, quarter, year, all)
// @Param start_date query string false "开始日期" format(date)
// @Param end_date query string false "结束日期" format(date)
// @Success 200 {object} model.Response{data=model.IncomeStatsResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /income/stats [get]
func (h *IncomeHandler) GetIncomeStats(c *gin.Context) {
	var req model.IncomeStatsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从上下文获取大师ID（实际应该从JWT token中获取）
	mentorID := c.GetString("mentor_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权访问",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.incomeService.GetIncomeStats(c.Request.Context(), mentorID, &req)
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

// GetIncomeTransactions 获取收入明细
// @Summary 获取收入明细
// @Description 获取大师的收入交易明细记录
// @Tags 收入管理
// @Accept json
// @Produce json
// @Param type query string false "收入类型" Enums(course_enrollment, appointment, refund)
// @Param status query string false "交易状态" Enums(completed, pending, failed)
// @Param start_date query string false "开始日期" format(date)
// @Param end_date query string false "结束日期" format(date)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.IncomeTransactionsResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /income/transactions [get]
func (h *IncomeHandler) GetIncomeTransactions(c *gin.Context) {
	var req model.IncomeTransactionsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从上下文获取大师ID
	mentorID := c.GetString("mentor_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权访问",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.incomeService.GetIncomeTransactions(c.Request.Context(), mentorID, &req)
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

// GetIncomeTrends 获取收入趋势
// @Summary 获取收入趋势
// @Description 获取大师的收入趋势数据
// @Tags 收入管理
// @Accept json
// @Produce json
// @Param period query string true "趋势周期" Enums(daily, weekly, monthly)
// @Param start_date query string true "开始日期" format(date)
// @Param end_date query string true "结束日期" format(date)
// @Success 200 {object} model.Response{data=model.IncomeTrendsResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /income/trends [get]
func (h *IncomeHandler) GetIncomeTrends(c *gin.Context) {
	var req model.IncomeTrendsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从上下文获取大师ID
	mentorID := c.GetString("mentor_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权访问",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.incomeService.GetIncomeTrends(c.Request.Context(), mentorID, &req)
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

// ExportIncomeReport 导出收入报告
// @Summary 导出收入报告
// @Description 导出大师的收入报告
// @Tags 收入管理
// @Accept json
// @Produce json
// @Param format query string true "导出格式" Enums(csv, excel, pdf)
// @Param start_date query string true "开始日期" format(date)
// @Param end_date query string true "结束日期" format(date)
// @Param type query string true "收入类型" Enums(all, course_enrollment, appointment)
// @Success 200 {object} model.Response{data=model.IncomeExportResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /income/export [get]
func (h *IncomeHandler) ExportIncomeReport(c *gin.Context) {
	var req model.IncomeExportRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从上下文获取大师ID
	mentorID := c.GetString("mentor_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权访问",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.incomeService.ExportIncomeReport(c.Request.Context(), mentorID, &req)
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
		Message:   "报告生成成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetWithdrawals 获取提现记录
// @Summary 获取提现记录
// @Description 获取大师的提现记录
// @Tags 收入管理
// @Accept json
// @Produce json
// @Param status query string false "提现状态" Enums(pending, completed, failed)
// @Param start_date query string false "开始日期" format(date)
// @Param end_date query string false "结束日期" format(date)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} model.Response{data=model.WithdrawalsResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /income/withdrawals [get]
func (h *IncomeHandler) GetWithdrawals(c *gin.Context) {
	var req model.WithdrawalsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从上下文获取大师ID
	mentorID := c.GetString("mentor_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权访问",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.incomeService.GetWithdrawals(c.Request.Context(), mentorID, &req)
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

// CreateWithdrawal 申请提现
// @Summary 申请提现
// @Description 大师申请提现
// @Tags 收入管理
// @Accept json
// @Produce json
// @Param request body model.CreateWithdrawalRequest true "提现申请信息"
// @Success 200 {object} model.Response{data=model.CreateWithdrawalResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /income/withdrawals [post]
func (h *IncomeHandler) CreateWithdrawal(c *gin.Context) {
	var req model.CreateWithdrawalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:      400,
			Message:   "请求参数错误",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	// 从上下文获取大师ID
	mentorID := c.GetString("mentor_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权访问",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.incomeService.CreateWithdrawal(c.Request.Context(), mentorID, &req)
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
		Message:   "提现申请提交成功",
		Data:      response,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

// GetAvailableIncome 获取可提现金额
// @Summary 获取可提现金额
// @Description 获取大师的可提现金额信息
// @Tags 收入管理
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=model.AvailableIncomeResponse}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /income/available [get]
func (h *IncomeHandler) GetAvailableIncome(c *gin.Context) {
	// 从上下文获取大师ID
	mentorID := c.GetString("mentor_id")
	if mentorID == "" {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:      401,
			Message:   "未授权访问",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	response, err := h.incomeService.GetAvailableIncome(c.Request.Context(), mentorID)
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
