package EarningRules

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

const pathEarningRulesV1 = "/v1/earning_rules"

type ListEarningRulesInput struct {
	Limit  *int
	Cursor *string
}

type EarningRule struct {
	ID                     int64          `json:"id"`
	Name                   string         `json:"name"`
	ImageURL               string         `json:"image_url"`
	ActionText             *string        `json:"action_text"`
	ActionURL              *string        `json:"action_url"`
	RestrictedToVipTierIDs []int64        `json:"restricted_to_vip_tier_ids"`
	Reward                 map[string]any `json:"reward"`
	RewardValue            map[string]any `json:"reward_value"`
	EarningLimit           map[string]any `json:"earning_limit"`
}

type PaginationMetadata struct {
	NextCursor     *string `json:"next_cursor"`
	PreviousCursor *string `json:"previous_cursor"`
}

type listEarningRulesResponse struct {
	EarningRules []EarningRule      `json:"earning_rules"`
	Metadata     PaginationMetadata `json:"metadata"`
}

// ListEarningRules 调用 Smile「List earning rules」：GET /v1/earning_rules。
func (e *EarningRules) ListEarningRules(ctx context.Context, in ListEarningRulesInput) ([]EarningRule, *PaginationMetadata, error) {
	if e == nil || e.client == nil {
		return nil, nil, fmt.Errorf("smile/apis/EarningRules: client 未初始化")
	}
	path := pathEarningRulesV1
	query := in.encode()
	if query != "" {
		path += "?" + query
	}
	var resp listEarningRulesResponse
	if err := e.client.GetJSON(ctx, path, &resp); err != nil {
		return nil, nil, err
	}
	return resp.EarningRules, &resp.Metadata, nil
}

func (in *ListEarningRulesInput) encode() string {
	q := url.Values{}
	if in.Limit != nil {
		q.Set("limit", strconv.Itoa(*in.Limit))
	}
	if in.Cursor != nil {
		q.Set("cursor", *in.Cursor)
	}
	return q.Encode()
}
