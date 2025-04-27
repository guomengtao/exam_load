package routes

import (
	"github.com/gin-gonic/gin"
	"gin-go-test/handlers" // ✅ 确保 module 名正确
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/api/hello", handlers.HelloWorld)

	// ✅ 添加数据库状态接口
	router.GET("/api/mysql", handlers.MySQLStatus)

	// ✅ 添加 Redis 状态接口
	router.GET("/api/redis", handlers.RedisStatus)
}