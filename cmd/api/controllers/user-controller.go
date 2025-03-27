package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/config"
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
	//userId := ctx.Param("id")
}
