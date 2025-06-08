package models

import (
	"encoding/json"
	"fmt"
	"sync"
)

// ActivityFactory 活动工厂接口
type ActivityFactory interface {
	Create(config ActivityConfigJSON) (ActivityInterface, error)
}

// activityFactoryRegistry 活动工厂注册表
var (
	activityFactoryRegistry = make(map[string]ActivityFactory)
	registryMutex           sync.RWMutex
)

// RegisterActivityFactory 注册活动工厂
func RegisterActivityFactory(category string, factory ActivityFactory) {
	registryMutex.Lock()
	defer registryMutex.Unlock()
	activityFactoryRegistry[category] = factory
}

// GetActivityFactory 获取活动工厂
func GetActivityFactory(category string) (ActivityFactory, bool) {
	registryMutex.RLock()
	defer registryMutex.RUnlock()
	factory, exists := activityFactoryRegistry[category]
	return factory, exists
}

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

	factory, exists := GetActivityFactory(config.Category)
	if !exists {
		return nil, fmt.Errorf("unsupported activity category: %s", config.Category)
	}

	return factory.Create(config)
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
