package Activities

import "github.com/lascyb/shopify-smile-golang/contract"

type Activities struct {
	client contract.SmileClient
}

func NewActivities(client contract.SmileClient) *Activities {
	return &Activities{client: client}
}
