package services

import (
	"testing"

	"github.com/fusic/skill-api/internal/constant"
	"github.com/fusic/skill-api/internal/entity"
	"github.com/fusic/skill-api/internal/models"
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

func TestLoanService_ApplyLoan(t *testing.T) {
	request := entity.ApplyLoanRequest{
		FullName:      "John Doe",
		MonthlyIncome: 10000,
		LoanAmount:    120000,
		LoanPurpose:   "home",
		Age:           20,
		PhoneNumber:   "0812345678",
		Email:         "john@example.com",
	}
	var savedApplication models.LoanApplication

	mockRepository := mocks.NewLoanRepository(t)

	mockRepository.
		On("ApplyLoan", mock.AnythingOfType("models.LoanApplication")).
		Run(func(args mock.Arguments) {
			savedApplication = args.Get(0).(models.LoanApplication)
		}).
		Return(nil).
		Once()

	s := NewLoanService(mockRepository)

	response, err := s.ApplyLoan(request)

	require.NoError(t, err)
	assert.NotEmpty(t, response.ApplicationID)
	assert.True(t, response.Eligible)
	assert.Equal(t, constant.ReasonEligible, response.Reason)
	assert.False(t, response.Timestamp.IsZero())

	assert.Equal(t, request.FullName, savedApplication.FullName)
	assert.Equal(t, request.MonthlyIncome, savedApplication.MonthlyIncome)
	assert.Equal(t, request.LoanAmount, savedApplication.LoanAmount)
	assert.Equal(t, request.LoanPurpose, savedApplication.LoanPurpose)
	assert.Equal(t, request.Age, savedApplication.Age)
	assert.Equal(t, request.PhoneNumber, savedApplication.PhoneNumber)
	assert.Equal(t, request.Email, savedApplication.Email)

	assert.Equal(t, response.ApplicationID, savedApplication.ApplicationID)
	assert.Equal(t, response.Eligible, savedApplication.Eligible)
	assert.Equal(t, response.Reason, savedApplication.Reason)
	assert.Equal(t, response.Timestamp, savedApplication.Timestamp)
}
