package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	controllers "gin-go-test/app/controllers"
	"gin-go-test/app/models"
	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
)

func TestRoleController_BatchCreateHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// 初始化数据库
	utils.InitGorm()
	db := utils.GormDB

	r := gin.Default()
	controllers.RegisterRoleRoutes(r, db)

	t.Run("正常批量创建", func(t *testing.T) {
		items := []*models.Role{
			{
				Name: "Role1",
				Desc: "Description1",
			},
		}
		body, _ := json.Marshal(items)
		req := httptest.NewRequest(http.MethodPost, "/api/role", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK && w.Code != http.StatusPartialContent {
			t.Errorf("期望状态码 200 或 206，实际是 %d", w.Code)
		}
	})

	t.Run("超过批量限制", func(t *testing.T) {
		items := make([]*models.Role, 31)
		for i := range items {
			items[i] = &models.Role{
				Name: "Role" + strconv.Itoa(i+1),
				Desc: "Description" + strconv.Itoa(i+1),
			}
		}
		body, _ := json.Marshal(items)
		req := httptest.NewRequest(http.MethodPost, "/api/role", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("期望状态码 400，实际是 %d", w.Code)
		}
	})

	t.Run("无效请求体", func(t *testing.T) {
		body := []byte(`{"invalid": "data"`)
		req := httptest.NewRequest(http.MethodPost, "/api/role", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("期望状态码 400，实际是 %d", w.Code)
		}
	})
}

/*
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
*/
