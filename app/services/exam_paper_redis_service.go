package services

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-go-test/utils"
)

// ExamPaperRedisService handles Redis-based exam paper data operations
type ExamPaperRedisService struct {
	redisHelper utils.RedisHelper
}

// NewExamPaperRedisService creates a new instance of ExamPaperRedisService
func NewExamPaperRedisService() *ExamPaperRedisService {
	return &ExamPaperRedisService{
		redisHelper: &utils.DefaultRedisHelper{},
	}
}

// GetExamPaperByUUID retrieves a single exam paper by UUID from Redis
func (s *ExamPaperRedisService) GetExamPaperByUUID(ctx context.Context, uuid string) (map[string]interface{}, error) {
	redisKey := fmt.Sprintf("exam_paper:%s", uuid)
	val, err := s.redisHelper.Get(ctx, redisKey)
	if err != nil {
		return nil, fmt.Errorf("试卷未找到: %v", err)
	}

	var paper map[string]interface{}
	if err := json.Unmarshal([]byte(val), &paper); err != nil {
		return nil, fmt.Errorf("解析失败: %v", err)
	}
	return paper, nil
}

// ListExamPapers retrieves a paginated list of exam papers from Redis
func (s *ExamPaperRedisService) ListExamPapers(ctx context.Context, page, limit int) (map[string]interface{}, error) {
	matchPattern := "exam_paper:*"
	keys, err := s.redisHelper.Scan(ctx, matchPattern)
	if err != nil {
		return nil, fmt.Errorf("Redis 扫描失败: %v", err)
	}

	start := (page - 1) * limit
	end := start + limit
	if start > len(keys) {
		start = len(keys)
	}
	if end > len(keys) {
		end = len(keys)
	}
	pagedKeys := keys[start:end]

	var papers []map[string]interface{}
	for _, key := range pagedKeys {
		val, err := s.redisHelper.Get(ctx, key)
		if err != nil {
			continue
		}
		var paper map[string]interface{}
		if err := json.Unmarshal([]byte(val), &paper); err != nil {
			continue
		}
		papers = append(papers, paper)
	}

	return map[string]interface{}{
		"list": papers,
		"pagination": map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": len(keys),
		},
	}, nil
} 