package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gin-go-test/utils"
	"log"
	"net/http"
	"github.com/google/uuid" // 用于生成UUID
	"time"
)

// 定义提交答题记录的结构体
type AnswerRequest struct {
	UUID    string          `json:"uuid" binding:"required"`   // 用户UUID
	ExamID  int64           `json:"exam_id" binding:"required"` // 试卷ID
	Answers json.RawMessage `json:"answers" binding:"required"` // 用户的答案（JSON格式）
}

func SubmitAnswer(c *gin.Context) {
	var answerReq AnswerRequest

	// 绑定请求中的 JSON 数据到结构体
	if err := c.ShouldBindJSON(&answerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 生成一个 UUID 作为唯一标识
	answerUID := uuid.New().String()

	// 计算总分
	var totalScore int
	var answerDetails []map[string]interface{}

	// 假设我们可以从请求的 answers 里计算每道题的得分
	// 这里我们使用答案计算一个假设的得分逻辑
	err := json.Unmarshal(answerReq.Answers, &answerDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse answers"})
		return
	}

	for _, answer := range answerDetails {
		// 假设每个问题都有一个 score 字段来保存每道题的得分
		score, ok := answer["score"].(int)
		if ok {
			totalScore += score
		}
	}

	// 生成答题记录
	answerData := map[string]interface{}{
		"answer_uid":  answerUID,
		"exam_id":    answerReq.ExamID,
		"user_uuid":  answerReq.UUID,
		"answers":    answerReq.Answers,
		"total_score": totalScore,
	}

	// 将数据存入 Redis
	redisKey := fmt.Sprintf("exam_answer:%s", answerUID)
	expiration := 7 * 24 * time.Hour // 设置过期时间为7天

	err = utils.RedisClient.HMSet(utils.Ctx, redisKey, answerData).Err()
	if err != nil {
		log.Println("Error writing to Redis:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save to Redis"})
		return
	}

	// 设置过期时间
	err = utils.RedisClient.Expire(utils.Ctx, redisKey, expiration).Err()
	if err != nil {
		log.Println("Error setting expiration:", err)
	}

	// 将数据存入 MySQL（在后台异步执行）
	go func() {
		// 这里是保存到数据库的逻辑
		query := `
		INSERT INTO ym_exam_answers (uuid, exam_id, uuid, answers, total_score, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, NOW(), NOW())`
		_, err := utils.DB.Exec(query, answerUID, answerReq.ExamID, answerReq.UUID, answerReq.Answers, totalScore)
		if err != nil {
			log.Println("Error inserting into MySQL:", err)
		}
	}()

	// 返回结果：将 UUID 和答题结果返回给前端
	c.JSON(http.StatusOK, gin.H{
		"message":    "Answer submitted successfully",
		"answer_uid": answerUID,
	})
}