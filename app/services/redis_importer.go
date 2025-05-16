package services

import (
    "context"
    "fmt"
    "strconv"
    "time"

    "github.com/redis/go-redis/v9"
    "gin-go-test/utils"
    "gin-go-test/app/models"
)

// RunRedisImportOnce 同步执行一次Redis导入逻辑
func RunRedisImportOnce() {
    ctx := context.Background()
    batchSize := int64(100)
    start := int64(0)
    totalProcessed := 0

    for {
        // 打印当前待处理数量（submitted 和 processed）
        totalPending, _ := utils.RedisClient.ZCard(ctx, "exam:submitted").Result()
        totalProcessedSet, _ := utils.RedisClient.ZCard(ctx, "exam:processed").Result()
        fmt.Printf("⏳ exam:submitted 当前待处理数量: %d, exam:processed 已处理数量: %d\n", totalPending, totalProcessedSet)

        ids, err := utils.RedisClient.ZRange(ctx, "exam:submitted", start, start+batchSize-1).Result()
        if err != nil {
            fmt.Println("❌ 获取待处理 ID 失败:", err)
            break
        }
        if len(ids) == 0 {
            break
        }

        var batch []models.ExamAnswer
        for _, id := range ids {
            data, err := utils.RedisClient.HGetAll(ctx, "exam_answer:"+id).Result()
            if err != nil || len(data) == 0 {
                continue
            }

            createdAtInt, err := strconv.ParseInt(data["created_at"], 10, 64)
            if err != nil {
                continue
            }
            createdAt := time.Unix(createdAtInt, 0)

            duration, _ := strconv.Atoi(data["duration"])
            score, _ := strconv.Atoi(data["score"])
            totalScore, _ := strconv.Atoi(data["total_score"])
            examID, _ := strconv.Atoi(data["exam_id"])

            answer := models.ExamAnswer{
                AnswerUID:  id,
                ExamUUID:   data["exam_uuid"],
                UserID:     data["user_id"],
                Username:   data["username"],
                CreatedAt:  createdAt,
                Duration:   duration,
                Score:      score,
                TotalScore: totalScore,
                Answers:    data["answers"],
                ExamID:     examID,
            }
            batch = append(batch, answer)
        }

        if len(batch) > 0 {
            result := utils.GormDB.CreateInBatches(&batch, 100)
            if result.Error != nil {
                fmt.Println("❌ 入库失败:", result.Error)
                break
            }

            successCount := 0
            for _, a := range batch {
                remCount, err := utils.RedisClient.ZRem(ctx, "exam:submitted", a.AnswerUID).Result()
                if err != nil || remCount == 0 {
                    continue
                }

                err = utils.RedisClient.ZAdd(ctx, "exam:processed", redis.Z{
                    Score:  float64(time.Now().Unix()),
                    Member: a.AnswerUID,
                }).Err()
                if err != nil {
                    continue
                }

                _ = utils.RedisClient.Expire(ctx, "exam_answer:"+a.AnswerUID, 7*24*time.Hour).Err()
                successCount++
            }

            newPending, _ := utils.RedisClient.ZCard(ctx, "exam:submitted").Result()
            newProcessed, _ := utils.RedisClient.ZCard(ctx, "exam:processed").Result()
            totalProcessed += successCount
            fmt.Printf("🎉 本批次处理成功 %d 条，累计处理 %d 条\n", successCount, totalProcessed)
            fmt.Printf("📊 处理后 exam:submitted 数量: %d, exam:processed 数量: %d\n", newPending, newProcessed)
        }

        if int64(len(ids)) < batchSize {
            break
        }
        start += batchSize
    }
}