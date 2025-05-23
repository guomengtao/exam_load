

package services

import (
	"gin-go-test/app/models"
	"gin-go-test/utils/generated/service"

	"gorm.io/gorm"
)

type UserService struct {
	skeleton *service.UserServiceSkeleton
	db       *gorm.DB
}
func (s *UserService) GetDB() *gorm.DB {
    return s.db
}

func NewUserService(db *gorm.DB) *UserService {
	if db == nil {
		panic("NewUserService got nil db")
	}
	return &UserService{
		skeleton: &service.UserServiceSkeleton{},
		db:       db,
	}
}

func (s *UserService) GetCount() (int64, error) {
	var count int64
	err := s.db.Model(&models.User{}).Count(&count).Error
	return count, err
}