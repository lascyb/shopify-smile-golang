package PointsProducts

import (
	"context"
	"fmt"
	"strconv"
)

type retrievePointsProductResponse struct {
	PointsProduct *PointsProduct `json:"points_product"`
}

// RetrievePointsProduct 调用 Smile「Retrieve a points product」：GET /v1/points_products/{id}。
func (p *PointsProducts) RetrievePointsProduct(ctx context.Context, id int64) (*PointsProduct, error) {
	if p == nil || p.client == nil {
		return nil, fmt.Errorf("smile/apis/PointsProducts: client 未初始化")
	}
	if id <= 0 {
		return nil, fmt.Errorf("smile/apis/PointsProducts: id 须大于 0")
	}
	var resp retrievePointsProductResponse
	if err := p.client.GetJSON(ctx, pathPointsProductsV1+"/"+strconv.FormatInt(id, 10), &resp); err != nil {
		return nil, err
	}
	if resp.PointsProduct == nil {
		return nil, fmt.Errorf("smile/apis/PointsProducts: 响应缺少 points_product")
	}
	return resp.PointsProduct, nil
}
