package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	Id          string `gorm:"type:uuid;primaryKey;unique"`
	EmployerId  uint
	Title       string
	Summary     string
	Location    string
	CompanyName string
	Type        string //on-site or remote
	IsActive    bool
	Industry    string
	Talents     []JobApplication `gorm:"foreignKey:JobId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type JobApplication struct {
	Id        string `gorm:"type:uuid;primaryKey"`
	TalentId  string `gorm:"type:uuid;not null"`
	JobId     string `gorm:"type:uuid;not null"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
