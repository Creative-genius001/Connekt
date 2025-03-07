package controllers

import (
	"net/http"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/config"
	"github.com/Creative-genius001/Connekt/utils"
	"github.com/gin-gonic/gin"
)

func GetUserData(ctx *gin.Context) {
	id := ctx.Param("id")

	var user models.User
	var talent models.Talent
	var company models.Company

	result := config.DB.Select("id, role").Where("id = ?", id).First(&user)
	if result.Error != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "User not found")
		return
	}

	if user.Role == "talent" {
		result = config.DB.Where("user_id = ?", id).First(&talent)
		if result.Error != nil {
			utils.ErrorResponse(ctx, http.StatusNotFound, "User not found")
			return
		}
		ctx.JSON(http.StatusOK, talent)
		return
	}

	if user.Role == "company" {
		result = config.DB.Where("user_id = ?", id).First(&company)
		if result.Error != nil {
			utils.ErrorResponse(ctx, http.StatusNotFound, "User not found")
			return
		}
		ctx.JSON(http.StatusOK, company)
		return
	}

	utils.ErrorResponse(ctx, http.StatusBadRequest, "User not found")

}
