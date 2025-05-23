package services

import (
	"testing"

	"gin-go-test/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserService_GetCount(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("数据库连接失败: %v", err)
	}

	// 使用真实模型 User 来建表
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		t.Fatalf("建表失败: %v", err)
	}

	service := NewUserService(db)
	count, err := service.GetCount()
	if err != nil {
		t.Errorf("GetCount 执行失败: %v", err)
	}

	t.Logf("当前记录数: %d", count)
}