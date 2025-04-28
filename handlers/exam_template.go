package handlers

import (
	"encoding/json"
	"gin-go-test/utils"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

 

 // 创建用的结构体
type ExamRequest struct {
	Title       string          `json:"title" binding:"required"`
	Description string          `json:"description"`
	CoverImage  string          `json:"cover_image"`
	TotalScore  int             `json:"total_score" binding:"required"`
	Questions   json.RawMessage `json:"questions" binding:"required"`
	CategoryID  int64           `json:"category_id"`
	PublishTime int             `json:"publish_time"`
	Status      int             `json:"status"`
	Creator     string          `json:"creator"`
}

// 更新用的结构体
type ExamUpdateRequest struct {
	ID          int64           `json:"id" binding:"required"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	CoverImage  string          `json:"cover_image"`
	TotalScore  int             `json:"total_score"`
	Questions   json.RawMessage `json:"questions"`
	CategoryID  int64           `json:"category_id"`
	PublishTime int             `json:"publish_time"`
	Status      int             `json:"status"`
	Creator     string          `json:"creator"`
}

// 查询时用的结构体
type ExamTemplate struct {
	ID          int64           `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	CoverImage  string          `json:"cover_image"`
	TotalScore  int             `json:"total_score"`
	Questions   json.RawMessage `json:"questions"`
	CategoryID  int64           `json:"category_id"`
	PublishTime int             `json:"publish_time"`
	Status      int             `json:"status"`
	Creator     string          `json:"creator"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}

// 新增试卷模板接口
func CreateExam(c *gin.Context) {
	var examReq ExamRequest

	// 绑定 JSON 请求到结构体
	if err := c.ShouldBindJSON(&examReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	// 读取表前缀
	tablePrefix := os.Getenv("TABLE_PREFIX")
	if tablePrefix == "" {
		tablePrefix = "ym_"
	}

	// 生成插入 SQL
	query := `
		INSERT INTO ` + tablePrefix + `exam_template (title, description, cover_image, total_score, questions, category_id, publish_time, status, creator, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
	`

	// 执行插入
	_, err := utils.DB.Exec(query, examReq.Title, examReq.Description, examReq.CoverImage, examReq.TotalScore, examReq.Questions, examReq.CategoryID, examReq.PublishTime, examReq.Status, examReq.Creator)
	if err != nil {
		log.Println("Error inserting exam:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to create exam",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Exam created successfully",
		"data": nil,
	})
}

// 查询试卷模板接口（支持单条/分页）
func GetExamTemplate(c *gin.Context) {
	// 读取表前缀
	tablePrefix := os.Getenv("TABLE_PREFIX")
	if tablePrefix == "" {
		tablePrefix = "ym_"
	}

	// 获取参数
	idParam := c.Query("id")
	pageParam := c.DefaultQuery("page", "1")
	limitParam := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	db := utils.DB // 你之前的utils.DB连接

	if idParam != "" {
		// 查询单条
		query := `
			SELECT id, title, description, cover_image, total_score, questions, category_id, publish_time, status, creator, created_at, updated_at
			FROM ` + tablePrefix + `exam_template
			WHERE id = ?
			LIMIT 1
		`
		var template ExamTemplate
		err := db.QueryRow(query, idParam).Scan(
			&template.ID, &template.Title, &template.Description, &template.CoverImage, &template.TotalScore,
			&template.Questions, &template.CategoryID, &template.PublishTime, &template.Status,
			&template.Creator, &template.CreatedAt, &template.UpdatedAt,
		)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "模板不存在",
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": template,
		})
	} else {
		// 查询分页
		countQuery := `SELECT COUNT(*) FROM ` + tablePrefix + `exam_template`
		var total int
		err := db.QueryRow(countQuery).Scan(&total)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "查询总数失败",
				"data": nil,
			})
			return
		}

		listQuery := `
			SELECT id, title, description, cover_image, total_score, questions, category_id, publish_time, status, creator, created_at, updated_at
			FROM ` + tablePrefix + `exam_template
			LIMIT ? OFFSET ?
		`
		rows, err := db.Query(listQuery, limit, offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "查询失败",
				"data": nil,
			})
			return
		}
		defer rows.Close()

		var templates []ExamTemplate
		for rows.Next() {
			var template ExamTemplate
			err := rows.Scan(
				&template.ID, &template.Title, &template.Description, &template.CoverImage, &template.TotalScore,
				&template.Questions, &template.CategoryID, &template.PublishTime, &template.Status,
				&template.Creator, &template.CreatedAt, &template.UpdatedAt,
			)
			if err == nil {
				templates = append(templates, template)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"list": templates,
				"pagination": gin.H{
					"page":  page,
					"limit": limit,
					"total": total,
				},
			},
		})
	}
}

 

func UpdateExamTemplate(c *gin.Context) {
	var updateReq ExamUpdateRequest

	// 绑定 JSON
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	// 读取表前缀
	tablePrefix := os.Getenv("TABLE_PREFIX")
	if tablePrefix == "" {
		tablePrefix = "ym_"
	}

	// 执行 UPDATE SQL
	query := `
		UPDATE ` + tablePrefix + `exam_template
		SET title = ?, description = ?, cover_image = ?, total_score = ?, questions = ?, 
			category_id = ?, publish_time = ?, status = ?, creator = ?, updated_at = NOW()
		WHERE id = ?
	`

	db := utils.DB

	_, err := db.Exec(query, updateReq.Title, updateReq.Description, updateReq.CoverImage,
		updateReq.TotalScore, updateReq.Questions, updateReq.CategoryID,
		updateReq.PublishTime, updateReq.Status, updateReq.Creator, updateReq.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to update exam template",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Exam template updated successfully",
		"data": nil,
	})
}