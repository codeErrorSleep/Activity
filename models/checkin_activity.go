package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CheckinActivity 签到活动
type CheckinActivity struct {
	MetaActivity
	GameList []GameInterface
}

// CheckinGame 签到玩法
type CheckinGame struct {
	Name_  string             `json:"-"`
	Prize  *DiscountCodePrize `json:"prize"`
	State  GameState          `json:"state"`
	Config CheckinConfig      `json:"config"`
}

// CheckinConfig 签到配置
type CheckinConfig struct {
	RequiredDays int64 `json:"required_days"` // 需要连续签到的天数
	CheckinDays  int64 `json:"checkin_days"`  // 当前已签到天数
}

// Name 返回玩法名称
func (p CheckinGame) Name(ctx context.Context) string {
	return p.Name_
}

// Perform 执行签到
func (p CheckinGame) Perform(ctx context.Context, user User, action ActionInterface) (ResultInterface, error) {
	// 1. 检查玩法状态
	if p.GameState(ctx) != GameStateOPEN {
		return nil, fmt.Errorf("game is not open")
	}

	// 2. 检查用户状态
	if p.UserState(ctx) != UserStateOPEN {
		return nil, fmt.Errorf("user cannot participate")
	}

	// 3. 更新签到天数
	p.Config.CheckinDays++

	// 4. 检查是否达到要求天数
	if p.Config.CheckinDays >= p.Config.RequiredDays {
		// 5. 发放奖励
		if p.Prize != nil {
			err := p.Prize.WinPrize(ctx, user)
			if err != nil {
				return nil, fmt.Errorf("failed to give prize: %w", err)
			}
		}
	}

	return &CheckinResult{
		GameName:     p.Name_,
		CheckinDays:  p.Config.CheckinDays,
		RequiredDays: p.Config.RequiredDays,
		Prize:        p.Prize,
	}, nil
}

// GameState 返回玩法状态
func (p CheckinGame) GameState(ctx context.Context) GameState {
	return p.State
}

// UserState 返回用户状态
func (p CheckinGame) UserState(ctx context.Context) UserState {
	// 如果已经达到要求天数，则不能再参与
	if p.Config.CheckinDays >= p.Config.RequiredDays {
		return UserStateCLOSED
	}
	return UserStateOPEN
}

// ValidateConfig 验证配置
func (p CheckinGame) ValidateConfig(ctx context.Context) error {
	if p.Prize == nil {
		return fmt.Errorf("prize is not configured")
	}
	if p.Config.RequiredDays <= 0 {
		return fmt.Errorf("required days must be greater than 0")
	}
	return nil
}

// Actions 返回支持的操作
func (p CheckinGame) Actions(ctx context.Context) []ActionInterface {
	return []ActionInterface{
		&CheckinAction{}, // 定义签到动作
	}
}

// Results 返回支持的结果
func (p CheckinGame) Results(ctx context.Context) []ResultInterface {
	return []ResultInterface{
		&CheckinResult{}, // 定义结果类型
	}
}

// MarshalJSON 实现json.Marshaler接口
func (p CheckinGame) MarshalJSON() ([]byte, error) {
	type Alias CheckinGame
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(&p),
	})
}

// UnmarshalJSON 实现json.Unmarshaler接口
func (p *CheckinGame) UnmarshalJSON(data []byte) error {
	type Alias CheckinGame
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

// CheckinAction 签到动作
type CheckinAction struct {
	CheckinTime time.Time `json:"checkin_time"` // 签到时间
}

func (a CheckinAction) Target(ctx context.Context) string {
	return "checkin"
}

// CheckinResult 签到结果
type CheckinResult struct {
	GameName     string             `json:"game_name"`
	CheckinDays  int64              `json:"checkin_days"`
	RequiredDays int64              `json:"required_days"`
	Prize        *DiscountCodePrize `json:"prize"`
}

func (r CheckinResult) Target(ctx context.Context) string {
	return r.GameName
}
