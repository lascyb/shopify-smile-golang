package PointsProducts

import "github.com/lascyb/shopify-smile-golang/contract"

type PointsProducts struct {
	client contract.SmileClient
}

func NewPointsProducts(client contract.SmileClient) *PointsProducts {
	return &PointsProducts{client: client}
}
