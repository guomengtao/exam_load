package controllers

import (
    "context"
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/jmoiron/sqlx"

    "gin-go-test/utils"
)

var sqlxDB *sqlx.DB

func init() {
    // 初始化数据库连接，保证只在这里调用
    utils.InitDBX()
    sqlxDB = utils.DBX
}

// StatusController 状态控制器
type StatusController struct{}

// NewStatusController 创建状态控制器实例
func NewStatusController() *StatusController {
	return &StatusController{}
}

// GetStatus 获取系统状态
// @Summary 获取系统状态信息
// @Description 返回当前系统的运行状态和健康检查信息
// @Tags 系统信息
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/status [get]
func (c *StatusController) GetStatus(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":    "running",
		"uptime":    "24h",
		"memory":    "256MB",
		"cpu_usage": "5%",
	})
}

// StatusHandler 返回 Redis 导入和 MySQL 写入状态
func StatusHandler(c *gin.Context) {
    ctx := context.Background()
    redis := utils.RedisClient

    submittedCount, err1 := redis.ZCard(ctx, "exam:submitted").Result()
    processedCount, err2 := redis.ZCard(ctx, "exam:processed").Result()

    examAnswerCount := int64(0)
    keys, err3 := redis.Keys(ctx, "exam_answer:*").Result()
    if err3 == nil {
        examAnswerCount = int64(len(keys))
    }

    now := time.Now().Unix()
    oneHourAgo := now - 3600
    processedLastHour, err4 := redis.ZCount(ctx, "exam:processed", strconv.FormatInt(oneHourAgo, 10), strconv.FormatInt(now, 10)).Result()

    // 查询 MySQL 写入总数和最近一小时写入数
    writtenTotal := int64(0)
    writtenLastHour := int64(0)

    err5 := sqlxDB.Get(&writtenTotal, "SELECT COUNT(*) FROM tm_exam_answers")
    if err5 == nil {
        oneHourAgoTime := time.Now().Add(-1 * time.Hour)
        err5 = sqlxDB.Get(&writtenLastHour, "SELECT COUNT(*) FROM tm_exam_answers WHERE created_at >= ?", oneHourAgoTime)
    }

    response := gin.H{
        "code":    200,
        "message": "success",
        "data": gin.H{
            "redis": gin.H{
                "submitted":           submittedCount,
                "processed":           processedCount,
                "exam_answer_total":   examAnswerCount,
                "processed_last_hour": processedLastHour,
            },
            "mysql": gin.H{
                "written_total":     writtenTotal,
                "written_last_hour": writtenLastHour,
            },
        },
    }
    if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
        response["code"] = 500
        response["message"] = "error"
        response["error"] = gin.H{
            "submitted_error":   err1,
            "processed_error":   err2,
            "keys_error":        err3,
            "count_error":       err4,
            "mysql_write_error": err5,
        }
    }
    c.JSON(http.StatusOK, response)
}