package controllers

import "github.com/gin-gonic/gin"

// 假设 DBX 是 interface{}，实际可根据需要调整

type DBInfoController struct{
	db interface{}
}

func NewDBInfoController(db interface{}) *DBInfoController {
	return &DBInfoController{db: db}
}

func (c *DBInfoController) GetDBInfo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"dbinfo": "mock info"})
} 