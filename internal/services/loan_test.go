package services

import (
	"errors"
	"testing"
	"time"

	"github.com/fusic/skill-api/internal/constant"
	"github.com/fusic/skill-api/internal/entity"
	"github.com/fusic/skill-api/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestIsEligibleMonthlyIncome_Eligible(t *testing.T) {
	income := float64(10000)
	actual := isEligibleMonthlyIncome(income)
	assert.True(t, actual)
}

func TestIsEligibleMonthlyIncome_Ineligible(t *testing.T) {
	income := float64(9999)
	actual := isEligibleMonthlyIncome(income)
	assert.False(t, actual)
}

func TestIsEligibleAge_Eligible(t *testing.T) {
	tests := []struct {
		name string
		age  int
	}{
		{name: "minimum age limit", age: 20},
		{name: "maximum age limit", age: 60},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isEligibleAge(tt.age)
			assert.True(t, actual)
		})
	}

}

func TestIsEligibleAge_Ineligible(t *testing.T) {

	tests := []struct {
		name string
		age  int
	}{
		{name: "age too low", age: 19},
		{name: "age too high", age: 61},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isEligibleAge(tt.age)
			assert.False(t, actual)
		})
	}
}

func TestIsEligibleLoanPurpose_Eligible(t *testing.T) {
	purpose := "home"
	actual := isEligibleLoanPurpose(purpose)
	assert.True(t, actual)
}

func TestIsEligibleLoanPurpose_Ineligible(t *testing.T) {
	purpose := "business"
	actual := isEligibleLoanPurpose(purpose)
	assert.False(t, actual)
}

func TestIsEligibleLoanAmount_Eligible(t *testing.T) {
	loanAmount, income := float64(120000), float64(10000)
	actual := isEligibleLoanAmount(loanAmount, income)
	assert.True(t, actual)
}

func TestIsEligibleLoanAmount_Ineligible(t *testing.T) {
	loanAmount, income := float64(120001), float64(10000)
	actual := isEligibleLoanAmount(loanAmount, income)
	assert.False(t, actual)
}

func TestLoanService_ApplyLoan_Success(t *testing.T) {
	fixedID := "3fa85f64-5717-4562-b3fc-2c963f66afa6"
	fixedTime := time.Date(
		2026, time.August, 31,
		10, 30, 0, 0,
		time.Local,
	)

	request := entity.ApplyLoanRequest{
		FullName:      "John Doe",
		MonthlyIncome: 10000,
		LoanAmount:    120000,
		LoanPurpose:   "home",
		Age:           20,
		PhoneNumber:   "0812345678",
		Email:         "john@example.com",
	}
	//expectedresponse
	expectedResponse := entity.ApplyLoanResponse{
		ApplicationID: fixedID,
		Eligible:      true,
		Reason:        constant.ReasonEligible,
		Timestamp:     fixedTime,
	}

	mockRepository := mocks.NewLoanRepository(t)

	mockRepository.
		On("ApplyLoan", mock.AnythingOfType("models.LoanApplication")).
		Return(nil).
		Once()

	s := &loanService{
		loanRepository: mockRepository,
		generateID:     func() string { return fixedID },
		currentTime:    func() time.Time { return fixedTime },
	}
	actualResponse, err := s.ApplyLoan(request)

	require.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestLoanService_ApplyLoan_RepositoryError(t *testing.T) {
	request := entity.ApplyLoanRequest{
		FullName:      "John Doe",
		MonthlyIncome: 10000,
		LoanAmount:    120000,
		LoanPurpose:   "home",
		Age:           20,
		PhoneNumber:   "0812345678",
		Email:         "john@example.com",
	}

	expectedError := errors.New("failed to save loan application")
	expectedResponse := entity.ApplyLoanResponse{}

	mockRepository := mocks.NewLoanRepository(t)

	mockRepository.
		On("ApplyLoan", mock.AnythingOfType("models.LoanApplication")).
		Return(expectedError).
		Once()

	service := NewLoanService(mockRepository)

	actualResponse, err := service.ApplyLoan(request)

	require.ErrorIs(t, err, expectedError)
	assert.Equal(t, expectedResponse, actualResponse)
}
