package RewardFulfillments

import "smile/contract"

type RewardFulfillments struct {
	client contract.SmileClient
}

func NewRewardFulfillments(client contract.SmileClient) *RewardFulfillments {
	return &RewardFulfillments{client: client}
}
