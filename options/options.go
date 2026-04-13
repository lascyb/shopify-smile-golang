// Package options 提供 Smile REST 客户端的可选配置（函数式选项）。
package options

import (
	"net/http"
	"strings"
	"time"
)

const (
	defaultHTTPTimeout = 30 * time.Second
	// defaultBaseURL Smile REST 文档中的 API 根地址（无版本路径）；未调用 WithBaseURL 时使用。
	defaultBaseURL = "https://api.smile.io"
)

// Config 为 NewClient 合并后的可选配置。
type Config struct {
	BaseURL             string
	HTTPClient          *http.Client
	RateLimitMaxRetries int // HTTP 429 限流最大重试次数，0 表示不重试
}

// Option 用于修改 Config。
type Option func(*Config)

// NewConfig 应用 opt 并填充默认值（BaseURL、HTTP 客户端）。
func NewConfig(opt ...Option) *Config {
	c := &Config{
		BaseURL: defaultBaseURL,
	}
	for _, o := range opt {
		if o == nil {
			continue
		}
		o(c)
	}
	c.BaseURL = strings.TrimSpace(c.BaseURL)
	c.BaseURL = strings.TrimRight(c.BaseURL, "/")
	if c.BaseURL == "" {
		c.BaseURL = defaultBaseURL
	}
	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{Timeout: defaultHTTPTimeout}
	}
	return c
}

// WithBaseURL 设置 API 根地址（无尾斜杠、不含 /v1 等版本路径）；传空或未设置时使用 defaultBaseURL。
func WithBaseURL(baseURL string) Option {
	return func(c *Config) {
		if c == nil {
			return
		}
		c.BaseURL = strings.TrimSpace(baseURL)
	}
}

// WithHttpClient 指定 HTTP 客户端；不传则使用默认超时（defaultHTTPTimeout）。
func WithHttpClient(httpClient *http.Client) Option {
	return func(c *Config) {
		if c == nil {
			return
		}
		c.HTTPClient = httpClient
	}
}

// WithRateLimitRetry 启用 HTTP 429 限流自动重试；maxRetries 为最大重试次数，backoff 为基础退避间隔（实际等待 = backoff * attempt）。
func WithRateLimitRetry(maxRetries int, backoff time.Duration) Option {
	return func(c *Config) {
		if c == nil {
			return
		}
		c.RateLimitMaxRetries = maxRetries
	}
}
