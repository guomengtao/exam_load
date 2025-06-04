package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-go-test/app/services"
	"github.com/google/uuid"
	"sort"
	"strconv"
	"time"
)

// AnswerRequest represents the request structure for submitting answers
type AnswerRequest struct {
	UUID      string                     `json:"uuid" binding:"required"`
	ExamID    int64                      `json:"exam_id" binding:"required"`
	ExamUUID  string                     `json:"exam_uuid"`
	Answers   map[string]json.RawMessage `json:"answers" binding:"required"`
	Username  string                     `json:"username"`
	UserID    string                     `json:"user_id"`
	Duration  int                        `json:"duration"`
	FullScore int                        `json:"full_score"`
}

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

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// ExamPaper represents the exam paper structure
type ExamPaper struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
}

// Question represents a question in the exam paper
type Question struct {
	ID                   int         `json:"id"`
	Title                string      `json:"title"`
	Options              []string    `json:"options"`
	Type                 string      `json:"type"` // single/multi/judge
	Score                int         `json:"score"`
	CorrectAnswerBitmask int         `json:"correct_answer_bitmask"`
	CorrectAnswer        interface{} `json:"correct_answer"`
	Analysis             string      `json:"analysis"`
}

// FullAnswerResponse represents the complete answer response with questions
type FullAnswerResponse struct {
	RecordID    string               `json:"record_id"`
	ExamID      interface{}          `json:"exam_id"`
	UserUUID    string               `json:"user_uuid"`
	TotalScore  int                  `json:"total_score"`
	CreatedAt   int64                `json:"created_at"`
	Questions   []QuestionWithAnswer `json:"questions"`
	Username    string               `json:"username"`
	UserID      string               `json:"user_id"`
	Duration    int                  `json:"duration"`
	Score       int                  `json:"score"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
}

// QuestionWithAnswer represents a question with user's answer
type QuestionWithAnswer struct {
	ID            int         `json:"id"`
	Title         string      `json:"title"`
	Options       []string    `json:"options"`
	Type          string      `json:"type"`
	Score         int         `json:"score"`
	CorrectAnswer interface{} `json:"correct_answer"`
	UserAnswer    interface{} `json:"user_answer"`
	IsCorrect     bool        `json:"is_correct"`
	Analysis      string      `json:"analysis"`
}

// AnswerBiz handles the business logic for exam answers
type AnswerBiz struct {
	answerService *services.AnswerService
}

// NewAnswerBiz creates a new instance of AnswerBiz
func NewAnswerBiz(answerService *services.AnswerService) *AnswerBiz {
	return &AnswerBiz{
		answerService: answerService,
	}
}

// SubmitAnswer handles the submission of exam answers
func (b *AnswerBiz) SubmitAnswer(ctx context.Context, req *AnswerRequest) (*AnswerResponse, error) {
	if len(req.Answers) == 0 {
		return nil, fmt.Errorf("答案不能为空")
	}

	// Calculate total score
	totalScore := 0
	for _, answer := range req.Answers {
		var detail struct{ Score int }
		if err := json.Unmarshal(answer, &detail); err == nil {
			totalScore += detail.Score
		}
	}

	// Generate record ID and timestamp
	recordID := uuid.New().String()
	createdAt := time.Now().Unix()

	// Prepare record data
	record := map[string]interface{}{
		"answer_uid":  recordID,
		"exam_id":     req.ExamID,
		"exam_uuid":   req.ExamUUID,
		"user_uuid":   req.UUID,
		"answers":     req.Answers,
		"total_score": totalScore,
		"created_at":  createdAt,
		"username":    req.Username,
		"user_id":     req.UserID,
		"duration":    req.Duration,
		"score":       totalScore,
	}

	// Save to Redis
	if err := b.answerService.SaveToRedis(ctx, record); err != nil {
		return nil, fmt.Errorf("保存到Redis失败: %v", err)
	}

	// Async save to database
	go b.answerService.AsyncSaveToDatabase(ctx, record)

	// Prepare response
	response := &AnswerResponse{
		UUID:       recordID,
		ExamID:     req.ExamID,
		ExamUUID:   req.ExamUUID,
		UserUUID:   req.UUID,
		Answers:    req.Answers,
		TotalScore: totalScore,
		CreatedAt:  createdAt,
		Username:   req.Username,
		UserID:     req.UserID,
		Duration:   req.Duration,
		Score:      totalScore,
	}

	return response, nil
}

// GetAnswerResult retrieves the exam answer result
func (b *AnswerBiz) GetAnswerResult(ctx context.Context, recordID string) (*AnswerResponse, error) {
	record, err := b.answerService.GetAnswerRecord(ctx, recordID)
	if err != nil {
		return nil, err
	}
	// 将 services.AnswerResponse 转换为 AnswerResponse
	return &AnswerResponse{
		UUID:       record.UUID,
		ExamID:     record.ExamID,
		ExamUUID:   record.ExamUUID,
		UserUUID:   record.UserUUID,
		Answers:    record.Answers,
		TotalScore: record.TotalScore,
		CreatedAt:  record.CreatedAt,
		Username:   record.Username,
		UserID:     record.UserID,
		Duration:   record.Duration,
		Score:      record.Score,
	}, nil
}

// GetFullAnswerResult retrieves the complete exam answer result with questions
func (b *AnswerBiz) GetFullAnswerResult(ctx context.Context, recordID string) (*FullAnswerResponse, error) {
	// Get answer record
	record, err := b.answerService.GetAnswerRecord(ctx, recordID)
	if err != nil {
		return nil, err
	}

	// Get exam paper
	paper, err := b.answerService.GetExamPaper(ctx, record.ExamUUID)
	if err != nil {
		return nil, err
	}

	// 将 services.AnswerResponse 和 services.ExamPaper 转换为 AnswerResponse 和 ExamPaper
	answerResponse := &AnswerResponse{
		UUID:       record.UUID,
		ExamID:     record.ExamID,
		ExamUUID:   record.ExamUUID,
		UserUUID:   record.UserUUID,
		Answers:    record.Answers,
		TotalScore: record.TotalScore,
		CreatedAt:  record.CreatedAt,
		Username:   record.Username,
		UserID:     record.UserID,
		Duration:   record.Duration,
		Score:      record.Score,
	}

	// 将 paper.ID 转换为 int，并将 paper.Questions 转换为 []Question
	examPaper := &ExamPaper{
		ID:          int(paper.ID),
		Title:       paper.Title,
		Description: paper.Description,
		Questions:   convertQuestions(paper.Questions),
	}

	// Build full response
	return b.buildFullResponse(answerResponse, examPaper)
}

// convertQuestions 将 []services.Question 转换为 []Question
func convertQuestions(questions []services.Question) []Question {
	result := make([]Question, len(questions))
	for i, q := range questions {
		// 将 q.Options 类型断言为 []string
		options, ok := q.Options.([]string)
		if !ok {
			options = []string{} // 如果断言失败，使用空切片
		}

		// 将 CorrectAnswerBitmask 类型断言为 int
		bitmask, ok := q.CorrectAnswerBitmask.(int)
		if !ok {
			bitmask = 0 // 如果断言失败，使用默认值 0
		}

		result[i] = Question{
			ID:                   int(q.ID),
			Title:                q.Title,
			Options:              options,
			Type:                 q.Type,
			Score:                q.Score,
			CorrectAnswerBitmask: bitmask,
			CorrectAnswer:        q.CorrectAnswer,
			Analysis:             q.Analysis,
		}
	}
	return result
}

// buildFullResponse builds the complete response with questions and answers
func (b *AnswerBiz) buildFullResponse(record *AnswerResponse, paper *ExamPaper) (*FullAnswerResponse, error) {
	// Parse user answers
	userAnswers, ok := record.Answers.(map[string]json.RawMessage)
	if !ok {
		return nil, fmt.Errorf("无效的答案格式")
	}

	// Build questions with answers
	questions := make([]QuestionWithAnswer, 0, len(paper.Questions))
	for _, q := range paper.Questions {
		// Get user's answer for this question
		userAnswer, exists := userAnswers[strconv.Itoa(q.ID)]
		if !exists {
			continue
		}

		// Parse user's answer
		var answerDetail struct {
			Answer interface{} `json:"answer"`
			Score  int         `json:"score"`
		}
		if err := json.Unmarshal(userAnswer, &answerDetail); err != nil {
			return nil, fmt.Errorf("解析用户答案失败: %v", err)
		}

		// Check if answer is correct
		isCorrect := b.isAnswerCorrect(q, answerDetail.Answer)

		// Add to questions list
		questions = append(questions, QuestionWithAnswer{
			ID:            q.ID,
			Title:         q.Title,
			Options:       q.Options,
			Type:          q.Type,
			Score:         q.Score,
			CorrectAnswer: q.CorrectAnswer,
			UserAnswer:    answerDetail.Answer,
			IsCorrect:     isCorrect,
			Analysis:      q.Analysis,
		})
	}

	// Build full response
	response := &FullAnswerResponse{
		RecordID:    record.UUID,
		ExamID:      record.ExamID,
		UserUUID:    record.UserUUID,
		TotalScore:  record.TotalScore,
		CreatedAt:   record.CreatedAt,
		Questions:   questions,
		Username:    record.Username,
		UserID:      record.UserID,
		Duration:    record.Duration,
		Score:       record.Score,
		Title:       paper.Title,
		Description: paper.Description,
	}

	return response, nil
}

// isAnswerCorrect checks if the user's answer is correct
func (b *AnswerBiz) isAnswerCorrect(q Question, userAnswer interface{}) bool {
	// Normalize answers to []int for comparison
	correctAnswers := b.normalizeAnswer(q.CorrectAnswer)
	userAnswers := b.normalizeAnswer(userAnswer)

	// For single choice questions
	if q.Type == "single" {
		if len(correctAnswers) != 1 || len(userAnswers) != 1 {
			return false
		}
		return correctAnswers[0] == userAnswers[0]
	}

	// For multiple choice questions
	if q.Type == "multi" {
		if len(correctAnswers) != len(userAnswers) {
			return false
		}
		// Sort both slices for comparison
		sort.Ints(correctAnswers)
		sort.Ints(userAnswers)
		for i := range correctAnswers {
			if correctAnswers[i] != userAnswers[i] {
				return false
			}
		}
		return true
	}

	// For judgment questions
	if q.Type == "judge" {
		if len(correctAnswers) != 1 || len(userAnswers) != 1 {
			return false
		}
		return correctAnswers[0] == userAnswers[0]
	}

	return false
}

// normalizeAnswer converts different answer formats to []int
func (b *AnswerBiz) normalizeAnswer(answer interface{}) []int {
	switch v := answer.(type) {
	case []int:
		return v
	case []interface{}:
		result := make([]int, 0, len(v))
		for _, item := range v {
			if num, ok := item.(float64); ok {
				result = append(result, int(num))
			}
		}
		return result
	case float64:
		return []int{int(v)}
	case int:
		return []int{v}
	default:
		return nil
	}
}
