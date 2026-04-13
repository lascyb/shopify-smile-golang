package CustomerIdentities

import "smile/contract"

type CustomerIdentities struct {
	client contract.SmileClient
}

func NewCustomerIdentities(client contract.SmileClient) *CustomerIdentities {
	return &CustomerIdentities{client: client}
}
