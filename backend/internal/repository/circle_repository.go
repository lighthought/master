package repository

import (
	"context"

	"master-guide-backend/internal/model"

	"gorm.io/gorm"
)

// CircleRepository 圈子数据访问接口
type CircleRepository interface {
	GetCircles(ctx context.Context, domain string, page, pageSize int) ([]*model.Circle, int64, error)
	GetRecommendedCircles(ctx context.Context, userID string, limit int) ([]*model.Circle, error)
	GetCircleByID(ctx context.Context, circleID string) (*model.Circle, error)
	JoinCircle(ctx context.Context, userID, circleID string) error
	LeaveCircle(ctx context.Context, userID, circleID string) error
	IsUserJoinedCircle(ctx context.Context, userID, circleID string) (bool, error)
	GetUserJoinedCircles(ctx context.Context, userID string) ([]*model.Circle, error)
}

// circleRepository 圈子数据访问实现
type circleRepository struct {
	db *gorm.DB
}

// NewCircleRepository 创建圈子数据访问实例
func NewCircleRepository(db *gorm.DB) CircleRepository {
	return &circleRepository{db: db}
}

// GetCircles 获取圈子列表
func (r *circleRepository) GetCircles(ctx context.Context, domain string, page, pageSize int) ([]*model.Circle, int64, error) {
	var circles []*model.Circle
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Circle{}).Where("status = ?", "active")

	if domain != "" {
		query = query.Where("domain = ?", domain)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("member_count DESC, created_at DESC").Offset(offset).Limit(pageSize).Find(&circles).Error

	return circles, total, err
}

// GetRecommendedCircles 获取推荐圈子
func (r *circleRepository) GetRecommendedCircles(ctx context.Context, userID string, limit int) ([]*model.Circle, error) {
	var circles []*model.Circle

	// 基于用户兴趣和圈子活跃度推荐
	query := r.db.WithContext(ctx).Model(&model.Circle{}).
		Where("status = ?", "active").
		Order("member_count DESC, created_at DESC").
		Limit(limit)

	err := query.Find(&circles).Error
	return circles, err
}

// GetCircleByID 根据ID获取圈子
func (r *circleRepository) GetCircleByID(ctx context.Context, circleID string) (*model.Circle, error) {
	var circle model.Circle
	err := r.db.WithContext(ctx).Where("id = ?", circleID).First(&circle).Error
	if err != nil {
		return nil, err
	}
	return &circle, nil
}

// JoinCircle 加入圈子
func (r *circleRepository) JoinCircle(ctx context.Context, userID, circleID string) error {
	// 检查是否已经加入
	exists, err := r.IsUserJoinedCircle(ctx, userID, circleID)
	if err != nil {
		return err
	}
	if exists {
		return gorm.ErrRecordNotFound // 使用这个错误表示已加入
	}

	// 创建成员记录
	member := &model.CircleMember{
		CircleID: circleID,
		UserID:   userID,
		Role:     "member",
	}

	err = r.db.WithContext(ctx).Create(member).Error
	if err != nil {
		return err
	}

	// 更新圈子成员数量
	err = r.db.WithContext(ctx).Model(&model.Circle{}).
		Where("id = ?", circleID).
		UpdateColumn("member_count", gorm.Expr("member_count + 1")).Error

	return err
}

// LeaveCircle 退出圈子
func (r *circleRepository) LeaveCircle(ctx context.Context, userID, circleID string) error {
	// 删除成员记录
	result := r.db.WithContext(ctx).Where("circle_id = ? AND user_id = ?", circleID, userID).
		Delete(&model.CircleMember{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	// 更新圈子成员数量
	err := r.db.WithContext(ctx).Model(&model.Circle{}).
		Where("id = ?", circleID).
		UpdateColumn("member_count", gorm.Expr("member_count - 1")).Error

	return err
}

// IsUserJoinedCircle 检查用户是否已加入圈子
func (r *circleRepository) IsUserJoinedCircle(ctx context.Context, userID, circleID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.CircleMember{}).
		Where("circle_id = ? AND user_id = ?", circleID, userID).
		Count(&count).Error
	return count > 0, err
}

// GetUserJoinedCircles 获取用户加入的圈子
func (r *circleRepository) GetUserJoinedCircles(ctx context.Context, userID string) ([]*model.Circle, error) {
	var circles []*model.Circle

	err := r.db.WithContext(ctx).
		Joins("JOIN circle_members ON circles.id = circle_members.circle_id").
		Where("circle_members.user_id = ? AND circles.status = ?", userID, "active").
		Order("circles.member_count DESC, circles.created_at DESC").
		Find(&circles).Error

	return circles, err
}
