package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"gin-go-test/app/biz"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// MockExamAnswerBiz is a mock implementation of ExamAnswerBiz
type MockExamAnswerBiz struct {
	mock.Mock
}

func (m *MockExamAnswerBiz) SaveAnswer(ctx context.Context, data map[string]interface{}) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *MockExamAnswerBiz) GetAnswerRecord(ctx context.Context, recordID string) (*biz.AnswerResponse, error) {
	args := m.Called(ctx, recordID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*biz.AnswerResponse), args.Error(1)
}

func (m *MockExamAnswerBiz) GetExamPaper(ctx context.Context, examUUID string) (*biz.ExamPaper, error) {
	args := m.Called(ctx, examUUID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*biz.ExamPaper), args.Error(1)
}

func setupTestController() (*ExamAnswerController, *MockExamAnswerBiz) {
	mockBiz := new(MockExamAnswerBiz)
	controller := NewExamAnswerController(mockBiz)
	return controller, mockBiz
}

func TestSaveAnswer(t *testing.T) {
	controller, mockBiz := setupTestController()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/exam/answer", controller.SaveAnswer)

	tests := []struct {
		name           string
		payload        map[string]interface{}
		mockError      error
		expectedStatus int
	}{
		{
			name: "successful save",
			payload: map[string]interface{}{
				"answer_uid": "test-uuid",
				"exam_id":    1,
				"exam_uuid":  "exam-uuid",
				"answers":    map[string]interface{}{"1": map[string]interface{}{"answer": 1, "score": 10}},
				"score":      10,
				"created_at": time.Now(),
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name: "biz error",
			payload: map[string]interface{}{
				"answer_uid": "test-uuid",
				"exam_id":    1,
				"exam_uuid":  "exam-uuid",
			},
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockBiz.On("SaveAnswer", mock.Anything, mock.Anything).Return(tt.mockError).Once()

			// Create request
			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest(http.MethodPost, "/api/exam/answer", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Execute
			router.ServeHTTP(w, req)

			// Assertions
			assert.Equal(t, tt.expectedStatus, w.Code)
			mockBiz.AssertExpectations(t)
		})
	}
}

func TestGetAnswerRecord(t *testing.T) {
	controller, mockBiz := setupTestController()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/exam/answer/:recordID", controller.GetAnswerRecord)

	tests := []struct {
		name           string
		recordID       string
		mockResponse   *biz.AnswerResponse
		mockError      error
		expectedStatus int
	}{
		{
			name:     "successful retrieval",
			recordID: "test-record",
			mockResponse: &biz.AnswerResponse{
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
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "record not found",
			recordID:       "non-existent",
			mockResponse:   nil,
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockBiz.On("GetAnswerRecord", mock.Anything, tt.recordID).Return(tt.mockResponse, tt.mockError).Once()

			// Create request
			req, _ := http.NewRequest(http.MethodGet, "/api/exam/answer/"+tt.recordID, nil)
			w := httptest.NewRecorder()

			// Execute
			router.ServeHTTP(w, req)

			// Assertions
			assert.Equal(t, tt.expectedStatus, w.Code)
			mockBiz.AssertExpectations(t)
		})
	}
}

func TestGetExamPaper(t *testing.T) {
	controller, mockBiz := setupTestController()
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/exam/paper/:examUUID", controller.GetExamPaper)

	tests := []struct {
		name           string
		examUUID       string
		mockPaper      *biz.ExamPaper
		mockError      error
		expectedStatus int
	}{
		{
			name:     "successful retrieval",
			examUUID: "exam-uuid",
			mockPaper: &biz.ExamPaper{
				ID:          1,
				Title:       "Test Paper",
				Description: "Test Description",
				Questions:   []biz.Question{},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "paper not found",
			examUUID:       "non-existent",
			mockPaper:      nil,
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockBiz.On("GetExamPaper", mock.Anything, tt.examUUID).Return(tt.mockPaper, tt.mockError).Once()

			// Create request
			req, _ := http.NewRequest(http.MethodGet, "/api/exam/paper/"+tt.examUUID, nil)
			w := httptest.NewRecorder()

			// Execute
			router.ServeHTTP(w, req)

			// Assertions
			assert.Equal(t, tt.expectedStatus, w.Code)
			mockBiz.AssertExpectations(t)
		})
	}
}
