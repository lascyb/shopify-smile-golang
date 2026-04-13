package PointsTransactions

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type ListPointsTransactionsInput struct {
	CustomerID   *int64
	UpdatedAtMin *string
	Limit        *int
	Cursor       *string
}

type PaginationMetadata struct {
	NextCursor     *string `json:"next_cursor"`
	PreviousCursor *string `json:"previous_cursor"`
}

type listPointsTransactionsResponse struct {
	PointsTransactions []PointsTransaction `json:"points_transactions"`
	Metadata           PaginationMetadata  `json:"metadata"`
}

// ListPointsTransactions 调用 Smile「List points transactions」：GET /v1/points_transactions。
func (p *PointsTransactions) ListPointsTransactions(ctx context.Context, in ListPointsTransactionsInput) ([]PointsTransaction, *PaginationMetadata, error) {
	if p == nil || p.client == nil {
		return nil, nil, fmt.Errorf("smile/apis/PointsTransactions: client 未初始化")
	}
	path := pathPointsTransactionsV1
	query := in.encode()
	if query != "" {
		path += "?" + query
	}
	var resp listPointsTransactionsResponse
	if err := p.client.GetJSON(ctx, path, &resp); err != nil {
		return nil, nil, err
	}
	return resp.PointsTransactions, &resp.Metadata, nil
}

func (in *ListPointsTransactionsInput) encode() string {
	q := url.Values{}
	if in.CustomerID != nil {
		q.Set("customer_id", strconv.FormatInt(*in.CustomerID, 10))
	}
	if in.UpdatedAtMin != nil {
		q.Set("updated_at_min", *in.UpdatedAtMin)
	}
	if in.Limit != nil {
		q.Set("limit", strconv.Itoa(*in.Limit))
	}
	if in.Cursor != nil {
		q.Set("cursor", *in.Cursor)
	}
	return q.Encode()
}
