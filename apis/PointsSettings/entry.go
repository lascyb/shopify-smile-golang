package PointsSettings

import "smile/contract"

type PointsSettings struct {
	client contract.SmileClient
}

func NewPointsSettings(client contract.SmileClient) *PointsSettings {
	return &PointsSettings{client: client}
}
