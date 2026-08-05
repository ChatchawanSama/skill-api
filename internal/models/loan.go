package models

import "time"

type LoanPurpose string

const (
	LoanPurposeEducation LoanPurpose = "education"
	LoanPurposeHome      LoanPurpose = "home"
	LoanPurposeCar       LoanPurpose = "car"
	LoanPurposeBusiness  LoanPurpose = "business"
	LoanPurposePersonal  LoanPurpose = "personal"
)

func ValidLoanPurposes() []LoanPurpose {
	return []LoanPurpose{
		LoanPurposeEducation,
		LoanPurposeHome,
		LoanPurposeCar,
		LoanPurposeBusiness,
		LoanPurposePersonal,
	}
}

type ApplyLoanRequest struct {
	FullName      string  `json:"fullName"`
	MonthlyIncome float64 `json:"monthlyIncome"`
	LoanAmount    float64 `json:"loanAmount"`
	LoanPurpose   string  `json:"loanPurpose"`
	Age           int     `json:"age"`
	PhoneNumber   string  `json:"phoneNumber"`
	Email         string  `json:"email"`
}

type LoanApplication struct {
	ApplicationID string      `json:"applicationId"`
	FullName      string      `json:"fullName"`
	MonthlyIncome float64     `json:"monthlyIncome"`
	LoanAmount    float64     `json:"loanAmount"`
	LoanPurpose   LoanPurpose `json:"loanPurpose"`
	Age           int         `json:"age"`
	PhoneNumber   string      `json:"phoneNumber"`
	Email         string      `json:"email"`
	Eligible      bool        `json:"eligible"`
	Reason        string      `json:"reason"`
	Timestamp     time.Time   `json:"timestamp"`
}

type ApplyLoanResponse struct {
	ApplicationID string    `json:"applicationId"`
	Eligible      bool      `json:"eligible"`
	Reason        string    `json:"reason"`
	Timestamp     time.Time `json:"timestamp"`
}

type ListLoansResponse struct {
	Applications []LoanApplication `json:"applications"`
	Page         int               `json:"page"`
	TotalPages   int               `json:"totalPages"`
}

type ListLoansQuery struct {
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
	Eligible *bool  `query:"eligible"`
	Purpose  string `query:"purpose"`
}
