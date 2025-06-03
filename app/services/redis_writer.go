package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// RedisWriterInterval 表示写入间隔（单位：毫秒）
var RedisWriterInterval = 1000

// RedisWriter 处理写入 Redis 的逻辑
type RedisWriter struct {
	redisHelper RedisHelper
}

// NewRedisWriter 创建新的 RedisWriter 实例
func NewRedisWriter(redisHelper RedisHelper) *RedisWriter {
	return &RedisWriter{
		redisHelper: redisHelper,
	}
}

// StartRedisWriter 启动一个 goroutine 持续写入模拟数据到 Redis
func StartRedisWriter() {
	panic("not implemented") // 临时屏蔽未定义内容，便于单元测试
}

func writeOneRecord() {
	panic("not implemented") // 临时屏蔽未定义内容，便于单元测试
}

func SimulateBurst(n int) {
	panic("not implemented") // 临时屏蔽未定义内容，便于单元测试
}

func (w *RedisWriter) WriteMockData(ctx context.Context) error {
	panic("not implemented") // 临时屏蔽未定义内容，便于单元测试
}

// WriteAnswer 将答题记录写入 Redis
func (w *RedisWriter) WriteAnswer(ctx context.Context, answer *Answer) error {
	redisKey := fmt.Sprintf("exam_answer:%s", answer.UUID)
	
	// 将 answers 转换为 JSON 字符串
	answers, err := json.Marshal(answer.Answers)
	if err != nil {
		return fmt.Errorf("序列化答案失败: %v", err)
	}

	// 准备 Redis 数据
	redisData := map[string]interface{}{
		"uuid":        answer.UUID,
		"exam_id":     answer.ExamID,
		"exam_uuid":   answer.ExamUUID,
		"user_uuid":   answer.UserUUID,
		"username":    answer.Username,
		"user_id":     answer.UserID,
		"total_score": answer.TotalScore,
		"score":       answer.Score,
		"created_at":  answer.CreatedAt,
		"duration":    answer.Duration,
		"answers":     string(answers),
	}

	// 保存到 Redis
	if err := w.redisHelper.HMSet(ctx, redisKey, redisData); err != nil {
		return fmt.Errorf("保存到Redis失败: %v", err)
	}

	// 设置过期时间（7天）
	if err := w.redisHelper.Expire(ctx, redisKey, 7*24*time.Hour); err != nil {
		return fmt.Errorf("设置过期时间失败: %v", err)
	}

	return nil
}