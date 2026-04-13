package Customers

import "smile/contract"

type Customers struct {
	client contract.SmileClient
}

func NewCustomers(client contract.SmileClient) *Customers {
	return &Customers{
		client: client,
	}
}
