package utils

import "github.com/go-playground/validator/v10"

func initValidator() *validator.Validate {
	return validator.New()
}
