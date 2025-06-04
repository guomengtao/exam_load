package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponseFormat(t *testing.T) {
	// 设置gin测试模式
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name     string
		handler  func(*gin.Context)
		expected Response
	}{
		{
			name: "成功响应-空数据",
			handler: func(c *gin.Context) {
				Success(c, nil)
			},
			expected: Response{
				Code: 0,
				Msg:  "success",
				Data: nil,
			},
		},
		{
			name: "成功响应-有数据",
			handler: func(c *gin.Context) {
				Success(c, map[string]string{"key": "value"})
			},
			expected: Response{
				Code: 0,
				Msg:  "success",
				Data: map[string]interface{}{"key": "value"},
			},
		},
		{
			name: "错误响应",
			handler: func(c *gin.Context) {
				Error(c, "test error")
			},
			expected: Response{
				Code: 1,
				Msg:  "test error",
				Data: nil,
			},
		},
		{
			name: "分页响应",
			handler: func(c *gin.Context) {
				PageSuccess(c, []string{"item1", "item2"}, 2)
			},
			expected: Response{
				Code: 0,
				Msg:  "success",
				Data: map[string]interface{}{
					"list":  []interface{}{"item1", "item2"},
					"total": float64(2),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建测试上下文
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// 执行处理函数
			tt.handler(c)

			// 验证状态码
			assert.Equal(t, http.StatusOK, w.Code)

			// 解析响应
			var response Response
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			// 验证响应内容
			assert.Equal(t, tt.expected.Code, response.Code)
			assert.Equal(t, tt.expected.Msg, response.Msg)
			assert.Equal(t, tt.expected.Data, response.Data)
		})
	}
}
