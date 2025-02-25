package controllers

import (
	"net/http"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/config"
	"github.com/gin-gonic/gin"
)

func GetJobListing(ctx *gin.Context) {

	var jobs []models.Job

	result := config.DB.Find(&jobs)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jobs": jobs})
}
