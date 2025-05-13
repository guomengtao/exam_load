package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
	"strconv"
	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
)

// 初始化数据库 & Gin 引擎
func setupRouter() *gin.Engine {
	utils.InitDBX() // 确保数据库已连接
	r := gin.Default()
	r.GET("/admin/list", GetAdminsHandler)
	return r
}

func TestGetAdminsHandler(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/admin/list", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("期望 200，实际得到 %d", resp.Code)
	}

	// 可选：你还可以断言返回的 JSON 内容
	t.Logf("接口响应：%s", resp.Body.String())
}

func TestUpdateAdminPassword(t *testing.T) {
	adminIDStr := os.Getenv("TEST_ADMIN_ID")
	newPassword := os.Getenv("TEST_NEW_PASSWORD")

	if adminIDStr == "" || newPassword == "" {
		t.Fatal("❌ 请设置环境变量 TEST_ADMIN_ID 和 TEST_NEW_PASSWORD")
	}

	adminID, err := strconv.Atoi(adminIDStr)
	if err != nil {
		t.Fatalf("❌ TEST_ADMIN_ID 必须是整数: %v", err)
	}

	router := gin.Default()
	router.PUT("/api/admins/:id/password", UpdateAdminPasswordHandler)

	payload := map[string]interface{}{
		"admin_id":     adminID,
		"new_password": newPassword,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPut, "/api/admins/1/password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("❌ 接口响应码错误，期望 200，得到 %d，响应内容：%s", resp.Code, resp.Body.String())
	} else {
		t.Logf("✅ 密码更新成功: %s", resp.Body.String())
	}
}