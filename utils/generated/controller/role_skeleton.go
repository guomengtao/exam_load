package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-go-test/app/biz"
	    "strconv"  // 👈 加上这个
)

// RoleSkeleton 骨架层
type RoleSkeleton struct {
	biz *biz.RoleBiz
}

// NewRoleSkeleton 构造函数，传入业务层实例
func NewRoleSkeleton(biz *biz.RoleBiz) *RoleSkeleton {
	return &RoleSkeleton{
		biz: biz,
	}
}

// Hello 返回默认信息
func (s *RoleSkeleton) Hello() string {
	return "hello123"
}

// CountHandler 查询总数，调用业务层
func (s *RoleSkeleton) CountHandler(c *gin.Context) {
	count, err := s.biz.GetCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

// ListHandler 分页列表接口，调用业务层
func (s *RoleSkeleton) ListHandler(c *gin.Context) {
	page := 1
	pageSize := 10

	if p := c.Query("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	if ps := c.Query("pageSize"); ps != "" {
		if parsedPageSize, err := strconv.Atoi(ps); err == nil && parsedPageSize > 0 {
			pageSize = parsedPageSize
		}
	}

	list, total, err := s.biz.List(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"list": list, "total": total})
}