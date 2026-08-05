package services

import (
	"errors"

	"github.com/fusic/skill-api/internal/models"
	"github.com/go-playground/validator"
)

var (
	ErrNotImplemented = errors.New("not implemented")

	ErrInvalidRequestBody  = errors.New("Invalid request body")
	ErrLoanNotFound        = errors.New("Loan application not found")
	ErrMonthlyIncomeLow    = errors.New("Monthly income is insufficient")
	ErrAgeOutOfRange       = errors.New("Age not in range (must be between 20-60)")
	ErrBusinessNotSupported = errors.New("Business loans not supported")
	ErrLoanAmountTooHigh   = errors.New("Loan amount cannot exceed 12 months of income")
)

const ReasonEligible = "Eligible under base rules"

func ValidateStruct(payload interface{}) error {
	v := validator.New()
	return v.Struct(payload)
}

func SuccessResponse(data interface{}) models.AppResponse {
	return models.AppResponse{Status: models.StatusSuccess, Data: data}
}

func ErrorResponse(message string) models.AppErrorResponse {
	return models.AppErrorResponse{Status: models.StatusError, Message: message}
}
