package biz

import (
	"gin-go-test/app/services"
	"gin-go-test/utils/generated/biz"
	"gin-go-test/app/models"
	"fmt"
)

// RoleBiz 业务逻辑层
type RoleBiz struct {
	service  *services.RoleService
	skeleton *biz.RoleBizSkeleton
}

// NewRoleBiz 构造函数
func NewRoleBiz(service *services.RoleService) *RoleBiz {
	return &RoleBiz{
		service:  service,
		skeleton: biz.NewRoleBizSkeleton(service.GetDB()),
	}
}

// GetCount 示例方法：调用骨架层获取总数
func (b *RoleBiz) GetCount() (int64, error) {
	if b.skeleton == nil {
		return 0, fmt.Errorf("skeleton is nil")
	}
	return b.skeleton.GetCount()
}

// List 获取分页数据，调用骨架层实现
func (b *RoleBiz) List(page int, pageSize int) ([]models.Role, int64, error) {
    return b.skeleton.List(page, pageSize)
}