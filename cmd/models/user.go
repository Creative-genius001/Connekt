package models

import (
	"time"
)

type User struct {
	Id        string `gorm:"type:uuid;;primaryKey;unique"`
	UserId    string
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
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
	Role         string `gorm:"type:varchar(20);not null"`
	Phone        string
	Password     string
	Website      *string
	Twitter      *string
	LinkedIn     *string
	Facebook     *string
	Github       *string
	AppliedJobs  []JobApplication `gorm:"foreignKey:TalentId"`
	ProfilePhoto *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Company struct {
	Id             string `gorm:"type:uuid;primaryKey;unique"`
	CompanyName    string
	Email          string   `gorm:"unique"`
	Location       Location `gorm:"foreignKey:EmployerId"`
	Jobs           []Job    `gorm:"foreignKey:EmployerId"`
	Phone          string
	Role           string `gorm:"type:varchar(20);not null"`
	CompanyAddress string
	EmployeeNumber uint64
	Industry       string
	About          string
	Password       string
	CompanyLogo    *string
	Website        *string
	Twitter        *string
	LinkedIn       *string
	Facebook       *string
	Github         *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
