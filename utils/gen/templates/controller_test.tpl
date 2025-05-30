package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"gin-go-test/app/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test{{.ControllerName}}Controller_BatchCreateHandler(t *testing.T) {
	// 设置测试模式
	gin.SetMode(gin.TestMode)

	// 创建测试路由
	router := gin.New()
	Register{{.ControllerName}}Routes(router, nil)

	// 测试用例1：正常批量创建
	t.Run("Successful batch create", func(t *testing.T) {
		// 准备测试数据
		items := []*models.{{.ModelName}}{
			{
				// 根据实际模型字段填充测试数据
			},
		}
		body, _ := json.Marshal(items)

		// 创建测试请求
		req := httptest.NewRequest(http.MethodPost, "/api/{{.RoutePath}}/{{.TableName}}", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// 执行请求
		router.ServeHTTP(w, req)

		// 验证响应
		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "success_count")
	})

	// 测试用例2：超过批量限制
	t.Run("Exceed batch limit", func(t *testing.T) {
		// 准备超过限制的测试数据
		items := make([]*models.{{.ModelName}}, 31)
		body, _ := json.Marshal(items)

		// 创建测试请求
		req := httptest.NewRequest(http.MethodPost, "/api/{{.RoutePath}}/{{.TableName}}", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// 执行请求
		router.ServeHTTP(w, req)

		// 验证响应
		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})

	// 测试用例3：无效的请求体
	t.Run("Invalid request body", func(t *testing.T) {
		// 准备无效的测试数据
		body := []byte(`{"invalid": "data"`)

		// 创建测试请求
		req := httptest.NewRequest(http.MethodPost, "/api/{{.RoutePath}}/{{.TableName}}", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// 执行请求
		router.ServeHTTP(w, req)

		// 验证响应
		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})
} 