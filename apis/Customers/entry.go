package Customers

import "github.com/lascyb/shopify-smile-golang/contract"

type Customers struct {
	client contract.SmileClient
}

func NewCustomers(client contract.SmileClient) *Customers {
	return &Customers{
		client: client,
	}
}
