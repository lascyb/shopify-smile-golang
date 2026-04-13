# shopify-smile-golang

`shopify-smile-golang` 是面向 Shopify + Smile.io 场景的 Go SDK，用于封装 Smile.io REST API，统一 Bearer 鉴权、错误处理与资源接口调用。

## 安装

```bash
go get github.com/lascyb/shopify-smile-golang
```

## 快速开始

```go
package main

import (
	"context"
	"fmt"

	smile "github.com/lascyb/shopify-smile-golang"
    options "github.com/lascyb/shopify-smile-golang/options"
)

func main() {
	client, err := smile.NewClient(
		"your_private_api_key",
		options.WithBaseURL("https://api.smile.io"),
		options.WithRateLimitRetry(3),
	)
	if err != nil {
		panic(err)
	}

	customers, _, err := client.Apis.Customers.ListCustomers(context.Background(), struct {
		Email        *string
		State        *string
		UpdatedAtMin *string
		Limit        *int
		Cursor       *string
		Include      *string
	}{})
	if err != nil {
		panic(err)
	}

	fmt.Println("customers:", len(customers))
}
```

## 已封装接口

- `Activities`
  - `CreateActivity` -> `POST /v1/activities`
- `CustomerIdentities`
  - `CreateOrUpdateIdentity` -> `POST /v1/customer_identities/create_or_update`
- `Customers`
  - `ListCustomers` -> `GET /v1/customers`
  - `RetrieveCustomer` -> `GET /v1/customers/{id}`
- `EarningRules`
  - `ListEarningRules` -> `GET /v1/earning_rules`
- `PointsProducts`
  - `ListPointsProducts` -> `GET /v1/points_products`
  - `RetrievePointsProduct` -> `GET /v1/points_products/{id}`
  - `PurchasePointsProduct` -> `POST /v1/points_products/{id}/purchase`
- `PointsSettings`
  - `GetPointsSettings` -> `GET /v1/points_settings`
- `PointsTransactions`
  - `CreateAPointsTransaction` -> `POST /v1/points_transactions`
  - `ListPointsTransactions` -> `GET /v1/points_transactions`
  - `RetrievePointsTransaction` -> `GET /v1/points_transactions/{id}`
- `RewardFulfillments`
  - `ListRewardFulfillments` -> `GET /v1/reward_fulfillments`
- `VipTiers`
  - `ListVipTiers` -> `GET /v1/vip_tiers`

## 错误处理

SDK 在非 2xx 响应时返回 `*smile.APIError`，可使用 `errors.As` 判断。

```go
if smile.IsRateLimitAPIError(err) {
	// HTTP 429
}
if smile.IsInsufficientPointsAPIError(err) {
	// 负积分余额
}
```

## 说明

- 文档参考：<https://dev.smile.io/api/>
- 默认 API 根地址：`https://api.smile.io`
- 支持通过 `options.WithRateLimitRetry` 配置 429 重试次数
