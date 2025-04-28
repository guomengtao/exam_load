package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetVersion 返回当前的版本号
func GetVersion(c *gin.Context) {
	// 你可以根据需求设置返回的格式，这里是 JSON 格式
	c.JSON(http.StatusOK, gin.H{
		"version": "1.0.0",
	})
}