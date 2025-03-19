package models

import (
	"context"
	"time"
)

type ActivityInterface interface {
	Category() string       // 活动类型
	Version() string        // 活动版本
	Name() string           // 活动名称，具有唯一约束
	Games() []GameInterface // 活动下玩法
	StartAt() int64         // 活动开始时间戳
	EndAt() int64           // 活动结束时间戳
	Status() int64          // 活动状态
}

type ResultInterface interface {
	Target(ctx context.Context) string // 当有多个相同玩法时，用于标识当前响应是来自哪个具体玩法
}

type ActionInterface interface {
	Target(ctx context.Context) string // 当有多个相同玩法时，用于标识当前请求指定的具体玩法
}

type MetaActivity struct {
	ID             int64          `db:"id"`              // 主键
	CreatedAt      time.Time      `db:"created_at"`      // 创建时间
	UpdatedAt      time.Time      `db:"updated_at"`      // 更新时间
	Category       string         `db:"category"`        // 活动类型
	Version        string         `db:"version"`         // 活动的版本
	ActivityConfig ActivityConfig `db:"activity_config"` // 活动的JSON配置
	StartAt        int64          `db:"start_at"`        // 活动开始时间戳
	EndAt          int64          `db:"end_at"`          // 活动结束时间戳
	Status         int64          `db:"status"`          // 0-draft; 1-online
}
type ActivityConfig struct {
	Activity ActivityInterface
}

type User struct {
	Uid string
}

type UserState = string

const (
	UserStateUNKNOWN UserState = "UNKNOWN" // 占位用
	UserStatePENDING UserState = "PENDING" // 请优先判断玩法状态；玩法开放的前提下，表示用户未参加;
	UserStateOPEN    UserState = "OPEN"    // 请优先判断玩法状态；玩法开放的前提下，还可以继续参加;
	UserStateCLOSED  UserState = "CLOSED"  // 请优先判断玩法状态；玩法开放的前提下，不能参加了，已经有结果
)
