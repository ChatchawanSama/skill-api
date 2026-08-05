package services

//go:generate mockery --name LoanService --output ./mocks

import (
	"context"

	"github.com/fusic/skill-api/internal/models"
	"github.com/fusic/skill-api/internal/repositories/adaptor"
)

type LoanService interface {
	Apply(ctx context.Context, req models.ApplyLoanRequest) (models.ApplyLoanResponse, error)
	GetByID(ctx context.Context, applicationID string) (models.LoanApplication, error)
	List(ctx context.Context, query models.ListLoansQuery) (models.ListLoansResponse, error)
}

type loanService struct {
	repo adaptor.LoanRepository
}

func NewLoanService(repo adaptor.LoanRepository) LoanService {
	return &loanService{repo: repo}
}

// TODO: implement. Generate UUID applicationId, evaluate eligibility rules (see PROJECT.md),
// set timestamp in Asia/Bangkok, persist via repo.Apply, return ApplyLoanResponse.
// Eligibility rules:
//   - monthlyIncome >= 10000 else ErrMonthlyIncomeLow
//   - age 20..60 else ErrAgeOutOfRange
//   - loanPurpose != "business" else ErrBusinessNotSupported
//   - loanAmount <= 12*monthlyIncome else ErrLoanAmountTooHigh
// All pass -> eligible=true, reason=ReasonEligible.
func (s *loanService) Apply(ctx context.Context, req models.ApplyLoanRequest) (models.ApplyLoanResponse, error) {
	return models.ApplyLoanResponse{}, ErrNotImplemented
}

// TODO: implement. Call repo.GetByID. Return ErrLoanNotFound when missing.
func (s *loanService) GetByID(ctx context.Context, applicationID string) (models.LoanApplication, error) {
	return models.LoanApplication{}, ErrNotImplemented
}

// TODO: implement. Apply pagination defaults (page=1, limit=10), call repo.List,
// build ListLoansResponse with applications + page + totalPages.
func (s *loanService) List(ctx context.Context, query models.ListLoansQuery) (models.ListLoansResponse, error) {
	return models.ListLoansResponse{}, ErrNotImplemented
}
