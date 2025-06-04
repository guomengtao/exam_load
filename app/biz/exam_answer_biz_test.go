package biz

import (
	"context"
	"gin-go-test/app/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// MockAnswerService is a mock implementation of AnswerService
type MockAnswerService struct {
	mock.Mock
}

func (m *MockAnswerService) SaveToRedis(ctx context.Context, data map[string]interface{}) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *MockAnswerService) AsyncSaveToDatabase(ctx context.Context, data map[string]interface{}) {
	m.Called(ctx, data)
}

func (m *MockAnswerService) GetAnswerRecord(ctx context.Context, recordID string) (*services.AnswerResponse, error) {
	args := m.Called(ctx, recordID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*services.AnswerResponse), args.Error(1)
}

func (m *MockAnswerService) GetExamPaper(ctx context.Context, examUUID string) (*services.ExamPaper, error) {
	args := m.Called(ctx, examUUID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*services.ExamPaper), args.Error(1)
}

func setupTestBiz() (*ExamAnswerBiz, *MockAnswerService) {
	mockService := new(MockAnswerService)
	biz := NewExamAnswerBiz(mockService)
	return biz, mockService
}

func TestSaveAnswer(t *testing.T) {
	biz, mockService := setupTestBiz()
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
			mockService.On("SaveToRedis", ctx, tt.data).Return(tt.mockError).Once()
			if tt.mockError == nil {
				mockService.On("AsyncSaveToDatabase", ctx, tt.data).Return().Once()
			}

			// Execute
			err := biz.SaveAnswer(ctx, tt.data)

			// Assertions
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestGetAnswerRecord(t *testing.T) {
	biz, mockService := setupTestBiz()
	ctx := context.Background()

	tests := []struct {
		name           string
		recordID       string
		mockResponse   *services.AnswerResponse
		mockError      error
		expectedError  bool
		expectedRecord *services.AnswerResponse
	}{
		{
			name:     "successful retrieval",
			recordID: "test-record",
			mockResponse: &services.AnswerResponse{
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
			mockError:     nil,
			expectedError: false,
			expectedRecord: &services.AnswerResponse{
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
			mockResponse:   nil,
			mockError:      assert.AnError,
			expectedError:  true,
			expectedRecord: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockService.On("GetAnswerRecord", ctx, tt.recordID).Return(tt.mockResponse, tt.mockError).Once()

			// Execute
			record, err := biz.GetAnswerRecord(ctx, tt.recordID)

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

			mockService.AssertExpectations(t)
		})
	}
}

func TestGetExamPaper(t *testing.T) {
	biz, mockService := setupTestBiz()
	ctx := context.Background()

	tests := []struct {
		name          string
		examUUID      string
		mockPaper     *services.ExamPaper
		mockError     error
		expectedError bool
		expectedPaper *services.ExamPaper
	}{
		{
			name:     "successful retrieval",
			examUUID: "exam-uuid",
			mockPaper: &services.ExamPaper{
				ID:          1,
				Title:       "Test Paper",
				Description: "Test Description",
				Questions:   []services.Question{},
			},
			mockError:     nil,
			expectedError: false,
			expectedPaper: &services.ExamPaper{
				ID:          1,
				Title:       "Test Paper",
				Description: "Test Description",
				Questions:   []services.Question{},
			},
		},
		{
			name:          "paper not found",
			examUUID:      "non-existent",
			mockPaper:     nil,
			mockError:     assert.AnError,
			expectedError: true,
			expectedPaper: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockService.On("GetExamPaper", ctx, tt.examUUID).Return(tt.mockPaper, tt.mockError).Once()

			// Execute
			paper, err := biz.GetExamPaper(ctx, tt.examUUID)

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

			mockService.AssertExpectations(t)
		})
	}
}
