package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-go-test/biz"
	"gin-go-test/app/services"
)

// AnswerController handles exam answer related HTTP requests
type AnswerController struct {
	answerBiz *biz.AnswerBiz
}

// NewAnswerController creates a new instance of AnswerController
func NewAnswerController() *AnswerController {
	return &AnswerController{
		answerBiz: biz.NewAnswerBiz(services.NewAnswerService()),
	}
}

// SubmitAnswer handles the submission of exam answers
// @Summary 提交用户的答题记录
// @Description 用户完成答题后提交记录，并保存到 Redis 和数据库
// @Tags exam_answer
// @Accept json
// @Produce json
// @Param request body biz.AnswerRequest true "答题请求"
// @Success 200 {object} biz.AnswerResponse "返回答题记录"
// @Failure 400 {object} biz.ErrorResponse "请求参数错误"
// @Failure 500 {object} biz.ErrorResponse "服务器错误"
// @Router /api/exam/answer [post]
func (c *AnswerController) SubmitAnswer(ctx *gin.Context) {
	var req biz.AnswerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.sendErrorResponse(ctx, http.StatusBadRequest, "请求参数错误", err)
		return
	}

	response, err := c.answerBiz.SubmitAnswer(ctx, &req)
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "提交答题失败", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "提交成功",
		"data":    response,
	})
}

// GetAnswerResult retrieves the exam answer result
// @Summary 获取用户的答题记录
// @Description 通过答题记录ID获取用户的答题结果
// @Tags exam_answer
// @Accept json
// @Produce json
// @Param record_id path string true "答题记录ID"
// @Success 200 {object} biz.AnswerResponse "返回用户答题记录"
// @Failure 400 {object} biz.ErrorResponse "请求参数错误"
// @Failure 404 {object} biz.ErrorResponse "未找到答题记录"
// @Failure 500 {object} biz.ErrorResponse "服务器错误"
// @Router /api/user/answer/{record_id} [get]
func (c *AnswerController) GetAnswerResult(ctx *gin.Context) {
	recordID := ctx.Param("record_id")
	if recordID == "" {
		c.sendErrorResponse(ctx, http.StatusBadRequest, "记录ID不能为空", nil)
		return
	}

	response, err := c.answerBiz.GetAnswerResult(ctx, recordID)
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "获取答题记录失败", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    response,
	})
}

// GetFullAnswerResult retrieves the complete exam answer result with questions
// @Summary 获取完整的答题记录（包含题目信息）
// @Description 通过答题记录ID获取完整的答题结果，包括题目详情
// @Tags exam_answer
// @Accept json
// @Produce json
// @Param record_id path string true "答题记录ID"
// @Success 200 {object} biz.FullAnswerResponse "返回完整答题记录"
// @Failure 400 {object} biz.ErrorResponse "请求参数错误"
// @Failure 404 {object} biz.ErrorResponse "未找到答题记录"
// @Failure 500 {object} biz.ErrorResponse "服务器错误"
// @Router /api/user/answer/{record_id}/full [get]
func (c *AnswerController) GetFullAnswerResult(ctx *gin.Context) {
	recordID := ctx.Param("record_id")
	if recordID == "" {
		c.sendErrorResponse(ctx, http.StatusBadRequest, "记录ID不能为空", nil)
		return
	}

	response, err := c.answerBiz.GetFullAnswerResult(ctx, recordID)
	if err != nil {
		c.sendErrorResponse(ctx, http.StatusInternalServerError, "获取完整答题记录失败", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    response,
	})
}

// sendErrorResponse sends an error response to the client
func (c *AnswerController) sendErrorResponse(ctx *gin.Context, code int, message string, err error) {
	response := biz.ErrorResponse{
		Code:    code,
		Message: message,
	}
	if err != nil {
		response.Error = err.Error()
	}
	ctx.JSON(code, response)
} 