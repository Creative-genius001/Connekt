package models

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Country    string
	State      string
	EmployerID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
