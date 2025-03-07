package routes

import (
	"github.com/Creative-genius001/Connekt/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRouter := router.Group("/api/user")
	{
		userRouter.GET("/:id", controllers.GetUserData)
	}
}
