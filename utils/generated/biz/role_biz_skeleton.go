package biz

import (
	"fmt"
	"gin-go-test/app/services"
)

type RoleBizSkeleton struct {
	service *services.RoleService
}

func NewRoleBizSkeleton(service *services.RoleService) *RoleBizSkeleton {
	return &RoleBizSkeleton{service: service}
}

// GetCount 示例方法：调用 Server 层获取总数
func (b *RoleBizSkeleton) GetCount() (int64, error) {
	if b.service == nil {
		return 0, fmt.Errorf("service is nil")
	}
	return b.service.GetCount()
}

// List 示例方法：调用 Server 层获取分页数据
func (b *RoleBizSkeleton) List(page int, pageSize int) ([]interface{}, int64, error) {
	if b.service == nil {
		return nil, 0, fmt.Errorf("service is nil")
	}
	data, total, err := b.service.List(page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	list, ok := data.([]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("invalid data type returned from service.List")
	}
	return list, total, nil
}