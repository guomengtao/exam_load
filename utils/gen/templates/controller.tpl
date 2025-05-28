package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-go-test/utils/generated/controller"
	"gin-go-test/app/biz"
	"gin-go-test/app/services"
	"gorm.io/gorm"
	"strconv"
)

// {{.ControllerName}}Controller 控制器示例
type {{.ControllerName}}Controller struct {
	skeleton *controller.{{.ControllerName}}Skeleton
}

// Register{{.ControllerName}}Routes 注册 {{.ControllerName}} 相关路由
func Register{{.ControllerName}}Routes(router *gin.Engine, db *gorm.DB) {
	group := router.Group("/api/{{.RoutePath}}")
	ctrl := &{{.ControllerName}}Controller{
		skeleton: controller.New{{.ControllerName}}Skeleton(
			biz.New{{.ControllerName}}Biz(
				services.New{{.ControllerName}}Service(db),
			),
		),
	}

	group.GET("/hello", ctrl.HelloHandler)
	group.GET("/count", ctrl.CountHandler)
	group.GET("/list", ctrl.ListHandler)
	// 这里以后根据需要添加接口，比如分页列表，增删改查等
}

// HelloHandler 示例接口，调用骨架层对应方法
func (ctrl *{{.ControllerName}}Controller) HelloHandler(c *gin.Context) {
	result := ctrl.skeleton.Hello()
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

// CountHandler 查询真实总数，调用骨架层方法
func (ctrl *{{.ControllerName}}Controller) CountHandler(c *gin.Context) {
	count, err := ctrl.skeleton.GetCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

// ListHandler 分页列表接口，调用骨架层 List 方法
func (ctrl *{{.ControllerName}}Controller) ListHandler(c *gin.Context) {
	page := 1
	pageSize := 10
	if p := c.Query("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	if ps := c.Query("pageSize"); ps != "" {
		if v, err := strconv.Atoi(ps); err == nil && v > 0 {
			pageSize = v
		}
	}
	list, total, err := ctrl.skeleton.List(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list":  list,
		"total": total,
		"page":  page,
		"pageSize": pageSize,
	})
}