package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gin-go-test/utils"
)

func TestGetRolesHandler(t *testing.T) {
	// 初始化数据库连接
	utils.InitGorm()

	// 设置 Gin 为测试模式
	gin.SetMode(gin.TestMode)

	// 初始化路由
	router := gin.Default()
	router.GET("/api/roles", GetRolesHandler)

	// 创建测试请求
	req, _ := http.NewRequest(http.MethodGet, "/api/roles", nil)
	resp := httptest.NewRecorder()

	// 执行请求
	router.ServeHTTP(resp, req)

	// 验证响应
	if resp.Code != http.StatusOK {
		t.Errorf("期望状态码 200，实际是 %d", resp.Code)
	}

	t.Logf("响应内容：%s", resp.Body.String())
}