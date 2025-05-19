package services

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "time"
 
 
    "github.com/google/uuid"
    "github.com/redis/go-redis/v9"
    "gin-go-test/utils"
)

// RedisWriterInterval 表示写入间隔（单位：毫秒）
var RedisWriterInterval = 1000

// StartRedisWriter 启动一个 goroutine 持续写入模拟数据到 Redis
func StartRedisWriter() {
    ticker := time.NewTicker(time.Duration(RedisWriterInterval) * time.Millisecond)
    go func() {
        for range ticker.C {
            writeOneRecord()
        }
    }()
    fmt.Println("✅ Redis 写入模拟器已启动，每", RedisWriterInterval, "ms 写入一次")
}

func writeOneRecord() {
 
    now := time.Now().Unix()

    // 从 Redis 中随机读取一个用户名和用户ID
    poolMember, err := utils.RedisClient.SRandMember(utils.Ctx, "mock:user_pool").Result()
    if err != nil {
        fmt.Println("❌ 获取模拟用户失败:", err)
        return
    }

    var user struct {
        UserID   string `json:"user_id"`
        Username string `json:"username"`
    }
    if err := json.Unmarshal([]byte(poolMember), &user); err != nil {
        fmt.Println("❌ 模拟用户解析失败:", err)
        return
    }

    userID := user.UserID
    username := user.Username

    // 随机选择多选题答案组合 [0,1,2,3] 的非空子集
    options := []int{0, 1, 2, 3}
    rand.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
    selected := options[:rand.Intn(len(options))+1]

    // 随机单选题答案 [0,1,2,3]
    singleChoice := []int{0, 1, 2, 3}
    singleAnswer := singleChoice[rand.Intn(len(singleChoice))]

    answersMap := map[string]interface{}{
        "1746969239826": map[string]interface{}{
            "answer": selected,
            "score":  5,
        },
        "1747002892004": map[string]interface{}{
            "answer": singleAnswer,
            "score":  7,
        },
    }
    answersBytes, err := json.Marshal(answersMap)
    if err != nil {
        fmt.Println("❌ JSON 序列化失败:", err)
        return
    }

    answers := string(answersBytes)
    answerUID := uuid.New().String()

    data := map[string]interface{}{
        "answer_uid":  answerUID,
        "user_uuid":   uuid.New().String(),
        "user_id":     userID,
        "username":    username,
        "exam_id":     "15",
        "exam_uuid":   "40a89290-0ff9-4c01-a803-57155a24985c",
        "score":       "12",
        "total_score": "12",
        "duration":    "120",
        "created_at":  fmt.Sprintf("%d", now),
        "answers":     answers,
    }

    key := "exam_answer:" + answerUID

    // 写入 Hash（详情）
    err = utils.RedisClient.HSet(utils.Ctx, key, data).Err()
    if err != nil {
        fmt.Println("❌ Hash 写入失败:", err)
        return
    }

    // 写入 ZSet（待处理集合）
    err = utils.RedisClient.ZAdd(utils.Ctx, "exam:submitted", redis.Z{
        Score:  float64(now),
        Member: answerUID,
    }).Err()
    if err != nil {
        fmt.Println("❌ ZSet 写入失败:", err)
        return
    }

    fmt.Println("✅ Redis 写入成功：", answerUID)
}

func SimulateBurst(n int) {
    fmt.Printf("🚀 开始高并发模拟提交 %d 份试卷\n", n)
    for i := 0; i < n; i++ {
        go writeOneRecord()
   
        if i%100 == 0 {
            time.Sleep(10 * time.Millisecond)
        }
		if i%1000 == 0 {
			fmt.Printf("✅ 已启动第 %d 条写入\n", i)
		}
    }
 
}