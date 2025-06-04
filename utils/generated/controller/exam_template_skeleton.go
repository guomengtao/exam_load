// ⚠️ 本文件为控制器骨架模板，禁止直接修改任何生成器生成的文件！
// 如需调整，请修改本模板，并通过 go run utils/gen/gen.go -table=表名 -cmd=c 等命令重新生成覆盖。

package controller

import (
	"gin-go-test/app/biz"
	"gin-go-test/app/models"
	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
)

// ExamTemplateSkeleton 控制器骨架
type ExamTemplateSkeleton struct {
	biz *biz.ExamTemplateBiz
}

// NewExamTemplateSkeleton 创建新的ExamTemplateSkeleton实例
func NewExamTemplateSkeleton(biz *biz.ExamTemplateBiz) *ExamTemplateSkeleton {
	return &ExamTemplateSkeleton{biz: biz}
}

// CountHandler 获取记录总数
func (s *ExamTemplateSkeleton) CountHandler(c *gin.Context) {
	count, err := s.biz.GetCount()
	if err != nil {
		utils.Error(c, "获取记录总数失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{"count": count})
}

// ListHandler 获取记录列表
func (s *ExamTemplateSkeleton) ListHandler(c *gin.Context) {
	page := utils.GetPage(c)
	pageSize := utils.GetPageSize(c)
	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "desc")
	items, errs := s.biz.ListWithOrder(page, pageSize, sort, order)
	if len(errs) > 0 {
		utils.Error(c, "获取记录列表失败: "+errs[0].Error())
		return
	}
	utils.PageSuccess(c, items, int64(len(items)))
}

// BatchCreateHandler 批量创建记录
func (s *ExamTemplateSkeleton) BatchCreateHandler(c *gin.Context) {
	var items []*models.ExamTemplate
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
func (s *ExamTemplateSkeleton) BatchUpdateHandler(c *gin.Context) {
	var items []*models.ExamTemplate
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
func (s *ExamTemplateSkeleton) BatchDeleteHandler(c *gin.Context) {
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

// GetDetail 获取单条详情
func (s *ExamTemplateSkeleton) GetDetail(c *gin.Context) {
	id := c.Param("id")
	item, err := s.biz.GetDetail(id)
	if err != nil {
		utils.Error(c, "未找到记录: "+err.Error())
		return
	}
	utils.Success(c, item)
}

// 路由注册
group.GET("/:id", s.GetDetail)