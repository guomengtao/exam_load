package controllers

import (
	"os"
	"testing"
	"gin-go-test/utils"
)

func TestMain(m *testing.M) {
	// 不需要再手动加载 .env，config 包已经自动加载了

	// 初始化数据库连接
	utils.InitDBX()

	// 执行测试
	code := m.Run()
	os.Exit(code)
}