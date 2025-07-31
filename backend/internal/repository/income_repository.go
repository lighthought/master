package repository

import (
	"context"
	"time"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// IncomeRepository 收入数据访问接口
type IncomeRepository interface {
	GetIncomeStats(ctx context.Context, mentorID, period string, startDate, endDate time.Time) (*model.IncomeStats, error)
	GetIncomeTransactions(ctx context.Context, mentorID, transactionType, status string, startDate, endDate time.Time, page, pageSize int) ([]*model.IncomeTransaction, int64, error)
	GetIncomeTrends(ctx context.Context, mentorID, period string, startDate, endDate time.Time) ([]*model.IncomeTrend, error)
	GetWithdrawals(ctx context.Context, mentorID, status string, startDate, endDate time.Time, page, pageSize int) ([]*model.Withdrawal, int64, error)
	CreateWithdrawal(ctx context.Context, withdrawal *model.WithdrawalModel) error
	GetAvailableIncome(ctx context.Context, mentorID string) (*model.AvailableIncome, error)
}

// incomeRepository 收入数据访问实现
type incomeRepository struct {
	db *gorm.DB
}

// NewIncomeRepository 创建收入数据访问实例
func NewIncomeRepository(db *gorm.DB) IncomeRepository {
	return &incomeRepository{db: db}
}

// GetIncomeStats 获取收入统计
func (r *incomeRepository) GetIncomeStats(ctx context.Context, mentorID, period string, startDate, endDate time.Time) (*model.IncomeStats, error) {
	var stats model.IncomeStats

	// 构建基础查询
	baseQuery := r.db.WithContext(ctx).Table("income_transactions")
	if mentorID != "" {
		baseQuery = baseQuery.Where("mentor_id = ?", mentorID)
	}

	// 根据周期设置时间范围
	if !startDate.IsZero() && !endDate.IsZero() {
		baseQuery = baseQuery.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// 获取总收入
	err := baseQuery.Select("COALESCE(SUM(amount), 0)").Scan(&stats.TotalIncome).Error
	if err != nil {
		return nil, err
	}

	// 获取总交易数
	err = baseQuery.Count(&stats.TotalTransactions).Error
	if err != nil {
		return nil, err
	}

	// 计算平均每笔交易金额
	if stats.TotalTransactions > 0 {
		stats.AveragePerTransaction = stats.TotalIncome / float64(stats.TotalTransactions)
	}

	// 获取按来源分类的收入
	var incomeBySource model.IncomeBySource
	sourceQuery := r.db.WithContext(ctx).Table("income_transactions")
	if mentorID != "" {
		sourceQuery = sourceQuery.Where("mentor_id = ?", mentorID)
	}
	if !startDate.IsZero() && !endDate.IsZero() {
		sourceQuery = sourceQuery.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}
	err = sourceQuery.Select(`
		COALESCE(SUM(CASE WHEN transaction_type = 'course_enrollment' THEN amount ELSE 0 END), 0) as course_enrollments,
		COALESCE(SUM(CASE WHEN transaction_type = 'appointment' THEN amount ELSE 0 END), 0) as appointments
	`).Scan(&incomeBySource).Error
	if err == nil {
		stats.IncomeBySource = &incomeBySource
	}

	// 获取按周期分类的收入
	var incomeByPeriod model.IncomeByPeriod
	now := time.Now()

	// 当前月
	currentMonthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	currentMonthQuery := r.db.WithContext(ctx).Table("income_transactions")
	if mentorID != "" {
		currentMonthQuery = currentMonthQuery.Where("mentor_id = ?", mentorID)
	}
	err = currentMonthQuery.Where("created_at >= ?", currentMonthStart).Select("COALESCE(SUM(amount), 0)").Scan(&incomeByPeriod.CurrentMonth).Error
	if err != nil {
		return nil, err
	}

	// 上个月
	lastMonthStart := currentMonthStart.AddDate(0, -1, 0)
	lastMonthEnd := currentMonthStart.Add(-time.Second)
	lastMonthQuery := r.db.WithContext(ctx).Table("income_transactions")
	if mentorID != "" {
		lastMonthQuery = lastMonthQuery.Where("mentor_id = ?", mentorID)
	}
	err = lastMonthQuery.Where("created_at BETWEEN ? AND ?", lastMonthStart, lastMonthEnd).Select("COALESCE(SUM(amount), 0)").Scan(&incomeByPeriod.PreviousMonth).Error
	if err != nil {
		return nil, err
	}

	// 当前季度
	quarter := (now.Month()-1)/3 + 1
	quarterStart := time.Date(now.Year(), time.Month((quarter-1)*3+1), 1, 0, 0, 0, 0, now.Location())
	quarterQuery := r.db.WithContext(ctx).Table("income_transactions")
	if mentorID != "" {
		quarterQuery = quarterQuery.Where("mentor_id = ?", mentorID)
	}
	err = quarterQuery.Where("created_at >= ?", quarterStart).Select("COALESCE(SUM(amount), 0)").Scan(&incomeByPeriod.CurrentQuarter).Error
	if err != nil {
		return nil, err
	}

	// 当前年
	yearStart := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	yearQuery := r.db.WithContext(ctx).Table("income_transactions")
	if mentorID != "" {
		yearQuery = yearQuery.Where("mentor_id = ?", mentorID)
	}
	err = yearQuery.Where("created_at >= ?", yearStart).Select("COALESCE(SUM(amount), 0)").Scan(&incomeByPeriod.CurrentYear).Error
	if err != nil {
		return nil, err
	}

	stats.IncomeByPeriod = &incomeByPeriod

	// 计算增长率
	if incomeByPeriod.PreviousMonth > 0 {
		stats.GrowthRate = (incomeByPeriod.CurrentMonth - incomeByPeriod.PreviousMonth) / incomeByPeriod.PreviousMonth * 100
	}

	// 获取热门课程收入
	var topCourses []*model.TopCourseIncome
	topCoursesQuery := r.db.WithContext(ctx).
		Table("income_transactions it").
		Select(`
			it.course_id,
			c.title,
			COALESCE(SUM(it.amount), 0) as income,
			COUNT(DISTINCT it.student_id) as enrollment_count
		`).
		Joins("LEFT JOIN courses c ON it.course_id = c.id").
		Where("it.transaction_type = ?", "course_enrollment").
		Where("it.course_id IS NOT NULL")

	if mentorID != "" {
		topCoursesQuery = topCoursesQuery.Where("it.mentor_id = ?", mentorID)
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		topCoursesQuery = topCoursesQuery.Where("it.created_at BETWEEN ? AND ?", startDate, endDate)
	}

	err = topCoursesQuery.Group("it.course_id, c.title").
		Order("income DESC").
		Limit(10).
		Find(&topCourses).Error
	if err == nil {
		stats.TopCourses = topCourses
	}

	return &stats, nil
}

// GetIncomeTransactions 获取收入明细
func (r *incomeRepository) GetIncomeTransactions(ctx context.Context, mentorID, transactionType, status string, startDate, endDate time.Time, page, pageSize int) ([]*model.IncomeTransaction, int64, error) {
	var transactions []*model.IncomeTransaction
	var total int64

	query := r.db.WithContext(ctx).
		Table("income_transactions it").
		Select(`
			it.id, it.transaction_type as type, it.amount, it.status, it.description,
			u.email as student_name, c.title as course_title,
			it.created_at, it.completed_at, it.platform_fee, it.net_income
		`).
		Joins("LEFT JOIN users u ON it.student_id = u.id").
		Joins("LEFT JOIN courses c ON it.course_id = c.id")

	if mentorID != "" {
		query = query.Where("it.mentor_id = ?", mentorID)
	}

	if transactionType != "" {
		query = query.Where("it.transaction_type = ?", transactionType)
	}

	if status != "" {
		query = query.Where("it.status = ?", status)
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("it.created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("it.created_at DESC").Offset(offset).Limit(pageSize).Find(&transactions).Error

	return transactions, total, err
}

// GetIncomeTrends 获取收入趋势
func (r *incomeRepository) GetIncomeTrends(ctx context.Context, mentorID, period string, startDate, endDate time.Time) ([]*model.IncomeTrend, error) {
	var trends []*model.IncomeTrend

	// 根据周期调整日期格式
	var dateFormat string
	switch period {
	case "daily":
		dateFormat = "DATE(it.created_at)"
	case "weekly":
		dateFormat = "DATE_TRUNC('week', it.created_at)"
	case "monthly":
		dateFormat = "DATE_TRUNC('month', it.created_at)"
	default:
		dateFormat = "DATE(it.created_at)"
	}

	query := r.db.WithContext(ctx).
		Table("income_transactions it").
		Select(`
			` + dateFormat + ` as date,
			COALESCE(SUM(it.amount), 0) as income,
			COUNT(*) as transactions,
			COUNT(CASE WHEN it.transaction_type = 'course_enrollment' THEN 1 END) as course_enrollments,
			COUNT(CASE WHEN it.transaction_type = 'appointment' THEN 1 END) as appointments
		`)

	if mentorID != "" {
		query = query.Where("it.mentor_id = ?", mentorID)
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("it.created_at BETWEEN ? AND ?", startDate, endDate)
	}

	query = query.Group(dateFormat).Order("date")

	err := query.Find(&trends).Error

	return trends, err
}

// GetWithdrawals 获取提现记录
func (r *incomeRepository) GetWithdrawals(ctx context.Context, mentorID, status string, startDate, endDate time.Time, page, pageSize int) ([]*model.Withdrawal, int64, error) {
	var withdrawals []*model.Withdrawal
	var total int64

	query := r.db.WithContext(ctx).
		Table("withdrawals w").
		Select(`
			w.id, w.amount, w.status, w.bank_account, w.created_at, w.completed_at, w.fee, w.net_amount
		`)

	if mentorID != "" {
		query = query.Where("w.mentor_id = ?", mentorID)
	}

	if status != "" {
		query = query.Where("w.status = ?", status)
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("w.created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("w.created_at DESC").Offset(offset).Limit(pageSize).Find(&withdrawals).Error

	return withdrawals, total, err
}

// CreateWithdrawal 创建提现申请
func (r *incomeRepository) CreateWithdrawal(ctx context.Context, withdrawal *model.WithdrawalModel) error {
	return r.db.WithContext(ctx).Create(withdrawal).Error
}

// GetAvailableIncome 获取可提现金额
func (r *incomeRepository) GetAvailableIncome(ctx context.Context, mentorID string) (*model.AvailableIncome, error) {
	var available model.AvailableIncome

	// 获取总收入
	err := r.db.WithContext(ctx).
		Table("income_transactions").
		Where("mentor_id = ? AND status = ?", mentorID, "completed").
		Select("COALESCE(SUM(net_income), 0)").
		Scan(&available.TotalEarned).Error
	if err != nil {
		return nil, err
	}

	// 获取已提现总额
	err = r.db.WithContext(ctx).
		Table("withdrawals").
		Where("mentor_id = ? AND status = ?", mentorID, "completed").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&available.TotalWithdrawn).Error
	if err != nil {
		return nil, err
	}

	// 获取待处理提现总额
	err = r.db.WithContext(ctx).
		Table("withdrawals").
		Where("mentor_id = ? AND status = ?", mentorID, "pending").
		Select("COALESCE(SUM(amount), 0)").
		Scan(&available.PendingAmount).Error
	if err != nil {
		return nil, err
	}

	// 计算可提现金额
	available.AvailableAmount = available.TotalEarned - available.TotalWithdrawn - available.PendingAmount

	// 设置提现限制
	available.MinWithdrawal = 100.0
	available.MaxWithdrawal = available.AvailableAmount

	return &available, nil
}
