package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// SourceController 数据源控制器
type SourceController struct {
	db *sqlx.DB
}

// NewSourceController 创建数据源控制器实例
func NewSourceController(db *sqlx.DB) *SourceController {
	return &SourceController{db: db}
}

// CheckSource 检查数据源
// @Summary 检查数据源状态
// @Description 检查数据库连接和数据源状态
// @Tags 系统信息
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/source/check [get]
func (c *SourceController) CheckSource(ctx *gin.Context) {
	var result int
	c.db.Get(&result, "SELECT 1")

	if result == 1 {
		ctx.JSON(200, gin.H{
			"status":  "ok",
			"message": "数据源连接正常",
		})
	} else {
		ctx.JSON(500, gin.H{
			"status":  "error",
			"message": "数据源连接异常",
		})
	}
}
