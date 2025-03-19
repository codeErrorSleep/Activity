package models

import (
	"context"
	"encoding/json"
)

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

type GameState = string

const (
	GameStateUNKNOWN GameState = "UNKNOWN" // 占位用
	GameStateOPEN    GameState = "OPEN"    // 还可以继续参加
	GameStateCLOSED  GameState = "CLOSED"  // 玩法关闭
)
