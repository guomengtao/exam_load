package controllers

import (
	"bytes"
	"encoding/json"
	"gin-go-test/biz"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockAnswerBiz is a mock implementation of biz.AnswerBiz
type MockAnswerBiz struct {
	mock.Mock
}

func (m *MockAnswerBiz) SubmitAnswer(ctx *gin.Context, req *biz.AnswerRequest) (*biz.AnswerResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*biz.AnswerResponse), args.Error(1)
}

func (m *MockAnswerBiz) GetAnswerResult(ctx *gin.Context, recordID string) (*biz.AnswerResponse, error) {
	args := m.Called(ctx, recordID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*biz.AnswerResponse), args.Error(1)
}

func (m *MockAnswerBiz) GetFullAnswerResult(ctx *gin.Context, recordID string) (*biz.FullAnswerResponse, error) {
	args := m.Called(ctx, recordID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*biz.FullAnswerResponse), args.Error(1)
}

func setupTestRouter() (*gin.Engine, *MockAnswerBiz) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	mockBiz := new(MockAnswerBiz)
	controller := &AnswerController{answerBiz: mockBiz}

	// Setup routes
	router.POST("/api/user/answer", controller.SubmitAnswer)
	router.GET("/api/user/answer/:record_id", controller.GetAnswerResult)
	router.GET("/api/user/answer/:record_id/full", controller.GetFullAnswerResult)

	return router, mockBiz
}

func TestSubmitAnswer(t *testing.T) {
	router, mockBiz := setupTestRouter()

	tests := []struct {
		name           string
		requestBody    biz.AnswerRequest
		mockResponse   *biz.AnswerResponse
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "successful submission",
			requestBody: biz.AnswerRequest{
				UUID:     "test-uuid",
				ExamID:   1,
				ExamUUID: "exam-uuid",
				Answers: map[string]json.RawMessage{
					"1": json.RawMessage(`{"answer": 1, "score": 10}`),
				},
			},
			mockResponse: &biz.AnswerResponse{
				UUID:       "test-uuid",
				ExamID:     1,
				ExamUUID:   "exam-uuid",
				TotalScore: 10,
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"code":    200,
				"message": "提交成功",
			},
		},
		{
			name: "invalid request",
			requestBody: biz.AnswerRequest{
				UUID:   "test-uuid",
				ExamID: 1,
			},
			mockResponse:   nil,
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"code":    500,
				"message": "提交答题失败",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockBiz.On("SubmitAnswer", mock.Anything, &tt.requestBody).Return(tt.mockResponse, tt.mockError)

			// Create request
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/api/user/answer", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(w, req)

			// Assertions
			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			for k, v := range tt.expectedBody {
				assert.Equal(t, v, response[k])
			}
		})
	}
}

func TestGetAnswerResult(t *testing.T) {
	router, mockBiz := setupTestRouter()

	tests := []struct {
		name           string
		recordID       string
		mockResponse   *biz.AnswerResponse
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:     "successful retrieval",
			recordID: "test-record",
			mockResponse: &biz.AnswerResponse{
				UUID:       "test-record",
				ExamID:     1,
				TotalScore: 10,
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"code":    200,
				"message": "获取成功",
			},
		},
		{
			name:           "record not found",
			recordID:       "non-existent",
			mockResponse:   nil,
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"code":    500,
				"message": "获取答题记录失败",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockBiz.On("GetAnswerResult", mock.Anything, tt.recordID).Return(tt.mockResponse, tt.mockError)

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/api/user/answer/"+tt.recordID, nil)
			w := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(w, req)

			// Assertions
			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			for k, v := range tt.expectedBody {
				assert.Equal(t, v, response[k])
			}
		})
	}
}

func TestGetFullAnswerResult(t *testing.T) {
	router, mockBiz := setupTestRouter()

	tests := []struct {
		name           string
		recordID       string
		mockResponse   *biz.FullAnswerResponse
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:     "successful retrieval",
			recordID: "test-record",
			mockResponse: &biz.FullAnswerResponse{
				RecordID:   "test-record",
				TotalScore: 10,
				Questions:  []biz.QuestionWithAnswer{},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"code":    200,
				"message": "获取成功",
			},
		},
		{
			name:           "record not found",
			recordID:       "non-existent",
			mockResponse:   nil,
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"code":    500,
				"message": "获取完整答题记录失败",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockBiz.On("GetFullAnswerResult", mock.Anything, tt.recordID).Return(tt.mockResponse, tt.mockError)

			// Create request
			req := httptest.NewRequest(http.MethodGet, "/api/user/answer/"+tt.recordID+"/full", nil)
			w := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(w, req)

			// Assertions
			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			for k, v := range tt.expectedBody {
				assert.Equal(t, v, response[k])
			}
		})
	}
}
