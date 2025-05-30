package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code int         `json:"code"` // 0 表示成功，非0表示错误
	Msg  string      `json:"msg"`  // 响应消息
	Data interface{} `json:"data"` // 响应数据
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// Error 错误响应
func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 1,
		Msg:  msg,
		Data: nil,
	})
}

// PageSuccess 分页成功响应
func PageSuccess(c *gin.Context, list interface{}, total int64) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: gin.H{
			"list":  list,
			"total": total,
		},
	})
}

// GetPage 获取页码
func GetPage(c *gin.Context) int {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}
	return page
}

// GetPageSize 获取每页数量
func GetPageSize(c *gin.Context) int {
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	return pageSize
}