package biz

import (
	"context"
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gin-go-test/services"
)

// MockAnswerService is a mock implementation of services.AnswerService
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

func (m *MockAnswerService) GetAnswerRecord(ctx context.Context, recordID string) (*AnswerResponse, error) {
	args := m.Called(ctx, recordID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*AnswerResponse), args.Error(1)
}

func (m *MockAnswerService) GetExamPaper(ctx context.Context, examUUID string) (*ExamPaper, error) {
	args := m.Called(ctx, examUUID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ExamPaper), args.Error(1)
}

func setupTestBiz() (*AnswerBiz, *MockAnswerService) {
	mockService := new(MockAnswerService)
	biz := NewAnswerBiz(mockService)
	return biz, mockService
}

func TestSubmitAnswer(t *testing.T) {
	biz, mockService := setupTestBiz()
	ctx := context.Background()

	tests := []struct {
		name           string
		request        *AnswerRequest
		mockError      error
		expectedError  bool
		expectedScore  int
	}{
		{
			name: "successful submission",
			request: &AnswerRequest{
				UUID:     "test-uuid",
				ExamID:   1,
				ExamUUID: "exam-uuid",
				Answers: map[string]json.RawMessage{
					"1": json.RawMessage(`{"answer": 1, "score": 10}`),
					"2": json.RawMessage(`{"answer": 2, "score": 20}`),
				},
			},
			mockError:     nil,
			expectedError: false,
			expectedScore: 30,
		},
		{
			name: "empty answers",
			request: &AnswerRequest{
				UUID:     "test-uuid",
				ExamID:   1,
				ExamUUID: "exam-uuid",
				Answers:  map[string]json.RawMessage{},
			},
			mockError:     nil,
			expectedError: true,
			expectedScore: 0,
		},
		{
			name: "redis error",
			request: &AnswerRequest{
				UUID:     "test-uuid",
				ExamID:   1,
				ExamUUID: "exam-uuid",
				Answers: map[string]json.RawMessage{
					"1": json.RawMessage(`{"answer": 1, "score": 10}`),
				},
			},
			mockError:     assert.AnError,
			expectedError: true,
			expectedScore: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			if !tt.expectedError {
				mockService.On("SaveToRedis", ctx, mock.Anything).Return(nil)
				mockService.On("AsyncSaveToDatabase", ctx, mock.Anything).Return()
			} else if tt.request.Answers == nil {
				mockService.On("SaveToRedis", ctx, mock.Anything).Return(tt.mockError)
			}

			// Execute
			response, err := biz.SubmitAnswer(ctx, tt.request)

			// Assertions
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, tt.expectedScore, response.TotalScore)
				assert.Equal(t, tt.expectedScore, response.Score)
			}
		})
	}
}

func TestGetAnswerResult(t *testing.T) {
	biz, mockService := setupTestBiz()
	ctx := context.Background()

	tests := []struct {
		name           string
		recordID       string
		mockResponse   *AnswerResponse
		mockError      error
		expectedError  bool
	}{
		{
			name:     "successful retrieval",
			recordID: "test-record",
			mockResponse: &AnswerResponse{
				UUID:       "test-record",
				ExamID:     1,
				TotalScore: 10,
			},
			mockError:     nil,
			expectedError: false,
		},
		{
			name:           "record not found",
			recordID:       "non-existent",
			mockResponse:   nil,
			mockError:      assert.AnError,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockService.On("GetAnswerRecord", ctx, tt.recordID).Return(tt.mockResponse, tt.mockError)

			// Execute
			response, err := biz.GetAnswerResult(ctx, tt.recordID)

			// Assertions
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, tt.mockResponse, response)
			}
		})
	}
}

func TestGetFullAnswerResult(t *testing.T) {
	biz, mockService := setupTestBiz()
	ctx := context.Background()

	tests := []struct {
		name           string
		recordID       string
		mockAnswer     *AnswerResponse
		mockPaper      *ExamPaper
		mockError      error
		expectedError  bool
	}{
		{
			name:     "successful retrieval",
			recordID: "test-record",
			mockAnswer: &AnswerResponse{
				UUID:       "test-record",
				ExamID:     1,
				ExamUUID:   "exam-uuid",
				TotalScore: 10,
			},
			mockPaper: &ExamPaper{
				ID:          1,
				Title:       "Test Paper",
				Description: "Test Description",
				Questions:   []Question{},
			},
			mockError:     nil,
			expectedError: false,
		},
		{
			name:           "answer not found",
			recordID:       "non-existent",
			mockAnswer:     nil,
			mockPaper:      nil,
			mockError:      assert.AnError,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockService.On("GetAnswerRecord", ctx, tt.recordID).Return(tt.mockAnswer, tt.mockError)
			if tt.mockAnswer != nil {
				mockService.On("GetExamPaper", ctx, tt.mockAnswer.ExamUUID).Return(tt.mockPaper, nil)
			}

			// Execute
			response, err := biz.GetFullAnswerResult(ctx, tt.recordID)

			// Assertions
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, response)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, tt.mockAnswer.UUID, response.RecordID)
				assert.Equal(t, tt.mockPaper.Title, response.Title)
			}
		})
	}
}

func TestIsAnswerCorrect(t *testing.T) {
	biz, _ := setupTestBiz()

	tests := []struct {
		name           string
		question       Question
		userAnswer     interface{}
		expectedResult bool
	}{
		{
			name: "single choice correct",
			question: Question{
				Type:          "single",
				CorrectAnswer: 1,
			},
			userAnswer:     1,
			expectedResult: true,
		},
		{
			name: "single choice incorrect",
			question: Question{
				Type:          "single",
				CorrectAnswer: 1,
			},
			userAnswer:     2,
			expectedResult: false,
		},
		{
			name: "multiple choice correct",
			question: Question{
				Type:          "multi",
				CorrectAnswer: []int{1, 2, 3},
			},
			userAnswer:     []int{1, 2, 3},
			expectedResult: true,
		},
		{
			name: "multiple choice incorrect order",
			question: Question{
				Type:          "multi",
				CorrectAnswer: []int{1, 2, 3},
			},
			userAnswer:     []int{3, 2, 1},
			expectedResult: true,
		},
		{
			name: "multiple choice wrong",
			question: Question{
				Type:          "multi",
				CorrectAnswer: []int{1, 2, 3},
			},
			userAnswer:     []int{1, 2, 4},
			expectedResult: false,
		},
		{
			name: "judge correct",
			question: Question{
				Type:          "judge",
				CorrectAnswer: 1,
			},
			userAnswer:     1,
			expectedResult: true,
		},
		{
			name: "judge incorrect",
			question: Question{
				Type:          "judge",
				CorrectAnswer: 1,
			},
			userAnswer:     0,
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := biz.isAnswerCorrect(tt.question, tt.userAnswer)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
} 