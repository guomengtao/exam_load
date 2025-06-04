package controller

import (
	"gin-go-test/app/biz"
	"gin-go-test/app/models"
	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
)

// BadmintonGameSkeleton 控制器骨架
type BadmintonGameSkeleton struct {
	biz *biz.BadmintonGameBiz
}

// NewBadmintonGameSkeleton 创建新的BadmintonGameSkeleton实例
func NewBadmintonGameSkeleton(biz *biz.BadmintonGameBiz) *BadmintonGameSkeleton {
	return &BadmintonGameSkeleton{biz: biz}
}

// CountHandler 获取记录总数
func (s *BadmintonGameSkeleton) CountHandler(c *gin.Context) {
	count, err := s.biz.GetCount()
	if err != nil {
		utils.Error(c, "获取记录总数失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{"count": count})
}

// ListHandler 获取记录列表
func (s *BadmintonGameSkeleton) ListHandler(c *gin.Context) {
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
func (s *BadmintonGameSkeleton) BatchCreateHandler(c *gin.Context) {
	var items []*models.BadmintonGame
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
func (s *BadmintonGameSkeleton) BatchUpdateHandler(c *gin.Context) {
	var items []*models.BadmintonGame
	if err := c.ShouldBindJSON(&items); err != nil {
		utils.Error(c, "请求参数错误: "+err.Error())
		return
	}
	err := s.biz.BatchUpdate(items)
	if err != nil {
		utils.Error(c, "批量更新记录失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{"message": "批量更新成功"})
}

// BatchDeleteHandler 批量删除记录
func (s *BadmintonGameSkeleton) BatchDeleteHandler(c *gin.Context) {
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
