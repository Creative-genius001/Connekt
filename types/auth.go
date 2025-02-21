package types

import "database/sql"

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type JobSeekerForm struct {
	FirstName    string `form:"firstName" json:"firstName" binding:"required"`
	LastName     string `form:"lastName" json:"lasrName" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required"`
	Gender       string `form:"gender" json:"gender" binding:"required"`
	Country      string `form:"country" json:"country" binding:"required"`
	State        string `form:"state" json:"state" binding:"required"`
	About        string `form:"about" json:"about" binding:"required"`
	Phone        string `form:"phone" json:"phone" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	Experience   uint8  `form:"experience" json:"experience" binding:"required"`
	ProfilePhoto sql.NullString
	CV           sql.NullString
}

type EmployerForm struct {
	Name               string
	Email              string `form:"email" json:"email" binding:"required"`
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
	ProfilePhoto       sql.NullString
	Password           string `form:"password" json:"password" binding:"required"`
	Experience         uint8
}
