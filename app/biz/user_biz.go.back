package biz

import (
	"gin-go-test/app/services"
	"gin-go-test/utils/generated/biz"
)

// UserBiz 业务逻辑层
type UserBiz struct {
	service  *services.UserService
	skeleton *biz.UserBizSkeleton
}

// NewUserBiz 构造函数
func NewUserBiz(service *services.UserService) *UserBiz {
	return &UserBiz{
		service:  service,
		skeleton: biz.NewUserBizSkeleton(service.GetDB()),
	}
}

// GetCount 示例方法：调用 service 获取总数
func (b *UserBiz) GetCount() (int64, error) {
	return b.service.GetCount()
}