package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-go-test/utils/generated/controller"
)

// {{.ControllerName}}Controller 控制器示例
type {{.ControllerName}}Controller struct {
	skeleton *controller.{{.ControllerName}}Skeleton
}

// Register{{.ControllerName}}Routes 注册 {{.ControllerName}} 相关路由
func Register{{.ControllerName}}Routes(router *gin.Engine) {
	group := router.Group("/api/{{.RoutePath}}")
	ctrl := &{{.ControllerName}}Controller{
		skeleton: &controller.{{.ControllerName}}Skeleton{},
	}

	group.GET("/hello", ctrl.HelloHandler)
}

// HelloHandler 示例接口，调用骨架层对应方法
func (ctrl *{{.ControllerName}}Controller) HelloHandler(c *gin.Context) {
	result := ctrl.skeleton.Hello()
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}