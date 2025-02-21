package routes

import (
	"github.com/Creative-genius001/Connekt/cmd/api/routes"
	"github.com/Creative-genius001/Connekt/cmd/auth"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	routes.JobRoutes(router)
	auth.AuthRoutes(router)
}
