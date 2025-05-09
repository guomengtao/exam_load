package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-go-test/utils"
)
// @Summary 获取MySQL状态
// @Description 检查MySQL数据库连接状态
// @Tags 数据库
// @Produce json
// @Success 200 {object} map[string]interface{} "数据库状态"
// @Router /api/mysql [get]
func MySQLStatus(c *gin.Context) {
	status := utils.CheckMySQLStatus()
	if status {
		c.JSON(http.StatusOK, gin.H{"status": "✅ MySQL connected"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "❌ MySQL connection failed"})
	}
}

func RedisStatus(c *gin.Context) {
	status := utils.CheckRedisStatus()
	if status {
		c.JSON(http.StatusOK, gin.H{"status": "✅ Redis connected"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "❌ Redis connection failed"})
	}
}