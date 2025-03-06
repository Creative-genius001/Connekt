package models

import (
	"time"
)

type Location struct {
	Id        string `gorm:"type:uuid;primaryKey;unique"`
	Country   string
	State     string
	City      string
	CompanyId string
	CreatedAt time.Time
	UpdatedAt time.Time
}
