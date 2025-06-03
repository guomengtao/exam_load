package controllers

import "github.com/gin-gonic/gin"

type SourceController struct{
	db interface{}
}

func NewSourceController(db interface{}) *SourceController {
	return &SourceController{db: db}
}

func (c *SourceController) CheckSource(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"source": "ok"})
} 