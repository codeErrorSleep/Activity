package entity

import (
	"time"

	"gorm.io/gorm"
)

// Activity 活动表实体
type Activity struct {
	ID        int64          `gorm:"primaryKey;autoIncrement"`
	Category  string         `gorm:"type:varchar(50);not null;index:idx_category"`
	Version   string         `gorm:"type:varchar(20);not null"`
	Name      string         `gorm:"type:varchar(100);not null"`
	Config    string         `gorm:"type:json;not null"`
	StartAt   int64          `gorm:"not null;index:idx_status_time"`
	EndAt     int64          `gorm:"not null;index:idx_status_time"`
	Status    int64          `gorm:"type:tinyint;not null;default:0;index:idx_status_time"`
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName 指定表名
func (Activity) TableName() string {
	return "activities"
}

// ActivityParticipation 用户参与记录表实体
type ActivityParticipation struct {
	ID         int64          `gorm:"primaryKey;autoIncrement"`
	ActivityID int64          `gorm:"not null;index:idx_activity_user"`
	UserID     string         `gorm:"type:varchar(50);not null;index:idx_activity_user,idx_user_state"`
	GameType   string         `gorm:"type:varchar(50);not null"`
	GameTarget string         `gorm:"type:varchar(50);not null"`
	State      string         `gorm:"type:varchar(20);not null;index:idx_user_state"`
	Extra      string         `gorm:"type:json"`
	CreatedAt  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// TableName 指定表名
func (ActivityParticipation) TableName() string {
	return "activity_participations"
}

// PrizeRecord 奖品发放记录表实体
type PrizeRecord struct {
	ID         int64          `gorm:"primaryKey;autoIncrement"`
	ActivityID int64          `gorm:"not null;index:idx_activity_user"`
	UserID     string         `gorm:"type:varchar(50);not null;index:idx_activity_user"`
	PrizeType  string         `gorm:"type:varchar(50);not null"`
	PrizeID    string         `gorm:"type:varchar(50);not null"`
	Status     int64          `gorm:"type:tinyint;not null;default:0;index:idx_status"`
	CreatedAt  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// TableName 指定表名
func (PrizeRecord) TableName() string {
	return "prize_records"
}
