package PointsTransactions

import "github.com/lascyb/shopify-smile-golang/contract"

type PointsTransactions struct {
	client contract.SmileClient
}

func NewPointsTransactions(client contract.SmileClient) *PointsTransactions {
	return &PointsTransactions{
		client: client,
	}
}
