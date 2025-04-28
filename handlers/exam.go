package handlers

import (
    "context"
    "net/http"
    "strconv"  // 加上它！
    "github.com/gin-gonic/gin"
    "gin-go-test/utils" // 根据你的模块路径调整
)

func SubmitAnswers(c *gin.Context) {
    examID := c.Param("id")

    var input struct {
        StudentID int    `json:"student_id"`
        Answers   string `json:"answers"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        return
    }

    // 正确转换 student_id
    answerKey := "answers:" + examID + ":" + strconv.Itoa(input.StudentID)

    err := utils.RedisClient.Set(context.Background(), answerKey, input.Answers, 0).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save answer"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Answer saved successfully"})
}

 

// 读取学生答案 (GET /api/exam/:id?student_id=xxx)
func GetExam(c *gin.Context) {
    examID := c.Param("id")
    studentID := c.Query("student_id")

    answerKey := "answers:" + examID + ":" + studentID

    answer, err := utils.RedisClient.Get(context.Background(), answerKey).Result()
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Answer not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "exam_id":    examID,
        "student_id": studentID,
        "answers":    answer,
    })
}