package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloController 控制器示例
type HelloController struct{}

// RegisterHelloRoutes 注册 Hello 相关路由
func RegisterHelloRoutes(router *gin.Engine) {
	group := router.Group("/api/hello")
	ctrl := &HelloController{}

	group.GET("/hello", ctrl.HelloHandler)
}

// HelloHandler 示例接口
func (ctrl *HelloController) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from HelloController!",
	})
}