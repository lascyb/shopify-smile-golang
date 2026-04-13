// Package smile 封装 Smile.io REST API（https://dev.smile.io/api）的客户端与通用请求能力；资源接口通过 client.Apis 访问（如 Apis.PointsTransactions）。
package smile

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/lascyb/shopify-smile-golang/apis"

	"github.com/lascyb/shopify-smile-golang/options"
)

// Client Smile REST 客户端；baseURL 来自 options（默认 https://api.smile.io）。
type Client struct {
	baseURL             string
	apiKey              string
	httpClient          *http.Client
	rateLimitMaxRetries int
	Apis                *apis.Apis // 资源 API 入口，由 NewClient 初始化
}

// NewClient 使用私有 API Key（Bearer）创建客户端；根地址与其它可选项通过 options.Option 配置（见 options.WithBaseURL 等）。
func NewClient(apiKey string, opt ...options.Option) (*Client, error) {
	key := strings.TrimSpace(apiKey)
	if key == "" {
		return nil, fmt.Errorf("smile: apiKey 不能为空")
	}
	cfg := options.NewConfig(opt...)
	c := &Client{
		baseURL:             cfg.BaseURL,
		apiKey:              key,
		httpClient:          cfg.HTTPClient,
		rateLimitMaxRetries: cfg.RateLimitMaxRetries,
	}
	c.Apis = apis.NewApis(c)
	return c, nil
}

func (c *Client) SmileSDK() {}

// GetJSON 向 baseURL+path 发送 GET（Bearer 鉴权），并将 2xx 响应体解码到 out（out 为 nil 时不解码）。
func (c *Client) GetJSON(ctx context.Context, path string, out any) error {
	return c.doJSON(ctx, http.MethodGet, path, nil, out)
}

// PostJSON 向 baseURL+path 发送 JSON POST（Bearer 鉴权），并将 2xx 响应体解码到 out（out 为 nil 时不解码）。
func (c *Client) PostJSON(ctx context.Context, path string, body any, out any) error {
	raw, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("smile: 编码请求体失败：%w", err)
	}
	return c.doJSON(ctx, http.MethodPost, path, raw, out)
}

// doJSON 通用请求方法：构建请求 → Bearer 鉴权 → 执行 → 429 自动重试 → 状态码检查 → 响应解码。
func (c *Client) doJSON(ctx context.Context, method, path string, bodyBytes []byte, out any) error {
	if c == nil {
		return fmt.Errorf("smile: client 未初始化")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	maxAttempts := 1 + c.rateLimitMaxRetries
	if maxAttempts < 1 {
		maxAttempts = 1
	}

	var lastErr error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		lastErr = c.doOnce(ctx, method, path, bodyBytes, out)
		if lastErr == nil {
			return nil
		}
		if !IsRateLimitAPIError(lastErr) || attempt == maxAttempts-1 {
			return lastErr
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Second * time.Duration(min(attempt, 3))):
		}
	}
	return lastErr
}

func (c *Client) doOnce(ctx context.Context, method, path string, bodyBytes []byte, out any) error {
	var body io.Reader
	if bodyBytes != nil {
		body = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, body)
	if err != nil {
		return fmt.Errorf("smile: 构建请求失败：%w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	if bodyBytes != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("smile: 请求失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("smile: 读取响应失败：%w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return newAPIErrorFromResponse(resp.StatusCode, respBytes)
	}

	if out == nil {
		return nil
	}
	if err = json.Unmarshal(respBytes, out); err != nil {
		return fmt.Errorf("smile: 解析响应失败：%w body=%s", err, string(respBytes))
	}
	return nil
}
