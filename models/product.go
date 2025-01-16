package models

import "context"

// Product 商品
type ProductPrize struct {
	Sku         string `json:"sku"`
	Title       string `json:"title"`
	Probability int64  `json:"probability"` // 中奖概率
}

func (p ProductPrize) WinPrize(ctx context.Context, user User) error {
	return nil
}

func (p ProductPrize) WinProbability() int64 {
	return p.Probability
}
