package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"gin-go-test/app/services"
)

func GetRolesHandler(c *gin.Context) {
	// 默认值 page=1, page_size=10
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	result, err := services.GetRolesPaginated(page, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取角色失败"})
		return
	}

	// 返回格式：data + meta
	c.JSON(200, gin.H{
		"data": result.List,
		"meta": gin.H{
			"page":      result.Page,
			"page_size": result.PageSize,
			"total":     result.Total,
		},
	})
}