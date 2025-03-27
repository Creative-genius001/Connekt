package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/cmd/services"
	"github.com/Creative-genius001/Connekt/config"
	"github.com/Creative-genius001/Connekt/types"
	"github.com/Creative-genius001/Connekt/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserData(ctx *gin.Context) {
	id := ctx.Param("id")

	var user models.User

	// Fetch user with role and preload associated Talent or Company data
	err := config.DB.
		Preload("Talent").
		Preload("Company").
		Where("id = ?", id).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse(ctx, http.StatusNotFound, "User not found")
			return
		}
		log.Printf("Error fetching user: %s", err)
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
		return
	}
	switch user.Role {
	case "talent":
		ctx.JSON(http.StatusOK, gin.H{
			"id":     user.Id,
			"email":  user.Email,
			"role":   user.Role,
			"talent": user.Talent,
		})
	case "company":
		ctx.JSON(http.StatusOK, gin.H{
			"id":      user.Id,
			"email":   user.Email,
			"role":    user.Role,
			"company": user.Company,
		})
	default:
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user role")
	}
}

func UpdateUserData(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusForbidden, "Role not determined")
		return
	}

	ID, exists := ctx.Get("id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid token")
		return
	}

	IDStr, ok := ID.(string)
	if !ok {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Invalid talent ID format")
		return
	}

	userID := ctx.Param("id")

	switch role {
	case "talent":
		var talentForm types.UpdateTalentForm
		if err := ctx.ShouldBindJSON(&talentForm); err != nil {
			utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input data")
			return
		}
		err := services.UpdateTalentData(talentForm, userID, IDStr)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusBadRequest, "update failed try again")
			return
		}
		ctx.JSON(200, gin.H{"message": "Updated successfully"})

	case "company":
		var companyForm types.UpdateCompanyForm
		if err := ctx.ShouldBindJSON(&companyForm); err != nil {
			utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid input data")
			return
		}
		err := services.UpdateCompanyData(companyForm, userID, IDStr)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusBadRequest, "update failed try again")
			return
		}
		ctx.JSON(200, gin.H{"message": "Updated successfully"})

	default:
		utils.ErrorResponse(ctx, http.StatusForbidden, "update failed")
	}
}

// func UpdateEmail(ctx *gin.Context) {
// 	type Form struct {
// 		Email *string `json:"email" binding:"required"`
// 	}
// }
