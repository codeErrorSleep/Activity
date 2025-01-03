package models

import (
	"context"
	"fmt"
)

// 简单活动，只有一个玩法
// 玩法就是发帖，发帖完成之后就可以发奖品了

type CommunityPostGame struct {
	Name  string         `json:"name"`
	Prize PrizeInterface // 以旧换新的奖品
}

func (p CommunityPostGame) Game(parent string) string {
	return parent + p.Name
}

func (p CommunityPostGame) Perform(user User, action ActionInterface) (ResultInterface, error) {
	// TODO: select db
	fmt.Println(user)
	if user.Uid == "" {
		return nil, fmt.Errorf("user is nil")
	}

	// 这里判断一下是否需要发奖
	// 成功后发奖
	if p.Prize != nil {
		err := p.Prize.WinPrize(context.Background(), user)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (p CommunityPostGame) Actions() []ActionInterface {
	return []ActionInterface{}
}

func (p CommunityPostGame) Results() []ResultInterface {
	return []ResultInterface{}
}

func (p CommunityPostGame) ValidateConfig() error {
	//TODO implement me
	if p.Name == "" {
		return fmt.Errorf("name is empty")
	}
	return nil
}

func (p CommunityPostGame) UnmarshalJSON(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}

func (p CommunityPostGame) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}
