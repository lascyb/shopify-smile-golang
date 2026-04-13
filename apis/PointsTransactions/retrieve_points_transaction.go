package PointsTransactions

import (
	"context"
	"fmt"
	"strconv"
)

type retrievePointsTransactionResponse struct {
	PointsTransaction *PointsTransaction `json:"points_transaction"`
}

// RetrievePointsTransaction 调用 Smile「Retrieve a points transaction」：GET /v1/points_transactions/{id}。
func (p *PointsTransactions) RetrievePointsTransaction(ctx context.Context, id int64) (*PointsTransaction, error) {
	if p == nil || p.client == nil {
		return nil, fmt.Errorf("smile/apis/PointsTransactions: client 未初始化")
	}
	if id <= 0 {
		return nil, fmt.Errorf("smile/apis/PointsTransactions: id 须大于 0")
	}
	path := pathPointsTransactionsV1 + "/" + strconv.FormatInt(id, 10)
	var resp retrievePointsTransactionResponse
	if err := p.client.GetJSON(ctx, path, &resp); err != nil {
		return nil, err
	}
	if resp.PointsTransaction == nil {
		return nil, fmt.Errorf("smile/apis/PointsTransactions: 响应缺少 points_transaction")
	}
	return resp.PointsTransaction, nil
}
