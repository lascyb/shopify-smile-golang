package smile

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// APIError Smile REST 非 2xx 时常见的 JSON 形态，例如扣减积分导致负余额时 HTTP 400：
//
//	{"error":{"message":"...","request_id":"..."}}
//
// 可用 errors.As(err, new(*smile.APIError)) 判断并读取 Message、RequestID。
type APIError struct {
	StatusCode int    // HTTP 状态码，如 400
	Message    string // error.message
	RequestID  string // error.request_id
	Body       string // 原始响应体
}

// Error 实现 error 接口。
func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	if e.Message != "" {
		s := fmt.Sprintf("smile: HTTP %d: %s", e.StatusCode, e.Message)
		if e.RequestID != "" {
			s += fmt.Sprintf(" (request_id=%s)", e.RequestID)
		}
		return s
	}
	if e.Body != "" {
		return fmt.Sprintf("smile: HTTP %d body=%s", e.StatusCode, e.Body)
	}
	return fmt.Sprintf("smile: HTTP %d", e.StatusCode)
}

type smileErrorResponse struct {
	Error *struct {
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	} `json:"error"`
}

// newAPIErrorFromResponse 从非 2xx 响应构造 *APIError。
func newAPIErrorFromResponse(statusCode int, body []byte) *APIError {
	raw := string(body)
	var wrap smileErrorResponse
	if err := json.Unmarshal(body, &wrap); err == nil && wrap.Error != nil && wrap.Error.Message != "" {
		return &APIError{
			StatusCode: statusCode,
			Message:    wrap.Error.Message,
			RequestID:  wrap.Error.RequestID,
			Body:       raw,
		}
	}
	return &APIError{
		StatusCode: statusCode,
		Body:       raw,
	}
}

// IsInsufficientPointsAPIError 判断是否为 Smile 扣减积分导致负余额（HTTP 400，message 含 negative points balance）。
func IsInsufficientPointsAPIError(err error) bool {
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr == nil {
		return false
	}
	if apiErr.StatusCode != http.StatusBadRequest {
		return false
	}
	return strings.Contains(strings.ToLower(apiErr.Message), "negative points balance")
}

// IsRateLimitAPIError 判断是否为 Smile HTTP 429（限流等），常用于提示用户稍后重试。
func IsRateLimitAPIError(err error) bool {
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr == nil {
		return false
	}
	return apiErr.StatusCode == http.StatusTooManyRequests
}
