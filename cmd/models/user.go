package models

import (
	"time"
)

type User struct {
	Id           string `gorm:"type:uuid;;primaryKey;unique"`
	UserId       string
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	FirstName    string
	LastName     string
	ProfilePhoto *string
	Role         string `gorm:"type:varchar(20);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

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
	Id                 string `gorm:"type:uuid;primaryKey;unique"`
	FirstName          string
	LastName           string
	Email              string `gorm:"unique"`
	Gender             string
	Location           Location `gorm:"foreignKey:EmployerId"` // One-to-One Relationship
	Jobs               []Job    `gorm:"foreignKey:EmployerId"` // One Employer → Many Jobs
	Phone              string
	CompanyName        string
	CompanyAddress     string
	RegistrationNumber string
	EmployeeNumber     uint64
	Industry           string
	About              string
	Password           string
	ProfilePhoto       *string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
