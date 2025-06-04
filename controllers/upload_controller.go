package controllers

import "github.com/gin-gonic/gin"

type UploadController struct {
	dir string
}

func NewUploadController(dir string) *UploadController {
	return &UploadController{dir: dir}
}

func (c *UploadController) UploadFile(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"upload": "success"})
}
