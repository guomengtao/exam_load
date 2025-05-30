package service

import (
	"fmt"
	"time"
)

// ErrorResponse 用于统一错误返回格式
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Time    string `json:"timestamp"`
}

// NewErrorResponse 创建新的错误响应
func NewErrorResponse(code int, message string, details string) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
		Time:    time.Now().Format(time.RFC3339),
	}
}

// isZero 判断一个字段是否为零值
func isZero(v interface{}) bool {
	switch val := v.(type) {
	case string:
		return val == ""
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%v", val) == "0"
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%v", val) == "0"
	case float32, float64:
		return fmt.Sprintf("%v", val) == "0" || fmt.Sprintf("%v", val) == "0.0"
	case bool:
		return !val
	case nil:
		return true
	default:
		return v == nil
	}
} 