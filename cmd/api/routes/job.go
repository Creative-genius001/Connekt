package routes

import (
	"github.com/Creative-genius001/Connekt/cmd/api/controllers"
	"github.com/Creative-genius001/Connekt/cmd/middleware"
	"github.com/gin-gonic/gin"
)

func JobRoutes(router *gin.Engine) {
	jobRouter := router.Group("/api/job")

	{
		jobRouter.GET("/listings", controllers.GetAllJobs)
		jobRouter.GET("/:id", controllers.GetSingleJob)
		jobRouter.POST("/create", middleware.JWTAuthMiddleware(), controllers.CreateJob)
		jobRouter.PUT("/update/:id", middleware.JWTAuthMiddleware(), controllers.UpdateJob)
		jobRouter.POST("/apply/:jobId", middleware.JWTAuthMiddleware(), controllers.ApplyToJob)
		jobRouter.GET("/applicants/:jobId", middleware.JWTAuthMiddleware(), controllers.GetJobApplicants)
		jobRouter.GET("/me", middleware.JWTAuthMiddleware(), controllers.GetMyApplications)
	}
}
