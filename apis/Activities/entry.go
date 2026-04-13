package Activities

import "smile/contract"

type Activities struct {
	client contract.SmileClient
}

func NewActivities(client contract.SmileClient) *Activities {
	return &Activities{client: client}
}
