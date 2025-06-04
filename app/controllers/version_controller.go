package controllers

import (
	"github.com/gin-gonic/gin"
)

// VersionController 版本控制器
type VersionController struct{}

// NewVersionController 创建版本控制器实例
func NewVersionController() *VersionController {
	return &VersionController{}
}

// GetVersion 获取版本信息
// @Summary 获取系统版本信息
// @Description 返回当前系统的版本号和其他版本相关信息
// @Tags 系统信息
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/version [get]
func (c *VersionController) GetVersion(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"version": "1.0.0",
		"build":   "2024-03-21",
		"status":  "stable",
	})
}
