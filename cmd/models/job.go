package models

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	EmployerID  uint
	Title       string
	Summary     string
	Location    string
	CompanyName string
	IsActive    bool
	Industry    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
