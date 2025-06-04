package controllers

import (
	"gin-go-test/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TaskControlHandler(c *gin.Context) {
	// 读取参数 name=importer, action=stop/start
	var req struct {
		Name   string `json:"name" binding:"required"`
		Action string `json:"action" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	switch req.Action {
	case "stop":
		services.StopTask(req.Name)
	case "start":
		// 暂不实现重启，如需实现可调用 RestartTask
		c.JSON(http.StatusNotImplemented, gin.H{"message": "暂不支持启动任务"})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的操作"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "操作成功"})
}
