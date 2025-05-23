package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-go-test/app/biz"
)

// UserSkeleton 骨架层
type UserSkeleton struct {
	biz *biz.UserBiz
}

// NewUserSkeleton 构造函数，传入业务层实例
func NewUserSkeleton(biz *biz.UserBiz) *UserSkeleton {
	return &UserSkeleton{
		biz: biz,
	}
}

// Hello 返回默认信息
func (s *UserSkeleton) Hello() string {
	return "hello123"
}

// CountHandler 查询总数，调用业务层
func (s *UserSkeleton) CountHandler(c *gin.Context) {
	count, err := s.biz.GetCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}