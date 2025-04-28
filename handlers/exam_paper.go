package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 定义请求体结构
type CreateExamPaperRequest struct {
	TemplateID int64 `json:"template_id" binding:"required"`
}

// CreateExamPaper 用于根据模板ID生成试卷
func CreateExamPaper(c *gin.Context) {
	var req CreateExamPaperRequest

	// 绑定请求JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	// 获取表前缀
	tablePrefix := os.Getenv("TABLE_PREFIX")
	if tablePrefix == "" {
		tablePrefix = "ym_"
	}

	// 查询模板数据
	var template struct {
		Title       string
		Description string
		CoverImage  string
		TotalScore  int
		Questions   string
		CategoryID  int64
		PublishTime int
		Status      int
		Creator     string
	}
	queryTemplate := fmt.Sprintf(`SELECT title, description, cover_image, total_score, questions, category_id, publish_time, status, creator FROM %sexam_template WHERE id = ? LIMIT 1`, tablePrefix)
	err := utils.DB.QueryRow(queryTemplate, req.TemplateID).Scan(
		&template.Title, &template.Description, &template.CoverImage, &template.TotalScore,
		&template.Questions, &template.CategoryID, &template.PublishTime, &template.Status,
		&template.Creator,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "试卷模板不存在",
			"data": nil,
		})
		return
	}

	newUUID := uuid.New().String()

	queryInsert := fmt.Sprintf(`INSERT INTO %sexam_papers (title, description, cover_image, total_score, questions, category_id, publish_time, status, creator, template_id, uuid, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`, tablePrefix)
	result, err := utils.DB.Exec(queryInsert,
		template.Title, template.Description, template.CoverImage,
		template.TotalScore, template.Questions, template.CategoryID,
		template.PublishTime, template.Status, template.Creator, req.TemplateID, newUUID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "生成试卷失败",
			"data": nil,
		})
		return
	}

	// 获取新生成的试卷ID
	examPaperID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "获取试卷ID失败",
			"data": nil,
		})
		return
	}

	// 同步保存到Redis
	examPaperData := map[string]interface{}{
		"id":            examPaperID,
		"uuid":          newUUID,
		"title":         template.Title,
		"description":   template.Description,
		"cover_image":   template.CoverImage,
		"total_score":   template.TotalScore,
		"questions":     template.Questions,
		"category_id":   template.CategoryID,
		"publish_time":  template.PublishTime,
		"status":        template.Status,
		"creator":       template.Creator,
		"template_id":   req.TemplateID,
	}

	jsonData, err := json.Marshal(examPaperData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "试卷生成失败（JSON序列化失败）",
			"data": nil,
		})
		return
	}

	redisKey := fmt.Sprintf("exam_paper:%s", newUUID)
	err = utils.RedisClient.Set(utils.Ctx, redisKey, jsonData, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "试卷生成失败（写入Redis失败）",
			"data": nil,
		})
		return
	}

	// 返回成功
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "试卷生成成功",
		"data": gin.H{
			"id":   examPaperID,
			"uuid": newUUID,
		},
	})
}