package handlers

import (
	"github.com/gin-gonic/gin"
)

// @Summary 测试接口
// @Description 返回Hello World测试信息
// @Tags 测试
// @Produce json
// @Success 200 {string} string "Hello World"
// @Router /api/hellobay [get]
func HelloWorld(c *gin.Context) {
    c.String(200, "Hello World")
}