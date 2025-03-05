package routes

import (
	"github.com/Creative-genius001/Connekt/cmd/api/controllers"
	"github.com/Creative-genius001/Connekt/cmd/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("/api/user")
	userRouter.Use(middleware.JWTAuthMiddleware())
	{
		userRouter.GET("/:id", controllers.GetUserData)
	}
}
