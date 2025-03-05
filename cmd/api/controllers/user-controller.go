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
	role := ctx.Query("role")

	var talent models.Talent
	var employer models.Employer

	switch role {
	case "talent":
		result := config.DB.Where("id = ?", id).First(&talent)
		if result.Error != nil {
			utils.ErrorResponse(ctx, http.StatusNotFound, "User not found")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"talent": talent})
	case "employer":
		result := config.DB.Where("id = ?", id).First(&employer)
		if result.Error != nil {
			utils.ErrorResponse(ctx, http.StatusNotFound, "User not found")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"employer": employer})
	default:
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
	}

}
