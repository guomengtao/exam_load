package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloaController 控制器示例
type HelloaController struct{}

// RegisterHelloaRoutes 注册 Helloa 相关路由
func RegisterHelloaRoutes(router *gin.Engine) {
	group := router.Group("/api/helloa")
	ctrl := &HelloaController{}

	group.GET("/hello", ctrl.HelloHandler)
}

// HelloHandler 示例接口
func (ctrl *HelloaController) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from HelloaController!",
	})
}