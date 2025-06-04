package biz

import (
	"context"
	"strconv"
	"gin-go-test/app/services"
)

// ExamPaperRedisBiz handles the business logic for Redis-based exam paper operations
type ExamPaperRedisBiz struct {
	paperService *services.ExamPaperRedisService
}

// NewExamPaperRedisBiz creates a new instance of ExamPaperRedisBiz
func NewExamPaperRedisBiz() *ExamPaperRedisBiz {
	return &ExamPaperRedisBiz{
		paperService: services.NewExamPaperRedisService(),
	}
}

// ListExamPapersFromRedis 从 Redis 中读取试卷列表或单条（通过 UUID）
func (b *ExamPaperRedisBiz) ListExamPapersFromRedis(ctx context.Context, pageParam, limitParam, uuidParam string) (interface{}, error) {
	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 5
	}

	if uuidParam != "" {
		return b.paperService.GetExamPaperByUUID(ctx, uuidParam)
	}
	return b.paperService.ListExamPapers(ctx, page, limit)
} 