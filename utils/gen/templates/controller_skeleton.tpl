// ⚠️ 本文件为控制器骨架模板，禁止直接修改任何生成器生成的文件！
// 如需调整，请修改本模板，并通过 go run utils/gen/gen.go -table=表名 -cmd=c 等命令重新生成覆盖。

package controller

import (
	"gin-go-test/app/biz"
	"gin-go-test/app/models"
	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
)

// {{.ModelName}}Skeleton 控制器骨架
type {{.ModelName}}Skeleton struct {
	biz *biz.{{.ModelName}}Biz
}

// New{{.ModelName}}Skeleton 创建新的{{.ModelName}}Skeleton实例
func New{{.ModelName}}Skeleton(biz *biz.{{.ModelName}}Biz) *{{.ModelName}}Skeleton {
	return &{{.ModelName}}Skeleton{biz: biz}
}

// CountHandler 获取记录总数
func (s *{{.ModelName}}Skeleton) CountHandler(c *gin.Context) {
	count, err := s.biz.GetCount()
	if err != nil {
		utils.Error(c, "获取记录总数失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{"count": count})
}

// ListHandler 获取记录列表
func (s *{{.ModelName}}Skeleton) ListHandler(c *gin.Context) {
	page := utils.GetPage(c)
	pageSize := utils.GetPageSize(c)
	items, err := s.biz.List(page, pageSize)
	if err != nil {
		utils.Error(c, "获取记录列表失败: "+err.Error())
		return
	}
	utils.PageSuccess(c, items, int64(len(items)))
}

// BatchCreateHandler 批量创建记录
func (s *{{.ModelName}}Skeleton) BatchCreateHandler(c *gin.Context) {
	var items []*models.{{.ModelName}}
	if err := c.ShouldBindJSON(&items); err != nil {
		utils.Error(c, "请求参数错误: "+err.Error())
		return
	}
	createdItems, errs := s.biz.BatchCreate(items)
	if len(errs) > 0 {
		utils.Error(c, "批量创建记录失败: "+errs[0].Error())
		return
	}
	utils.Success(c, gin.H{"items": createdItems})
}

// BatchUpdateHandler 批量更新记录
func (s *{{.ModelName}}Skeleton) BatchUpdateHandler(c *gin.Context) {
	var items []*models.{{.ModelName}}
	if err := c.ShouldBindJSON(&items); err != nil {
		utils.Error(c, "请求参数错误: "+err.Error())
		return
	}
	updatedItems, errs := s.biz.BatchUpdate(items)
	if len(errs) > 0 {
		utils.Error(c, "批量更新记录失败: "+errs[0].Error())
		return
	}
	utils.Success(c, gin.H{"items": updatedItems})
}

// BatchDeleteHandler 批量删除记录
func (s *{{.ModelName}}Skeleton) BatchDeleteHandler(c *gin.Context) {
	var ids []int
	if err := c.ShouldBindJSON(&ids); err != nil {
		utils.Error(c, "请求参数错误: "+err.Error())
		return
	}
	errs := s.biz.BatchDelete(ids)
	if len(errs) > 0 {
		utils.Error(c, "批量删除记录失败: "+errs[0].Error())
		return
	}
	utils.Success(c, gin.H{"message": "删除成功"})
}