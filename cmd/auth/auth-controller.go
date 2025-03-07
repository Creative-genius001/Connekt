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
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	email := form.Email
	password := form.Password

	//search database for user with email
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid Credentials")
		return
	}

	//compare hashes
	psw := utils.CheckPasswordHash(password, user.Password)
	if !psw {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid Credentials")
		return
	}

	token, err := utils.CreateToken(user.Role)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Server Error")
		return
	}

	//return success and user details
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":    user.Id,
			"email": user.Email,
			"role":  user.Role,
		},
		"token": token,
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

	talent := models.Talent{
		Id:           uuid.New().String(),
		FirstName:    form.FirstName,
		LastName:     form.LastName,
		Country:      form.Country,
		About:        form.About,
		State:        form.State,
		City:         form.City,
		Gender:       form.Gender,
		Phone:        form.Phone,
		ProfilePhoto: form.ProfilePhoto,
		Website:      form.Website,
		Twitter:      form.Twitter,
		Linkedin:     form.Linkedin,
		Github:       form.Github,
	}

	user := models.User{
		Id:       uuid.New().String(),
		Talent:   talent,
		Email:    form.Email,
		Password: form.Password,
		Role:     "talent",
	}

	//check if user already exist in db
	var count int64
	config.DB.Model(&user).Where("email = ?", form.Email).Count(&count)
	if count > 0 {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Email already exists")
		return
	} else {
		//hash password
		hashedPassword, err := utils.HashPassword(form.Password)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Server Error")
			return
		}
		user.Password = hashedPassword

		//generate jwt token
		token, err := utils.CreateToken("talent")
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Signup Failed! Server Error")
			return
		}

		//add user to database
		res := config.DB.Create(&user)
		if res.Error != nil {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Signup Failed! Server Error")
			log.Fatalf("creating user in db failed: %v", res.Error)
			return
		}

		//send user email for successful registration

		//send token and user data back
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "User registered successfully",
			"user": gin.H{
				"id":           user.Id,
				"userId":       talent.Id,
				"firstName":    talent.FirstName,
				"lastName":     talent.LastName,
				"email":        user.Email,
				"profilePhoto": form.ProfilePhoto,
				"role":         user.Role,
			},
			"token": token,
		})
	}
}

// Register company
func RegisterAsCompany(ctx *gin.Context) {

	var form types.CompanyForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	id := uuid.New().String()

	company := models.Company{
		Id:             id,
		About:          form.About,
		Phone:          form.Phone,
		CompanyLogo:    form.CompanyLogo,
		CompanyName:    form.CompanyName,
		CompanyAddress: form.CompanyAddress,
		EmployeeNumber: form.EmployeeNumber,
		Industry:       form.Industry,
		Website:        form.Website,
		Twitter:        form.Twitter,
		Linkedin:       form.Linkedin,
		Location: models.Location{
			Id:        uuid.New().String(),
			Country:   form.Country,
			State:     form.State,
			City:      form.City,
			CompanyId: id,
		},
	}

	user := models.User{
		Id:       uuid.New().String(),
		Email:    form.Email,
		Password: form.Password,
		Role:     "company",
		Company:  company,
	}

	//check if user already exist in db
	var count int64
	config.DB.Model(&user).Where("email = ?", form.Email).Count(&count)
	if count > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is already in use"})
		return
	} else {
		//hash password
		hashedPassword, err := utils.HashPassword(form.Password)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Server Error")
			return
		}
		user.Password = hashedPassword

		//generate jwt token
		token, err := utils.CreateToken("company")
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Server Error")
			return
		}

		//add user to database
		res := config.DB.Create(&user)
		if res.Error != nil {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Signup Failed! Server Error")
			return
		}

		//send user email for successful registration

		//send token and user data back
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "User registered successfully",
			"user": gin.H{
				"id":     user.Id,
				"userId": company.Id,
				"email":  user.Email,
				"role":   user.Role,
			},
			"token": token,
		})
	}
}
