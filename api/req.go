package api

import (
	"time"
)

// ActivityRequest 活动请求基类
type ActivityRequest struct {
	ActivityID int64  `json:"activity_id" binding:"required"`
	UserID     string `json:"user_id" binding:"required"`
}

// CreateActivityRequest 创建活动请求
type CreateActivityRequest struct {
	Category string `json:"category" binding:"required"`
	Version  string `json:"version" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Config   string `json:"config" binding:"required"`
	StartAt  int64  `json:"start_at" binding:"required"`
	EndAt    int64  `json:"end_at" binding:"required"`
	Status   int64  `json:"status" binding:"required"`
}

// UpdateActivityRequest 更新活动请求
type UpdateActivityRequest struct {
	ActivityID int64  `json:"activity_id" binding:"required"`
	Config     string `json:"config" binding:"required"`
	Status     int64  `json:"status" binding:"required"`
}

// ParticipateRequest 参与活动请求
type ParticipateRequest struct {
	ActivityRequest
	GameType   string `json:"game_type" binding:"required"`
	GameTarget string `json:"game_target" binding:"required"`
}

// ActivityResponse 活动响应
type ActivityResponse struct {
	ID        int64     `json:"id"`
	Category  string    `json:"category"`
	Version   string    `json:"version"`
	Name      string    `json:"name"`
	Config    string    `json:"config"`
	StartAt   int64     `json:"start_at"`
	EndAt     int64     `json:"end_at"`
	Status    int64     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ParticipationResponse 参与记录响应
type ParticipationResponse struct {
	ID         int64     `json:"id"`
	ActivityID int64     `json:"activity_id"`
	UserID     string    `json:"user_id"`
	GameType   string    `json:"game_type"`
	GameTarget string    `json:"game_target"`
	State      string    `json:"state"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// PrizeResponse 奖品响应
type PrizeResponse struct {
	ID         int64     `json:"id"`
	ActivityID int64     `json:"activity_id"`
	UserID     string    `json:"user_id"`
	PrizeType  string    `json:"prize_type"`
	PrizeID    string    `json:"prize_id"`
	Status     int64     `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

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
