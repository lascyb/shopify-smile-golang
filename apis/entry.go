package apis

import (
	"smile/apis/Activities"
	"smile/apis/CustomerIdentities"
	"smile/apis/Customers"
	"smile/apis/EarningRules"
	"smile/apis/PointsProducts"
	"smile/apis/PointsSettings"
	"smile/apis/PointsTransactions"
	"smile/apis/RewardFulfillments"
	"smile/apis/VipTiers"
	"smile/contract"
)

type Apis struct {
	client             contract.SmileClient
	Activities         *Activities.Activities
	CustomerIdentities *CustomerIdentities.CustomerIdentities
	Customers          *Customers.Customers
	EarningRules       *EarningRules.EarningRules
	PointsProducts     *PointsProducts.PointsProducts
	PointsSettings     *PointsSettings.PointsSettings
	PointsTransactions *PointsTransactions.PointsTransactions
	RewardFulfillments *RewardFulfillments.RewardFulfillments
	VipTiers           *VipTiers.VipTiers
}

func NewApis(client contract.SmileClient) *Apis {
	return &Apis{
		client:             client,
		Activities:         Activities.NewActivities(client),
		CustomerIdentities: CustomerIdentities.NewCustomerIdentities(client),
		Customers:          Customers.NewCustomers(client),
		EarningRules:       EarningRules.NewEarningRules(client),
		PointsProducts:     PointsProducts.NewPointsProducts(client),
		PointsSettings:     PointsSettings.NewPointsSettings(client),
		PointsTransactions: PointsTransactions.NewPointsTransactions(client),
		RewardFulfillments: RewardFulfillments.NewRewardFulfillments(client),
		VipTiers:           VipTiers.NewVipTiers(client),
	}
}
