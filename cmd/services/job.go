package services

import (
	"errors"
	"log"

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
		} else {
			return err
		}
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

func GetJobApplicants(jobId, companyId string) ([]models.JobApplication, error) {
	//query job table using job_id and check if the company_id is the same as that in the job column else return unauthorized
	var count int64

	err := config.DB.Model(&models.Job{}).
		Where("id = ? AND company_id = ?", jobId, companyId).
		Count(&count).Error

	if err != nil {
		return nil, errors.New("error checking if user is authorized")
	}

	if count == 0 {
		return nil, errors.New("unauthorized: company does not own this job")
	}

	var jobApplications []models.JobApplication
	if err := config.DB.Where("job_id = ?", jobId).Find(&jobApplications).Error; err != nil {
		return nil, err
	}

	return jobApplications, nil
}

func GetMyApplications(talentId string) ([]models.Job, error) {
	var jobs []models.Job

	if err := config.DB.
		Joins("JOIN job_applications ON job_applications.job_id = jobs.id").
		Joins("JOIN companies ON companies.id = jobs.company_id").
		Joins("JOIN salaries ON salaries.job_id = jobs.id").
		Where("job_applications.talent_id = ?", talentId).
		Preload("Company").
		Preload("Salary").
		Find(&jobs).Error; err != nil {
		log.Printf("Error getting jobs applied to by user: %s", err)
		return nil, err
	}
	return jobs, nil
}
