package model

import "time"

// IncomeStatsRequest 获取收入统计请求
type IncomeStatsRequest struct {
	Period    string    `json:"period" form:"period" binding:"oneof=week month quarter year all"`
	StartDate time.Time `json:"start_date" form:"start_date" time:"2006-01-02"`
	EndDate   time.Time `json:"end_date" form:"end_date" time:"2006-01-02"`
}

// IncomeTransactionsRequest 获取收入明细请求
type IncomeTransactionsRequest struct {
	PaginationRequest
	Type      string    `json:"type" form:"type" binding:"oneof=course_enrollment appointment refund"`
	Status    string    `json:"status" form:"status" binding:"oneof=completed pending failed"`
	StartDate time.Time `json:"start_date" form:"start_date" time:"2006-01-02"`
	EndDate   time.Time `json:"end_date" form:"end_date" time:"2006-01-02"`
}

// IncomeTrendsRequest 获取收入趋势请求
type IncomeTrendsRequest struct {
	Period    string    `json:"period" form:"period" binding:"required,oneof=daily weekly monthly"`
	StartDate time.Time `json:"start_date" form:"start_date" binding:"required" time:"2006-01-02"`
	EndDate   time.Time `json:"end_date" form:"end_date" binding:"required" time:"2006-01-02"`
}

// IncomeExportRequest 导出收入报告请求
type IncomeExportRequest struct {
	Format    string    `json:"format" form:"format" binding:"required,oneof=csv excel pdf"`
	StartDate time.Time `json:"start_date" form:"start_date" binding:"required" time:"2006-01-02"`
	EndDate   time.Time `json:"end_date" form:"end_date" binding:"required" time:"2006-01-02"`
	Type      string    `json:"type" form:"type" binding:"required,oneof=all course_enrollment appointment"`
}

// WithdrawalsRequest 获取提现记录请求
type WithdrawalsRequest struct {
	PaginationRequest
	Status    string    `json:"status" form:"status" binding:"oneof=pending completed failed"`
	StartDate time.Time `json:"start_date" form:"start_date" time:"2006-01-02"`
	EndDate   time.Time `json:"end_date" form:"end_date" time:"2006-01-02"`
}

// CreateWithdrawalRequest 申请提现请求
type CreateWithdrawalRequest struct {
	Amount      float64 `json:"amount" binding:"required,min=100"`
	BankAccount string  `json:"bank_account" binding:"required"`
	BankName    string  `json:"bank_name" binding:"required"`
}
