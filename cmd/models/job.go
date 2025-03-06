package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	Id            string `gorm:"type:uuid;primaryKey;unique"`
	CompanyId     string
	Company       Company `gorm:"foreignKey:CompanyId"`
	Title         string
	Description   string
	Remote        bool
	IsActive      bool
	Industry      string
	Salary        *Salary          `gorm:"foreignKey:JobId"`
	Applicantions []JobApplication `gorm:"foreignKey:JobId"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Salary struct {
	MinValue string
	MaxValue string
	Currency string
	JobId    string `gorm:"type:uuid;unique;not null"`
	Job      *Job   `gorm:"constraint:OnDelete:CASCADE;"`
}

type JobApplication struct {
	Id        string `gorm:"type:uuid;primaryKey"`
	TalentId  string `gorm:"type:uuid;not null"`
	JobId     string `gorm:"type:uuid;not null"`
	Job       Job    `gorm:"foreignKey:JobId"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
