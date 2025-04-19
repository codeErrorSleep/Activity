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
// @Description 创建活动请求参数
type CreateActivityRequest struct {
	// @Description 活动名称
	Name string `json:"name" binding:"required"`
	// @Description 活动类型
	Category string `json:"category" binding:"required"`
	// @Description 活动版本
	Version string `json:"version" binding:"required"`
	// @Description 活动开始时间
	StartAt int64 `json:"start_at" binding:"required"`
	// @Description 活动结束时间
	EndAt int64 `json:"end_at" binding:"required"`
	// @Description 活动状态
	Status int `json:"status" binding:"required"`
}

// CreateActivityResponse 创建活动响应
// @Description 创建活动响应数据
type CreateActivityResponse struct {
	// @Description 活动ID
	ID int64 `json:"id"`
}

// UpdateActivityRequest 更新活动请求
// @Description 更新活动请求参数
type UpdateActivityRequest struct {
	// @Description 活动名称
	Name string `json:"name"`
	// @Description 活动类型
	Category string `json:"category"`
	// @Description 活动版本
	Version string `json:"version"`
	// @Description 活动开始时间
	StartAt int64 `json:"start_at"`
	// @Description 活动结束时间
	EndAt int64 `json:"end_at"`
	// @Description 活动状态
	Status int `json:"status"`
}

// UpdateActivityResponse 更新活动响应
// @Description 更新活动响应数据
type UpdateActivityResponse struct {
	// @Description 是否更新成功
	Success bool `json:"success"`
}

// GetActivityResponse 获取活动响应
// @Description 获取活动响应数据
type GetActivityResponse struct {
	// @Description 活动ID
	ID int64 `json:"id"`
	// @Description 活动名称
	Name string `json:"name"`
	// @Description 活动类型
	Category string `json:"category"`
	// @Description 活动版本
	Version string `json:"version"`
	// @Description 活动开始时间
	StartAt int64 `json:"start_at"`
	// @Description 活动结束时间
	EndAt int64 `json:"end_at"`
	// @Description 活动状态
	Status int `json:"status"`
}

// ParticipateRequest 参与活动请求
// @Description 参与活动请求参数
type ParticipateRequest struct {
	// @Description 用户ID
	UserID string `json:"user_id" binding:"required"`
}

// ParticipateResponse 参与活动响应
// @Description 参与活动响应数据
type ParticipateResponse struct {
	// @Description 参与ID
	ParticipationID int64 `json:"participation_id"`
}

// GetParticipationResponse 获取参与记录响应
// @Description 获取参与记录响应数据
type GetParticipationResponse struct {
	// @Description 参与ID
	ParticipationID int64 `json:"participation_id"`
	// @Description 用户ID
	UserID string `json:"user_id"`
	// @Description 参与时间
	CreatedAt int64 `json:"created_at"`
}

// ParticipateGameReq 参与玩法请求
// @Description 参与玩法请求参数
type ParticipateGameReq struct {
	// @Description 活动ID
	ActivityID string `json:"activity_id" binding:"required"`
	// @Description 玩法名称
	GameName string `json:"game_name" binding:"required"`
	// @Description 用户ID
	UserID string `json:"user_id" binding:"required"`
}

// ParticipateGameResponse 参与玩法响应
// @Description 参与玩法响应数据
type ParticipateGameResponse struct {
	// @Description 是否参与成功
	Success bool `json:"success"`
	// @Description 是否获得奖品
	HasPrize bool `json:"has_prize"`
	// @Description 奖品信息
	Prize *PrizeInfo `json:"prize,omitempty"`
}

// GetGameStatusReq 获取玩法状态请求
// @Description 获取玩法状态请求参数
type GetGameStatusReq struct {
	// @Description 活动ID
	ActivityID string `form:"activity_id" binding:"required"`
	// @Description 玩法名称
	GameName string `form:"game_name" binding:"required"`
}

// GetGameStatusResponse 获取玩法状态响应
// @Description 获取玩法状态响应数据
type GetGameStatusResponse struct {
	// @Description 玩法状态
	Status string `json:"status"`
	// @Description 参与次数
	ParticipationCount int `json:"participation_count"`
}

// GetUserPrizeReq 获取用户奖品请求
// @Description 获取用户奖品请求参数
type GetUserPrizeReq struct {
	// @Description 活动ID
	ActivityID string `form:"activity_id" binding:"required"`
	// @Description 玩法名称
	GameName string `form:"game_name" binding:"required"`
}

// GetUserPrizeResponse 获取用户奖品响应
// @Description 获取用户奖品响应数据
type GetUserPrizeResponse struct {
	// @Description 奖品列表
	Prizes []*PrizeInfo `json:"prizes"`
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
