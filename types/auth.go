package types

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type TalentForm struct {
	FirstName    string  `form:"firstName" json:"firstName" binding:"required"`
	LastName     string  `form:"lastName" json:"lastName" binding:"required"`
	Email        string  `gorm:"type:unique" form:"email" json:"email" binding:"required,email"`
	Gender       string  `form:"gender" json:"gender" binding:"required"`
	Country      string  `form:"country" json:"country" binding:"required"`
	State        string  `form:"state" json:"state" binding:"required"`
	About        string  `form:"about" json:"about" binding:"required"`
	Phone        string  `form:"phone" json:"phone" binding:"required"`
	Password     string  `form:"password" json:"password" binding:"required"`
	Experience   uint8   `form:"experience" json:"experience" binding:"required"`
	ProfilePhoto *string `form:"profilePhoto" json:"profilePhoto"`
	CV           *string `form:"cv" json:"cv"`
}

type EmployerForm struct {
	Name               string
	Email              string `gorm:"type:unique" form:"email" json:"email" binding:"required"`
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
	ProfilePhoto       *string
	Password           string `form:"password" json:"password" binding:"required"`
	Experience         uint8
}
