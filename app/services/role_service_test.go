package services

import (
	"testing"

	"gin-go-test/app/models"
	"gin-go-test/utils/generated/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("无法连接测试数据库: %v", err)
	}

	err = db.AutoMigrate(&models.Role{})
	if err != nil {
		t.Fatalf("自动迁移失败: %v", err)
	}

	// 插入模拟数据
	db.Create(&models.Role{Name: "管理员", Desc: "系统管理员"})
	db.Create(&models.Role{Name: "用户", Desc: "普通用户"})

	return db
}

func TestRoleServiceSkeleton_GetCount(t *testing.T) {
	db := setupTestDB(t)
	skeleton := service.NewRoleServiceSkeleton(db)

	count, err := skeleton.GetCount()
	if err != nil {
		t.Errorf("获取总数失败: %v", err)
	}
	if count != 2 {
		t.Errorf("期望总数为 2，实际为 %d", count)
	}
}

func TestRoleServiceSkeleton_List(t *testing.T) {
	db := setupTestDB(t)
	skeleton := service.NewRoleServiceSkeleton(db)

	items, total, err := skeleton.List(1, 10)
	if err != nil {
		t.Errorf("分页查询失败: %v", err)
	}
	if total != 2 {
		t.Errorf("期望总数为 2，实际为 %d", total)
	}
	if len(items) != 2 {
		t.Errorf("期望返回 2 条数据，实际为 %d", len(items))
	}
}
