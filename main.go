package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"gin-go-test/routes"
	"gin-go-test/utils" // 引入 utils 初始化 DB 和 Redis
)

func main() {
	// ✅ 初始化数据库和 Redis
	utils.InitDB()
	utils.InitRedis()

	router := gin.Default()

	// ✅ 加上 CORS 中间件
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://47.120.38.206:8086"}, // 允许前端访问
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ✅ 挂载所有路由（包括 /api/mysql 和 /api/redis）
	routes.SetupRoutes(router)

	// ✅ 启动服务器
	router.Run(":8081")
}