package models

import (
	"time"

	"gorm.io/gorm"
)

type Talent struct {
	Id           string `gorm:"type:uuid;primaryKey;unique"`
	FirstName    string
	LastName     string
	Email        string `gorm:"unique"`
	Gender       string
	Country      string
	State        string
	About        string
	Phone        string
	Password     string
	Experience   uint8
	CV           *string
	Jobs         []JobApplication `gorm:"foreignKey:TalentId"`
	ProfilePhoto *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Employer struct {
	gorm.Model
	FirstName          string
	LastName           string
	Email              string `gorm:"unique"`
	Gender             string
	Location           Location `gorm:"foreignKey:EmployerID"` // One-to-One Relationship
	Jobs               []Job    `gorm:"foreignKey:EmployerID"` // One Employer → Many Jobs
	Phone              string
	CompanyName        string
	CompanyAddress     string
	RegistrationNumber string
	EmployeeNumber     uint64
	Industry           string
	EmployerType       string
	About              string
	Password           string
	ProfilePhoto       string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
