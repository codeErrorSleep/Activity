package models

import "context"

// 折扣码的奖品

type DiscountCodePrize struct {
	DiscountCode string `json:"discount_code"`
	PriceRuleID  int64  `json:"price_rule_id"`
	Probability  int64  `json:"probability"` // 中奖概率
}

func (p DiscountCodePrize) WinPrize(ctx context.Context, user User) error {
	// 找到一个空的，直接分配给用户J

	//

	return nil
}

func (p DiscountCodePrize) WinProbability() int64 {
	//TODO implement me
	return p.Probability
}
