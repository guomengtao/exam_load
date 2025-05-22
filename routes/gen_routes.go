package routes

import (
	"github.com/gin-gonic/gin"
	"gin-go-test/app/controllers/member"

)

// RegisterGeneratedRoutes 统一注册所有生成的路由
func RegisterGeneratedRoutes(router *gin.Engine) {
	member.RegisterMemberRoutes(router)

}
