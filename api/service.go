package api

import (
	"Activity/models"
	"Activity/storage/mysql/entity"
	"Activity/storage/mysql/repository"
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

// ActivityService 活动服务接口
type ActivityService interface {
	// 活动管理
	CreateActivity(ctx context.Context, req *CreateActivityRequest) (*ActivityResponse, error)
	UpdateActivity(ctx context.Context, req *UpdateActivityRequest) (*ActivityResponse, error)
	GetActivity(ctx context.Context, activityID int64) (*ActivityResponse, error)

	// 活动参与
	Participate(ctx context.Context, req *ParticipateRequest) (*ParticipationResponse, error)
	GetParticipation(ctx context.Context, activityID int64, userID string) (*ParticipationResponse, error)

	// 奖品管理
	DistributePrize(ctx context.Context, activityID int64, userID string, prizeType string, prizeID string) (*PrizeResponse, error)
}

// gameService 玩法服务实现
type gameService struct {
	activityRepo ActivityRepository
}

// activityService 活动服务实现
type activityService struct {
	activityRepo repository.ActivityRepository
}

func NewGameService(activityRepo ActivityRepository) GameService {
	return &gameService{
		activityRepo: activityRepo,
	}
}

// NewActivityService 创建活动服务实例
func NewActivityService(activityRepo repository.ActivityRepository) ActivityService {
	return &activityService{
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

	// 7. 保存用户参与记录
	if err := s.saveUserGameRecord(ctx, user, activityID, gameName, result); err != nil {
		return nil, fmt.Errorf("failed to save user game record: %w", err)
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
	switch g := game.(type) {
	case *models.CommunityPostGame:
		if g.Prize != nil {
			remainNum = g.Prize.RemainNum
			totalNum = g.Prize.TotalNum
		}
	case *models.CheckinGame:
		if g.Prize != nil {
			remainNum = g.Prize.RemainNum
			totalNum = g.Prize.TotalNum
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

// saveUserGameRecord 保存用户参与记录
func (s *gameService) saveUserGameRecord(ctx context.Context, user models.User, activityID, gameName string, result models.ResultInterface) error {
	// TODO: 实现用户参与记录的保存
	return nil
}

// ActivityRepository 活动仓库接口
type ActivityRepository interface {
	GetActivity(ctx context.Context, activityID string) (models.ActivityInterface, error)
}

// CreateActivity 创建活动
func (s *activityService) CreateActivity(ctx context.Context, req *CreateActivityRequest) (*ActivityResponse, error) {
	// 验证时间范围
	if req.StartAt >= req.EndAt {
		return nil, ErrInvalidParam
	}

	// 创建活动实体
	activity := &entity.Activity{
		Category: req.Category,
		Version:  req.Version,
		Name:     req.Name,
		// Config:   req.Config,
		StartAt: req.StartAt,
		EndAt:   req.EndAt,
		// Status:   req.Status,
	}

	// 保存到数据库
	if err := s.activityRepo.Create(ctx, activity); err != nil {
		return nil, ErrSystem
	}

	// 转换为响应
	return &ActivityResponse{
		ID:        activity.ID,
		Category:  activity.Category,
		Version:   activity.Version,
		Name:      activity.Name,
		Config:    activity.Config,
		StartAt:   activity.StartAt,
		EndAt:     activity.EndAt,
		Status:    activity.Status,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
	}, nil
}

// UpdateActivity 更新活动
func (s *activityService) UpdateActivity(ctx context.Context, req *UpdateActivityRequest) (*ActivityResponse, error) {
	// 获取活动
	activity, err := s.activityRepo.FindByID(ctx, 0)
	if err != nil {
		return nil, ErrActivityNotFound
	}

	// // 更新活动信息
	// activity.Config = req.Config
	// activity.Status = req.Status

	// 保存更新
	if err := s.activityRepo.Update(ctx, activity); err != nil {
		return nil, ErrSystem
	}

	// 转换为响应
	return &ActivityResponse{
		ID:        activity.ID,
		Category:  activity.Category,
		Version:   activity.Version,
		Name:      activity.Name,
		Config:    activity.Config,
		StartAt:   activity.StartAt,
		EndAt:     activity.EndAt,
		Status:    activity.Status,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
	}, nil
}

// GetActivity 获取活动信息
func (s *activityService) GetActivity(ctx context.Context, activityID int64) (*ActivityResponse, error) {
	// 获取活动
	activity, err := s.activityRepo.FindByID(ctx, activityID)
	if err != nil {
		return nil, ErrActivityNotFound
	}

	// 转换为响应
	return &ActivityResponse{
		ID:        activity.ID,
		Category:  activity.Category,
		Version:   activity.Version,
		Name:      activity.Name,
		Config:    activity.Config,
		StartAt:   activity.StartAt,
		EndAt:     activity.EndAt,
		Status:    activity.Status,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
	}, nil
}

// Participate 参与活动
func (s *activityService) Participate(ctx context.Context, req *ParticipateRequest) (*ParticipationResponse, error) {
	// 获取活动
	activity, err := s.activityRepo.FindByID(ctx, 0)
	if err != nil {
		return nil, ErrActivityNotFound
	}

	// 检查活动状态
	now := time.Now().Unix()
	if now < activity.StartAt {
		return nil, ErrActivityNotStarted
	}
	if now > activity.EndAt {
		return nil, ErrActivityEnded
	}
	if activity.Status != 1 {
		return nil, ErrGameClosed
	}

	// TODO: 实现具体的参与逻辑
	// 这里需要根据不同的活动类型和玩法类型实现不同的参与逻辑
	// 建议将这部分逻辑抽象到独立的GameService中

	return nil, nil
}

// GetParticipation 获取参与记录
func (s *activityService) GetParticipation(ctx context.Context, activityID int64, userID string) (*ParticipationResponse, error) {
	// TODO: 实现获取参与记录的逻辑
	return nil, nil
}

// DistributePrize 发放奖品
func (s *activityService) DistributePrize(ctx context.Context, activityID int64, userID string, prizeType string, prizeID string) (*PrizeResponse, error) {
	// TODO: 实现奖品发放逻辑
	return nil, nil
}
