package biz

import (
	"gin-go-test/app/services"
	"gin-go-test/utils/generated/biz"
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

// GetCount 示例方法：调用 service 获取总数
func (b *RoleBiz) GetCount() (int64, error) {
	return b.service.GetCount()
}