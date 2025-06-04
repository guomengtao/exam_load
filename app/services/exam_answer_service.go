package services

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-go-test/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// RedisHelper defines the interface for Redis operations
type RedisHelper interface {
	HSet(ctx context.Context, key string, field string, value interface{}) error
	HGet(ctx context.Context, key string, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HMSet(ctx context.Context, key string, fields map[string]interface{}) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
}

// DefaultRedisHelper implements RedisHelper using utils package
type DefaultRedisHelper struct{}

func (h *DefaultRedisHelper) HSet(ctx context.Context, key string, field string, value interface{}) error {
	return utils.RedisHSet(key, field, value)
}

func (h *DefaultRedisHelper) HGet(ctx context.Context, key string, field string) (string, error) {
	return utils.RedisHGet(key, field)
}

func (h *DefaultRedisHelper) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return utils.RedisHGetAll(key)
}

func (h *DefaultRedisHelper) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	return utils.RedisHMSet(key, fields)
}

func (h *DefaultRedisHelper) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return utils.RedisExpire(key, expiration)
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
	Title                string      `json:"title"`
	Content              string      `json:"content"`
	Options              interface{} `json:"options"`
	CorrectAnswer        interface{} `json:"correct_answer"`
	CorrectAnswerBitmask interface{} `json:"correct_answer_bitmask"`
	Score                int         `json:"score"`
	Analysis             string      `json:"analysis"`
}

// AnswerService handles the data operations for exam answers
type AnswerService struct {
	db          *gorm.DB
	redisHelper RedisHelper
}

// NewAnswerService creates a new instance of AnswerService
func NewAnswerService() *AnswerService {
	return &AnswerService{
		db:          utils.GormDB,
		redisHelper: &DefaultRedisHelper{},
	}
}

// NewAnswerServiceWithRedis creates a new instance of AnswerService with custom Redis helper
func NewAnswerServiceWithRedis(db *gorm.DB, redisHelper RedisHelper) *AnswerService {
	return &AnswerService{
		db:          db,
		redisHelper: redisHelper,
	}
}

// SaveToRedis saves the answer record to Redis
func (s *AnswerService) SaveToRedis(ctx context.Context, data map[string]interface{}) error {
	redisKey := fmt.Sprintf("exam_answer:%s", data["answer_uid"])

	// Convert answers to JSON string
	answers, err := json.Marshal(data["answers"])
	if err != nil {
		return fmt.Errorf("序列化答案失败: %v", err)
	}
	data["answers"] = string(answers)

	// Convert all values to strings for Redis
	redisData := make(map[string]interface{})
	for k, v := range data {
		switch val := v.(type) {
		case string:
			redisData[k] = val
		case int:
			redisData[k] = strconv.Itoa(val)
		case int64:
			redisData[k] = strconv.FormatInt(val, 10)
		default:
			redisData[k] = fmt.Sprintf("%v", val)
		}
	}

	// Save to Redis with expiration
	if err := s.redisHelper.HMSet(ctx, redisKey, redisData); err != nil {
		return fmt.Errorf("保存到Redis失败: %v", err)
	}

	// Set expiration time (e.g., 7 days)
	if err := s.redisHelper.Expire(ctx, redisKey, 7*24*time.Hour); err != nil {
		return fmt.Errorf("设置过期时间失败: %v", err)
	}

	return nil
}

// AsyncSaveToDatabase saves the answer record to database asynchronously
func (s *AnswerService) AsyncSaveToDatabase(ctx context.Context, data map[string]interface{}) {
	// Create a new context for the goroutine
	go func() {
		// Create a new context with timeout
		_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Start a transaction
		tx := s.db.Begin()
		if tx.Error != nil {
			fmt.Printf("Failed to start transaction: %v\n", tx.Error)
			return
		}

		// Create answer record
		answerRecord := map[string]interface{}{
			"uuid":        data["answer_uid"],
			"exam_id":     data["exam_id"],
			"exam_uuid":   data["exam_uuid"],
			"user_uuid":   data["user_uuid"],
			"answers":     data["answers"],
			"total_score": data["total_score"],
			"created_at":  data["created_at"],
			"username":    data["username"],
			"user_id":     data["user_id"],
			"duration":    data["duration"],
			"score":       data["score"],
		}

		// Save to database
		if err := tx.Table("exam_answers").Create(answerRecord).Error; err != nil {
			tx.Rollback()
			fmt.Printf("Failed to save answer record: %v\n", err)
			return
		}

		// Commit transaction
		if err := tx.Commit().Error; err != nil {
			fmt.Printf("Failed to commit transaction: %v\n", err)
			return
		}

		fmt.Printf("Successfully saved answer record to database: %s\n", data["answer_uid"])
	}()
}

// GetAnswerRecord retrieves an answer record from Redis
func (s *AnswerService) GetAnswerRecord(ctx context.Context, recordID string) (*AnswerResponse, error) {
	redisKey := fmt.Sprintf("exam_answer:%s", recordID)
	result, err := s.redisHelper.HGetAll(ctx, redisKey)
	if err != nil {
		return nil, fmt.Errorf("获取答题记录失败: %v", err)
	}
	if result == nil || len(result) == 0 {
		return nil, fmt.Errorf("答题记录不存在或为空")
	}

	requiredFields := []string{"exam_id", "total_score", "created_at", "duration", "score", "exam_uuid", "user_uuid", "username", "user_id", "answers"}
	for _, field := range requiredFields {
		val, exists := result[field]
		if !exists || val == "" {
			return nil, fmt.Errorf("答题记录数据不完整: 缺少字段 %s", field)
		}
	}

	examID, err := strconv.ParseInt(result["exam_id"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("解析考试ID失败: %v", err)
	}
	totalScore, err := strconv.Atoi(result["total_score"])
	if err != nil {
		return nil, fmt.Errorf("解析总分失败: %v", err)
	}
	createdAt, err := strconv.ParseInt(result["created_at"], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("解析创建时间失败: %v", err)
	}
	duration, err := strconv.Atoi(result["duration"])
	if err != nil {
		return nil, fmt.Errorf("解析答题时长失败: %v", err)
	}
	score, err := strconv.Atoi(result["score"])
	if err != nil {
		return nil, fmt.Errorf("解析得分失败: %v", err)
	}

	var answers interface{}
	if err := json.Unmarshal([]byte(result["answers"]), &answers); err != nil {
		return nil, fmt.Errorf("解析答案数据失败: %v", err)
	}

	return &AnswerResponse{
		UUID:       recordID,
		ExamID:     examID,
		ExamUUID:   result["exam_uuid"],
		UserUUID:   result["user_uuid"],
		Answers:    answers,
		TotalScore: totalScore,
		CreatedAt:  createdAt,
		Username:   result["username"],
		UserID:     result["user_id"],
		Duration:   duration,
		Score:      score,
	}, nil
}

// GetExamPaper retrieves the exam paper information
func (s *AnswerService) GetExamPaper(ctx context.Context, examUUID string) (*ExamPaper, error) {
	redisKey := fmt.Sprintf("exam_paper:%s", examUUID)
	result, err := s.redisHelper.HGetAll(ctx, redisKey)
	if err != nil {
		return nil, fmt.Errorf("获取试卷信息失败: %v", err)
	}
	if result == nil || len(result) == 0 {
		return nil, fmt.Errorf("试卷信息不存在或为空")
	}

	data, exists := result["data"]
	if !exists || data == "" {
		return nil, fmt.Errorf("试卷数据不完整: 缺少字段 data")
	}

	paper := &ExamPaper{}
	if err := json.Unmarshal([]byte(data), paper); err != nil {
		return nil, fmt.Errorf("解析试卷数据失败: %v", err)
	}
	if paper.ID == 0 || paper.Title == "" {
		return nil, fmt.Errorf("试卷数据不完整: 缺少 ID 或标题")
	}

	return paper, nil
}

// AnswerServiceInterface 便于 mock 的接口
//go:generate mockery --name=AnswerServiceInterface --output=../mocks --case=underscore
// 你可以用 mockery 工具自动生成 mock

type AnswerServiceInterface interface {
	SaveToRedis(ctx context.Context, data map[string]interface{}) error
	AsyncSaveToDatabase(ctx context.Context, data map[string]interface{})
	GetAnswerRecord(ctx context.Context, recordID string) (*AnswerResponse, error)
	GetExamPaper(ctx context.Context, examUUID string) (*ExamPaper, error)
}

var _ AnswerServiceInterface = (*AnswerService)(nil)
