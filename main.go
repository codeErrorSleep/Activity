package main

import (
	"Activity/api"
	"Activity/models"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Gin引擎
	r := gin.Default()

	// 创建依赖
	activityRepo := NewActivityRepository() // 需要实现
	gameService := api.NewGameService(activityRepo)
	handler := api.NewHandler(gameService)

	// 注册路由
	handler.RegisterRoutes(r)

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// ActivityRepository 活动仓库实现
type activityRepository struct {
	// TODO: 添加数据库连接等依赖
}

func NewActivityRepository() api.ActivityRepository {
	return &activityRepository{}
}

func (r *activityRepository) GetActivity(ctx context.Context, activityID string) (models.ActivityInterface, error) {
	// TODO: 从数据库获取活动信息
	// 这里先返回一个模拟数据
	return &models.CommunityActivity{
		MetaActivity: models.MetaActivity{
			ID:       1,
			Category: "community",
			Version:  "v1",
			StartAt:  1679000000,
			EndAt:    1679086400,
			Status:   1,
		},
		GameList: []models.GameInterface{
			&models.CommunityPostGame{
				Name_: "发帖奖励",
				Prize: &models.DiscountCodePrize{
					DiscountCode: "COMMUNITY_2024",
					PriceRuleID:  123,
					Probability:  100,
					TotalNum:     1000,
					RemainNum:    1000,
				},
				State: models.GameStateOPEN,
			},
		},
	}, nil
}
