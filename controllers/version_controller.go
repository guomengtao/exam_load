package controllers

import "github.com/gin-gonic/gin"

type VersionController struct{}

func NewVersionController() *VersionController {
	return &VersionController{}
}

func (c *VersionController) GetVersion(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"version": "v0.0.1"})
} 