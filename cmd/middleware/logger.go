package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware logs request details
func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		logrus.WithFields(logrus.Fields{
			"status":    ctx.Writer.Status(),
			"method":    ctx.Request.Method,
			"path":      ctx.Request.URL.Path,
			"client_IP": ctx.ClientIP(),
			"duration":  time.Since(startTime),
		}).Info("Request received")
	}
}
