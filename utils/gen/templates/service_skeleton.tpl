package service

import (
	"fmt"
	"gin-go-test/app/models"
	"gorm.io/gorm"
)

type {{ .ServiceName }}ServiceSkeleton struct {
	db *gorm.DB
}

func New{{ .ServiceName }}ServiceSkeleton(db *gorm.DB) *{{ .ServiceName }}ServiceSkeleton {
	return &{{ .ServiceName }}ServiceSkeleton{
		db: db,
	}
}

// GetCount 返回数据库中该模型的总记录数
func (s *{{ .ServiceName }}ServiceSkeleton) GetCount() (int64, error) {
	var count int64
	if err := s.db.Model(&models.{{ .ModelName }}{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to get count: %w", err)
	}
	return count, nil
}

// List 根据分页参数返回数据列表和总数
func (s *{{ .ServiceName }}ServiceSkeleton) List(page int, pageSize int) ([]models.{{ .ModelName }}, int64, error) {
	var items []models.{{ .ModelName }}
	var total int64

	offset := (page - 1) * pageSize

	err := s.db.Model(&models.{{ .ModelName }}{}).
		Count(&total).
		Limit(pageSize).
		Offset(offset).
		Find(&items).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list items: %w", err)
	}

	return items, total, nil
}