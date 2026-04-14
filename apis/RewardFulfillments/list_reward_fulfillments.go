package RewardFulfillments

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const pathRewardFulfillmentsV1 = "/v1/reward_fulfillments"

type ListRewardFulfillmentsInput struct {
	CustomerID        *int64
	FulfillmentStatus *string
	UsageStatus       *string
	UpdatedAtMin      *string
	Limit             *int
	Cursor            *string
}

type RewardFulfillment struct {
	ID                 int64     `json:"id"`
	Name               string    `json:"name"`
	Code               string    `json:"code"`
	CustomerID         int64     `json:"customer_id"`
	FulfillmentStatus  string    `json:"fulfillment_status"`
	ImageURL           string    `json:"image_url"`
	ActionText         *string   `json:"action_text"`
	ActionURL          *string   `json:"action_url"`
	UsageInstructions  *string   `json:"usage_instructions"`
	TermsAndConditions *string   `json:"terms_and_conditions"`
	ExpiresAt          *string   `json:"expires_at"`
	UsageStatus        string    `json:"usage_status"`
	UsedAt             *string   `json:"used_at"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type PaginationMetadata struct {
	NextCursor     *string `json:"next_cursor"`
	PreviousCursor *string `json:"previous_cursor"`
}

type listRewardFulfillmentsResponse struct {
	RewardFulfillments []RewardFulfillment `json:"reward_fulfillments"`
	Metadata           PaginationMetadata  `json:"metadata"`
}

// ListRewardFulfillments 调用 Smile「List reward fulfillments」：GET /v1/reward_fulfillments。
func (r *RewardFulfillments) ListRewardFulfillments(ctx context.Context, in ListRewardFulfillmentsInput) ([]RewardFulfillment, *PaginationMetadata, error) {
	if r == nil || r.client == nil {
		return nil, nil, fmt.Errorf("smile/apis/RewardFulfillments: client 未初始化")
	}
	path := pathRewardFulfillmentsV1
	query := in.encode()
	if query != "" {
		path += "?" + query
	}
	var resp listRewardFulfillmentsResponse
	if err := r.client.GetJSON(ctx, path, &resp); err != nil {
		return nil, nil, err
	}
	return resp.RewardFulfillments, &resp.Metadata, nil
}

func (in *ListRewardFulfillmentsInput) encode() string {
	q := url.Values{}
	if in.CustomerID != nil {
		q.Set("customer_id", strconv.FormatInt(*in.CustomerID, 10))
	}
	if in.FulfillmentStatus != nil {
		q.Set("fulfillment_status", *in.FulfillmentStatus)
	}
	if in.UsageStatus != nil {
		q.Set("usage_status", *in.UsageStatus)
	}
	if in.UpdatedAtMin != nil {
		q.Set("updated_at_min", *in.UpdatedAtMin)
	}
	if in.Limit != nil {
		q.Set("limit", strconv.Itoa(*in.Limit))
	}
	if in.Cursor != nil {
		q.Set("cursor", *in.Cursor)
	}
	return q.Encode()
}
