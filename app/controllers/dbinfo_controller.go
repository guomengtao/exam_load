package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// DBInfoController 数据库信息控制器
type DBInfoController struct {
	db *sqlx.DB
}

// NewDBInfoController 创建数据库信息控制器实例
func NewDBInfoController(db *sqlx.DB) *DBInfoController {
	return &DBInfoController{db: db}
}

// GetDBInfo 获取数据库信息
// @Summary 获取数据库结构信息
// @Description 返回数据库表结构和字段信息
// @Tags 系统信息
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/dbinfo [get]
func (c *DBInfoController) GetDBInfo(ctx *gin.Context) {
	var tables []string
	c.db.Select(&tables, "SHOW TABLES")

	tableInfo := make(map[string]interface{})
	for _, table := range tables {
		var columns []map[string]interface{}
		c.db.Select(&columns, "SHOW COLUMNS FROM "+table)
		tableInfo[table] = columns
	}

	ctx.JSON(200, gin.H{
		"tables": tableInfo,
	})
}
