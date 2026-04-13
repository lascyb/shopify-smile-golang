package PointsProducts

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

const pathPointsProductsV1 = "/v1/points_products"

type ListPointsProductsInput struct {
	ExchangeType *string
	PageSize     *int
	Page         *int
}

type listPointsProductsResponse struct {
	PointsProducts []PointsProduct `json:"points_products"`
}

// ListPointsProducts 调用 Smile「List points products」：GET /v1/points_products。
func (p *PointsProducts) ListPointsProducts(ctx context.Context, in ListPointsProductsInput) ([]PointsProduct, error) {
	if p == nil || p.client == nil {
		return nil, fmt.Errorf("smile/apis/PointsProducts: client 未初始化")
	}
	path := pathPointsProductsV1
	query := in.encode()
	if query != "" {
		path += "?" + query
	}
	var resp listPointsProductsResponse
	if err := p.client.GetJSON(ctx, path, &resp); err != nil {
		return nil, err
	}
	return resp.PointsProducts, nil
}

func (in *ListPointsProductsInput) encode() string {
	q := url.Values{}
	if in.ExchangeType != nil {
		q.Set("exchange_type", *in.ExchangeType)
	}
	if in.PageSize != nil {
		q.Set("page_size", strconv.Itoa(*in.PageSize))
	}
	if in.Page != nil {
		q.Set("page", strconv.Itoa(*in.Page))
	}
	return q.Encode()
}
