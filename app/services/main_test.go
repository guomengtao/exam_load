package services

import (
	"os"
	"testing"
	"gin-go-test/utils"
)

func TestMain(m *testing.M) {
	// 初始化数据库（config 包会自动加载 .env）
	utils.InitDBX()

	utils.InitGorm() // 初始化 GORM
	 

	// 执行测试
	code := m.Run()
	os.Exit(code)
}