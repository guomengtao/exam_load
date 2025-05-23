package biz

import (
	"gorm.io/gorm"
	"gin-go-test/app/models"
)

type UserBizSkeleton struct {
	db *gorm.DB
}

// NewUserBizSkeleton 构造函数
func NewUserBizSkeleton(db *gorm.DB) *UserBizSkeleton {
	return &UserBizSkeleton{db: db}
}

// SayHello 示例骨架方法
func (s *UserBizSkeleton) SayHello() string {
	return "hello from UserBizSkeleton"
}

// GetCount 真实统计表记录数
func (s *UserBizSkeleton) GetCount() int64 {
	var count int64
	s.db.Model(&models.User{}).Count(&count)
	return count
}
