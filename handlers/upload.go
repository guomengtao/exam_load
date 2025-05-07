package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadExamImage(c *gin.Context) {
	// 1. 获取上传文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请选择上传文件",
		})
		return
	}

	// 2. 基础文件类型验证
	ext := filepath.Ext(file.Filename)
	switch ext {
	case ".jpg", ".jpeg", ".png":
		// 允许的格式
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "仅支持jpg/jpeg/png格式图片",
		})
		return
	}

	// 3. 生成唯一文件名
	filename := "exam_" + uuid.New().String() + ext
	savePath := filepath.Join("static", "uploads", "images", "exam", filename)

	// 4. 保存文件
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "文件保存失败",
			"detail": err.Error(),
		})
		return
	}

	// 5. 返回结果
	c.JSON(http.StatusOK, gin.H{
		"url":  "/static/uploads/images/exam/" + filename,
		"name": filename,
		"size": file.Size,
	})
}