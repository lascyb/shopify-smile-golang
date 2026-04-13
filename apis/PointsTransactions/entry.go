package PointsTransactions

import "smile/contract"

type PointsTransactions struct {
	client contract.SmileClient
}

func NewPointsTransactions(client contract.SmileClient) *PointsTransactions {
	return &PointsTransactions{
		client: client,
	}
}
