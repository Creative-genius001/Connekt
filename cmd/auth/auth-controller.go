package auth

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/config"
	"github.com/Creative-genius001/Connekt/types"
	"github.com/Creative-genius001/Connekt/utils"
	"github.com/gin-gonic/gin"
)

// Login function
func Login(ctx *gin.Context) {
	var form types.LoginForm

	//Bind the request body to a type
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	email := form.Email
	password := form.Password

	user := types.LoginForm{
		Email:    email,
		Password: password,
	}
	config.DB.Create(&user)
	//search database for user with email
	//hash the password
	//compare hashes
	//return success and user details
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successfullyy", "form": form})
}

// Register as job seeker
func RegisterAsJobSeeker(ctx *gin.Context) {

	//get signup details from body
	var form types.JobSeekerForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	//hash password
	hashedPassword, err := utils.HashPassword(form.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	//add user to database
	jobseeker := models.JobSeeker{
		FirstName:    form.FirstName,
		LastName:     form.LastName,
		Email:        form.Email,
		Password:     hashedPassword,
		Country:      form.Country,
		About:        form.About,
		State:        form.State,
		Gender:       form.Gender,
		Phone:        form.Phone,
		Experience:   form.Experience,
		CV:           form.CV,
		ProfilePhoto: form.ProfilePhoto,
	}

	result := config.DB.Create(&jobseeker)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Signup Failed. User could not be created"})
		log.Fatalf("creating user in db failed: %v", result.Error)
		return
	}

	//generate jwt token
	//jwt secretKey
	userID := strconv.FormatUint(uint64(jobseeker.ID), 10)
	token, err := utils.CreateToken(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		log.Fatalf("error creating jwt: %v", err)
		return
	}

	//send user email for successful registration

	//send token and user data back

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":           userID,
			"firstName":    jobseeker.FirstName,
			"lastName":     jobseeker.LastName,
			"email":        jobseeker.Email,
			"country":      jobseeker.Country,
			"about":        jobseeker.About,
			"state":        form.State,
			"gender":       form.Gender,
			"phone":        form.Phone,
			"experience":   form.Experience,
			"cv":           form.CV,
			"profilePhoto": form.ProfilePhoto,
			"token":        token,
		}})
}

// Register employer
func RegisterAsEmployer(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered"})
}
