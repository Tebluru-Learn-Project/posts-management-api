package helper

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string][]string {
	errors := make(map[string][]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors["body"] = []string{"Invalid request payload."}
		return errors
	}

	for _, err := range validationErrors {
		field := strings.ToLower(err.Field())

		switch err.Tag() {
		case "required":
			errors[field] = append(errors[field], "The "+field+" field is required.")
		case "email":
			errors[field] = append(errors[field], "The "+field+" field must be a valid email address.")
		case "min":
			errors[field] = append(errors[field], "The "+field+" field must be at least "+err.Param()+" characters.")
		default:
			errors[field] = append(errors[field], "The "+field+" field is invalid.")
		}
	}

	return errors
}