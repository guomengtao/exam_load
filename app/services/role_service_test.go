package services

import (
	"os"
	"strconv"
	"testing"

 )

 

func TestGetRolesPaginated(t *testing.T) {
	// 从环境变量读取 page 和 pageSize，默认 page=1, pageSize=10
	page := getEnvAsInt("TEST_PAGE", 1)
	pageSize := getEnvAsInt("TEST_PAGE_SIZE", 10)

	result, err := GetRolesPaginated(page, pageSize)
	if err != nil {
		t.Fatalf("获取角色分页失败: %v", err)
	}

	t.Logf("✅ 获取到 %d 个角色（共 %d 条） - 当前页: %d", len(result.List), result.Total, page)
}

// 工具函数：读取环境变量并转换为 int
func getEnvAsInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}