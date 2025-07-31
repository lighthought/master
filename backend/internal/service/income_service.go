package service

import (
	"context"
	"errors"
	"math"
	"time"

	"master-guide-backend/internal/model"
	"master-guide-backend/internal/repository"
)

// IncomeService 收入服务接口
type IncomeService interface {
	GetIncomeStats(ctx context.Context, mentorID string, req *model.IncomeStatsRequest) (*model.IncomeStatsResponse, error)
	GetIncomeTransactions(ctx context.Context, mentorID string, req *model.IncomeTransactionsRequest) (*model.IncomeTransactionsResponse, error)
	GetIncomeTrends(ctx context.Context, mentorID string, req *model.IncomeTrendsRequest) (*model.IncomeTrendsResponse, error)
	ExportIncomeReport(ctx context.Context, mentorID string, req *model.IncomeExportRequest) (*model.IncomeExportResponse, error)
	GetWithdrawals(ctx context.Context, mentorID string, req *model.WithdrawalsRequest) (*model.WithdrawalsResponse, error)
	CreateWithdrawal(ctx context.Context, mentorID string, req *model.CreateWithdrawalRequest) (*model.CreateWithdrawalResponse, error)
	GetAvailableIncome(ctx context.Context, mentorID string) (*model.AvailableIncomeResponse, error)
}

// incomeService 收入服务实现
type incomeService struct {
	incomeRepo repository.IncomeRepository
}

// NewIncomeService 创建收入服务实例
func NewIncomeService(incomeRepo repository.IncomeRepository) IncomeService {
	return &incomeService{
		incomeRepo: incomeRepo,
	}
}

// GetIncomeStats 获取收入统计
func (s *incomeService) GetIncomeStats(ctx context.Context, mentorID string, req *model.IncomeStatsRequest) (*model.IncomeStatsResponse, error) {
	// 设置默认时间范围
	startDate, endDate := s.getDateRange(req.Period, req.StartDate, req.EndDate)

	stats, err := s.incomeRepo.GetIncomeStats(ctx, mentorID, req.Period, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return &model.IncomeStatsResponse{
		Stats: stats,
	}, nil
}

// GetIncomeTransactions 获取收入明细
func (s *incomeService) GetIncomeTransactions(ctx context.Context, mentorID string, req *model.IncomeTransactionsRequest) (*model.IncomeTransactionsResponse, error) {
	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	transactions, total, err := s.incomeRepo.GetIncomeTransactions(ctx, mentorID, req.Type, req.Status, req.StartDate, req.EndDate, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.IncomeTransactionsResponse{
		Transactions: transactions,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// GetIncomeTrends 获取收入趋势
func (s *incomeService) GetIncomeTrends(ctx context.Context, mentorID string, req *model.IncomeTrendsRequest) (*model.IncomeTrendsResponse, error) {
	trends, err := s.incomeRepo.GetIncomeTrends(ctx, mentorID, req.Period, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	return &model.IncomeTrendsResponse{
		Trends: trends,
	}, nil
}

// ExportIncomeReport 导出收入报告
func (s *incomeService) ExportIncomeReport(ctx context.Context, mentorID string, req *model.IncomeExportRequest) (*model.IncomeExportResponse, error) {
	// 生成报告文件名
	fileName := s.generateReportFileName(req.Format, req.StartDate, req.EndDate)

	// 生成下载URL（这里简化处理，实际应该上传到文件存储服务）
	downloadURL := "https://example.com/reports/" + fileName

	// 设置过期时间（7天后）
	expiresAt := time.Now().AddDate(0, 0, 7)

	return &model.IncomeExportResponse{
		DownloadURL: downloadURL,
		ExpiresAt:   expiresAt,
	}, nil
}

// GetWithdrawals 获取提现记录
func (s *incomeService) GetWithdrawals(ctx context.Context, mentorID string, req *model.WithdrawalsRequest) (*model.WithdrawalsResponse, error) {
	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	withdrawals, total, err := s.incomeRepo.GetWithdrawals(ctx, mentorID, req.Status, req.StartDate, req.EndDate, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	// 计算分页信息
	totalPages := int(math.Ceil(float64(total) / float64(req.PageSize)))

	return &model.WithdrawalsResponse{
		Withdrawals: withdrawals,
		Pagination: &model.PaginationResponse{
			Total:      total,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: totalPages,
		},
	}, nil
}

// CreateWithdrawal 申请提现
func (s *incomeService) CreateWithdrawal(ctx context.Context, mentorID string, req *model.CreateWithdrawalRequest) (*model.CreateWithdrawalResponse, error) {
	// 检查可提现金额
	available, err := s.incomeRepo.GetAvailableIncome(ctx, mentorID)
	if err != nil {
		return nil, err
	}

	if req.Amount > available.AvailableAmount {
		return nil, errors.New("可提现金额不足")
	}

	if req.Amount < available.MinWithdrawal {
		return nil, errors.New("提现金额低于最小提现限制")
	}

	if req.Amount > available.MaxWithdrawal {
		return nil, errors.New("提现金额超过最大提现限制")
	}

	// 计算手续费（简化处理，实际应该有复杂的费率计算）
	fee := s.calculateWithdrawalFee(req.Amount)
	netAmount := req.Amount - fee

	// 创建提现记录
	withdrawal := &model.WithdrawalModel{
		MentorID:    mentorID,
		Amount:      req.Amount,
		Fee:         fee,
		NetAmount:   netAmount,
		Status:      "pending",
		BankAccount: req.BankAccount,
		BankName:    req.BankName,
	}

	err = s.incomeRepo.CreateWithdrawal(ctx, withdrawal)
	if err != nil {
		return nil, err
	}

	// 估算完成时间（2-3个工作日）
	estimatedCompletionTime := time.Now().AddDate(0, 0, 3)

	return &model.CreateWithdrawalResponse{
		WithdrawalID:            withdrawal.ID,
		EstimatedCompletionTime: estimatedCompletionTime,
	}, nil
}

// GetAvailableIncome 获取可提现金额
func (s *incomeService) GetAvailableIncome(ctx context.Context, mentorID string) (*model.AvailableIncomeResponse, error) {
	available, err := s.incomeRepo.GetAvailableIncome(ctx, mentorID)
	if err != nil {
		return nil, err
	}

	return &model.AvailableIncomeResponse{
		AvailableAmount: available.AvailableAmount,
		PendingAmount:   available.PendingAmount,
		TotalEarned:     available.TotalEarned,
		TotalWithdrawn:  available.TotalWithdrawn,
		MinWithdrawal:   available.MinWithdrawal,
		MaxWithdrawal:   available.MaxWithdrawal,
	}, nil
}

// getDateRange 根据周期获取时间范围
func (s *incomeService) getDateRange(period string, startDate, endDate time.Time) (time.Time, time.Time) {
	now := time.Now()

	if !startDate.IsZero() && !endDate.IsZero() {
		return startDate, endDate
	}

	switch period {
	case "week":
		weekStart := now.AddDate(0, 0, -int(now.Weekday()))
		return weekStart, now
	case "month":
		monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		return monthStart, now
	case "quarter":
		quarter := (now.Month()-1)/3 + 1
		quarterStart := time.Date(now.Year(), time.Month((quarter-1)*3+1), 1, 0, 0, 0, 0, now.Location())
		return quarterStart, now
	case "year":
		yearStart := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		return yearStart, now
	default:
		// all - 返回一个很早的时间到当前时间
		return time.Date(2020, 1, 1, 0, 0, 0, 0, now.Location()), now
	}
}

// generateReportFileName 生成报告文件名
func (s *incomeService) generateReportFileName(format string, startDate, endDate time.Time) string {
	startStr := startDate.Format("20060102")
	endStr := endDate.Format("20060102")
	return "income_report_" + startStr + "_" + endStr + "." + format
}

// calculateWithdrawalFee 计算提现手续费
func (s *incomeService) calculateWithdrawalFee(amount float64) float64 {
	// 简化费率计算：金额的0.2%，最低1元，最高50元
	fee := amount * 0.002
	if fee < 1 {
		fee = 1
	}
	if fee > 50 {
		fee = 50
	}
	return fee
}
