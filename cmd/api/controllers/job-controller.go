package controllers

import (
	"fmt"
	"net/http"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/cmd/services"
	"github.com/Creative-genius001/Connekt/config"
	"github.com/Creative-genius001/Connekt/types"
	"github.com/Creative-genius001/Connekt/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	//"gorm.io/gorm"
)

// get all jobs data
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

// get single job data
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

// create a job posting
func CreateJob(ctx *gin.Context) {

	companyId := ctx.Query("companyId")

	role, exists := ctx.Get("role")
	if !exists || role != "company" {
		utils.ErrorResponse(ctx, http.StatusForbidden, "User Unauthorized!")
		return
	}

	var form types.CreateJobForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data")
		return
	}

	id := uuid.New().String()

	salary := models.Salary{
		Id:       uuid.New().String(),
		JobId:    id,
		MaxValue: form.MaxValue,
		MinValue: form.MinValue,
		Currency: form.Currency,
	}

	job := models.Job{
		Id:          id,
		CompanyId:   companyId,
		State:       form.State,
		Country:     form.Country,
		Title:       form.Title,
		Description: form.Description,
		Remote:      form.Remote,
		City:        form.City,
		IsActive:    true,
		Industry:    form.Industry,
		Salary:      salary,
	}

	if err := config.DB.Create(&job).Error; err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Unable to create job")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Job created successfully"})
}

// UpdateJob updates an existing job
func UpdateJob(ctx *gin.Context) {
	jobId := ctx.Param("id")
	companyId := ctx.Query("companyId")

	_, exists := ctx.Get("role")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var job models.Job
	if err := config.DB.Preload("Salary").Where("id = ?", jobId).First(&job).Error; err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Job not found")
		return
	}

	if job.CompanyId != companyId {
		utils.ErrorResponse(ctx, http.StatusForbidden, "User Unauthorized")
		return
	}

	var updateData types.UpdateJobForm
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data")
		return
	}

	job.Title = updateData.Title
	job.Description = updateData.Description
	job.Remote = updateData.Remote
	job.City = updateData.City
	job.State = updateData.State
	job.Country = updateData.Country
	job.Industry = updateData.Industry
	job.Salary.MaxValue = updateData.MaxValue
	job.Salary.MinValue = updateData.MinValue
	job.Salary.Currency = updateData.Currency

	tx := config.DB.Begin()

	if err := tx.Save(&job).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error updating job")
		return
	}

	// Commit transaction
	tx.Commit()

	// if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&job).Error; err != nil {
	// 	utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error updating job")
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{"message": "Job updated successfully"})
}

func ApplyToJob(ctx *gin.Context) {
	jobID := ctx.Param("jobId")

	fmt.Println("this is the job-id", jobID)

	talentId, exists := ctx.Get("id")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}
	talentID, ok := talentId.(string)
	if !ok {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Invalid talent ID format")
		return
	}

	var applicationData types.JobApplicationForm

	if err := ctx.ShouldBindJSON(&applicationData); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data")
		return
	}

	if err := services.ApplyToJob(jobID, talentID, applicationData.Resume, applicationData.Coverletter); err != nil {
		if err.Error() == "job not found" {
			utils.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		} else if err.Error() == "already applied for this job" {
			utils.ErrorResponse(ctx, http.StatusConflict, err.Error())
		} else {
			utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error applying to job")
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Application submitted successfully"})
}
