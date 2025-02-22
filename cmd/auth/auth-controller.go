package auth

import (
	"log"
	"net/http"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/config"
	"github.com/Creative-genius001/Connekt/types"
	"github.com/Creative-genius001/Connekt/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var form types.TalentForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	talent := models.Talent{
		Id:           uuid.New().String(),
		FirstName:    form.FirstName,
		LastName:     form.LastName,
		Email:        form.Email,
		Password:     form.Password,
		Country:      form.Country,
		About:        form.About,
		State:        form.State,
		Gender:       form.Gender,
		Phone:        form.Phone,
		Experience:   form.Experience,
		CV:           form.CV,
		ProfilePhoto: form.ProfilePhoto,
	}

	//check if user already exist in db
	var count int64
	config.DB.Model(&talent).Where("email = ?", form.Email).Count(&count)
	if count > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is already in use"})
		return
	} else {
		//hash password
		hashedPassword, err := utils.HashPassword(form.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		talent.Password = hashedPassword

		//add user to database
		result := config.DB.Create(&talent)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Signup Failed. User could not be created"})
			log.Fatalf("creating user in db failed: %v", result.Error)
			return
		}

		//generate jwt token
		token, err := utils.CreateToken(talent.Id)
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
				"id":           talent.Id,
				"firstName":    talent.FirstName,
				"lastName":     talent.LastName,
				"email":        talent.Email,
				"country":      talent.Country,
				"about":        talent.About,
				"state":        form.State,
				"gender":       form.Gender,
				"phone":        form.Phone,
				"experience":   form.Experience,
				"cv":           form.CV,
				"profilePhoto": form.ProfilePhoto,
				"token":        token,
			}})
	}
}

// Register employer
func RegisterAsEmployer(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered"})
}
