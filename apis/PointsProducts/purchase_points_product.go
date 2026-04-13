package PointsProducts

import (
	"context"
	"fmt"
	"strconv"
)

type PurchasePointsProductInput struct {
	CustomerID    uint64 `json:"customer_id"`
	PointsToSpend *int   `json:"points_to_spend,omitempty"`
}

type purchasePointsProductResponse struct {
	PointsPurchase *PointsPurchase `json:"points_purchase"`
}

// PurchasePointsProduct 调用 Smile「Purchase a points product」：POST /v1/points_products/{id}/purchase。
func (p *PointsProducts) PurchasePointsProduct(ctx context.Context, id int64, in PurchasePointsProductInput) (*PointsPurchase, error) {
	if p == nil || p.client == nil {
		return nil, fmt.Errorf("smile/apis/PointsProducts: client 未初始化")
	}
	if id <= 0 {
		return nil, fmt.Errorf("smile/apis/PointsProducts: id 须大于 0")
	}
	if in.CustomerID == 0 {
		return nil, fmt.Errorf("smile/apis/PointsProducts: customer_id 须大于 0")
	}
	path := pathPointsProductsV1 + "/" + strconv.FormatInt(id, 10) + "/purchase"
	var resp purchasePointsProductResponse
	if err := p.client.PostJSON(ctx, path, in, &resp); err != nil {
		return nil, err
	}
	if resp.PointsPurchase == nil {
		return nil, fmt.Errorf("smile/apis/PointsProducts: 响应缺少 points_purchase")
	}
	return resp.PointsPurchase, nil
}
