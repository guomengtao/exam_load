package biz

import (
	"context"
	"gin-go-test/app/services"
)

// AnswerResponse represents the response structure for answer records
type AnswerResponse struct {
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

// ExamPaper represents the exam paper structure
type ExamPaper struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
}

// Question represents a question in the exam paper
type Question struct {
	ID                   int64       `json:"id"`
	Type                 string      `json:"type"`
	Content              string      `json:"content"`
	Options              interface{} `json:"options"`
	CorrectAnswer        interface{} `json:"correct_answer"`
	CorrectAnswerBitmask int         `json:"correct_answer_bitmask"`
	Score                int         `json:"score"`
}

// ExamAnswerBiz 处理考试答题相关的业务逻辑
type ExamAnswerBiz struct {
	service services.AnswerServiceInterface
}

// NewExamAnswerBiz 创建新的 ExamAnswerBiz 实例
func NewExamAnswerBiz(service services.AnswerServiceInterface) *ExamAnswerBiz {
	return &ExamAnswerBiz{
		service: service,
	}
}

// SaveAnswer 保存答题记录
func (b *ExamAnswerBiz) SaveAnswer(ctx context.Context, data map[string]interface{}) error {
	// 1. 保存到 Redis
	if err := b.service.SaveToRedis(ctx, data); err != nil {
		return err
	}

	// 2. 异步保存到数据库
	b.service.AsyncSaveToDatabase(ctx, data)

	return nil
}

// GetAnswerRecord 获取答题记录
func (b *ExamAnswerBiz) GetAnswerRecord(ctx context.Context, recordID string) (*AnswerResponse, error) {
	svcRecord, err := b.service.GetAnswerRecord(ctx, recordID)
	if err != nil {
		return nil, err
	}
	if svcRecord == nil {
		return nil, nil
	}
	return &AnswerResponse{
		UUID:       svcRecord.UUID,
		ExamID:     svcRecord.ExamID,
		ExamUUID:   svcRecord.ExamUUID,
		UserUUID:   svcRecord.UserUUID,
		Answers:    svcRecord.Answers,
		TotalScore: svcRecord.TotalScore,
		CreatedAt:  svcRecord.CreatedAt,
		Username:   svcRecord.Username,
		UserID:     svcRecord.UserID,
		Duration:   svcRecord.Duration,
		Score:      svcRecord.Score,
	}, nil
}

// GetExamPaper 获取试卷信息
func (b *ExamAnswerBiz) GetExamPaper(ctx context.Context, examUUID string) (*ExamPaper, error) {
	svcPaper, err := b.service.GetExamPaper(ctx, examUUID)
	if err != nil {
		return nil, err
	}
	if svcPaper == nil {
		return nil, nil
	}
	questions := make([]Question, len(svcPaper.Questions))
	for i, q := range svcPaper.Questions {
		var correctAnswerBitmask int
		switch v := q.CorrectAnswerBitmask.(type) {
		case int:
			correctAnswerBitmask = v
		case int64:
			correctAnswerBitmask = int(v)
		case float64:
			correctAnswerBitmask = int(v)
		default:
			correctAnswerBitmask = 0
		}

		questions[i] = Question{
			ID:                   q.ID,
			Type:                 q.Type,
			Content:              q.Content,
			Options:              q.Options,
			CorrectAnswer:        q.CorrectAnswer,
			CorrectAnswerBitmask: correctAnswerBitmask,
			Score:                q.Score,
		}
	}
	return &ExamPaper{
		ID:          svcPaper.ID,
		Title:       svcPaper.Title,
		Description: svcPaper.Description,
		Questions:   questions,
	}, nil
}

// ExamAnswerBizInterface 便于 mock 的接口
type ExamAnswerBizInterface interface {
	SaveAnswer(ctx context.Context, data map[string]interface{}) error
	GetAnswerRecord(ctx context.Context, recordID string) (*AnswerResponse, error)
	GetExamPaper(ctx context.Context, examUUID string) (*ExamPaper, error)
}

var _ ExamAnswerBizInterface = (*ExamAnswerBiz)(nil)
