package repository

import (
	"Activity/storage/mysql/entity"
	"context"
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
