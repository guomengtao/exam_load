package services

import (
    "context"
    "fmt"
    "testing"
    "time"

    "gin-go-test/utils"
)

func TestRunRedisImportOnce(t *testing.T) {
    t.Setenv("TEST_RUN_UNIQUE", fmt.Sprintf("%d", time.Now().UnixNano()))
    utils.InitRedis()
    utils.InitGorm()

    // Run synchronous import on existing Redis data
    RunRedisImportOnce()

    submittedCount, _ := utils.RedisClient.ZCard(context.Background(), "exam:submitted").Result()
    processedCount, _ := utils.RedisClient.ZCard(context.Background(), "exam:processed").Result()
    hashKeys, _ := utils.RedisClient.Keys(context.Background(), "exam_answer:*").Result()

    t.Logf("📊 Redis 状态统计：")
    t.Logf("  exam:submitted 剩余：%d 条", submittedCount)
    t.Logf("  exam:processed 已处理：%d 条", processedCount)
    t.Logf("  exam_answer:* HashSet 总数：%d 条", len(hashKeys))
}