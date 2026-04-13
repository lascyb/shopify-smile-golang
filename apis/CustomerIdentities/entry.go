package CustomerIdentities

import "github.com/lascyb/shopify-smile-golang/contract"

type CustomerIdentities struct {
	client contract.SmileClient
}

func NewCustomerIdentities(client contract.SmileClient) *CustomerIdentities {
	return &CustomerIdentities{client: client}
}
