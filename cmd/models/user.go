package models

import (
	"time"
)

type JobSeekerModel struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Email      string `gorm:"unique"`
	Gender     string
	Country    string
	State      string
	Phone      string
	Password   string
	Experience uint8
	CV         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type EmployerModel struct {
	ID                 uint `gorm:"primaryKey"`
	Name               string
	Email              string `gorm:"unique"`
	Gender             string
	Country            string
	State              string
	Phone              string
	CompanyName        string
	CompanyAddress     string
	RegistrationNumber string
	EmployeeNumber     uint64
	Industry           string
	EmployerType       string
	About              string
	Password           string
	Experience         uint8
	CV                 string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
