package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-go-test/utils/generated/controller"
	"gin-go-test/app/services"
	"gin-go-test/app/biz"
	"gorm.io/gorm"
)

// BoyController 控制器示例
type BoyController struct {
	skeleton *controller.BoySkeleton
	biz      *biz.UserBiz
}

// RegisterBoyRoutes 注册 Boy 相关路由
func RegisterBoyRoutes(router *gin.Engine, db *gorm.DB) {
	group := router.Group("/api/boy")
	ctrl := &BoyController{
		skeleton: &controller.BoySkeleton{},
		biz:      biz.NewUserBiz(services.NewUserService(db)),
	}

	group.GET("/hello", ctrl.HelloHandler)
	group.GET("/count", ctrl.CountHandler)
}

// HelloHandler 示例接口，调用骨架层对应方法
func (ctrl *BoyController) HelloHandler(c *gin.Context) {
	result := ctrl.skeleton.Hello()
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func (ctrl *BoyController) CountHandler(c *gin.Context) {
	count, err := ctrl.biz.GetCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取总数失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}