package routes

import (
	"github.com/gin-gonic/gin"
	"gin-go-test/app/controllers"

)

// RegisterGeneratedRoutes 统一注册所有生成的路由
func RegisterGeneratedRoutes(router *gin.Engine) {
 
 
	controllers.RegisterHelloRoutes(router)
	controllers.RegisterHelloaRoutes(router)
	controllers.RegisterKingRoutes(router)
	controllers.RegisterBoyRoutes(router)
// === GENERATED ROUTES END ===
}
