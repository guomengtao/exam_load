package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExportAnswersHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/api/export_answers", ExportAnswersHandler)

	t.Run("有数据时返回200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/export_answers?exam_uuid=40a89290-0ff9-4c01-a803-57155a24985c&school=德胜实验小学&limit=1000", nil)
		assert.NoError(t, err)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		body := resp.Body.String()
		t.Logf("响应内容: %s", body)
		assert.Contains(t, body, "导出成功")
	})

	t.Run("无数据时返回404", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/export_answers?exam_uuid=00000000-0000-0000-0000-000000000000&limit=10", nil)
		assert.NoError(t, err)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		body := resp.Body.String()
		t.Logf("响应内容: %s", body)
		assert.True(t, strings.Contains(body, "没有对应数据") || strings.Contains(body, "not found"))
	})
}