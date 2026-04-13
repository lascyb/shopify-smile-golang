package VipTiers

import "smile/contract"

type VipTiers struct {
	client contract.SmileClient
}

func NewVipTiers(client contract.SmileClient) *VipTiers {
	return &VipTiers{client: client}
}
