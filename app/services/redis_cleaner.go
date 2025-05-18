

package services

import (
    "context"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
    "gin-go-test/utils"
)

// CleanOldProcessedData 清理 Redis 中 exam:processed 超过 7 天的成员
func CleanOldProcessedData() error {
    ctx := context.Background()
    client := utils.RedisClient
    zsetKey := "exam:processed"

    // 7 天前的时间戳
    sevenDaysAgo := float64(time.Now().Add(-7 * 24 * time.Hour).Unix())

    // 获取旧数据
    members, err := client.ZRangeByScore(ctx, zsetKey, &redis.ZRangeBy{
        Min: "0",
        Max: fmt.Sprintf("%f", sevenDaysAgo),
    }).Result()
    if err != nil {
        return fmt.Errorf("获取旧数据失败: %v", err)
    }

    if len(members) == 0 {
        fmt.Println("🧹 没有需要清理的旧数据")
        return nil
    }

    // 删除这些旧成员
    deleted, err := client.ZRem(ctx, zsetKey, members).Result()
    if err != nil {
        return fmt.Errorf("清理旧数据失败: %v", err)
    }

    fmt.Printf("🧹 清理成功: %d 条过期 processed 数据\n", deleted)
    return nil
}