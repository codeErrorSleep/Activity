package constant

// API 路径常量
const (
	// 活动相关
	ActivityBasePath = "/api/v1/activity"
	// 玩法相关
	GameBasePath = "/api/v1/game"
)

// 错误码
const (
	// 系统错误码
	ErrSystem = 10000
	// 参数错误
	ErrInvalidParam = 10001
	// 活动不存在
	ErrActivityNotFound = 10002
	// 活动已结束
	ErrActivityEnded = 10003
	// 活动未开始
	ErrActivityNotStarted = 10004
	// 玩法不存在
	ErrGameNotFound = 10005
	// 玩法已关闭
	ErrGameClosed = 10006
	// 用户已参与
	ErrUserAlreadyParticipated = 10007
	// 奖品库存不足
	ErrPrizeStockEmpty = 10008
	// 用户未发帖
	ErrUserNotPosted = 10009
	// 用户未签到
	ErrUserNotCheckedIn = 10010
)

// 错误消息
const (
	ErrMsgSystem                  = "系统错误"
	ErrMsgInvalidParam            = "参数错误"
	ErrMsgActivityNotFound        = "活动不存在"
	ErrMsgActivityEnded           = "活动已结束"
	ErrMsgActivityNotStarted      = "活动未开始"
	ErrMsgGameNotFound            = "玩法不存在"
	ErrMsgGameClosed              = "玩法已关闭"
	ErrMsgUserAlreadyParticipated = "用户已参与"
	ErrMsgPrizeStockEmpty         = "奖品库存不足"
	ErrMsgUserNotPosted           = "用户未发帖"
	ErrMsgUserNotCheckedIn        = "用户未签到"
)
