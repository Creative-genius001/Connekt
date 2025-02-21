package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JobRoutes(router *gin.Engine) {
	jobRouter := router.Group("/api/job")
	{
		jobRouter.GET("/listings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Get job listing routes"})
		})
	}
}
