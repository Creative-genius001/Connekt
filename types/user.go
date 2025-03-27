package types

type UpdateCompanyForm struct {
	CompanyName    *string `json:"companyName,omitempty"`
	Phone          *string `json:"phone,omitempty"`
	CompanyAddress *string `json:"companyAddress,omitempty"`
	EmployeeNumber *uint64 `json:"employeeNumber,omitempty"`
	Industry       *string `json:"industry,omitempty"`
	About          *string `json:"about,omitempty"`
	CompanyLogo    *string `json:"companyLogo,omitempty"`
	Website        *string `json:"website,omitempty"`
	Twitter        *string `json:"twitter,omitempty"`
	Linkedin       *string `json:"linkedin,omitempty"`
}

type UpdateTalentForm struct {
	FirstName    *string `json:"firstName,omitempty"`
	LastName     *string `json:"lastName,omitempty"`
	Gender       *string `json:"gender,omitempty"`
	Country      *string `json:"country,omitempty"`
	State        *string `json:"state,omitempty"`
	City         *string `json:"city,omitempty"`
	About        *string `json:"about,omitempty"`
	Phone        *string `json:"phone,omitempty"`
	Website      *string `json:"website,omitempty"`
	Twitter      *string `json:"twitter,omitempty"`
	Linkedin     *string `json:"linkedin,omitempty"`
	Github       *string `json:"github,omitempty"`
	ProfilePhoto *string `json:"profilePhoto,omitempty"`
}
