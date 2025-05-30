// app/services/admin_service_test.go
 

package services

import (
	"testing"
    "os"
	"fmt"
 
	"gin-go-test/utils"
	"golang.org/x/crypto/bcrypt"
)

func TestAdminService_List(t *testing.T) {
	service := NewAdminService(utils.GormDB)
	admins, err := service.List(1, 10)
	if err != nil {
		t.Fatalf("获取管理员失败: %v", err)
	}
	if len(admins) == 0 {
		t.Errorf("返回管理员数量应大于 0，但得到 0")
	}
	t.Logf("获取成功，返回 %d 条记录", len(admins))
}

func TestAdminService_UpdatePassword(t *testing.T) {
	utils.InitDBX()

	adminID := os.Getenv("TEST_ADMIN_ID")
	newPlainPassword := os.Getenv("TEST_NEW_PASSWORD")

	if adminID == "" || newPlainPassword == "" {
		t.Fatal("❌ 必须设置 TEST_ADMIN_ID 和 TEST_NEW_PASSWORD 环境变量")
	}

	// 查询旧密码（加密的）
	var oldEncryptedPassword string
	err := utils.DBX.Get(&oldEncryptedPassword, "SELECT password FROM "+utils.PrefixTable("admin")+" WHERE id = ?", adminID)
	if err != nil {
		t.Fatalf("❌ 查询旧密码失败: %v", err)
	}

	// 生成新密码加密值
	newEncryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newPlainPassword), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("❌ 密码加密失败: %v", err)
	}

	// 打印信息
	fmt.Println("🧩 原始旧密码（明文）: 无法还原（bcrypt 不可逆）")
	fmt.Println("🔐 原始旧密码（加密）:", oldEncryptedPassword)
	fmt.Println("🆕 新密码测试输出（Hello World 123）:", newPlainPassword)
	fmt.Println("🔐 新密码（加密）:", string(newEncryptedPassword))

	// 更新数据库密码
	_, err = utils.DBX.Exec("UPDATE "+utils.PrefixTable("admin")+" SET password = ? WHERE id = ?", string(newEncryptedPassword), adminID)
	if err != nil {
		t.Fatalf("❌ 更新密码失败: %v", err)
	}

	fmt.Println("✅ 密码更新成功")
}

func TestGetJWTInfo(t *testing.T) {
	token := os.Getenv("TEST_JWT_TOKEN")
	if token == "" {
		t.Fatal("❌ 必须设置 TEST_JWT_TOKEN 环境变量")
	}

	claims, err := utils.GetJWTInfo(token)
	if err != nil {
		t.Fatalf("❌ 获取 JWT 信息失败: %v", err)
	}

	fmt.Println("✅ JWT 信息解析成功:")
	for k, v := range claims {
		fmt.Printf("  %s: %v\n", k, v)
	}
}