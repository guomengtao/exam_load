package services

import (
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
		skeleton: &service.{{.ServiceName}}ServiceSkeleton{},
		db:       db,
	}
}

func (s *{{.ServiceName}}Service) GetCount() (int64, error) {
	var count int64
	err := s.db.Model(&models.{{.ServiceName}}{}).Count(&count).Error
	return count, err
}

func (s *{{.ServiceName}}Service) GetDB() *gorm.DB {
	return s.db
}