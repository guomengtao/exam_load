package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "gin-go-test/app/models"       // 这里是示范，你可以按实际路径调整
)

// HelloController 控制器示例
func HelloController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello fr++om HelloController!",
	})
}