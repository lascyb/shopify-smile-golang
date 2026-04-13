package RewardFulfillments

import "github.com/lascyb/shopify-smile-golang/contract"

type RewardFulfillments struct {
	client contract.SmileClient
}

func NewRewardFulfillments(client contract.SmileClient) *RewardFulfillments {
	return &RewardFulfillments{client: client}
}
