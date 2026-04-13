package PointsSettings

import "github.com/lascyb/shopify-smile-golang/contract"

type PointsSettings struct {
	client contract.SmileClient
}

func NewPointsSettings(client contract.SmileClient) *PointsSettings {
	return &PointsSettings{client: client}
}
