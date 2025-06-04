package routes

import (
	"gin-go-test/app/controllers"
	"gin-go-test/auth"
	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// 公共路由（无需认证）
	public := router.Group("/api")
	{
		utils.InitGorm()

		// 系统信息接口
		versionController := controllers.NewVersionController()
		public.GET("/version", versionController.GetVersion)

		statusController := controllers.NewStatusController()
		public.GET("/status", statusController.GetStatus)

		dbInfoController := controllers.NewDBInfoController(utils.DBX)
		public.GET("/dbinfo", dbInfoController.GetDBInfo)

		sourceController := controllers.NewSourceController(utils.DBX)
		public.GET("/source/check", sourceController.CheckSource)

		// 认证相关
		public.POST("/login", auth.LoginHandler)
	}

	// 需要认证的API组
	authGroup := router.Group("/api")
	authGroup.Use(auth.AuthMiddleware()) // 应用JWT认证中间件
	{
		// 文件上传
		uploadController := controllers.NewUploadController("uploads")
		authGroup.POST("/upload", uploadController.UploadFile)

		// 考试模板管理（需要 exam:manage 权限）
		examTemplate := authGroup.Group("/exam_template")
		examTemplate.Use(auth.PermissionMiddleware("exam:manage"))
		{
			// 删除 examTemplateController 相关的三行路由注册代码
		}

		// 试卷管理
		// 已用生成器统一注册，无需手动注册 examPaperController 相关路由
	}

	// 路由管理相关接口
	routeController := controllers.NewRouteController(utils.GormDB, router)
	routeGroup := router.Group("/api/route")
	{
		routeGroup.POST("/refresh", routeController.RefreshRoutes)               // 刷新路由列表
		routeGroup.GET("/list", routeController.GetRoutes)                       // 获取路由列表
		routeGroup.GET("/:id", routeController.GetRouteDetail)                   // 获取路由详情
		routeGroup.PUT("/:id/status", routeController.UpdateRouteStatus)         // 更新路由状态
		routeGroup.GET("/groups", routeController.GetRouteGroups)                // 获取路由分组
		routeGroup.PUT("/:id/group", routeController.UpdateRouteGroup)           // 更新路由分组
		routeGroup.PUT("/:id/permission", routeController.UpdateRoutePermission) // 更新路由权限
		routeGroup.GET("/stats", routeController.GetRouteStats)                  // 获取路由统计
	}

	// Auto Gen Router
	RegisterGeneratedRoutes(router)
}
