package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "{{.ModuleName}}/app/models"       // 这里是示范，你可以按实际路径调整
)

// {{.ControllerName}} 控制器示例
func {{.ControllerName}}(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello fr++om {{.ControllerName}}!",
	})
}