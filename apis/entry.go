package apis

import (
	"github.com/lascyb/shopify-smile-golang/apis/Activities"
	"github.com/lascyb/shopify-smile-golang/apis/CustomerIdentities"
	"github.com/lascyb/shopify-smile-golang/apis/Customers"
	"github.com/lascyb/shopify-smile-golang/apis/EarningRules"
	"github.com/lascyb/shopify-smile-golang/apis/PointsProducts"
	"github.com/lascyb/shopify-smile-golang/apis/PointsSettings"
	"github.com/lascyb/shopify-smile-golang/apis/PointsTransactions"
	"github.com/lascyb/shopify-smile-golang/apis/RewardFulfillments"
	"github.com/lascyb/shopify-smile-golang/apis/VipTiers"
	"github.com/lascyb/shopify-smile-golang/contract"
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
