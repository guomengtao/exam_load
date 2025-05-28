package biz

import (
	"gin-go-test/app/services"
	"gin-go-test/utils/generated/biz"
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
		skeleton: biz.New{{ .ModelName }}BizSkeleton(service),
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
func (b *{{ .ModelName }}Biz) List(page int, pageSize int) ([]interface{}, int64, error) {
	if b.skeleton == nil {
		return nil, 0, fmt.Errorf("skeleton is nil")
	}
	return b.skeleton.List(page, pageSize)
}