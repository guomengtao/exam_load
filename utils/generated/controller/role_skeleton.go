// ⚠️ 本文件为控制器骨架模板，禁止直接修改任何生成器生成的文件！
// 如需调整，请修改本模板，并通过 go run utils/gen/gen.go -table=表名 -cmd=c 等命令重新生成覆盖。

package controller

import (
	"gin-go-test/app/biz"
	"gin-go-test/app/services"
	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RoleSkeleton 控制器骨架
type RoleSkeleton struct {
	biz *biz.RoleBiz
}

// NewRoleSkeleton 创建新的RoleSkeleton实例
func NewRoleSkeleton(biz *biz.RoleBiz) *RoleSkeleton {
	return &RoleSkeleton{biz: biz}
}

// CountHandler 获取记录总数
func (s *RoleSkeleton) CountHandler(c *gin.Context) {
	count, err := s.biz.GetCount()
	if err != nil {
		utils.Error(c, "获取记录总数失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{"count": count})
}

// ListHandler 获取记录列表
func (s *RoleSkeleton) ListHandler(c *gin.Context) {
	page := utils.GetPage(c)
	pageSize := utils.GetPageSize(c)
	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "desc")
	items, err := s.biz.ListWithOrder(page, pageSize, sort, order)
	if err != nil {
		utils.Error(c, "获取记录列表失败: "+err.Error())
		return
	}
	utils.PageSuccess(c, items, int64(len(items)))
}

// BatchCreateHandler 批量创建记录
func (s *RoleSkeleton) BatchCreateHandler(c *gin.Context) {
	var items []*models.Role
	if err := c.ShouldBindJSON(&items); err != nil {
		utils.Error(c, "请求参数错误: "+err.Error())
		return
	}
	createdItems, err := s.biz.BatchCreate(items)
	if err != nil {
		utils.Error(c, "批量创建记录失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{"items": createdItems})
}

// BatchUpdateHandler 批量更新记录
func (s *RoleSkeleton) BatchUpdateHandler(c *gin.Context) {
	var items []*models.Role
	if err := c.ShouldBindJSON(&items); err != nil {
		utils.Error(c, "请求参数错误: "+err.Error())
		return
	}
	if err := s.biz.BatchUpdate(items); err != nil {
		utils.Error(c, "批量更新记录失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{"message": "更新成功"})
}

// BatchDeleteHandler 批量删除记录
func (s *RoleSkeleton) BatchDeleteHandler(c *gin.Context) {
	var ids []int
	if err := c.ShouldBindJSON(&ids); err != nil {
		utils.Error(c, "请求参数错误: "+err.Error())
		return
	}
	if err := s.biz.BatchDelete(ids); err != nil {
		utils.Error(c, "批量删除记录失败: "+err.Error())
		return
	}
	utils.Success(c, gin.H{"message": "删除成功"})
}

// GetDetail 获取单条详情
func (s *RoleSkeleton) GetDetail(c *gin.Context) {
	id := c.Param("id")
	item, err := s.biz.GetDetail(id)
	if err != nil {
		utils.Error(c, "未找到记录: "+err.Error())
		return
	}
	utils.Success(c, item)
}

// RegisterRoleRoutes 注册 Role 相关路由
func RegisterRoleRoutes(router *gin.Engine, db *gorm.DB) {
	skeleton := NewRoleSkeleton(biz.NewRoleBiz(services.NewRoleService(db)))
	// 注册路由
	router.GET("/role", skeleton.ListHandler)
	router.POST("/role", skeleton.BatchCreateHandler)
	router.PUT("/role/:id", skeleton.BatchUpdateHandler)
	router.DELETE("/role/:id", skeleton.BatchDeleteHandler)
	router.GET("/role/:id", skeleton.GetDetail)
}