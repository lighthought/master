package model

import "time"

// IncomeStats 收入统计
type IncomeStats struct {
	TotalIncome           float64            `json:"total_income"`
	TotalTransactions     int64              `json:"total_transactions"`
	AveragePerTransaction float64            `json:"average_per_transaction"`
	IncomeBySource        *IncomeBySource    `json:"income_by_source"`
	IncomeByPeriod        *IncomeByPeriod    `json:"income_by_period"`
	GrowthRate            float64            `json:"growth_rate"`
	TopCourses            []*TopCourseIncome `json:"top_courses"`
}

// IncomeBySource 按来源分类的收入
type IncomeBySource struct {
	CourseEnrollments float64 `json:"course_enrollments"`
	Appointments      float64 `json:"appointments"`
}

// IncomeByPeriod 按周期分类的收入
type IncomeByPeriod struct {
	CurrentMonth   float64 `json:"current_month"`
	PreviousMonth  float64 `json:"previous_month"`
	CurrentQuarter float64 `json:"current_quarter"`
	CurrentYear    float64 `json:"current_year"`
}

// TopCourseIncome 热门课程收入
type TopCourseIncome struct {
	CourseID        string  `json:"course_id"`
	Title           string  `json:"title"`
	Income          float64 `json:"income"`
	EnrollmentCount int64   `json:"enrollment_count"`
}

// IncomeStatsResponse 收入统计响应
type IncomeStatsResponse struct {
	Stats *IncomeStats `json:"stats"`
}

// IncomeTransaction 收入交易记录
type IncomeTransaction struct {
	ID          string     `json:"id"`
	Type        string     `json:"type"`
	Amount      float64    `json:"amount"`
	Status      string     `json:"status"`
	Description string     `json:"description"`
	StudentName string     `json:"student_name"`
	CourseTitle string     `json:"course_title"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
	PlatformFee float64    `json:"platform_fee"`
	NetIncome   float64    `json:"net_income"`
}

// IncomeTransactionsResponse 收入明细响应
type IncomeTransactionsResponse struct {
	Transactions []*IncomeTransaction `json:"transactions"`
	Pagination   *PaginationResponse  `json:"pagination"`
}

// IncomeTrend 收入趋势
type IncomeTrend struct {
	Date              time.Time `json:"date"`
	Income            float64   `json:"income"`
	Transactions      int       `json:"transactions"`
	CourseEnrollments int       `json:"course_enrollments"`
	Appointments      int       `json:"appointments"`
}

// IncomeTrendsResponse 收入趋势响应
type IncomeTrendsResponse struct {
	Trends []*IncomeTrend `json:"trends"`
}

// IncomeExportResponse 导出收入报告响应
type IncomeExportResponse struct {
	DownloadURL string    `json:"download_url"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// Withdrawal 提现记录
type Withdrawal struct {
	ID          string     `json:"id"`
	Amount      float64    `json:"amount"`
	Status      string     `json:"status"`
	BankAccount string     `json:"bank_account"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
	Fee         float64    `json:"fee"`
	NetAmount   float64    `json:"net_amount"`
}

// WithdrawalsResponse 提现记录响应
type WithdrawalsResponse struct {
	Withdrawals []*Withdrawal       `json:"withdrawals"`
	Pagination  *PaginationResponse `json:"pagination"`
}

// CreateWithdrawalResponse 申请提现响应
type CreateWithdrawalResponse struct {
	WithdrawalID            string    `json:"withdrawal_id"`
	EstimatedCompletionTime time.Time `json:"estimated_completion_time"`
}

// AvailableIncome 可提现金额信息
type AvailableIncome struct {
	AvailableAmount float64 `json:"available_amount"`
	PendingAmount   float64 `json:"pending_amount"`
	TotalEarned     float64 `json:"total_earned"`
	TotalWithdrawn  float64 `json:"total_withdrawn"`
	MinWithdrawal   float64 `json:"min_withdrawal"`
	MaxWithdrawal   float64 `json:"max_withdrawal"`
}

// AvailableIncomeResponse 可提现金额响应
type AvailableIncomeResponse struct {
	AvailableAmount float64 `json:"available_amount"`
	PendingAmount   float64 `json:"pending_amount"`
	TotalEarned     float64 `json:"total_earned"`
	TotalWithdrawn  float64 `json:"total_withdrawn"`
	MinWithdrawal   float64 `json:"min_withdrawal"`
	MaxWithdrawal   float64 `json:"max_withdrawal"`
}
