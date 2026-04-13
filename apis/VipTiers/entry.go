package VipTiers

import "github.com/lascyb/shopify-smile-golang/contract"

type VipTiers struct {
	client contract.SmileClient
}

func NewVipTiers(client contract.SmileClient) *VipTiers {
	return &VipTiers{client: client}
}
