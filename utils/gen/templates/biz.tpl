package biz

import (
	"gin-go-test/app/services"
	"gin-go-test/utils/generated/biz"
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

// GetCount 示例方法：调用 service 获取总数
func (b *{{ .ModelName }}Biz) GetCount() (int64, error) {
	return b.service.GetCount()
}