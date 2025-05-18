package services

import (
    "testing"
	"gin-go-test/utils"

)

func TestCleanOldProcessedData(t *testing.T) {
	utils.InitRedis() // ✅ 初始化 Redis 客户端

    err := CleanOldProcessedData()
    if err != nil {
        t.Errorf("清理旧数据失败: %v", err)
    }
}