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

	//search database for user with email
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	//compare hashes
	psw := utils.CheckPasswordHash(password, user.Password)
	if !psw {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.CreateToken(user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is already in use"})
		return
	} else {
		//hash password
		hashedPassword, err := utils.HashPassword(form.Password)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to hash password")
			return
		}
		user.Password = hashedPassword

		//add user to database
		res := config.DB.Create(&user)
		if res.Error != nil {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Signup Failed. User could not be created")
			log.Fatalf("creating user in db failed: %v", res.Error)
			return
		}

		// result := config.DB.Create(&talent)
		// if result.Error != nil {
		// 	utils.ErrorResponse(ctx, http.StatusInternalServerError, "Signup Failed. User could not be created")
		// 	return
		// }

		//generate jwt token
		token, err := utils.CreateToken("talent")
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate token")
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
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
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to hash password")
			return
		}
		user.Password = hashedPassword

		//add user to database
		res := config.DB.Create(&user)
		if res.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Signup Failed. User could not be created"})
			log.Fatalf("creating user in db failed: %v", res.Error)
			return
		}

		// result := config.DB.Create(&company)
		// if result.Error != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Signup Failed. User could not be created"})
		// 	log.Fatalf("creating user in db failed: %v", result.Error)
		// 	return
		// }

		// ltn := config.DB.Create(&location)
		// if ltn.Error != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Signup Failed. User could not be created"})
		// 	log.Fatalf("creating user in db failed: %v", ltn.Error)
		// 	return
		// }

		//generate jwt token
		token, err := utils.CreateToken("company")
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
				"id":    company.Id,
				"email": user.Email,
				"role":  user.Role,
			},
			"token": token,
		})
	}
}
