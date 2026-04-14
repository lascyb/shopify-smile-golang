package VipTiers

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

const pathVipTiersV1 = "/v1/vip_tiers"

type ListVipTiersInput struct {
	Include *string
}

type Reward struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type VipTierPerk struct {
	Name string `json:"name"`
}

type VipTier struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	ImageURL     string        `json:"image_url"`
	Milestone    float64       `json:"milestone"`
	Perks        []VipTierPerk `json:"perks"`
	EntryRewards []Reward      `json:"entry_rewards"`
}

type listVipTiersResponse struct {
	VipTiers []VipTier `json:"vip_tiers"`
}

// ListVipTiers 调用 Smile「List VIP tiers」：GET /v1/vip_tiers。
func (v *VipTiers) ListVipTiers(ctx context.Context, in ListVipTiersInput) ([]VipTier, error) {
	if v == nil || v.client == nil {
		return nil, fmt.Errorf("smile/apis/VipTiers: client 未初始化")
	}
	path := pathVipTiersV1
	if in.Include != nil {
		q := url.Values{}
		q.Set("include", *in.Include)
		path += "?" + q.Encode()
	}
	var resp listVipTiersResponse
	if err := v.client.GetJSON(ctx, path, &resp); err != nil {
		return nil, err
	}
	return resp.VipTiers, nil
}
