package utils

import (
	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
	})
}
