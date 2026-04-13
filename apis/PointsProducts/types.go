package PointsProducts

type Reward struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type PointsProduct struct {
	ID                            int64  `json:"id"`
	ExchangeType                  string `json:"exchange_type"`
	ExchangeDescription           string `json:"exchange_description"`
	PointsPrice                   *int   `json:"points_price"`
	VariablePointsStep            *int   `json:"variable_points_step"`
	VariablePointsStepRewardValue *int   `json:"variable_points_step_reward_value"`
	VariablePointsMin             *int   `json:"variable_points_min"`
	VariablePointsMax             *int   `json:"variable_points_max"`
	Reward                        Reward `json:"reward"`
	CreatedAt                     string `json:"created_at"`
	UpdatedAt                     string `json:"updated_at"`
}

type RewardFulfillment struct {
	ID                 int64   `json:"id"`
	Name               string  `json:"name"`
	Code               string  `json:"code"`
	CustomerID         int64   `json:"customer_id"`
	FulfillmentStatus  string  `json:"fulfillment_status"`
	ImageURL           string  `json:"image_url"`
	ActionText         *string `json:"action_text"`
	ActionURL          *string `json:"action_url"`
	UsageInstructions  *string `json:"usage_instructions"`
	TermsAndConditions *string `json:"terms_and_conditions"`
	ExpiresAt          *string `json:"expires_at"`
	UsageStatus        string  `json:"usage_status"`
	UsedAt             *string `json:"used_at"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
}

type PointsPurchase struct {
	ID                int64             `json:"id"`
	CustomerID        int64             `json:"customer_id"`
	PointsProductID   int64             `json:"points_product_id"`
	PointsSpent       int               `json:"points_spent"`
	RewardFulfillment RewardFulfillment `json:"reward_fulfillment"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
}
