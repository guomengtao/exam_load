package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-go-test/utils"
)

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