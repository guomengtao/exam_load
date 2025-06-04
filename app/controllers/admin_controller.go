package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminController struct{}

func (a *AdminController) UpdatePassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Password updated (mock)"})
}
