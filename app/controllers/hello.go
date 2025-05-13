package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// HelloHandler 是一个简单的 Hello World 控制器
func HelloHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello from app/controllers!",
    })
}