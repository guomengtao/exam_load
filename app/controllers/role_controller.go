package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-go-test/utils/generated/controller"
	"gin-go-test/app/biz"
	"gin-go-test/app/services"
	"gorm.io/gorm"
)

// RoleController 控制器示例
type RoleController struct {
	skeleton *controller.RoleSkeleton
}

// RegisterRoleRoutes 注册 Role 相关路由
func RegisterRoleRoutes(router *gin.Engine, db *gorm.DB) {
	group := router.Group("/api/role")
	ctrl := &RoleController{
		skeleton: controller.NewRoleSkeleton(
			biz.NewRoleBiz(
				services.NewRoleService(db),
			),
		),
	}

	group.GET("/hello", ctrl.HelloHandler)
	group.GET("/count", ctrl.CountHandler)
}

// HelloHandler 示例接口，调用骨架层对应方法
func (ctrl *RoleController) HelloHandler(c *gin.Context) {
	result := ctrl.skeleton.Hello()
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

// CountHandler 查询真实总数，调用 Biz 层
func (ctrl *RoleController) CountHandler(c *gin.Context) {
	ctrl.skeleton.CountHandler(c)
}

// GetRolesHandler 是兼容旧接口的独立函数
func GetRolesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from app/controllers!",
	})
}