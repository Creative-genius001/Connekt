package services

import (
	"errors"
	"log"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/config"
	"github.com/Creative-genius001/Connekt/types"
	"gorm.io/gorm"
)

func UpdateTalentData(Tdata types.UpdateTalentForm, userID, talentID string) error {
	var user models.User

	// Start transaction
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Fetch user data
	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		log.Printf("Error fetching user: %v", err)
		return errors.New("internal server error")
	}

	updates := make(map[string]interface{})

	if Tdata.FirstName != nil {
		updates["first_name"] = *Tdata.FirstName
	}
	if Tdata.LastName != nil {
		updates["last_name"] = *Tdata.LastName
	}
	if Tdata.Gender != nil {
		updates["gender"] = *Tdata.Gender
	}
	if Tdata.Country != nil {
		updates["country"] = *Tdata.Country
	}
	if Tdata.State != nil {
		updates["state"] = *Tdata.State
	}
	if Tdata.City != nil {
		updates["city"] = *Tdata.City
	}
	if Tdata.About != nil {
		updates["about"] = *Tdata.About
	}
	if Tdata.Phone != nil {
		updates["phone"] = *Tdata.Phone
	}

	if Tdata.Website != nil {
		updates["website"] = Tdata.Website
	}
	if Tdata.Twitter != nil {
		updates["twitter"] = Tdata.Twitter
	}
	if Tdata.Linkedin != nil {
		updates["linkedin"] = Tdata.Linkedin
	}
	if Tdata.Github != nil {
		updates["github"] = Tdata.Github
	}
	if Tdata.ProfilePhoto != nil {
		updates["profile_photo"] = Tdata.ProfilePhoto
	}

	if len(updates) == 0 {
		tx.Rollback()
		return errors.New("no data provided")
	}

	var existingTalent models.Talent
	if err := tx.Where("id = ? AND user_id = ?", talentID, userID).First(&existingTalent).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("company not found or not owned by user")
		}
		log.Printf("Database Error: %v", err)
		return errors.New("database error")
	}

	// Update talent data
	err := tx.Model(&models.Talent{}).Where("user_id = ? AND id = ?", userID, talentID).Updates(updates).Error

	if err != nil {
		tx.Rollback()
		log.Printf("Error updating talent data: %v", err)
		return errors.New("failed to update talent data")
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
		return errors.New("failed to complete update")
	}
	return nil
}

func UpdateCompanyData(Cdata types.UpdateCompanyForm, userID, companyID string) error {
	var user models.User

	// Start transaction
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Fetch user data
	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		log.Printf("Error fetching user: %v", err)
		return errors.New("internal server error")
	}

	updates := make(map[string]interface{})

	if Cdata.CompanyName != nil {
		updates["company_name"] = *Cdata.CompanyName
	}
	if Cdata.Phone != nil {
		updates["phone"] = *Cdata.Phone
	}
	if Cdata.CompanyAddress != nil {
		updates["company_address"] = *Cdata.CompanyAddress
	}
	if Cdata.EmployeeNumber != nil {
		updates["employee_number"] = *Cdata.EmployeeNumber
	}
	if Cdata.Industry != nil {
		updates["industry"] = *Cdata.Industry
	}
	if Cdata.About != nil {
		updates["about"] = *Cdata.About
	}

	if Cdata.CompanyLogo != nil {
		updates["company_logo"] = Cdata.CompanyLogo
	}
	if Cdata.Website != nil {
		updates["website"] = Cdata.Website
	}
	if Cdata.Twitter != nil {
		updates["twitter"] = Cdata.Twitter
	}
	if Cdata.Linkedin != nil {
		updates["linkedin"] = Cdata.Linkedin
	}

	if len(updates) == 0 {
		tx.Rollback()
		return errors.New("no data provided")
	}

	// First verify the company exists and belongs to the user
	var company models.Company
	if err := tx.Where("id = ? AND user_id = ?", companyID, userID).First(&company).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("company not found or not owned by user")
		}
		log.Printf("Database Error: %v", err)
		return errors.New("database error")
	}

	// Update company data
	err := tx.Model(&models.Company{}).Where("id = ?", companyID).Updates(updates).Error

	if err != nil {
		tx.Rollback()
		log.Printf("Error updating talent data: %v", err)
		return errors.New("failed to update talent data")
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
		return errors.New("failed to complete update")
	}
	return nil
}

func UpdateEmail() {
	// Update user email if provided
	// if Tdata.Email != "" && Tdata.Email != &user.Email {
	// 	if err := tx.Model(&user).Update("email", Tdata.Email).Error; err != nil {
	// 		tx.Rollback()
	// 		log.Printf("Error updating user email: %s", err)
	// 		return errors.New("failed to update email")
	// 	}
	// }
}
