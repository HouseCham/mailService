package validator

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

// SetUpValidations sets up the custom validations for the validator
func SetUpValidations() {
	customValidator := validator.New()
	Validate = customValidator
}