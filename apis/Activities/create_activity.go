package Activities

import (
	"context"
	"fmt"
)

const pathActivitiesV1 = "/v1/activities"

type CreateActivityInput struct {
	Token             string  `json:"token"`
	CustomerID        *uint64 `json:"customer_id,omitempty"`
	CustomerEmail     *string `json:"customer_email,omitempty"`
	DistinctID        *string `json:"distinct_id,omitempty"`
	CreatedOnOriginAt *string `json:"created_on_origin_at,omitempty"`
}

type Activity struct {
	ID                int64   `json:"id"`
	CustomerID        int64   `json:"customer_id"`
	Token             string  `json:"token"`
	DistinctID        *string `json:"distinct_id"`
	CreatedOnOriginAt *string `json:"created_on_origin_at"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}

type createActivityRequest struct {
	Activity CreateActivityInput `json:"activity"`
}

type createActivityResponse struct {
	Activity *Activity `json:"activity"`
}

// CreateActivity 调用 Smile「Create an activity」：POST /v1/activities。
func (a *Activities) CreateActivity(ctx context.Context, in CreateActivityInput) (*Activity, error) {
	if a == nil || a.client == nil {
		return nil, fmt.Errorf("smile/apis/Activities: client 未初始化")
	}
	if in.Token == "" {
		return nil, fmt.Errorf("smile/apis/Activities: token 不能为空")
	}
	if in.CustomerID == nil && in.CustomerEmail == nil {
		return nil, fmt.Errorf("smile/apis/Activities: customer_id 与 customer_email 必须二选一")
	}
	req := createActivityRequest{Activity: in}
	var resp createActivityResponse
	if err := a.client.PostJSON(ctx, pathActivitiesV1, req, &resp); err != nil {
		return nil, err
	}
	if resp.Activity == nil {
		return nil, fmt.Errorf("smile/apis/Activities: 响应缺少 activity")
	}
	return resp.Activity, nil
}
