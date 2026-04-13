// Package contract 定义 Smile SDK 子包与 Client 之间的最小契约，避免 import cycle。
package contract

import "context"

// SmileClient 提供带 Bearer 的 JSON 请求与响应解码（由 *smile.Client 实现）；429 重试通过 options.SetRetry 配置。
type SmileClient interface {
	SmileSDK()
	GetJSON(ctx context.Context, path string, out any) error
	PostJSON(ctx context.Context, path string, body any, out any) error
}
