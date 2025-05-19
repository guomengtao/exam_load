package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gin-go-test/app/services"
)

// ExportAnswersHandler 是导出答题记录的接口
func ExportAnswersHandler(c *gin.Context) {
    examUUID := c.Query("exam_uuid")
    school := c.Query("school")
    limitStr := c.DefaultQuery("limit", "10000")

    if examUUID == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    400,
            "message": "缺少必要参数 exam_uuid",
        })
        return
    }

    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit <= 0 {
        limit = 10000
    }

    filepath, err := services.ExportAnswersToCSV(examUUID, school, limit, 0)
    if err != nil {
        // 如果是没有数据的错误，返回404
        if err.Error() == "没有对应数据，无导出文件" {
            c.JSON(http.StatusNotFound, gin.H{
                "code":    404,
                "message": err.Error(),
            })
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{
            "code":    500,
            "message": "导出失败: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code":     200,
        "message":  "导出成功",
        "filepath": filepath,
    })
}