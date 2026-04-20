package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (v *Validator) FormatErrors(err error) []ValidationError {
	var errors []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, ValidationError{
			Field:   err.Field(),
			Message: fmt.Sprintf("validation failed on '%s'", err.Tag()),
		})
	}
	return errors
}
