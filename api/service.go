package api

import (
	"Activity/models"
	"context"
	"fmt"
	"time"
)

// GameService 玩法服务接口
type GameService interface {
	ParticipateGame(ctx context.Context, user models.User, activityID, gameName string, action models.ActionInterface) (interface{}, error)
	GetGameStatus(ctx context.Context, user models.User, activityID, gameName string) (*GameStatusResp, error)
	GetUserPrize(ctx context.Context, user models.User, activityID, gameName string) (*UserPrizeResp, error)
}

// gameService 玩法服务实现
type gameService struct {
	activityRepo ActivityRepository
}

func NewGameService(activityRepo ActivityRepository) GameService {
	return &gameService{
		activityRepo: activityRepo,
	}
}

// ParticipateGame 参与玩法
func (s *gameService) ParticipateGame(ctx context.Context, user models.User, activityID, gameName string, action models.ActionInterface) (interface{}, error) {
	// 1. 获取活动信息
	activity, err := s.activityRepo.GetActivity(ctx, activityID)
	if err != nil {
		return nil, fmt.Errorf("failed to get activity: %w", err)
	}

	// 2. 检查活动状态
	if err := s.checkActivityStatus(activity); err != nil {
		return nil, err
	}

	// 3. 获取玩法
	game, err := s.getGameByName(activity, gameName)
	if err != nil {
		return nil, err
	}

	// 4. 检查玩法状态
	if game.GameState(ctx) != models.GameStateOPEN {
		return nil, fmt.Errorf("game is closed")
	}

	// 5. 检查用户状态
	if game.UserState(ctx) != models.UserStateOPEN {
		return nil, fmt.Errorf("user cannot participate")
	}

	// 6. 执行玩法逻辑
	result, err := game.Perform(ctx, user, action)
	if err != nil {
		return nil, fmt.Errorf("failed to perform game: %w", err)
	}

	return result, nil
}

// GetGameStatus 获取玩法状态
func (s *gameService) GetGameStatus(ctx context.Context, user models.User, activityID, gameName string) (*GameStatusResp, error) {
	// 1. 获取活动信息
	activity, err := s.activityRepo.GetActivity(ctx, activityID)
	if err != nil {
		return nil, fmt.Errorf("failed to get activity: %w", err)
	}

	// 2. 检查活动状态
	if err := s.checkActivityStatus(activity); err != nil {
		return nil, err
	}

	// 3. 获取玩法
	game, err := s.getGameByName(activity, gameName)
	if err != nil {
		return nil, err
	}

	// 4. 获取奖品信息
	var remainNum, totalNum int64
	if communityGame, ok := game.(*models.CommunityPostGame); ok {
		if communityGame.Prize != nil {
			remainNum = communityGame.Prize.RemainNum
			totalNum = communityGame.Prize.TotalNum
		}
	}

	return &GameStatusResp{
		GameState: string(game.GameState(ctx)),
		UserState: string(game.UserState(ctx)),
		StartTime: time.Unix(activity.StartAt(), 0),
		EndTime:   time.Unix(activity.EndAt(), 0),
		RemainNum: remainNum,
		TotalNum:  totalNum,
	}, nil
}

// GetUserPrize 获取用户奖品
func (s *gameService) GetUserPrize(ctx context.Context, user models.User, activityID, gameName string) (*UserPrizeResp, error) {
	// 1. 获取活动信息
	activity, err := s.activityRepo.GetActivity(ctx, activityID)
	if err != nil {
		return nil, fmt.Errorf("failed to get activity: %w", err)
	}

	// 2. 检查活动状态
	if err := s.checkActivityStatus(activity); err != nil {
		return nil, err
	}

	// 3. 获取玩法
	_, err = s.getGameByName(activity, gameName)
	if err != nil {
		return nil, err
	}

	// 4. 获取用户奖品信息
	// TODO: 从数据库获取用户奖品记录
	prize := &PrizeInfo{
		Type:         "discount_code",
		DiscountCode: "COMMUNITY_2024",
		PriceRuleID:  123,
	}

	return &UserPrizeResp{
		Prize:     *prize,
		CreatedAt: time.Now(),
	}, nil
}

// checkActivityStatus 检查活动状态
func (s *gameService) checkActivityStatus(activity models.ActivityInterface) error {
	now := time.Now().Unix()
	if now < activity.StartAt() {
		return fmt.Errorf("activity not started")
	}
	if now > activity.EndAt() {
		return fmt.Errorf("activity ended")
	}
	return nil
}

// getGameByName 根据名称获取玩法
func (s *gameService) getGameByName(activity models.ActivityInterface, gameName string) (models.GameInterface, error) {
	for _, game := range activity.Games() {
		if game.Name(context.Background()) == gameName {
			return game, nil
		}
	}
	return nil, fmt.Errorf("game not found")
}

// ActivityRepository 活动仓库接口
type ActivityRepository interface {
	GetActivity(ctx context.Context, activityID string) (models.ActivityInterface, error)
}
