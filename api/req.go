package api

import (
	"time"
)

// 请求结构体
type (
	// ParticipateGameReq 参与玩法请求
	ParticipateGameReq struct {
		ActivityID string `json:"activity_id" binding:"required"` // 活动ID
		GameName   string `json:"game_name" binding:"required"`   // 玩法名称
		GameType   string `json:"game_type" binding:"required"`   // 玩法类型：post/checkin
		PostID     string `json:"post_id"`                        // 发帖ID（社区发帖玩法需要）
	}

	// GetGameStatusReq 获取玩法状态请求
	GetGameStatusReq struct {
		ActivityID string `json:"activity_id" binding:"required"` // 活动ID
		GameName   string `json:"game_name" binding:"required"`   // 玩法名称
	}

	// GetUserPrizeReq 获取用户奖品请求
	GetUserPrizeReq struct {
		ActivityID string `json:"activity_id" binding:"required"` // 活动ID
		GameName   string `json:"game_name" binding:"required"`   // 玩法名称
	}
)

// 响应结构体
type (
	// BaseResp 基础响应
	BaseResp struct {
		Code    int         `json:"code"`    // 错误码
		Message string      `json:"message"` // 错误信息
		Data    interface{} `json:"data"`    // 响应数据
	}

	// GameStatusResp 玩法状态响应
	GameStatusResp struct {
		GameState string    `json:"game_state"` // 玩法状态
		UserState string    `json:"user_state"` // 用户状态
		StartTime time.Time `json:"start_time"` // 开始时间
		EndTime   time.Time `json:"end_time"`   // 结束时间
		RemainNum int64     `json:"remain_num"` // 剩余奖品数量
		TotalNum  int64     `json:"total_num"`  // 总奖品数量
	}

	// PrizeInfo 奖品信息
	PrizeInfo struct {
		Type         string `json:"type"`          // 奖品类型
		DiscountCode string `json:"discount_code"` // 折扣码（折扣码类型）
		PriceRuleID  int64  `json:"price_rule_id"` // 价格规则ID（折扣码类型）
		SKU          string `json:"sku"`           // 商品SKU（商品类型）
		Title        string `json:"title"`         // 商品标题（商品类型）
	}

	// UserPrizeResp 用户奖品响应
	UserPrizeResp struct {
		Prize     PrizeInfo `json:"prize"`      // 奖品信息
		CreatedAt time.Time `json:"created_at"` // 获得时间
	}
)
