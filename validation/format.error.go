package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) string {
	var sb strings.Builder

	for _, e := range err.(validator.ValidationErrors) {
		field := strings.ToLower(e.Field())
		tag := e.Tag()

		switch tag {
		case "required":
			sb.WriteString(field + " is required; ")
		case "email":
			sb.WriteString(field + " must be a valid email; ")
		case "min":
			sb.WriteString(field + " minimum length is " + e.Param() + "; ")
		default:
			sb.WriteString(field + " is invalid; ")
		}
	}

	return strings.TrimSuffix(sb.String(), "; ")
}
