package models

import "context"

// PrizeInterface 奖品的interface
type PrizeInterface interface {
	WinPrize(ctx context.Context, user User) error // 中奖后需要执行的逻辑
	WinProbability() int64                         // 中奖概率
}
