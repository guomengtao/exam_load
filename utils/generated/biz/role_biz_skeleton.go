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
func (s *RoleBizSkeleton) GetCount() int64 {
	var count int64
	s.db.Model(&models.Role{}).Count(&count)
	return count
}
