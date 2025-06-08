package models

import (
	"context"
	"encoding/json"
	"fmt"
)

// CommunityActivityFactory 社区活动工厂
type CommunityActivityFactory struct{}

func (f *CommunityActivityFactory) Create(config ActivityConfigJSON) (ActivityInterface, error) {
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

// init 注册活动工厂
func init() {
	RegisterActivityFactory("community", &CommunityActivityFactory{})
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

func (a *CommunityActivity) StartAt() int64 {
	return a.MetaActivity.StartAt
}

func (a *CommunityActivity) EndAt() int64 {
	return a.MetaActivity.EndAt
}

func (a *CommunityActivity) Status() int64 {
	return a.MetaActivity.Status
}

// CommunityPostGame 社区发帖玩法
type CommunityPostGame struct {
	Name_ string             `json:"-"` // 玩法名称，从GameConfig中获取
	Prize *DiscountCodePrize `json:"prize"`
	State GameState          `json:"state"`
}

// Name 返回玩法名称
func (p CommunityPostGame) Name(ctx context.Context) string {
	return p.Name_
}

// Perform 执行玩法逻辑
func (p CommunityPostGame) Perform(ctx context.Context, user User, action ActionInterface) (ResultInterface, error) {
	// 1. 检查玩法状态
	if p.GameState(ctx) != GameStateOPEN {
		return nil, fmt.Errorf("game is not open")
	}

	// 2. 检查用户状态
	if p.UserState(ctx) != UserStateOPEN {
		return nil, fmt.Errorf("user cannot participate")
	}

	// 3. 验证用户是否已发帖
	// TODO: 调用社区服务检查用户是否已发帖
	// checkUserPost(ctx, user.Uid)

	// 4. 发放折扣码奖励
	if p.Prize != nil {
		err := p.Prize.WinPrize(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("failed to give prize: %w", err)
		}
	}

	// 5. 记录用户参与状态
	// TODO: 更新用户参与记录
	// updateUserGameRecord(ctx, user.Uid, UserStateCLOSED)

	return &CommunityPostResult{
		GameName: p.Name_,
		Prize:    p.Prize,
	}, nil
}

// GameState 返回玩法状态
func (p CommunityPostGame) GameState(ctx context.Context) GameState {
	return p.State
}

// UserState 返回用户参与状态
func (p CommunityPostGame) UserState(ctx context.Context) UserState {
	// TODO:
	// 1. 检查用户是否已经参与过
	// 2. 检查用户是否已经获得奖励
	// 3. 返回对应状态
	return UserStateOPEN
}

// ValidateConfig 验证配置
func (p CommunityPostGame) ValidateConfig(ctx context.Context) error {
	if p.Prize == nil {
		return fmt.Errorf("prize is not configured")
	}
	return nil
}

// Actions 返回支持的操作
func (p CommunityPostGame) Actions(ctx context.Context) []ActionInterface {
	return []ActionInterface{
		&CommunityPostAction{}, // 定义发帖动作
	}
}

// Results 返回支持的结果
func (p CommunityPostGame) Results(ctx context.Context) []ResultInterface {
	return []ResultInterface{
		&CommunityPostResult{}, // 定义结果类型
	}
}

// MarshalJSON 实现json.Marshaler接口
func (p CommunityPostGame) MarshalJSON() ([]byte, error) {
	type Alias CommunityPostGame
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(&p),
	})
}

// UnmarshalJSON 实现json.Unmarshaler接口
func (p *CommunityPostGame) UnmarshalJSON(data []byte) error {
	type Alias CommunityPostGame
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

// CommunityPostAction 发帖动作
type CommunityPostAction struct {
	PostID string `json:"post_id"` // 用户发帖ID
}

func (a CommunityPostAction) Target(ctx context.Context) string {
	return "community_post"
}

// CommunityPostResult 发帖结果
type CommunityPostResult struct {
	GameName string             `json:"game_name"`
	Prize    *DiscountCodePrize `json:"prize"`
}

func (r CommunityPostResult) Target(ctx context.Context) string {
	return r.GameName
}
