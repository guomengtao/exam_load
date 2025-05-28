package services

import (
	"fmt"
	"gin-go-test/app/models"
	"gin-go-test/utils/generated/service"

	"gorm.io/gorm"
)

type {{.ServiceName}}Service struct {
	skeleton *service.{{.ServiceName}}ServiceSkeleton
	db       *gorm.DB
}

func New{{.ServiceName}}Service(db *gorm.DB) *{{.ServiceName}}Service {
	return &{{.ServiceName}}Service{
		skeleton: service.New{{.ServiceName}}ServiceSkeleton(db),
		db:       db,
	}
}

func (s *{{.ServiceName}}Service) GetCount() (int64, error) {
	if s.skeleton == nil {
		return 0, fmt.Errorf("skeleton is nil")
	}
	return s.skeleton.GetCount()
}

func (s *{{.ServiceName}}Service) List(page int, pageSize int) (interface{}, int64, error) {
	if s.skeleton == nil {
		return nil, 0, fmt.Errorf("skeleton is nil")
	}
	return s.skeleton.List(page, pageSize)
}

func (s *{{.ServiceName}}Service) GetDB() *gorm.DB {
	return s.db
}