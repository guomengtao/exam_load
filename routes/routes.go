package routes

import (
	"github.com/gin-gonic/gin"
	"gin-go-test/handlers" // ✅ 确保 module 名正确
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/api/hellobay", handlers.HelloWorld)

	// ✅ 添加数据库状态接口
	router.GET("/api/mysql", handlers.MySQLStatus)

	// ✅ 添加 Redis 状态接口
	router.GET("/api/redis", handlers.RedisStatus)

	// 注册 API 路由
	router.GET("/api/version", handlers.GetVersion)  
	// 路由 /api/exam/:id 将根据请求方法分别调用不同的处理函数
    router.GET("/api/exam/:id", handlers.GetExam)   // GET 方法：获取试卷
    router.POST("/api/exam/:id", handlers.SubmitAnswers) // POST 方法：提交学生回答

	router.GET("/api/dbinfo", handlers.GetDBTablesInfo)


		// 注册 dog 路由
		// router.GET("/api/dog", handlers.QueryDog)
		// 注册 cat 路由
		// router.GET("/api/cat", handlers.QueryCat)

		 
		
			// 路由配置
	router.POST("/api/exam_template", handlers.CreateExam)
	
	router.GET("/api/exam_template", handlers.GetExamTemplate)

	router.PUT("/api/exam_template", handlers.UpdateExamTemplate)


	router.POST("/api/exam_paper", handlers.CreateExamPaper)

	router.GET("/api/exam_paper/redis", handlers.ListExamPapersFromRedis)
		 

}