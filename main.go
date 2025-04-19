package main

import (
	"Activity/api"
	"Activity/config"
	"Activity/models"
	"Activity/storage/mysql"
	"Activity/storage/mysql/repository"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 设置gin模式
	gin.SetMode(cfg.API.Mode)

	// 初始化数据库连接
	db, err := mysql.NewDB(&mysql.Config{
		Host:     cfg.MySQL.Host,
		Port:     cfg.MySQL.Port,
		User:     cfg.MySQL.User,
		Password: cfg.MySQL.Password,
		Database: cfg.MySQL.Database,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 创建仓储实例
	activityRepo := repository.NewActivityRepository(db)

	// 创建服务实例
	activityService := api.NewActivityService(activityRepo)
	gameService := api.NewGameService(activityRepo)

	// 创建处理器
	handler := api.NewHandler(gameService, activityService)

	// 创建路由
	r := gin.Default()

	// 注册Swagger路由
	handler.RegisterSwagger(r)

	// 注册API路由
	handler.RegisterRoutes(r)

	// 启动服务
	addr := fmt.Sprintf(":%d", cfg.API.Port)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
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
