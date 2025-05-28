package service

import (
	"fmt"
	"gin-go-test/app/models"
	"gorm.io/gorm"
)

type RoleServiceSkeleton struct {
	db *gorm.DB
}

func NewRoleServiceSkeleton(db *gorm.DB) *RoleServiceSkeleton {
	return &RoleServiceSkeleton{
		db: db,
	}
}

// GetCount 返回数据库中该模型的总记录数
func (s *RoleServiceSkeleton) GetCount() (int64, error) {
	var count int64
	if err := s.db.Model(&models.Role{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to get count: %w", err)
	}
	return count, nil
}

// List 根据分页参数返回数据列表和总数
func (s *RoleServiceSkeleton) List(page int, pageSize int) ([]models.Role, int64, error) {
	var items []models.Role
	var total int64

	offset := (page - 1) * pageSize

	err := s.db.Model(&models.Role{}).
		Count(&total).
		Limit(pageSize).
		Offset(offset).
		Find(&items).Error
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list items: %w", err)
	}

	return items, total, nil
}