package controllers

import (
	"gin-go-test/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ExamPaperRedisController handles Redis-based exam paper operations
type ExamPaperRedisController struct {
	paperBiz *biz.ExamPaperRedisBiz
}

// NewExamPaperRedisController creates a new instance of ExamPaperRedisController
func NewExamPaperRedisController() *ExamPaperRedisController {
	return &ExamPaperRedisController{
		paperBiz: biz.NewExamPaperRedisBiz(),
	}
}

// ListExamPapersFromRedis 从 Redis 中读取试卷列表或单条（通过 UUID）
func (c *ExamPaperRedisController) ListExamPapersFromRedis(ctx *gin.Context) {
	pageParam := ctx.DefaultQuery("page", "1")
	limitParam := ctx.DefaultQuery("limit", "5")
	uuidParam := ctx.DefaultQuery("uuid", "")

	response, err := c.paperBiz.ListExamPapersFromRedis(ctx, pageParam, limitParam, uuidParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": response,
	})
}
