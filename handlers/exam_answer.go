package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gin-go-test/utils"
)

// 数据结构定义
type AnswerRequest struct {
	UUID       string                     `json:"uuid" binding:"required"`
	ExamID     int64                      `json:"exam_id" binding:"required"`
	ExamUUID   string                     `json:"exam_uuid"`
	Answers    map[string]json.RawMessage `json:"answers" binding:"required"`
	Username   string                     `json:"username"`
	UserID     string                     `json:"user_id"`
	Duration   int                        `json:"duration"`
	FullScore  int                        `json:"full_score"`
}

type AnswerResponse struct {
    UUID       string           `json:"uuid"`
    ExamID     int64            `json:"exam_id"` // 保持为int64
    ExamUUID   string           `json:"exam_uuid"`
    UserUUID   string           `json:"user_uuid"`
    Answers    json.RawMessage  `json:"answers"`
    TotalScore int              `json:"total_score"`
    CreatedAt  int64            `json:"created_at"`
    Username   string           `json:"username"`
    UserID     string           `json:"user_id"`
    Duration   int              `json:"duration"`
    Score      int              `json:"score"`  // 修改字段名为 score
}

// 新增数据结构
type ExamPaper struct {
    ID       int        `json:"id"`
    Title    string     `json:"title"`
    Questions []Question `json:"questions"`
}

type Question struct {
    ID                   int         `json:"id"`
    Title                string      `json:"title"`
    Options              []string    `json:"options"`
    Type                 string      `json:"type"` // single/multi/judge
    Score                int         `json:"score"`
    CorrectAnswerBitmask int         `json:"correct_answer_bitmask"` // bitmask 字段
    CorrectAnswer        interface{} `json:"correct_answer"`         // 存储解析后的正确答案
    Analysis             string      `json:"analysis"`
}

type FullAnswerResponse struct {
    RecordID    string               `json:"record_id"`
    ExamID      interface{}          `json:"exam_id"` // 兼容数字/字符串
    UserUUID    string               `json:"user_uuid"`
    TotalScore  int                  `json:"total_score"`
    CreatedAt   int64                `json:"created_at"`
    Questions   []QuestionWithAnswer `json:"questions"`
    Username    string               `json:"username"`
    UserID      string               `json:"user_id"`
    Duration    int                  `json:"duration"`
    Score       int                  `json:"score"`
}

type QuestionWithAnswer struct {
    ID           int         `json:"id"`
    Title        string      `json:"title"`
    Options      []string    `json:"options"`
    Type         string      `json:"type"`
    Score        int         `json:"score"`
    CorrectAnswer interface{} `json:"correct_answer"`
    UserAnswer   interface{} `json:"user_answer"`
    IsCorrect    bool        `json:"is_correct"`
    Analysis     string      `json:"analysis"`
}

// SubmitAnswer 提交答题记录
func SubmitAnswer(c *gin.Context) {
	var req AnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "请求参数错误", err)
		return
	}

	if len(req.Answers) == 0 {
		sendErrorResponse(c, http.StatusBadRequest, "答案不能为空", nil)
		return
	}

	totalScore := 0
	for _, answer := range req.Answers {
		var detail struct{ Score int }
		if err := json.Unmarshal(answer, &detail); err == nil {
			totalScore += detail.Score
		}
	}

	recordID := uuid.New().String()
	createdAt := time.Now().Unix()

	record := map[string]interface{}{
		"answer_uid":  recordID,
		"exam_id":    req.ExamID,
		"exam_uuid":  req.ExamUUID,
		"user_uuid":  req.UUID,
		"answers":    req.Answers,
		"total_score": totalScore,
		"created_at": createdAt,
		"username": req.Username,
		"user_id": req.UserID,
		"duration": req.Duration,
		"score": totalScore, // 修改为使用 score 字段
	}

	if err := saveToRedis(record); err != nil {
		log.Printf("Redis保存失败: %v", err)
		sendErrorResponse(c, http.StatusInternalServerError, "服务端存储错误", err)
		return
	}

	go asyncSaveToDatabase(record, c.Request.Context())

	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"message":   "提交成功",
		"record_id": recordID,
		"score":     totalScore,
		"timestamp": createdAt,
	})
}

// GetAnswerResult 获取答题结果
func GetAnswerResult(c *gin.Context) {
	recordID := c.Param("record_id")
	if recordID == "" {
		sendErrorResponse(c, http.StatusBadRequest, "记录ID不能为空", nil)
		return
	}

	// 从Redis获取数据
	redisKey := fmt.Sprintf("exam_answer:%s", recordID)
	result, err := utils.RedisClient.HGetAll(utils.Ctx, redisKey).Result()
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, "获取答题记录失败", err)
		return
	}

	if len(result) == 0 {
		sendErrorResponse(c, http.StatusNotFound, "答题记录不存在", nil)
		return
	}

	// 转换数据类型
	examID, _ := strconv.ParseInt(result["exam_id"], 10, 64)
	totalScore, _ := strconv.Atoi(result["total_score"])
	createdAt, _ := strconv.ParseInt(result["created_at"], 10, 64)

	username := result["username"]
	userID := result["user_id"]
	duration, _ := strconv.Atoi(result["duration"])
	score, _ := strconv.Atoi(result["score"])

	response := AnswerResponse{
		UUID:       recordID,
		ExamID:     examID,
		ExamUUID:   result["exam_uuid"],
		UserUUID:   result["user_uuid"],
		Answers:    []byte(result["answers"]),
		TotalScore: totalScore,
		CreatedAt:  createdAt,
		Username:   username,
		UserID:     userID,
		Duration:   duration,
		Score:      score,
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    response,
	})
}


// 辅助函数
func saveToRedis(data map[string]interface{}) error {
	redisData := make(map[string]string)
	for k, v := range data {
		switch val := v.(type) {
		case string:
			redisData[k] = val
		case int:
			redisData[k] = strconv.Itoa(val)
		case int64:
			redisData[k] = strconv.FormatInt(val, 10)
		default:
			jsonVal, err := json.Marshal(v)
			if err != nil {
				return fmt.Errorf("字段%s序列化失败: %v", k, err)
			}
			redisData[k] = string(jsonVal)
		}
	}

	redisKey := fmt.Sprintf("exam_answer:%s", data["answer_uid"])
	if err := utils.RedisClient.HMSet(utils.Ctx, redisKey, redisData).Err(); err != nil {
		return err
	}
	return utils.RedisClient.Expire(utils.Ctx, redisKey, 7 * 24*time.Hour).Err()
}

func asyncSaveToDatabase(data map[string]interface{}, ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("请求已取消，中止数据库保存")
		return
	default:
	}

	answersJSON, err := json.Marshal(data["answers"])
	if err != nil {
		log.Printf("答案序列化失败: %v", err)
		return
	}

	query := `INSERT INTO ym_exam_answers (
		uuid, exam_id, exam_uuid, user_uuid, 
		answers, total_score, created_at, 
		username, user_id, duration, score
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	
	_, err = utils.DB.ExecContext(
		ctx,
		query,
		data["answer_uid"],
		data["exam_id"],
		data["exam_uuid"],
		data["user_uuid"],
		answersJSON,
		data["total_score"],
		time.Unix(data["created_at"].(int64), 0),
		data["username"],
		data["user_id"],
		data["duration"],
		data["score"],
	)
	
	if err != nil {
		log.Printf("数据库保存失败: %v", err)
	}
}

func sendErrorResponse(c *gin.Context, code int, message string, err error) {
	response := gin.H{
		"code":    code,
		"message": message,
	}
	if err != nil {
		response["error"] = err.Error()
	}
	c.JSON(code, response)
}

// 解析 bitmask
func decodeCorrectAnswer(bitmask int) []int {
	correctAnswers := []int{}
	for i := 0; i < 32; i++ {
		if (bitmask & (1 << i)) != 0 {
			correctAnswers = append(correctAnswers, i)
		}
	}
	return correctAnswers
}

// GetFullAnswerResult 获取完整答题结果（包含题目信息和正确答案）
func GetFullAnswerResult(c *gin.Context) {
	recordID := c.Param("record_id")
	if recordID == "" {
		sendErrorResponse(c, http.StatusBadRequest, "记录ID不能为空", nil)
		return
	}

	// 获取答题记录
	record, err := getAnswerRecord(recordID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, "获取答题记录失败", err)
		return
	}

	// 获取试卷信息
	paper, err := getExamPaper(record.ExamUUID)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, "获取试卷信息失败", err)
		return
	}
	if paper == nil {
		sendErrorResponse(c, http.StatusNotFound, "试卷信息缺失，请联系管理员", nil)
		return
	}

	// 解析每个题目的正确答案
	for i, q := range paper.Questions {
		// 使用解码函数解析bitmask
		paper.Questions[i].CorrectAnswer = decodeCorrectAnswer(q.CorrectAnswerBitmask)
	}

	// 构建响应
	response := buildFullResponse(record, paper)
	if response == nil {
		sendErrorResponse(c, http.StatusInternalServerError, "构建响应失败", nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    response,
	})
}


// 构建完整答题结果
func buildFullResponse(record *AnswerResponse, paper *ExamPaper) *FullAnswerResponse {
    // 解析用户答案
    var userAnswers map[string]struct {
        Answer interface{} `json:"answer"`
        Score  int         `json:"score"`
    }
    if err := json.Unmarshal(record.Answers, &userAnswers); err != nil {
        log.Printf("解析用户答案失败: %v", err)
        return nil
    }

    // 构建题目列表
    var questions []QuestionWithAnswer
    for _, q := range paper.Questions {
        qid := strconv.Itoa(q.ID)
        if userAns, exists := userAnswers[qid]; exists {
            questions = append(questions, QuestionWithAnswer{
                ID:            q.ID,
                Title:         q.Title,
                Options:       q.Options,
                Type:          q.Type,
                Score:         q.Score,
                CorrectAnswer: q.CorrectAnswer,
                UserAnswer:    userAns.Answer,
                IsCorrect:     isAnswerCorrect(q, userAns.Answer),
                Analysis:      q.Analysis,
            })
        }
    }

    return &FullAnswerResponse{
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
    }
}

// 判断答案是否正确
func isAnswerCorrect(q Question, userAnswer interface{}) bool {
    switch q.Type {
    case "multi":
        userAns, ok1 := userAnswer.([]interface{})
        correctAns, ok2 := q.CorrectAnswer.([]interface{})
        if !ok1 || !ok2 {
            return false
        }
        return compareSlicesOrdered(userAns, correctAns)
    default:
        return userAnswer == q.CorrectAnswer
    }
}

// 比较多选题的答案
func compareSlicesOrdered(a, b []interface{}) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

// 修改 getAnswerRecord 返回 *AnswerResponse 而不是 map
func getAnswerRecord(recordID string) (*AnswerResponse, error) {
    redisKey := fmt.Sprintf("exam_answer:%s", recordID)
    result, err := utils.RedisClient.HGetAll(utils.Ctx, redisKey).Result()
    if err != nil {
        return nil, fmt.Errorf("获取答题记录失败: %v", err)
    }
    
    examID, _ := strconv.ParseInt(result["exam_id"], 10, 64)
    examUUID := result["exam_uuid"]
    totalScore, _ := strconv.Atoi(result["total_score"])
    createdAt, _ := strconv.ParseInt(result["created_at"], 10, 64)
    username := result["username"]
    userID := result["user_id"]
    duration, _ := strconv.Atoi(result["duration"])
    score, _ := strconv.Atoi(result["score"])
    
    return &AnswerResponse{
        UUID:       recordID,
        ExamID:     examID,
        ExamUUID:   examUUID,
        UserUUID:   result["user_uuid"],
        Answers:    []byte(result["answers"]),
        TotalScore: totalScore,
        CreatedAt:  createdAt,
        Username:   username,
        UserID:     userID,
        Duration:   duration,
        Score:      score,
    }, nil
}

// getExamPaper fetches the exam paper details from Redis using the provided exam UUID
func getExamPaper(examUUID string) (*ExamPaper, error) {
    // Create the Redis key to fetch the exam paper
    redisKey := fmt.Sprintf("exam_paper:%s", examUUID)

    // Fetch the data from Redis
    result, err := utils.RedisClient.Get(utils.Ctx, redisKey).Result()
    if err != nil {
        return nil, fmt.Errorf("获取试卷数据失败: %v", err)
    }

    // Parse the data into the ExamPaper struct
    var paper ExamPaper
    if err := json.Unmarshal([]byte(result), &paper); err != nil {
        // If it's a stringified JSON array, we need to handle it differently
        var rawPaper map[string]interface{}
        if err := json.Unmarshal([]byte(result), &rawPaper); err != nil {
            return nil, fmt.Errorf("解析试卷数据失败: %v", err)
        }
        // Assuming the 'questions' field is a stringified array in rawPaper, we can manually unmarshal it
        if questionsJSON, ok := rawPaper["questions"].(string); ok {
            var questions []Question
            if err := json.Unmarshal([]byte(questionsJSON), &questions); err != nil {
                return nil, fmt.Errorf("解析试题数据失败: %v", err)
            }
            paper.Questions = questions
        }
        // 其他字段赋值
        if id, ok := rawPaper["id"].(float64); ok {
            paper.ID = int(id)
        }
        if title, ok := rawPaper["title"].(string); ok {
            paper.Title = title
        }
    }

    return &paper, nil
}