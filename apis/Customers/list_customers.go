// Package Customers 封装 Smile Customers 相关 REST API。
package Customers

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const pathCustomersV1 = "/v1/customers"

// ListCustomersInput List customers 请求参数（均为可选）。
type ListCustomersInput struct {
	Email        *string // 按邮箱筛选
	State        *string // {candidate:候选 member:会员 disabled:已禁用}
	UpdatedAtMin *string // RFC3339 格式，仅返回该时间之后更新的客户
	Limit        *int    // 每页数量，1-250，默认 50
	Cursor       *string // 分页游标
	Include      *string // 关联对象，如 "vip_status"
}

// Customer Smile 客户对象。
type Customer struct {
	ID            int64   `json:"id"`
	FirstName     *string `json:"first_name"`
	LastName      *string `json:"last_name"`
	Email         string  `json:"email"`
	State         string  `json:"state"`
	DateOfBirth   *string `json:"date_of_birth"`
	PointsBalance int     `json:"points_balance"`
	ReferralURL   string  `json:"referral_url"`
	VipTierID     *int64  `json:"vip_tier_id"`
	VipStatus     *struct {
		VipTierID            *int64     `json:"vip_tier_id"`
		VipTierExpiresAt     *time.Time `json:"vip_tier_expires_at"`
		ProgressValue        *float64   `json:"progress_value"`
		CurrentVipPeriodEnd  *time.Time `json:"current_vip_period_end"`
		DeltaToRetainVipTier *float64   `json:"delta_to_retain_vip_tier"`
		NextVipTierID        *int64     `json:"next_vip_tier_id"`
		DeltaToNextVipTier   *float64   `json:"delta_to_next_vip_tier"`
	} `json:"vip_status,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PaginationMetadata 分页元数据。
type PaginationMetadata struct {
	NextCursor     *string `json:"next_cursor"`
	PreviousCursor *string `json:"previous_cursor"`
}

type listCustomersResponse struct {
	Customers []Customer         `json:"customers"`
	Metadata  PaginationMetadata `json:"metadata"`
}

// ListCustomers 调用 Smile「List customers」：GET /v1/customers。
// 文档：https://dev.smile.io/api/resources/customers/list-customers
func (c *Customers) ListCustomers(ctx context.Context, in ListCustomersInput) ([]Customer, *PaginationMetadata, error) {
	if c == nil || c.client == nil {
		return nil, nil, fmt.Errorf("smile/apis/Customers: client 未初始化")
	}

	path := pathCustomersV1 + "?" + in.encode()

	var resp listCustomersResponse
	if err := c.client.GetJSON(ctx, path, &resp); err != nil {
		return nil, nil, err
	}
	return resp.Customers, &resp.Metadata, nil
}

func (in *ListCustomersInput) encode() string {
	q := url.Values{}
	if in.Email != nil {
		q.Set("email", *in.Email)
	}
	if in.State != nil {
		q.Set("state", *in.State)
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
	if in.Include != nil {
		q.Set("include", *in.Include)
	}
	return q.Encode()
}
