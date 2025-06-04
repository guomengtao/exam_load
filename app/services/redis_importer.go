package services

import (
	"context"
	"encoding/json"
	"fmt"
)

// RunRedisImportOnce 同步执行一次Redis导入逻辑
func RunRedisImportOnce() {
	panic("not implemented") // 临时屏蔽未定义内容，便于单元测试
}

func (i *RedisImporter) ImportBatch(ctx context.Context, batchSize int) error {
	panic("not implemented") // 临时屏蔽未定义内容，便于单元测试
}

// RedisImporter 处理从 Redis 导入数据的逻辑
type RedisImporter struct {
	redisHelper RedisHelper
}

// NewRedisImporter 创建新的 RedisImporter 实例
func NewRedisImporter(redisHelper RedisHelper) *RedisImporter {
	return &RedisImporter{
		redisHelper: redisHelper,
	}
}

// Answer 表示答题记录
type Answer struct {
	UUID       string      `json:"uuid"`
	ExamID     int64       `json:"exam_id"`
	ExamUUID   string      `json:"exam_uuid"`
	UserUUID   string      `json:"user_uuid"`
	Answers    interface{} `json:"answers"`
	TotalScore int         `json:"total_score"`
	CreatedAt  int64       `json:"created_at"`
	Username   string      `json:"username"`
	UserID     string      `json:"user_id"`
	Duration   int         `json:"duration"`
	Score      int         `json:"score"`
}

// ImportAnswer 从 Redis 导入答题记录
func (i *RedisImporter) ImportAnswer(ctx context.Context, recordID string) (*Answer, error) {
	redisKey := fmt.Sprintf("exam_answer:%s", recordID)
	result, err := i.redisHelper.HGetAll(ctx, redisKey)
	if err != nil {
		return nil, fmt.Errorf("获取答题记录失败: %v", err)
	}
	if result == nil || len(result) == 0 {
		return nil, fmt.Errorf("答题记录不存在或为空")
	}

	answer := &Answer{
		UUID:     recordID,
		ExamUUID: result["exam_uuid"],
		UserUUID: result["user_uuid"],
		Username: result["username"],
		UserID:   result["user_id"],
	}

	// 解析其他字段
	if err := json.Unmarshal([]byte(result["answers"]), &answer.Answers); err != nil {
		return nil, fmt.Errorf("解析答案数据失败: %v", err)
	}

	return answer, nil
}
