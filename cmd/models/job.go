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
	State         string
	City          string
	Country       string
	IsActive      bool
	Industry      string
	Salary        *Salary          `gorm:"foreignKey:JobId"`
	Applicantions []JobApplication `gorm:"foreignKey:JobId"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Salary struct {
	Id       string `gorm:"type:uuid;primaryKey;unique"`
	MinValue float64
	MaxValue float64
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
