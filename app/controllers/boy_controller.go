package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-go-test/utils/generated/controller"
)

// BoyController 控制器示例
type BoyController struct {
	skeleton *controller.BoySkeleton
}

// RegisterBoyRoutes 注册 Boy 相关路由
func RegisterBoyRoutes(router *gin.Engine) {
	group := router.Group("/api/boy")
	ctrl := &BoyController{
		skeleton: &controller.BoySkeleton{},
	}

	group.GET("/hello", ctrl.HelloHandler)
}

// HelloHandler 示例接口，调用骨架层对应方法
func (ctrl *BoyController) HelloHandler(c *gin.Context) {
	result := ctrl.skeleton.Hello()
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}