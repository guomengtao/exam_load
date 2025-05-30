package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-go-test/utils/generated/controller"
	"gin-go-test/app/biz"
	"gin-go-test/app/services"
	"gorm.io/gorm"
)

// Register{{.ControllerName}}Routes 注册 {{.ControllerName}} 相关路由
func Register{{.ControllerName}}Routes(router *gin.Engine, db *gorm.DB) {
	group := router.Group("/api/{{.RoutePath}}")
	skeleton := controller.New{{.ControllerName}}Skeleton(biz.New{{.ControllerName}}Biz(services.New{{.ControllerName}}Service(db)))
	group.GET("/count", skeleton.CountHandler)
	group.GET("/list", skeleton.ListHandler)
	group.POST("", skeleton.BatchCreateHandler)
	group.PUT("", skeleton.BatchUpdateHandler)
	group.DELETE("", skeleton.BatchDeleteHandler)
}