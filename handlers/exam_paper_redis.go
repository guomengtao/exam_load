package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
)

// ListExamPapersFromRedis 从 Redis 中读取试卷列表或单条（通过 UUID）
func ListExamPapersFromRedis(c *gin.Context) {
	// 获取查询参数
	pageParam := c.DefaultQuery("page", "1")
	limitParam := c.DefaultQuery("limit", "5")
	uuidParam := c.DefaultQuery("uuid", "")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 5
	}

	// 如果有传 uuid 参数，查询指定试卷
	if uuidParam != "" {
		redisKey := fmt.Sprintf("exam_paper:%s", uuidParam)
		val, err := utils.RedisClient.Get(context.Background(), redisKey).Result()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "试卷未找到",
				"data": nil,
			})
			return
		}

		var paper map[string]interface{}
		if err := json.Unmarshal([]byte(val), &paper); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "解析失败",
				"data": nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": paper,
		})
		return
	}

	// 没有 uuid，执行分页查询所有
	var cursor uint64
	var keys []string
	matchPattern := "exam_paper:*"

	for {
		scannedKeys, nextCursor, err := utils.RedisClient.Scan(context.Background(), cursor, matchPattern, 100).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "Redis 扫描失败",
				"data": nil,
			})
			return
		}
		keys = append(keys, scannedKeys...)
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	// 分页
	start := (page - 1) * limit
	end := start + limit
	if start > len(keys) {
		start = len(keys)
	}
	if end > len(keys) {
		end = len(keys)
	}
	pagedKeys := keys[start:end]

	// 根据 key 批量获取数据
	var papers []map[string]interface{}
	for _, key := range pagedKeys {
		val, err := utils.RedisClient.Get(context.Background(), key).Result()
		if err != nil {
			continue
		}
		var paper map[string]interface{}
		if err := json.Unmarshal([]byte(val), &paper); err != nil {
			continue
		}
		papers = append(papers, paper)
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"list": papers,
			"pagination": gin.H{
				"page":  page,
				"limit": limit,
				"total": len(keys),
			},
		},
	})
}