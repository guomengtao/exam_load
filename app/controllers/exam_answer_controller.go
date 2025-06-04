package controllers

import (
	"gin-go-test/app/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ExamAnswerController handles HTTP requests for exam answers
type ExamAnswerController struct {
	biz biz.ExamAnswerBizInterface
}

// NewExamAnswerController creates a new instance of ExamAnswerController
func NewExamAnswerController(biz biz.ExamAnswerBizInterface) *ExamAnswerController {
	return &ExamAnswerController{
		biz: biz,
	}
}

// SaveAnswer handles the POST request to save an answer
func (c *ExamAnswerController) SaveAnswer(ctx *gin.Context) {
	var data map[string]interface{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.biz.SaveAnswer(ctx, data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Answer saved successfully"})
}

// GetAnswerRecord handles the GET request to retrieve an answer record
func (c *ExamAnswerController) GetAnswerRecord(ctx *gin.Context) {
	recordID := ctx.Param("recordID")
	record, err := c.biz.GetAnswerRecord(ctx, recordID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}

// GetExamPaper handles the GET request to retrieve an exam paper
func (c *ExamAnswerController) GetExamPaper(ctx *gin.Context) {
	examUUID := ctx.Param("examUUID")
	paper, err := c.biz.GetExamPaper(ctx, examUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, paper)
}
