package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gin-go-test/app/biz"
	"gin-go-test/app/services"
)

// 直接使用全局 Redis 客户端，不调用 NewRedisHelper
func setupIntegrationRouter() *gin.Engine {
	// 假设全局 Redis 客户端已初始化
	answerService := services.NewAnswerServiceWithRedis(nil, nil) // 这里传入 nil，实际项目中应传入真实 Redis 客户端
	answerBiz := biz.NewExamAnswerBiz(answerService)
	controller := NewAnswerController(answerBiz) // 直接构造 controller 对象，不调用 NewExamAnswerController

	r := gin.Default()
	r.POST("/api/exam/answer", controller.SaveAnswer)
	r.GET("/api/exam/answer/:recordID", controller.GetAnswerRecord)
	r.GET("/api/exam/paper/:examUUID", controller.GetExamPaper)
	return r
}

func TestAnswerIntegration(t *testing.T) {
	r := setupIntegrationRouter()

	// 1. 保存答题记录
	payload := map[string]interface{}{
		"answer_uid": "test-uuid",
		"exam_id":    1,
		"exam_uuid":  "exam-uuid",
		"answers":    map[string]interface{}{"1": map[string]interface{}{"answer": 1, "score": 10}},
		"score":      10,
		"created_at": time.Now().Unix(),
		"user_uuid":  "user-uuid",
		"username":   "test-user",
		"user_id":    "user-1",
		"total_score": 10,
		"duration":    30,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "/api/exam/answer", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// 2. 查询答题记录
	req2, _ := http.NewRequest(http.MethodGet, "/api/exam/answer/test-uuid", nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)
	// 你可以进一步断言返回内容
} 