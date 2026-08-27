package services

import (
	"time"

	"github.com/fusic/skill-api/internal/constant"
	"github.com/fusic/skill-api/internal/entity"
	"github.com/fusic/skill-api/internal/models"
	"github.com/fusic/skill-api/internal/repositories"
	"github.com/google/uuid"
)

//go:generate mockery --name LoanService --output ./mocks

type LoanService interface {
	ApplyLoan(request entity.ApplyLoanRequest) (entity.ApplyLoanResponse, error)
	GetLoanStatus(applicationId string) (entity.LoanApplication, error)
	GetAllLoans(query entity.ListLoansQuery) (entity.ListLoansResponse, error)
}

type loanService struct {
	loanRepository repositories.LoanRepository
}

func NewLoanService(loanRepository repositories.LoanRepository) LoanService {
	return &loanService{
		loanRepository: loanRepository,
	}
}

func checkEligibility(request entity.ApplyLoanRequest) (string, bool) {
	switch {
	case !isEligibleMonthlyIncome(request.MonthlyIncome):
		return constant.ReasonInsufficientIncome, false

	case !isEligibleAge(request.Age):
		return constant.ReasonInvalidAge, false

	case !isEligibleLoanPurpose(request.LoanPurpose):
		return constant.ReasonInvalidPurpose, false

	case !isEligibleLoanAmount(request.LoanAmount, request.MonthlyIncome):
		return constant.ReasonExcessRequest, false

	default:
		return constant.ReasonEligible, true
	}
}

func isEligibleMonthlyIncome(income float64) bool {
	return income >= 10000
}

func isEligibleAge(age int) bool {
	return age >= 20 && age <= 60
}

func isEligibleLoanPurpose(purpose string) bool {
	return purpose != "business"
}

func isEligibleLoanAmount(loanAmount, income float64) bool {
	return loanAmount <= income*12
}

func (s *loanService) ApplyLoan(request entity.ApplyLoanRequest) (entity.ApplyLoanResponse, error) {

	reason, eligibility := checkEligibility(request)

	response := entity.ApplyLoanResponse{
		ApplicationID: uuid.NewString(),
		Eligible:      eligibility,
		Reason:        reason,
		Timestamp:     time.Now(),
	}

	application := models.LoanApplication{
		ApplicationID: response.ApplicationID,
		FullName:      request.FullName,
		MonthlyIncome: request.MonthlyIncome,
		LoanAmount:    request.LoanAmount,
		LoanPurpose:   request.LoanPurpose,
		Age:           request.Age,
		PhoneNumber:   request.PhoneNumber,
		Email:         request.Email,
		Eligible:      response.Eligible,
		Reason:        response.Reason,
		Timestamp:     response.Timestamp,
	}

	if err := s.loanRepository.ApplyLoan(application); err != nil {
		return entity.ApplyLoanResponse{}, err
	}

	return response, nil
}

func (s *loanService) GetLoanStatus(applicationId string) (entity.LoanApplication, error) {
	application, err := s.loanRepository.GetLoanStatus(applicationId)
	if err != nil {
		return entity.LoanApplication{}, err
	}

	response := entity.LoanApplication{
		ApplicationID: application.ApplicationID,
		FullName:      application.FullName,
		MonthlyIncome: application.MonthlyIncome,
		LoanAmount:    application.LoanAmount,
		LoanPurpose:   application.LoanPurpose,
		Age:           application.Age,
		PhoneNumber:   application.PhoneNumber,
		Email:         application.Email,
		Eligible:      application.Eligible,
		Reason:        application.Reason,
		Timestamp:     application.Timestamp,
	}

	return response, nil
}

func (s *loanService) GetAllLoans(queryEntity entity.ListLoansQuery) (entity.ListLoansResponse, error) {

	if queryEntity.Page <= 0 {
		queryEntity.Page = 1
	}

	if queryEntity.Limit <= 0 {
		queryEntity.Limit = 10
	}

	queryModel := models.ListLoansQuery{
		Page:     queryEntity.Page,
		Limit:    queryEntity.Limit,
		Eligible: queryEntity.Eligible,
		Purpose:  queryEntity.Purpose,
	}

	applicationsModel, total, err := s.loanRepository.GetAllLoans(queryModel)
	if err != nil {
		return entity.ListLoansResponse{}, err
	}

	totalPages := 0

	if total > 0 {
		totalPages = int((total + int64(queryEntity.Limit) - 1) / int64(queryEntity.Limit))
	}

	var applicationsEntity []entity.LoanApplication

	for _, applicationModel := range applicationsModel {
		applicationEntity := entity.LoanApplication{
			ApplicationID: applicationModel.ApplicationID,
			FullName:      applicationModel.FullName,
			MonthlyIncome: applicationModel.MonthlyIncome,
			LoanAmount:    applicationModel.LoanAmount,
			LoanPurpose:   applicationModel.LoanPurpose,
			Age:           applicationModel.Age,
			PhoneNumber:   applicationModel.PhoneNumber,
			Email:         applicationModel.Email,
			Eligible:      applicationModel.Eligible,
			Reason:        applicationModel.Reason,
			Timestamp:     applicationModel.Timestamp,
		}
		applicationsEntity = append(applicationsEntity, applicationEntity)
	}

	response := entity.ListLoansResponse{
		Applications: applicationsEntity,
		Page:         queryEntity.Page,
		TotalPages:   totalPages,
	}

	return response, nil
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
