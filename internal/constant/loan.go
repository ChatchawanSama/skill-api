package constant

import "errors"

var (
	ErrInsufficientIncome = errors.New("Monthly income is insufficient")
	ErrInvalidAge         = errors.New("Age not in range (must be between 20-60)")
	ErrInvalidPurpose     = errors.New("Business loans not supported")
	ErrExcessRequest      = errors.New("Loan amount cannot exceed 12 months of income")
)

const (
	ReasonEligible           = "Eligible under base rules"
	ReasonInsufficientIncome = "Monthly income is insufficient"
	ReasonInvalidAge         = "Age not in range (must be between 20-60)"
	ReasonInvalidPurpose     = "Business loans not supported"
	ReasonExcessRequest      = "Loan amount cannot exceed 12 months of income"
)
