package services

import (
	"gin-go-test/app/models"
	"gin-go-test/utils/generated/service"

	"gorm.io/gorm"
)

type RoleService struct {
	skeleton *service.RoleServiceSkeleton
	db       *gorm.DB
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{
		skeleton: &service.RoleServiceSkeleton{},
		db:       db,
	}
}

func (s *RoleService) GetCount() (int64, error) {
	var count int64
	err := s.db.Model(&models.Role{}).Count(&count).Error
	return count, err
}

func (s *RoleService) GetDB() *gorm.DB {
	return s.db
}