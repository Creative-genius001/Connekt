package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	Id            string  `gorm:"type:uuid;primaryKey;unique"`
	CompanyId     string  `gorm:"foreignKey:CompanyId"`
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
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"DeletedAt"`
}

type Salary struct {
	MinValue string
	MaxValue string
	Currency string
	JobId    string `gorm:"type:uuid;unique;not null"`
	Job      *Job   `gorm:"constraint:OnDelete:CASCADE;"`
}

type JobApplication struct {
	Id        string         `gorm:"type:uuid;primaryKey"`
	TalentId  string         `gorm:"type:uuid;not null" json:"TalentId"`
	JobId     string         `gorm:"type:uuid;not null" json:"JobId"`
	Job       Job            `gorm:"foreignKey:JobId"`
	CreatedAt time.Time      `json:"CreatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"DeletedAt"`
}
