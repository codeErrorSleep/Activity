package models

import (
	"context"
	"encoding/json"
	"time"
)

type ActivityInterface interface {
	Category() string       // 活动类型
	Version() string        // 活动版本
	Name() string           // 活动名称，具有唯一约束
	Games() []GameInterface // 活动下玩法
}

// GameInterface 是对目前Shopping项目中所有玩法的公共抽象
// Method 中第一个参数为context的出发点是传递上下文，例如tracing等场景
type GameInterface interface {
	// Perform 当请求到来时，游戏实例需要执行的业务逻辑
	Perform(ctx context.Context, user User, action ActionInterface) (ResultInterface, error)
	// ValidateConfig 校验配置是否合法, ctx 用于传递活动和玩法的上下文信息
	ValidateConfig(ctx context.Context) error

	Name(ctx context.Context) string               // 游戏的名称，在同一个活动中，游戏名称必须保证唯一
	Actions(ctx context.Context) []ActionInterface // 游戏支持哪些请求
	Results(ctx context.Context) []ResultInterface // 游戏支持哪些响应
	GameState(ctx context.Context) GameState       // 玩法状态
	UserState(ctx context.Context) UserState       // 用户参与结果

	json.Unmarshaler // 非业务功能
	json.Marshaler   // 非业务功能
}

// PrizeInterface 奖品的interface
type PrizeInterface interface {
	WinPrize(ctx context.Context, user User) error // 中奖后需要执行的逻辑
	WinProbability() int64                         // 中奖概率
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
type GameState = string

const (
	GameStateUNKNOWN GameState = "UNKNOWN" // 占位用
	GameStateOPEN    GameState = "OPEN"    // 还可以继续参加
	GameStateCLOSED  GameState = "CLOSED"  // 玩法关闭
)

type UserState = string

const (
	UserStateUNKNOWN UserState = "UNKNOWN" // 占位用
	UserStatePENDING UserState = "PENDING" // 请优先判断玩法状态；玩法开放的前提下，表示用户未参加;
	UserStateOPEN    UserState = "OPEN"    // 请优先判断玩法状态；玩法开放的前提下，还可以继续参加;
	UserStateCLOSED  UserState = "CLOSED"  // 请优先判断玩法状态；玩法开放的前提下，不能参加了，已经有结果
)
