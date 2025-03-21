package types

type CreateJobForm struct {
	Title          string  `json:"title" binding:"required"`
	Description    string  `json:"description" binding:"required"`
	Remote         bool    `json:"remote"`
	IsActive       bool    `json:"isActive"`
	Industry       string  `json:"industry" binding:"required"`
	MinValue       float64 `json:"minValue" binding:"required"`
	MaxValue       float64 `json:"maxValue" binding:"required"`
	Currency       string  `json:"currency" binding:"required"`
	City           string  `json:"city" binding:"required"`
	State          string  `json:"state" binding:"required"`
	Country        string  `json:"country" binding:"required"`
	ApplicationUrl string  `json:"applicationUrl" binding:"required"`
}

type UpdateJobForm struct {
	Title          string  `json:"title" binding:"required"`
	Description    string  `json:"description" binding:"required"`
	Remote         bool    `json:"remote"`
	IsActive       bool    `json:"isActive"`
	Industry       string  `json:"industry" binding:"required"`
	MinValue       float64 `json:"minValue" binding:"required"`
	MaxValue       float64 `json:"maxValue" binding:"required"`
	Currency       string  `json:"currency" binding:"required"`
	City           string  `json:"city" binding:"required"`
	State          string  `json:"state" binding:"required"`
	Country        string  `json:"country" binding:"required"`
	ApplicationUrl string  `json:"applicationUrl" binding:"required"`
}

type JobApplicationForm struct {
	Coverletter string `json:"coverletter" binding:"required"`
	Resume      string `json:"resume" binding:"required"`
}
