package utils

import (
	"github.com/gin-gonic/gin"
)

func NotFoundErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
	})
	ctx.Abort()
}
