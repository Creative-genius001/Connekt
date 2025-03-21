package services

import (
	"errors"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/Creative-genius001/Connekt/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ApplyToJob allows a talent to apply for a job
func ApplyToJob(jobID, talentID, cv, coverLetter string) error {
	var job models.Job

	// Check if the job exists
	if err := config.DB.Where("id = ?", jobID).First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("job not found")
		}
		return err
	}

	// Check if the talent has already applied for this job
	var existingApplication models.JobApplication
	if err := config.DB.Where("job_id = ? AND talent_id = ?", jobID, talentID).First(&existingApplication).Error; err == nil {
		return errors.New("already applied for this job")
	}

	application := models.JobApplication{
		Id:          uuid.New().String(),
		JobId:       jobID,
		TalentId:    talentID,
		Resume:      cv,
		Coverletter: coverLetter,
		Status:      "pending",
	}

	if err := config.DB.Create(&application).Error; err != nil {
		return err
	}

	return nil
}
