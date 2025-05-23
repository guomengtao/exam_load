package {{ .Package }}

import (
	"gorm.io/gorm"
	"gin-go-test/app/models"
)

type {{ .ModelName }}BizSkeleton struct {
	db *gorm.DB
}

// New{{ .ModelName }}BizSkeleton 构造函数
func New{{ .ModelName }}BizSkeleton(db *gorm.DB) *{{ .ModelName }}BizSkeleton {
	return &{{ .ModelName }}BizSkeleton{db: db}
}

// SayHello 示例骨架方法
func (s *{{ .ModelName }}BizSkeleton) SayHello() string {
	return "hello from {{ .ModelName }}BizSkeleton"
}

// GetCount 真实统计表记录数
func (s *{{ .ModelName }}BizSkeleton) GetCount() int64 {
	var count int64
	s.db.Model(&models.{{ .ModelName }}{}).Count(&count)
	return count
}
