package routes

import (
	"gin-go-test/app/controllers"
	"gin-go-test/auth"
	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
)

// RegisterGeneratedRoutes 统一注册所有生成的路由
func RegisterGeneratedRoutes(router *gin.Engine) {
	controllers.RegisterUserRoutes(router, utils.GormDB)
	controllers.RegisterRoleRoutes(router, utils.GormDB)
	controllers.RegisterTeacherRoutes(router, utils.GormDB)
	controllers.RegisterFileInfoRoutes(router, utils.GormDB)
	controllers.RegisterBadmintonGameRoutes(router, utils.GormDB)
	controllers.RegisterExamTemplateRoutes(router, utils.GormDB)
	controllers.RegisterExamPaperRoutes(router, utils.GormDB)
	controllers.RegisterExamPapersRoutes(router, utils.GormDB)
}

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	// 登录路由
	r.POST("/api/login", auth.LoginHandler)

	// 注册所有生成的路由
	RegisterGeneratedRoutes(r)
}
