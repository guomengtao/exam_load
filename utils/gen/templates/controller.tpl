package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-go-test/utils/generated/controller"
	"gin-go-test/app/biz"
	"gin-go-test/app/services"
	"gorm.io/gorm"
)

// {{.ControllerName}}Controller 控制器示例
type {{.ControllerName}}Controller struct {
	skeleton *controller.{{.ControllerName}}Skeleton
}

// Register{{.ControllerName}}Routes 注册 {{.ControllerName}} 相关路由
func Register{{.ControllerName}}Routes(router *gin.Engine, db *gorm.DB) {
	group := router.Group("/api/{{.RoutePath}}")
	ctrl := &{{.ControllerName}}Controller{
		skeleton: controller.New{{.ControllerName}}Skeleton(
			biz.New{{.ControllerName}}Biz(
				services.New{{.ControllerName}}Service(db),
			),
		),
	}

	group.GET("/hello", ctrl.HelloHandler)
	group.GET("/count", ctrl.CountHandler)
}

// HelloHandler 示例接口，调用骨架层对应方法
func (ctrl *{{.ControllerName}}Controller) HelloHandler(c *gin.Context) {
	result := ctrl.skeleton.Hello()
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

// CountHandler 查询真实总数，调用 Biz 层
func (ctrl *{{.ControllerName}}Controller) CountHandler(c *gin.Context) {
	ctrl.skeleton.CountHandler(c)
}