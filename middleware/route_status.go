package middleware

import (
	"gin-go-test/app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// RouteStatusMiddleware 路由状态中间件
func RouteStatusMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前请求的方法和路径
		method := c.Request.Method
		path := c.Request.URL.Path

		// 查询路由状态
		var routeStatus models.RouteStatus
		result := db.Where("method = ? AND path = ?", method, path).First(&routeStatus)

		// 如果路由存在且状态为暂停或废弃，返回错误
		if result.Error == nil {
			switch routeStatus.Status {
			case "paused":
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"code":    503,
					"message": "该接口已暂停服务",
				})
				c.Abort()
				return
			case "deprecated":
				c.JSON(http.StatusGone, gin.H{
					"code":    410,
					"message": "该接口已废弃",
				})
				c.Abort()
				return
			}

			// 更新访问统计
			routeStatus.VisitCount++
			routeStatus.LastVisitedAt = time.Now()
			db.Save(&routeStatus)
		}

		c.Next()
	}
}
