// Package pointstransactions 封装 Smile Points Transactions 相关 REST API。
package PointsTransactions

import (
	"context"
	"fmt"
)

const pathPointsTransactionsV1 = "/v1/points_transactions"

// CreatePointsTransactionInput Create a points transaction 请求体（对应文档 points_transaction）。
type CreatePointsTransactionInput struct {
	CustomerID   uint64  `json:"customer_id"`             // 客户 ID，必填
	PointsChange int     `json:"points_change"`           // 积分增减，正为加、负为减，必填
	Description  *string `json:"description,omitempty"`   // 对客户可见的说明，可选
	InternalNote *string `json:"internal_note,omitempty"` // 仅商家可见备注，可选
}

// PointsTransaction Smile 返回的积分交易对象。
type PointsTransaction struct {
	ID           int64   `json:"id"`
	CustomerID   int64   `json:"customer_id"`
	PointsChange int     `json:"points_change"`
	Description  *string `json:"description"`
	InternalNote *string `json:"internal_note"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type createPointsTransactionRequest struct {
	PointsTransaction CreatePointsTransactionInput `json:"points_transaction"`
}

type createPointsTransactionResponse struct {
	PointsTransaction *PointsTransaction `json:"points_transaction"`
}

// CreateAPointsTransaction 调用 Smile「Create a points transaction」：POST /v1/points_transactions。
// 文档：https://dev.smile.io/api/resources/points-transactions/create-points-transaction
func (p *PointsTransactions) CreateAPointsTransaction(ctx context.Context, in CreatePointsTransactionInput) (*PointsTransaction, error) {
	if p == nil || p.client == nil {
		return nil, fmt.Errorf("smile/apis/PointsTransactions: JSONPoster 未初始化")
	}
	if in.CustomerID <= 0 {
		return nil, fmt.Errorf("smile/apis/PointsTransactions: customer_id 须大于 0")
	}
	reqBody := createPointsTransactionRequest{PointsTransaction: in}
	var resp createPointsTransactionResponse
	if err := p.client.PostJSON(ctx, pathPointsTransactionsV1, reqBody, &resp); err != nil {
		return nil, err
	}
	if resp.PointsTransaction == nil {
		return nil, fmt.Errorf("smile/apis/PointsTransactions: 响应缺少 points_transaction")
	}
	return resp.PointsTransaction, nil
}
