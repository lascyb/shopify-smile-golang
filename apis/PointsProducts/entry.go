package PointsProducts

import "smile/contract"

type PointsProducts struct {
	client contract.SmileClient
}

func NewPointsProducts(client contract.SmileClient) *PointsProducts {
	return &PointsProducts{client: client}
}
