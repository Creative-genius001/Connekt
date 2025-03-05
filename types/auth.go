package types

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type TalentForm struct {
	FirstName    string `form:"firstName" json:"firstName" binding:"required"`
	LastName     string `form:"lastName" json:"lastName" binding:"required"`
	Email        string `gorm:"type:unique" form:"email" json:"email" binding:"required,email"`
	Gender       string `form:"gender" json:"gender" binding:"required"`
	Country      string `form:"country" json:"country" binding:"required"`
	State        string `form:"state" json:"state" binding:"required"`
	About        string `form:"about" json:"about" binding:"required"`
	Phone        string `form:"phone" json:"phone" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	Website      *string
	Twitter      *string
	LinkedIn     *string
	Facebook     *string
	Github       *string
	ProfilePhoto *string `form:"profilePhoto" json:"profilePhoto"`
}

type CompanyForm struct {
	CompanyName        string  `form:"companyName" json:"companyName" binding:"required"`
	Email              string  `gorm:"type:unique" form:"email" json:"email" binding:"required,email"`
	About              string  `form:"about" json:"about" binding:"required"`
	Phone              string  `form:"phone" json:"phone" binding:"required"`
	Password           string  `form:"password" json:"password" binding:"required"`
	CompanyLogo        *string `form:"companyLogo"`
	CompanyAddress     string  `form:"companyAddress" json:"companyAddress" binding:"required"`
	RegistrationNumber string  `form:"registrationNumber" json:"registrationNumber" binding:"required"`
	EmployeeNumber     uint64  `form:"employeeNumber" json:"employeeNumber" binding:"required"`
	Industry           string  `form:"industry" json:"industry" binding:"required"`
	Website            *string
	Twitter            *string
	LinkedIn           *string
	Facebook           *string
	Github             *string
	Country            string
	State              string
}
