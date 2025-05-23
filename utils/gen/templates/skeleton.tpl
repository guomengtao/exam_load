package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-go-test/app/biz"
)

// {{.ControllerName}}Skeleton 骨架层
type {{.ControllerName}}Skeleton struct {
	biz *biz.{{.ControllerName}}Biz
}

// New{{.ControllerName}}Skeleton 构造函数，传入业务层实例
func New{{.ControllerName}}Skeleton(biz *biz.{{.ControllerName}}Biz) *{{.ControllerName}}Skeleton {
	return &{{.ControllerName}}Skeleton{
		biz: biz,
	}
}

// Hello 返回默认信息
func (s *{{.ControllerName}}Skeleton) Hello() string {
	return "{{.HelloMessage}}"
}

// CountHandler 查询总数，调用业务层
func (s *{{.ControllerName}}Skeleton) CountHandler(c *gin.Context) {
	count, err := s.biz.GetCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}