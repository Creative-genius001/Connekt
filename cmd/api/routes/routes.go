package routes

import (
	"github.com/Creative-genius001/Connekt/cmd/auth"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	JobRoutes(router)
	UserRoutes(router)
	auth.AuthRoutes(router)
}
