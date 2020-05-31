package utils

import (
	"github.com/go-playground/validator"
)

var (
	validate = validator.New()
)

func ValidateInput(input interface{}) error {
	if err := validate.Struct(input); err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}
