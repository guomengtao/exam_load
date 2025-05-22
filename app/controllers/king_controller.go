package controllers

import (
	"github.com/gin-gonic/gin"
	"gin-go-test/utils/generated/controller"
	"net/http"
)

type KingController struct {
	skeleton *controller.KingSkeleton
}

func NewKingController() *KingController {
	return &KingController{
		skeleton: &controller.KingSkeleton{},
	}
}

func (ctrl *KingController) Get(c *gin.Context) {
	result := ctrl.skeleton.Hello()
	result = "hello456"
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func RegisterKingRoutes(r *gin.Engine) {
	ctrl := NewKingController()
	group := r.Group("/api/king")
	group.GET("/", ctrl.Get)
}