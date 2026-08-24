package rest

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return &CustomValidator{
		validator: validate,
	}
}

func (cv *CustomValidator) Validate(value interface{}) error {
	return cv.validator.Struct(value)
}

func validationReason(err error) string {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok || len(validationErrors) == 0 {
		return "request contains invalid fields"
	}

	missingFields := make([]string, 0)

	for _, fieldError := range validationErrors {
		if fieldError.Tag() == "required" {
			missingFields = append(missingFields, fieldError.Field())
		}
	}

	if len(missingFields) > 0 {
		return "missing required fields: " + strings.Join(missingFields, ", ")
	}

	field := validationErrors[0].Field()

	return fmt.Sprintf("%s must be a valid %s", field, field)

}

// fieldError := validationErrors[0]

// switch fieldError.Field() {
// case "FullName":
// 	if fieldError.Tag() == "required" {
// 		return "fullName is required"
// 	}
// 	return "fullName must be between 2 and 255 characters"

// case "MonthlyIncome":
// 	if fieldError.Tag() == "required" {
// 		return "monthlyIncome is required"
// 	}
// 	return "monthlyIncome must be between 5000 and 5000000"

// case "LoanAmount":
// 	if fieldError.Tag() == "required" {
// 		return "loanAmount is required"
// 	}
// 	return "loanAmount must be between 1000 and 5000000"

// case "LoanPurpose":
// 	if fieldError.Tag() == "required" {
// 		return "loanPurpose is required"
// 	}
// 	return "loanPurpose must be one of education, home, car, business, personal"

// case "Age":
// 	if fieldError.Tag() == "required" {
// 		return "age is required"
// 	}
// 	return "age must be more than 0"

// case "PhoneNumber":
// 	if fieldError.Tag() == "required" {
// 		return "phoneNumber is required"
// 	}
// 	return "phoneNumber must contain exactly 10 numeric digits"

// case "Email":
// 	if fieldError.Tag() == "required" {
// 		return "email is required"
// 	}
// 	return "email must be a valid email"

// default:
// 	return fmt.Sprintf("%s is invalid", fieldError.Field())
