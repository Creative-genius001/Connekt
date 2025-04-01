package auth

import (
	"net/http"

	"github.com/Creative-genius001/Connekt/cmd/services"
	"github.com/Creative-genius001/Connekt/types"
	"github.com/Creative-genius001/Connekt/utils"
	"github.com/gin-gonic/gin"
)

// Login function
func Login(ctx *gin.Context) {
	var form types.LoginForm

	// Bind the request body to a struct
	if err := ctx.ShouldBind(&form); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	if !utils.IsValidEmail(form.Email) {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	// Call the login service
	userData, token, err := services.LoginService(form.Email, form.Password)
	if err != nil {
		if err.Error() == "invalid credentials" {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		} else {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// Return response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    userData,
		"token":   token,
	})
}

// Register as job seeker
func RegisterAsTalent(ctx *gin.Context) {

	//get signup details from body
	var form types.TalentForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	err := services.RegisterAsTalentService(form)
	if err != nil {
		if err.Error() == "email already exists" {
			utils.ErrorResponse(ctx, http.StatusConflict, err.Error())
			return
		} else {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
	//send user email for successful registration

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Register company
func RegisterAsCompany(ctx *gin.Context) {

	var form types.CompanyForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	err := services.RegisterAsCompanyService(form)
	if err != nil {
		if err.Error() == "email already exists" {
			utils.ErrorResponse(ctx, http.StatusConflict, err.Error())
			return
		} else {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
	//send user email for successful registration

	//send success message
	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
