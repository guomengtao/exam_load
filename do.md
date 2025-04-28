package handlers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"gin-go-test/utils"
	"gin-go-test/utils/redis"
)

// 查询Redis中Cat的值
func QueryCat(c *gin.Context) {
	queue := utils.NewQueue(3) // 设定最大排队数为3
	// 使用排队机制
	if err := queue.Enqueue(); err != nil {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
		return
	}
	defer queue.Dequeue() // 请求完成后释放排队位置

	// 查询Redis中Cat的值
	catValue, err := redis.GetFromRedis("cat") // 查询 Redis 中小写 "cat"
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询Redis失败"})
		return
	}

	// 返回查询结果
	c.JSON(http.StatusOK, gin.H{"result": catValue})

	// 通知队列状态
	queue.Notify(fmt.Sprintf("队列状态: %d", queue.CheckStatus()))
}