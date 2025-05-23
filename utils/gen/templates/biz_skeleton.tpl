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
func (s *{{ .ModelName }}BizSkeleton) GetCount() (int64, error) {
	var count int64
	err := s.db.Model(&models.{{ .ModelName }}{}).Count(&count).Error
	return count, err
}

// List 获取分页数据
func (s *{{ .ModelName }}BizSkeleton) List(page int, pageSize int) ([]models.{{ .ModelName }}, int64, error) {
	var items []models.{{ .ModelName }}
	var total int64

	offset := (page - 1) * pageSize
	err := s.db.Model(&models.{{ .ModelName }}{}).
		Count(&total).
		Limit(pageSize).
		Offset(offset).
		Find(&items).Error

	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}
