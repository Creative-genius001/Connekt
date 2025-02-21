package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"golang.org/x/crypto/bcrypt"
	"github.com/Creative-genius001/Connekt/types"
)

// Login function
func Login(ctx *gin.Context) {
	var form types.LoginForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid input"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successfullyy", "form": form})
}

// Register function
func Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered"})
}
