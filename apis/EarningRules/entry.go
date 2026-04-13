package EarningRules

import "github.com/lascyb/shopify-smile-golang/contract"

type EarningRules struct {
	client contract.SmileClient
}

func NewEarningRules(client contract.SmileClient) *EarningRules {
	return &EarningRules{client: client}
}
