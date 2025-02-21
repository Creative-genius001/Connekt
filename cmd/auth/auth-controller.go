package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login function
func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// Register function
func Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered"})
}
