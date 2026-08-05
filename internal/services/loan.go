package services

//go:generate mockery --name LoanService --output ./mocks

type LoanService interface {
}

type loanService struct {
}

func NewLoanService() LoanService {
	return &loanService{}
}

// TODO: implement. Generate UUID applicationId, evaluate eligibility rules (see PROJECT.md),
// set timestamp in Asia/Bangkok, persist via repo.Apply, return ApplyLoanResponse.
// Eligibility rules:
//   - monthlyIncome >= 10000 else ErrMonthlyIncomeLow
//   - age 20..60 else ErrAgeOutOfRange
//   - loanPurpose != "business" else ErrBusinessNotSupported
//   - loanAmount <= 12*monthlyIncome else ErrLoanAmountTooHigh
// All pass -> eligible=true, reason=ReasonEligible.

// TODO: implement. Call repo.GetByID. Return ErrLoanNotFound when missing.

// TODO: implement. Apply pagination defaults (page=1, limit=10), call repo.List,
// build ListLoansResponse with applications + page + totalPages.
