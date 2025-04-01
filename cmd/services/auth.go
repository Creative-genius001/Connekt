package services

import (
	"errors"
	"log"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/config"
	"github.com/Creative-genius001/Connekt/types"
	"github.com/Creative-genius001/Connekt/utils"
	"github.com/google/uuid"
)

func LoginService(email, password string) (map[string]interface{}, string, error) {

	var user models.User
	var talent models.Talent
	var company models.Company

	// Search for user in database
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Compare password hashes
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, "", errors.New("invalid credentials")
	}

	// Retrieve user-specific ID (either talent or company)
	var id string
	var additionalData interface{}

	switch user.Role {
	case "talent":
		if err := config.DB.Where("user_id = ?", user.Id).First(&talent).Error; err != nil {
			log.Printf("Error fetching talent ID: %s", err)
			return nil, "", errors.New("unable to get user data")
		}
		id = talent.Id
		additionalData = talent

	case "company":
		if err := config.DB.Where("user_id = ?", user.Id).First(&company).Error; err != nil {
			log.Printf("Error fetching company ID: %s", err)
			return nil, "", errors.New("unable to get user data")
		}
		id = company.Id
		additionalData = company
	default:
		return nil, "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.CreateToken(user.Role, id)
	if err != nil {
		return nil, "", errors.New("server error")
	}

	// Prepare user data response
	userData := map[string]interface{}{
		"id":    user.Id,
		"email": user.Email,
		"role":  user.Role,
		"data":  additionalData,
	}

	return userData, token, nil
}

func RegisterAsTalentService(form types.TalentForm) error {

	id := uuid.New().String()
	talent := models.Talent{
		Id:           id,
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
		return errors.New("email already exists")
	} else {
		//hash password
		hashedPassword, err := utils.HashPassword(form.Password)
		if err != nil {
			return errors.New("sever error")
		}
		user.Password = hashedPassword

		//add user to database
		res := config.DB.Create(&user)
		if res.Error != nil {
			log.Fatalf("creating user in db failed: %v", res.Error)
			return errors.New("signup failed! server error")
		}

	}
	return nil
}

func RegisterAsCompanyService(form types.CompanyForm) error {

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
		return errors.New("email already exists")
	} else {
		//hash password
		hashedPassword, err := utils.HashPassword(form.Password)
		if err != nil {
			return errors.New("sever error")
		}
		user.Password = hashedPassword

		//add user to database
		res := config.DB.Create(&user)
		if res.Error != nil {
			log.Fatalf("creating user in db failed: %v", res.Error)
			return errors.New("signup failed! server error")
		}

	}
	return nil
}
