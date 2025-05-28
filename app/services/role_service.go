package services

import (
	"fmt"
 	"gin-go-test/utils/generated/service"

	"gorm.io/gorm"
)

type RoleService struct {
	skeleton *service.RoleServiceSkeleton
	db       *gorm.DB
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{
		skeleton: service.NewRoleServiceSkeleton(db),
		db:       db,
	}
}

func (s *RoleService) GetCount() (int64, error) {
	if s.skeleton == nil {
		return 0, fmt.Errorf("skeleton is nil")
	}
	return s.skeleton.GetCount()
}

func (s *RoleService) List(page int, pageSize int) (interface{}, int64, error) {
	if s.skeleton == nil {
		return nil, 0, fmt.Errorf("skeleton is nil")
	}
	return s.skeleton.List(page, pageSize)
}

func (s *RoleService) GetDB() *gorm.DB {
	return s.db
}