package validation

import "github.com/go-playground/validator/v10"

func FieldValidator(fe validator.FieldError) string {
	fieldName := fe.Field()

	switch fe.Tag() {

	case "required":
		return fieldName + " is required"
	case "min":
		return fieldName + " must be at least " + fe.Param() + " characters"
	case "email":
		return "Invalid email format"
	case "eqfield":
		return "Password doesn't match"
	default:
		return "Invalid " + fieldName

	}

}