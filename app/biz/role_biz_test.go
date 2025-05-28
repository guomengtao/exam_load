package biz

import (
	"testing"

 	"gin-go-test/app/services"
	"gin-go-test/utils"
)

func setupRoleBiz() *RoleBiz {
	// 初始化 GORM 数据库连接
	utils.InitGorm()

	// 创建 RoleService 实例
	service := services.NewRoleService(utils.GormDB)

	// 创建 RoleBiz 实例
	return NewRoleBiz(service)
}

func TestRoleBiz_GetCount(t *testing.T) {
	biz := setupRoleBiz()

	// 调用 GetCount 方法
	count, err := biz.GetCount()
	if err != nil {
		t.Fatalf("调用 GetCount 失败: %v", err)
	}

	t.Logf("角色总数: %d", count)
}

func TestRoleBiz_List(t *testing.T) {
	biz := setupRoleBiz()

	// 调用 List 方法，分页参数：第 1 页，每页 10 条
	items, total, err := biz.List(1, 10)
	if err != nil {
		t.Fatalf("调用 List 失败: %v", err)
	}

	t.Logf("分页获取角色，共 %d 条记录，当前页返回 %d 条", total, len(items))
}
