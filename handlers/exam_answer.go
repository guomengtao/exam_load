
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
	"sort"
)

// AnswerMap 用于 Swagger 显示答题数据格式
type AnswerMap map[string]interface{}

// ErrorResponse 通用错误返回结构
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// SuccessResponse 通用成功返回结构
type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

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
	UUID       string      `json:"uuid"`
	ExamID     int64       `json:"exam_id"`
	ExamUUID   string      `json:"exam_uuid"`
	UserUUID   string      `json:"user_uuid"`
	Answers    interface{} `json:"answers"` // for Swagger compatibility
	TotalScore int         `json:"total_score"`
	CreatedAt  int64       `json:"created_at"`
	Username   string      `json:"username"`
	UserID     string      `json:"user_id"`
	Duration   int         `json:"duration"`
	Score      int         `json:"score"`
}

// 新增数据结构
type ExamPaper struct {
    ID          int        `json:"id"`
    Title       string     `json:"title"`
    Description string     `json:"description"`
    Questions   []Question `json:"questions"`
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
    Title       string               `json:"title"`
    Description string               `json:"description"`
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
// @Summary 提交用户的答题记录
// @Description 用户完成答题后提交记录，并保存到 Redis 和数据库
// @Tags exam_answer
// @Accept json
// @Produce json
// @Param uuid body string true "用户UUID"
// @Param exam_id body int64 true "试卷ID"
// @Param exam_uuid body string false "试卷UUID"
// @Param answers body AnswerMap true "用户答题数据"
// @Param username body string false "用户名"
// @Param user_id body string false "用户学号"
// @Param duration body int false "考试时长"
// @Param full_score body int false "试卷总分"
// @Success 200 {object} AnswerResponse "返回答题记录"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 500 {object} ErrorResponse "服务器错误"
// @Router /api/exam/answer [post]
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

// GetAnswerResult 获取答题记录
// @Summary 获取用户的答题记录
// @Description 通过答题记录ID获取用户的答题结果
// @Tags exam_answer
// @Accept json
// @Produce json
// @Param record_id path string true "答题记录ID"
// @Success 200 {object} AnswerResponse "返回用户答题记录"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 404 {object} ErrorResponse "未找到答题记录"
// @Failure 500 {object} ErrorResponse "服务器错误"
// @Router /api/user/answer/{record_id} [get]
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

	// 转换答案字段为字节切片
	var answers []byte
	if result["answers"] != "" {
		answers = []byte(result["answers"])
	}

	response := AnswerResponse{
		UUID:       recordID,
		ExamID:     examID,
		ExamUUID:   result["exam_uuid"],
		UserUUID:   result["user_uuid"],
		Answers:    answers, // ensure answers is []byte here
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


// saveToRedis 将答题记录保存到 Redis
// @Summary 将答题记录存储到 Redis
// @Description 将用户的答题记录保存到 Redis，以便后续查询
// @Tags exam_answer
// @Accept json
// @Produce json
// @Param data body map[string]interface{} true "答题记录"
// @Success 200 {object} SuccessResponse "保存成功"
// @Failure 500 {object} ErrorResponse "保存失败"
// @Router /api/redis/save [post]
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

// @Tags exam_answer
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

// decodeCorrectAnswer 解码 bitmask 为正确答案数组
// @Tags exam_answer
// @Description 将 bitmask 转换为正确答案的数组
// @Param bitmask body int true "bitmask"
// @Success 200 {array} int "返回正确答案数组"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Router /api/exam/bitmask [post]
func decodeCorrectAnswer(bitmask int) []int {
	var answers []int
	for i := 0; i < 32; i++ {
		if (bitmask & (1 << i)) != 0 {
			answers = append(answers, i)
		}
	}
	return answers
}

// GetFullAnswerResult 获取完整答题结果（包含题目信息和正确答案）
// @Summary 获取完整的答题记录，包含试卷信息、用户答案、正确答案等
// @Description 获取用户的完整答题结果，包括试卷标题、描述、问题、答案等详细信息
// @Tags exam_answer
// @Accept json
// @Produce json
// @Param record_id path string true "答题记录ID"
// @Success 200 {object} FullAnswerResponse "返回完整的答题记录"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 500 {object} ErrorResponse "服务器错误"
// @Router /api/user/answer/{record_id}/full [get]
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


// buildFullResponse 构建完整答题结果
// @Summary 构建包含用户答案和正确答案的详细答题记录
// @Description 通过用户的答题记录和试卷信息，构建完整的答题结果
// @Tags exam_answer
// @Accept json
// @Produce json
// @Param record body AnswerResponse true "答题记录"
// @Param paper body ExamPaper true "试卷信息"
// @Success 200 {object} FullAnswerResponse "返回完整的答题记录"
// @Failure 500 {object} ErrorResponse "服务器错误"
// @Router /api/exam/fullresult [post]
func buildFullResponse(record *AnswerResponse, paper *ExamPaper) *FullAnswerResponse {
    // 解析用户答案
    var userAnswers map[string]struct {
        Answer interface{} `json:"answer"`
        Score  int         `json:"score"`
    }
    // 断言 record.Answers 为 []byte
    answersBytes, ok := record.Answers.([]byte)
    if !ok {
        log.Printf("record.Answers 断言为 []byte 失败，实际类型: %T", record.Answers)
        return nil
    }
    if err := json.Unmarshal(answersBytes, &userAnswers); err != nil {
        log.Printf("解析用户答案失败: %v", err)
        return nil
    }

    // 构建题目列表
    var questions []QuestionWithAnswer
    for _, q := range paper.Questions {
        qid := strconv.Itoa(q.ID)
        if userAns, exists := userAnswers[qid]; exists {
            isCorrect := isAnswerCorrect(q, userAns.Answer)
            questions = append(questions, QuestionWithAnswer{
                ID:            q.ID,
                Title:         q.Title,
                Options:       q.Options,
                Type:          q.Type,
                Score:         q.Score,
                CorrectAnswer: q.CorrectAnswer,
                UserAnswer:    userAns.Answer,
                IsCorrect:     isCorrect,
                Analysis:      q.Analysis,
            })
        } else {
            // 用户未答该题，也加入列表，isCorrect为false，UserAnswer为nil
            questions = append(questions, QuestionWithAnswer{
                ID:            q.ID,
                Title:         q.Title,
                Options:       q.Options,
                Type:          q.Type,
                Score:         q.Score,
                CorrectAnswer: q.CorrectAnswer,
                UserAnswer:    nil,
                IsCorrect:     false,
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
        Title:       paper.Title,
        Description: paper.Description,
    }
}

// 标准化答案为 []int，兼容多种输入类型
func normalizeAnswer(answer interface{}) []int {
	switch v := answer.(type) {
	case []interface{}:
		var result []int
		for _, item := range v {
			switch iv := item.(type) {
			case float64:
				result = append(result, int(iv))
			case int:
				result = append(result, iv)
			case json.Number:
				if i, err := iv.Int64(); err == nil {
					result = append(result, int(i))
				}
			}
		}
		return result
	case []int:
		return v
	case []float64:
		var result []int
		for _, item := range v {
			result = append(result, int(item))
		}
		return result
	case float64:
		return []int{int(v)}
	case int:
		return []int{v}
	case json.Number:
		if i, err := v.Int64(); err == nil {
			return []int{int(i)}
		}
		return nil
	default:
		return nil
	}
}

// isAnswerCorrect 判断答案是否正确，兼容单选/多选，顺序无关
// @Summary 判断用户的答案是否正确
// @Description 判断用户答案和正确答案是否一致，兼容顺序不同的情况
// @Tags exam
// @Accept json
// @Produce json
// @Param question body Question true "试题信息"
// @Param userAnswer body interface{} true "用户的答案"
// @Success 200 {boolean} true "返回是否正确"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Router /api/exam/iscorrect [post]
func isAnswerCorrect(q Question, userAnswer interface{}) bool {
	correct := normalizeAnswer(q.CorrectAnswer)
	user := normalizeAnswer(userAnswer)
	if correct == nil || user == nil {
		return false
	}
	if len(correct) != len(user) {
		return false
	}
	// 排序后比较
	sort.Ints(correct)
	sort.Ints(user)
	for i, v := range correct {
		if user[i] != v {
			return false
		}
	}
	return true
}


// getAnswerRecord 获取答题记录
// @Summary 获取用户的答题记录
// @Description 通过答题记录ID获取用户答题记录
// @Tags exam_answer
// @Accept json
// @Produce json
// @Param record_id path string true "答题记录ID"
// @Success 200 {object} AnswerResponse "返回用户答题记录"
// @Failure 500 {object} ErrorResponse "服务器错误"
// @Router /api/exam/answerrecord/{record_id} [get]
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

// getExamPaper 获取试卷信息
// @Summary 获取试卷信息
// @Description 根据试卷UUID获取试卷信息
// @Tags exam_answer
// @Accept json
// @Produce json
// @Param exam_uuid path string true "试卷UUID"
// @Success 200 {object} ExamPaper "返回试卷信息"
// @Failure 500 {object} ErrorResponse "服务器错误"
// @Router /api/exam/paper/{exam_uuid} [get]
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
        // If it's a stringified JSON object, try to parse as map
        var rawPaper map[string]interface{}
        if err := json.Unmarshal([]byte(result), &rawPaper); err != nil {
            return nil, fmt.Errorf("解析试卷数据失败: %v", err)
        }
        // Parse questions field
        questionsRaw, ok := rawPaper["questions"]
        if !ok {
            return nil, fmt.Errorf("试卷数据缺少 questions 字段")
        }
        switch questionsVal := questionsRaw.(type) {
        case string:
            var questions []Question
            if err := json.Unmarshal([]byte(questionsVal), &questions); err != nil {
                return nil, fmt.Errorf("解析试题数据失败: %v", err)
            }
            paper.Questions = questions
        case []interface{}:
            // Convert []interface{} to []Question
            questionsBytes, err := json.Marshal(questionsVal)
            if err != nil {
                return nil, fmt.Errorf("序列化试题数据失败: %v", err)
            }
            var questions []Question
            if err := json.Unmarshal(questionsBytes, &questions); err != nil {
                return nil, fmt.Errorf("解析试题数据失败: %v", err)
            }
            paper.Questions = questions
        default:
            return nil, fmt.Errorf("未知的 questions 字段类型")
        }
        // 解析其他字段
        if id, ok := rawPaper["id"].(float64); ok {
            paper.ID = int(id)
        }
        if title, ok := rawPaper["title"].(string); ok {
            paper.Title = title
        }
        if description, ok := rawPaper["description"].(string); ok {
            paper.Description = description
        }
    }

    return &paper, nil
}