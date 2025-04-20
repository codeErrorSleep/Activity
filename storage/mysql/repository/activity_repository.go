package repository

import (
	"Activity/models"
	"Activity/storage/mysql/entity"
	"context"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// ActivityRepository 活动仓储接口
type ActivityRepository interface {
	Create(ctx context.Context, activity *entity.Activity) error
	Update(ctx context.Context, activity *entity.Activity) error
	FindByID(ctx context.Context, id int64) (*entity.Activity, error)
	FindByCategory(ctx context.Context, category string) ([]*entity.Activity, error)
	FindActive(ctx context.Context) ([]*entity.Activity, error)
	GetActivity(ctx context.Context, activityID string) (models.ActivityInterface, error)
}

// activityRepository 活动仓储实现
type activityRepository struct {
	db *gorm.DB
}

// NewActivityRepository 创建活动仓储实例
func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{db: db}
}

// Create 创建活动
func (r *activityRepository) Create(ctx context.Context, activity *entity.Activity) error {
	return r.db.WithContext(ctx).Create(activity).Error
}

// Update 更新活动
func (r *activityRepository) Update(ctx context.Context, activity *entity.Activity) error {
	return r.db.WithContext(ctx).Save(activity).Error
}

// FindByID 根据ID查找活动
func (r *activityRepository) FindByID(ctx context.Context, id int64) (*entity.Activity, error) {
	var activity entity.Activity
	err := r.db.WithContext(ctx).First(&activity, id).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

// FindByCategory 根据类型查找活动
func (r *activityRepository) FindByCategory(ctx context.Context, category string) ([]*entity.Activity, error) {
	var activities []*entity.Activity
	err := r.db.WithContext(ctx).Where("category = ?", category).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

// FindActive 查找当前有效的活动
func (r *activityRepository) FindActive(ctx context.Context) ([]*entity.Activity, error) {
	now := time.Now().Unix()
	var activities []*entity.Activity
	err := r.db.WithContext(ctx).
		Where("status = ? AND start_at <= ? AND end_at >= ?", 1, now, now).
		Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

// GetActivity 获取活动信息
func (r *activityRepository) GetActivity(ctx context.Context, activityID string) (models.ActivityInterface, error) {
	id, err := strconv.ParseInt(activityID, 10, 64)
	if err != nil {
		return nil, err
	}

	activity, err := r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 将数据库实体转换为领域模型
	return &models.CommunityActivity{
		MetaActivity: models.MetaActivity{
			ID:       activity.ID,
			Category: activity.Category,
			Version:  activity.Version,
			StartAt:  activity.StartAt,
			EndAt:    activity.EndAt,
			Status:   activity.Status,
		},
		GameList: []models.GameInterface{
			&models.CommunityPostGame{
				Name_: "发帖奖励",
				Prize: &models.DiscountCodePrize{
					DiscountCode: "COMMUNITY_2024",
					PriceRuleID:  123,
					Probability:  100,
					TotalNum:     1000,
					RemainNum:    1000,
				},
				State: models.GameStateOPEN,
			},
		},
	}, nil
}
