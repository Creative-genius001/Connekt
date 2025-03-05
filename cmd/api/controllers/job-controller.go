package controllers

import (
	"net/http"
	// "strconv"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/config"
	"github.com/Creative-genius001/Connekt/utils"
	"github.com/gin-gonic/gin"
)

func GetAllJobs(ctx *gin.Context) {

	id := ctx.Query("id")
	limit := 10

	db := config.DB.Order("id ASC").Limit(limit)
	// page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	// pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	// if page < 1 {
	// 	page = 1
	// }
	// if pageSize < 1 || pageSize > 100 {
	// 	pageSize = 10
	// }

	// offset := (page - 1) * pageSize

	// If a cursor is provided, fetch jobs with ID > cursor
	if id != "" {
		db = db.Where("id > ?", id)
	}

	var jobs []models.Job
	result := db.Find(&jobs)
	//result := config.DB.WithContext(ctx).Limit(pageSize).Offset(offset).Find(&jobs)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jobs": jobs})
}

func GetSingleJob(ctx *gin.Context) {

	id := ctx.Param("id")

	var job models.Job
	result := config.DB.WithContext(ctx).Where("id = ?", id).First(&job)

	if result.Error != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Job not found")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"job": job})

}
