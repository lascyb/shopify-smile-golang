package CustomerIdentities

import (
	"context"
	"fmt"
)

const pathCustomerIdentitiesCreateOrUpdateV1 = "/v1/customer_identities/create_or_update"

type CreateOrUpdateIdentityInput struct {
	FirstName  *string        `json:"first_name,omitempty"`
	LastName   *string        `json:"last_name,omitempty"`
	Email      string         `json:"email"`
	DistinctID string         `json:"distinct_id"`
	Properties map[string]any `json:"properties,omitempty"`
}

type CustomerIdentity struct {
	ID         int64          `json:"id"`
	FirstName  *string        `json:"first_name"`
	LastName   *string        `json:"last_name"`
	Email      string         `json:"email"`
	DistinctID string         `json:"distinct_id"`
	Properties map[string]any `json:"properties"`
	CustomerID int64          `json:"customer_id"`
}

type createOrUpdateIdentityRequest struct {
	CustomerIdentity CreateOrUpdateIdentityInput `json:"customer_identity"`
}

type createOrUpdateIdentityResponse struct {
	CustomerIdentity *CustomerIdentity `json:"customer_identity"`
}

// CreateOrUpdateIdentity 调用 Smile「Create or update a customer identity」：POST /v1/customer_identities/create_or_update。
func (c *CustomerIdentities) CreateOrUpdateIdentity(ctx context.Context, in CreateOrUpdateIdentityInput) (*CustomerIdentity, error) {
	if c == nil || c.client == nil {
		return nil, fmt.Errorf("smile/apis/CustomerIdentities: client 未初始化")
	}
	if in.Email == "" {
		return nil, fmt.Errorf("smile/apis/CustomerIdentities: email 不能为空")
	}
	if in.DistinctID == "" {
		return nil, fmt.Errorf("smile/apis/CustomerIdentities: distinct_id 不能为空")
	}
	req := createOrUpdateIdentityRequest{CustomerIdentity: in}
	var resp createOrUpdateIdentityResponse
	if err := c.client.PostJSON(ctx, pathCustomerIdentitiesCreateOrUpdateV1, req, &resp); err != nil {
		return nil, err
	}
	if resp.CustomerIdentity == nil {
		return nil, fmt.Errorf("smile/apis/CustomerIdentities: 响应缺少 customer_identity")
	}
	return resp.CustomerIdentity, nil
}
