package biz

import (
	"gorm.io/gorm"
	"gin-go-test/app/models"
)

type RoleBizSkeleton struct {
	db *gorm.DB
}

// NewRoleBizSkeleton 构造函数
func NewRoleBizSkeleton(db *gorm.DB) *RoleBizSkeleton {
	return &RoleBizSkeleton{db: db}
}

// SayHello 示例骨架方法
func (s *RoleBizSkeleton) SayHello() string {
	return "hello from RoleBizSkeleton"
}

// GetCount 真实统计表记录数
func (s *RoleBizSkeleton) GetCount() (int64, error) {
	var count int64
	err := s.db.Model(&models.Role{}).Count(&count).Error
	return count, err
}

// List 获取分页数据
func (s *RoleBizSkeleton) List(page int, pageSize int) ([]models.Role, int64, error) {
	var items []models.Role
	var total int64

	offset := (page - 1) * pageSize
	err := s.db.Model(&models.Role{}).
		Count(&total).
		Limit(pageSize).
		Offset(offset).
		Find(&items).Error

	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}
