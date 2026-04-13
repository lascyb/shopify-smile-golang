package PointsSettings

import (
	"context"
	"fmt"
)

const pathPointsSettingsV1 = "/v1/points_settings"

type PointsLabel struct {
	One   string `json:"one"`
	Other string `json:"other"`
}

type PointsSettingsData struct {
	PointsLabel PointsLabel `json:"points_label"`
}

type getPointsSettingsResponse struct {
	PointsSettings *PointsSettingsData `json:"points_settings"`
}

// GetPointsSettings 调用 Smile「Get points settings」：GET /v1/points_settings。
func (p *PointsSettings) GetPointsSettings(ctx context.Context) (*PointsSettingsData, error) {
	if p == nil || p.client == nil {
		return nil, fmt.Errorf("smile/apis/PointsSettings: client 未初始化")
	}
	var resp getPointsSettingsResponse
	if err := p.client.GetJSON(ctx, pathPointsSettingsV1, &resp); err != nil {
		return nil, err
	}
	if resp.PointsSettings == nil {
		return nil, fmt.Errorf("smile/apis/PointsSettings: 响应缺少 points_settings")
	}
	return resp.PointsSettings, nil
}
