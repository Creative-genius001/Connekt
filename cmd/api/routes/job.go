package routes

import (
	"github.com/Creative-genius001/Connekt/cmd/api/controllers"
	// "github.com/Creative-genius001/Connekt/cmd/middleware"
	"github.com/gin-gonic/gin"
)

func JobRoutes(router *gin.Engine) {
	jobRouter := router.Group("/api/job")
	// jobRouter.Use(middleware.JWTAuthMiddleware())

	{
		jobRouter.GET("/listings", controllers.GetAllJobs)
		jobRouter.GET("/:id", controllers.GetSingleJob)
	}
}
