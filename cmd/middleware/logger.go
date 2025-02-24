package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs request details
func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		log.Printf("[%s] %s %s %d %s",
			ctx.Request.Method,
			ctx.Request.RequestURI,
			ctx.ClientIP(),
			ctx.Writer.Status(),
			time.Since(startTime),
		)
	}
}
