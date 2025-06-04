package controllers

import "github.com/gin-gonic/gin"

type StatusController struct{}

func NewStatusController() *StatusController {
	return &StatusController{}
}

func (c *StatusController) GetStatus(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "ok"})
}
