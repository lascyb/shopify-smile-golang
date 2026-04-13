package Customers

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type RetrieveCustomerInput struct {
	Include *string
}

type retrieveCustomerResponse struct {
	Customer *Customer `json:"customer"`
}

// RetrieveCustomer 调用 Smile「Retrieve a customer」：GET /v1/customers/{id}。
func (c *Customers) RetrieveCustomer(ctx context.Context, id int64, in RetrieveCustomerInput) (*Customer, error) {
	if c == nil || c.client == nil {
		return nil, fmt.Errorf("smile/apis/Customers: client 未初始化")
	}
	if id <= 0 {
		return nil, fmt.Errorf("smile/apis/Customers: id 须大于 0")
	}
	path := pathCustomersV1 + "/" + strconv.FormatInt(id, 10)
	if in.Include != nil {
		q := url.Values{}
		q.Set("include", *in.Include)
		path = path + "?" + q.Encode()
	}
	var resp retrieveCustomerResponse
	if err := c.client.GetJSON(ctx, path, &resp); err != nil {
		return nil, err
	}
	if resp.Customer == nil {
		return nil, fmt.Errorf("smile/apis/Customers: 响应缺少 customer")
	}
	return resp.Customer, nil
}
