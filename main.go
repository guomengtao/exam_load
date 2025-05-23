// @title        Gin API 文档
// @version      1.0
// @description  这是一个基于 Gin 的接口文档示例
// @host         localhost:8081
// @BasePath     /
package main

import (
	"io"
	"os"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"gin-go-test/routes"
	"gin-go-test/utils"
	"github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
	_ "gin-go-test/docs"  // 修改为你的 go.mod 模块名
	// "gin-go-test/auth"
	"gin-go-test/app/services"
)

func main() {
	// 启用控制台彩色日志
	gin.ForceConsoleColor()

	 
	
	// 创建日志文件
	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 初始化数据库和Redis
	utils.InitDB()
	utils.InitRedis()

	utils.InitDBX()
	utils.InitGorm()  // 必须调用这个

	// auth.InitDB(utils.DB) // 初始化认证模块

	// 启动所有任务
    services.StartAllTasks()

	router := gin.Default()
	
	// CORS配置
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:   []string{"Content-Length"},
		MaxAge:          12 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 设置路由
	routes.SetupRoutes(router)

	// 静态文件服务
	router.Static("/static", "./static")
	router.StaticFile("/", "./static/index.html")

	// 启动服务器
	router.Run(":8081")
}