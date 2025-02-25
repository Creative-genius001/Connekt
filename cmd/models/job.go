package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	Id          string           `gorm:"type:uuid;primaryKey;unique"`
	EmployerId  string           `json:"EmployerId"`
	Title       string           `json:"Title"`
	Summary     string           `json:"Summary"`
	Location    string           `json:"Location"`
	CompanyName string           `json:"CompanyName"`
	Type        string           `json:"Type"` //on-site or remote
	IsActive    bool             `json:"IsActive"`
	Industry    string           `json:"Industry"`
	Talents     []JobApplication `gorm:"foreignKey:JobId" json:"Talents"`
	CreatedAt   time.Time        `json:"CreatedAt"`
	UpdatedAt   time.Time        `json:"UpdatedAt"`
	DeletedAt   gorm.DeletedAt   `gorm:"index" json:"DeletedAt"`
}

type JobApplication struct {
	Id        string         `gorm:"type:uuid;primaryKey"`
	TalentId  string         `gorm:"type:uuid;not null" json:"TalentId"`
	JobId     string         `gorm:"type:uuid;not null" json:"JobId"`
	CreatedAt time.Time      `json:"CreatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"DeletedAt"`
}
