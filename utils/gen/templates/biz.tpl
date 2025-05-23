package biz

import (
	"gin-go-test/app/services"
	"gin-go-test/utils/generated/biz"
	"gin-go-test/app/models"
	"fmt"
)

// {{ .ModelName }}Biz 业务逻辑层
type {{ .ModelName }}Biz struct {
	service  *services.{{ .ModelName }}Service
	skeleton *biz.{{ .ModelName }}BizSkeleton
}

// New{{ .ModelName }}Biz 构造函数
func New{{ .ModelName }}Biz(service *services.{{ .ModelName }}Service) *{{ .ModelName }}Biz {
	return &{{ .ModelName }}Biz{
		service:  service,
		skeleton: biz.New{{ .ModelName }}BizSkeleton(service.GetDB()),
	}
}

// GetCount 示例方法：调用骨架层获取总数
func (b *{{ .ModelName }}Biz) GetCount() (int64, error) {
	if b.skeleton == nil {
		return 0, fmt.Errorf("skeleton is nil")
	}
	return b.skeleton.GetCount()
}

// List 获取分页数据，调用骨架层实现
func (b *{{ .ModelName }}Biz) List(page int, pageSize int) ([]models.{{ .ModelName }}, int64, error) {
    return b.skeleton.List(page, pageSize)
}