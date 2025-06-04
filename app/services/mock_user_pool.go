package services

import (
	"context"
	"encoding/json"
	"fmt"

	"gin-go-test/utils"
	"github.com/jmoiron/sqlx"
)

const redisUserPoolKey = "mock:user_pool"

func LoadUsersToRedis(db *sqlx.DB) error {
	ctx := context.Background()

	// 清空旧的数据
	utils.RedisClient.Del(ctx, redisUserPoolKey)

	// 查询数据库
	rows, err := db.Query("SELECT user_id, username FROM tm_user WHERE user_id IS NOT NULL AND username IS NOT NULL LIMIT 10000")
	if err != nil {
		return err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		var userID, username string
		if err := rows.Scan(&userID, &username); err != nil {
			continue
		}
		entry := map[string]string{
			"user_id":  userID,
			"username": username,
		}
		jsonData, err := json.Marshal(entry)
		if err != nil {
			continue
		}
		utils.RedisClient.SAdd(ctx, redisUserPoolKey, string(jsonData))
		count++
	}

	fmt.Printf("✅ 写入 Redis 模拟用户池 %d 条记录。\n", count)
	return nil
}
