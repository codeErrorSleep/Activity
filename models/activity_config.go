package models

import (
	"encoding/json"
	"fmt"
)

// ActivityConfigJSON 活动配置JSON结构体
type ActivityConfigJSON struct {
	Category string       `json:"category"` // 活动类型
	Version  string       `json:"version"`  // 活动版本
	Name     string       `json:"name"`     // 活动名称
	StartAt  int64        `json:"start_at"` // 开始时间
	EndAt    int64        `json:"end_at"`   // 结束时间
	Games    []GameConfig `json:"games"`    // 玩法配置列表
}

// GameConfig 玩法配置结构体
type GameConfig struct {
	Type   string          `json:"type"`   // 玩法类型
	Name   string          `json:"name"`   // 玩法名称
	Config json.RawMessage `json:"config"` // 玩法具体配置
}

// NewActivityFromConfig 根据配置创建活动实例
func NewActivityFromConfig(configJSON []byte) (ActivityInterface, error) {
	var config ActivityConfigJSON
	if err := json.Unmarshal(configJSON, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// 根据活动类型创建对应的活动实例
	switch config.Category {
	case "community":
		return NewCommunityActivity(config)
	default:
		return nil, fmt.Errorf("unsupported activity category: %s", config.Category)
	}
}

// NewCommunityActivity 创建社区活动实例
func NewCommunityActivity(config ActivityConfigJSON) (ActivityInterface, error) {
	// 解析玩法配置
	var games []GameInterface
	for _, gameConfig := range config.Games {
		game, err := NewGameFromConfig(gameConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create game: %w", err)
		}
		games = append(games, game)
	}

	return &CommunityActivity{
		MetaActivity: MetaActivity{
			Category: config.Category,
			Version:  config.Version,
			StartAt:  config.StartAt,
			EndAt:    config.EndAt,
			Status:   1, // 默认上线状态
		},
		GameList: games,
	}, nil
}

// NewGameFromConfig 根据配置创建玩法实例
func NewGameFromConfig(config GameConfig) (GameInterface, error) {
	switch config.Type {
	case "post":
		var game CommunityPostGame
		if err := json.Unmarshal(config.Config, &game); err != nil {
			return nil, fmt.Errorf("failed to unmarshal game config: %w", err)
		}
		game.Name_ = config.Name
		return &game, nil
	default:
		return nil, fmt.Errorf("unsupported game type: %s", config.Type)
	}
}

// CommunityActivity 社区活动实现
type CommunityActivity struct {
	MetaActivity
	GameList []GameInterface
}

func (a *CommunityActivity) Category() string {
	return a.MetaActivity.Category
}

func (a *CommunityActivity) Version() string {
	return a.MetaActivity.Version
}

func (a *CommunityActivity) Name() string {
	return a.MetaActivity.ActivityConfig.Activity.Name()
}

func (a *CommunityActivity) Games() []GameInterface {
	return a.GameList
}
