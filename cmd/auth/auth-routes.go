package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/login", Login)
		authGroup.POST("/register/talent", RegisterAsJobSeeker)
		authGroup.POST("/register/company", RegisterAsCompany)
	}
}
