package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	age := 20
	actual := isEligibleAge(age)
	assert.True(t, actual)
}

func TestIsEligibleAge_Ineligible(t *testing.T) {
	age := 15
	actual := isEligibleAge(age)
	assert.False(t, actual)
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
