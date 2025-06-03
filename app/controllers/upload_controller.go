package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"time"
)

// UploadController 上传控制器
type UploadController struct {
	uploadDir string
}

// NewUploadController 创建上传控制器实例
func NewUploadController(uploadDir string) *UploadController {
	return &UploadController{
		uploadDir: uploadDir,
	}
}

// UploadFile 上传文件
// @Summary 上传文件
// @Description 处理文件上传请求
// @Tags 文件操作
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {object} map[string]string
// @Router /api/upload [post]
func (c *UploadController) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "文件上传失败",
		})
		return
	}

	// 生成文件名
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	filepath := filepath.Join(c.uploadDir, filename)

	// 保存文件
	if err := ctx.SaveUploadedFile(file, filepath); err != nil {
		ctx.JSON(500, gin.H{
			"error": "文件保存失败",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message":  "文件上传成功",
		"filename": filename,
		"path":     filepath,
	})
} 