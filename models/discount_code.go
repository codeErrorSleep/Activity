package models

import (
	"context"
	"fmt"
)

// 折扣码的奖品

// DiscountCodePrize 折扣码奖品
type DiscountCodePrize struct {
	DiscountCode string `json:"discount_code"` // 折扣码前缀
	PriceRuleID  int64  `json:"price_rule_id"` // 价格规则ID
	Probability  int64  `json:"probability"`   // 中奖概率
	TotalNum     int64  `json:"total_num"`     // 总数量
	RemainNum    int64  `json:"remain_num"`    // 剩余数量
}

func (p DiscountCodePrize) WinPrize(ctx context.Context, user User) error {
	// 1. 检查库存
	if p.RemainNum <= 0 {
		return fmt.Errorf("prize stock is empty")
	}

	// 2. 找到一个空的，直接分配给用户
	// TODO: 实现折扣码分配逻辑

	// 3. 更新库存
	p.RemainNum--
	return nil
}

func (p DiscountCodePrize) WinProbability() int64 {
	return p.Probability
}
