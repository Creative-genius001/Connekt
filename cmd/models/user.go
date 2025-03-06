package models

import (
	"time"
)

type User struct {
	Id       string `gorm:"type:uuid;;primaryKey;unique"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:varchar(20);not null"`

	Talent    Talent  `gorm:"foreignKey:UserId"`
	Company   Company `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Talent struct {
	Id           string `gorm:"type:uuid;primaryKey;unique"`
	UserId       string `gorm:"type:uuid;not null"`
	FirstName    string
	LastName     string
	Gender       string
	Country      string
	State        string
	City         string
	About        string
	Phone        string
	Website      *string
	Twitter      *string
	Linkedin     *string
	Github       *string
	AppliedJobs  []JobApplication `gorm:"foreignKey:TalentId"`
	ProfilePhoto *string
}

type Company struct {
	Id             string `gorm:"type:uuid;primaryKey;unique"`
	UserId         string `gorm:"type:uuid;not null"`
	CompanyName    string
	Location       Location `gorm:"foreignKey:CompanyId"`
	Jobs           []Job    `gorm:"foreignKey:CompanyId"`
	Phone          string
	CompanyAddress string
	EmployeeNumber uint64
	Industry       string
	About          string
	CompanyLogo    *string
	Website        *string
	Twitter        *string
	Linkedin       *string
}
