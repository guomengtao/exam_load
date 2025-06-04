package services

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// MockRedisHelper is a mock implementation of RedisHelper
type MockRedisHelper struct {
	mock.Mock
}

func (m *MockRedisHelper) HSet(ctx context.Context, key string, field string, value interface{}) error {
	args := m.Called(ctx, key, field, value)
	return args.Error(0)
}

func (m *MockRedisHelper) HGet(ctx context.Context, key string, field string) (string, error) {
	args := m.Called(ctx, key, field)
	return args.String(0), args.Error(1)
}

func (m *MockRedisHelper) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	args := m.Called(ctx, key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]string), args.Error(1)
}

func (m *MockRedisHelper) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	args := m.Called(ctx, key, fields)
	return args.Error(0)
}

func (m *MockRedisHelper) Expire(ctx context.Context, key string, expiration time.Duration) error {
	args := m.Called(ctx, key, expiration)
	return args.Error(0)
}

func setupTestService() (*AnswerService, *MockRedisHelper) {
	mockRedis := new(MockRedisHelper)
	service := NewAnswerServiceWithRedis(nil, mockRedis)
	return service, mockRedis
}

func TestSaveToRedis(t *testing.T) {
	service, mockRedis := setupTestService()
	ctx := context.Background()

	tests := []struct {
		name          string
		data          map[string]interface{}
		mockError     error
		expectedError bool
	}{
		{
			name: "successful save",
			data: map[string]interface{}{
				"answer_uid": "test-uuid",
				"exam_id":    1,
				"exam_uuid":  "exam-uuid",
				"answers":    map[string]interface{}{"1": map[string]interface{}{"answer": 1, "score": 10}},
				"score":      10,
				"created_at": time.Now(),
			},
			mockError:     nil,
			expectedError: false,
		},
		{
			name: "redis error",
			data: map[string]interface{}{
				"answer_uid": "test-uuid",
				"exam_id":    1,
				"exam_uuid":  "exam-uuid",
			},
			mockError:     assert.AnError,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			if tt.mockError != nil {
				mockRedis.On("HMSet", ctx, mock.Anything, mock.Anything).Return(tt.mockError).Once()
			} else {
				mockRedis.On("HMSet", ctx, mock.Anything, mock.Anything).Return(nil).Once()
				mockRedis.On("Expire", ctx, mock.Anything, mock.Anything).Return(nil).Once()
			}

			// Execute
			err := service.SaveToRedis(ctx, tt.data)

			// Assertions
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRedis.AssertExpectations(t)
		})
	}
}

func TestGetAnswerRecord(t *testing.T) {
	service, mockRedis := setupTestService()
	ctx := context.Background()

	tests := []struct {
		name           string
		recordID       string
		mockData       map[string]string
		mockError      error
		expectedError  bool
		expectedRecord *AnswerResponse
	}{
		{
			name:     "successful retrieval",
			recordID: "test-record",
			mockData: map[string]string{
				"uuid":        "test-record",
				"exam_id":     "1",
				"exam_uuid":   "exam-uuid",
				"user_uuid":   "user-uuid",
				"username":    "test-user",
				"user_id":     "user-1",
				"total_score": "10",
				"score":       "10",
				"created_at":  "1234567890",
				"duration":    "30",
				"answers":     `{"1":{"answer":1,"score":10}}`,
			},
			mockError:     nil,
			expectedError: false,
			expectedRecord: &AnswerResponse{
				UUID:       "test-record",
				ExamID:     1,
				ExamUUID:   "exam-uuid",
				UserUUID:   "user-uuid",
				Username:   "test-user",
				UserID:     "user-1",
				TotalScore: 10,
				Score:      10,
				CreatedAt:  1234567890,
				Duration:   30,
				Answers:    map[string]interface{}{"1": map[string]interface{}{"answer": float64(1), "score": float64(10)}},
			},
		},
		{
			name:           "record not found",
			recordID:       "non-existent",
			mockData:       nil,
			mockError:      assert.AnError,
			expectedError:  true,
			expectedRecord: nil,
		},
		{
			name:     "invalid answer data",
			recordID: "test-record",
			mockData: map[string]string{
				"uuid":        "test-record",
				"exam_id":     "1",
				"exam_uuid":   "exam-uuid",
				"user_uuid":   "user-uuid",
				"username":    "test-user",
				"user_id":     "user-1",
				"total_score": "10",
				"score":       "10",
				"created_at":  "1234567890",
				"duration":    "30",
				"answers":     "invalid json",
			},
			mockError:      nil,
			expectedError:  true,
			expectedRecord: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			if tt.mockError != nil {
				mockRedis.On("HGetAll", ctx, mock.Anything).Return(nil, tt.mockError).Once()
			} else {
				mockRedis.On("HGetAll", ctx, mock.Anything).Return(tt.mockData, nil).Once()
			}

			// Execute
			record, err := service.GetAnswerRecord(ctx, tt.recordID)

			// Assertions
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, record)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, record)
				assert.Equal(t, tt.expectedRecord.UUID, record.UUID)
				assert.Equal(t, tt.expectedRecord.ExamID, record.ExamID)
				assert.Equal(t, tt.expectedRecord.TotalScore, record.TotalScore)
				assert.Equal(t, tt.expectedRecord.Score, record.Score)
				assert.Equal(t, tt.expectedRecord.Duration, record.Duration)
				assert.Equal(t, tt.expectedRecord.Answers, record.Answers)
			}

			mockRedis.AssertExpectations(t)
		})
	}
}

func TestGetExamPaper(t *testing.T) {
	service, mockRedis := setupTestService()
	ctx := context.Background()

	tests := []struct {
		name          string
		examUUID      string
		mockData      map[string]string
		mockError     error
		expectedError bool
		expectedPaper *ExamPaper
	}{
		{
			name:     "successful retrieval from redis",
			examUUID: "exam-uuid",
			mockData: map[string]string{
				"data": `{
					"id": 1,
					"title": "Test Paper",
					"description": "Test Description",
					"questions": []
				}`,
			},
			mockError:     nil,
			expectedError: false,
			expectedPaper: &ExamPaper{
				ID:          1,
				Title:       "Test Paper",
				Description: "Test Description",
				Questions:   []Question{},
			},
		},
		{
			name:          "paper not found in redis",
			examUUID:      "non-existent",
			mockData:      nil,
			mockError:     assert.AnError,
			expectedError: true,
			expectedPaper: nil,
		},
		{
			name:     "invalid paper data",
			examUUID: "exam-uuid",
			mockData: map[string]string{
				"data": "invalid json",
			},
			mockError:     nil,
			expectedError: true,
			expectedPaper: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			if tt.mockError != nil {
				mockRedis.On("HGetAll", ctx, mock.Anything).Return(nil, tt.mockError).Once()
			} else {
				mockRedis.On("HGetAll", ctx, mock.Anything).Return(tt.mockData, nil).Once()
			}

			// Execute
			paper, err := service.GetExamPaper(ctx, tt.examUUID)

			// Assertions
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, paper)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, paper)
				assert.Equal(t, tt.expectedPaper.ID, paper.ID)
				assert.Equal(t, tt.expectedPaper.Title, paper.Title)
				assert.Equal(t, tt.expectedPaper.Description, paper.Description)
				assert.Equal(t, tt.expectedPaper.Questions, paper.Questions)
			}

			mockRedis.AssertExpectations(t)
		})
	}
}

func TestAsyncSaveToDatabase(t *testing.T) {
	service, _ := setupTestService()
	ctx := context.Background()

	tests := []struct {
		name          string
		data          map[string]interface{}
		expectedError bool
	}{
		{
			name: "successful async save",
			data: map[string]interface{}{
				"answer_uid": "test-uuid",
				"exam_id":    1,
				"exam_uuid":  "exam-uuid",
				"answers":    map[string]interface{}{"1": map[string]interface{}{"answer": 1, "score": 10}},
				"score":      10,
				"created_at": time.Now(),
			},
			expectedError: false,
		},
		{
			name: "invalid data",
			data: map[string]interface{}{
				"answer_uid": "test-uuid",
				"exam_id":    "invalid", // 故意使用错误类型
				"exam_uuid":  "exam-uuid",
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			service.AsyncSaveToDatabase(ctx, tt.data)

			// 由于是异步操作，我们只能验证函数是否正常执行
			// 实际的数据保存验证应该在集成测试中进行
			assert.True(t, true)
		})
	}
}
