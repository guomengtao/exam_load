package routes

import (
	"github.com/gin-gonic/gin"
	"gin-go-test/handlers"
	"gin-go-test/auth" // 新增认证模块
	"gin-go-test/app/controllers" // 修改为你的实际模块路径
	"gin-go-test/utils"
 )

func SetupRoutes(router *gin.Engine) {
	// 公共路由（无需认证）
	public := router.Group("/api")
	{
		utils.InitGorm();
		// 系统状态接口
		public.GET("/hello_world", handlers.HelloWorld)
		public.GET("/hello", controllers.HelloHandler)
		public.GET("/mysql", handlers.MySQLStatus)
		public.GET("/redis", handlers.RedisStatus)
		public.GET("/version", handlers.GetVersion)
		public.GET("/dbinfo", handlers.GetDBTablesInfo)
		public.GET("/source/check_all", handlers.CheckAllSources)

		// 认证相关
		public.POST("/login", auth.LoginHandler)

		// 写入作答记录
		public.POST("/user/answer", handlers.SubmitAnswer)

		public.GET("user/answer/:record_id", handlers.GetAnswerResult)

		public.GET("user/answer/:record_id/full", handlers.GetFullAnswerResult)


		// 临时开放 
		public.GET("/admins", controllers.GetAdminsHandler)
		public.PUT("/admin/:id/password", controllers.UpdateAdminPasswordHandler)
 		public.GET("/status", controllers.StatusHandler)
		public.POST("/task/control", controllers.TaskControlHandler)
		public.GET("/export_answers", controllers.ExportAnswersHandler)
		public.POST("/import_students", controllers.ImportStudentsHandler)
		


		public.GET("/roles", controllers.GetRolesHandler)


		
	}

	// 需要认证的API组
	authGroup := router.Group("/api")
	authGroup.Use(auth.AuthMiddleware()) // 应用JWT认证中间件
	{
		// 考试模板管理（需要 exam:manage 权限）
		examTemplate := authGroup.Group("/exam_template")
		examTemplate.Use(auth.PermissionMiddleware("exam:manage"))
		{
			examTemplate.GET("", handlers.GetExamTemplate)
			examTemplate.POST("", handlers.CreateExam)
			examTemplate.PUT("", handlers.UpdateExamTemplate)
		}

		// 试卷管理
		examPaper := authGroup.Group("/exam_paper")
		{
			examPaper.GET("", handlers.GetExam)
			examPaper.POST("", handlers.CreateExamPaper)
			examPaper.GET("/redis", handlers.ListExamPapersFromRedis)
		}

		// 答案提交与查询
		answer := authGroup.Group("/answer")
		{
			answer.POST("", handlers.SubmitAnswer)
			answer.GET("/:record_id", handlers.GetAnswerResult)
			answer.GET("/:record_id/full", handlers.GetFullAnswerResult)
			answer.POST("/:id/submit", handlers.SubmitAnswers) // 保持原POST路径兼容
		}

		// 文件上传
		authGroup.POST("/upload", handlers.UploadExamImage)
	}

	// 兼容旧路由（逐步迁移时可保留）
	router.GET("/api/exam", handlers.GetExam) // 兼容旧GET请求
}