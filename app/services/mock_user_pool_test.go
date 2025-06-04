package services

import (
	"gin-go-test/utils"
	"testing"
)

func TestLoadUsersToRedis(t *testing.T) {
	// 初始化数据库连接
	utils.InitDBX()
	utils.InitRedis() // 添加 Redis 初始化
	err := LoadUsersToRedis(utils.DBX)
	if err != nil {
		t.Fatalf("❌ LoadUsersToRedis 执行失败: %v", err)
	}

	// 验证 Redis 中是否写入成功
	ctx := utils.Ctx
	count, err := utils.RedisClient.SCard(ctx, "mock:user_pool").Result()
	if err != nil {
		t.Fatalf("❌ Redis 获取 mock:user_pool 失败: %v", err)
	}

	if count == 0 {
		t.Fatal("❌ Redis 中 mock:user_pool 没有写入任何数据")
	}

	t.Logf("✅ Redis mock:user_pool 写入成功，共 %d 条数据", count)
}
