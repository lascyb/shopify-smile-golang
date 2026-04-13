package EarningRules

import "smile/contract"

type EarningRules struct {
	client contract.SmileClient
}

func NewEarningRules(client contract.SmileClient) *EarningRules {
	return &EarningRules{client: client}
}
