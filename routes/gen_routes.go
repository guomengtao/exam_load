package routes

import (
	"github.com/gin-gonic/gin"
	"gin-go-test/app/controllers"
	"gin-go-test/utils"
	"gin-go-test/auth"
)

// RegisterGeneratedRoutes 统一注册所有生成的路由
func RegisterGeneratedRoutes(router *gin.Engine) {
	controllers.RegisterHelloRoutes(router)
	controllers.RegisterHelloaRoutes(router)
	controllers.RegisterKingRoutes(router)
	controllers.RegisterBoyRoutes(router, utils.GormDB)

	controllers.RegisterUserRoutes(router, utils.GormDB)
	controllers.RegisterRoleRoutes(router, utils.GormDB)
	controllers.RegisterTeacherRoutes(router, utils.GormDB)
	controllers.RegisterFileInfoRoutes(router, utils.GormDB)
	controllers.RegisterBadmintonGameRoutes(router, utils.GormDB)
}

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	// 登录路由
	r.POST("/api/login", auth.LoginHandler)

	// 注册所有生成的路由
	RegisterGeneratedRoutes(r)
}
