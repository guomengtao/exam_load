package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-go-test/utils/generated/controller"
	"gin-go-test/app/biz"
	"gin-go-test/app/services"
	"gorm.io/gorm"
)

// UserController 控制器示例
type UserController struct {
	skeleton *controller.UserSkeleton
}

// RegisterUserRoutes 注册 User 相关路由
func RegisterUserRoutes(router *gin.Engine, db *gorm.DB) {
	group := router.Group("/api/user")
	ctrl := &UserController{
		skeleton: controller.NewUserSkeleton(
			biz.NewUserBiz(
				services.NewUserService(db),
			),
		),
	}

	group.GET("/hello", ctrl.HelloHandler)
	group.GET("/count", ctrl.CountHandler)
}

// HelloHandler 示例接口，调用骨架层对应方法
func (ctrl *UserController) HelloHandler(c *gin.Context) {
	result := ctrl.skeleton.Hello()
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

// CountHandler 查询真实总数，调用 Biz 层
func (ctrl *UserController) CountHandler(c *gin.Context) {
	ctrl.skeleton.CountHandler(c)
}