package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
)

func TestUpdateOwnPassword(t *testing.T) {
	adminController := &AdminController{}
	utils.InitRedis() // Initialize Redis client before testing

	token := os.Getenv("TEST_JWT_TOKEN")
	newPassword := os.Getenv("TEST_NEW_PASSWORD")

	if token == "" || newPassword == "" {
		t.Fatal("❌ 请设置环境变量 TEST_JWT_TOKEN 和 TEST_NEW_PASSWORD")
	}

	router := gin.Default()
	router.PUT("/admin/password", adminController.UpdatePassword)

	payload := map[string]string{
		"new_password": newPassword,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPut, "/admin/password", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("❌ 接口响应码错误，期望 200，得到 %d，响应内容：%s", resp.Code, resp.Body.String())
	} else {
		t.Logf("✅ 修改密码成功: %s", resp.Body.String())
	}
}
