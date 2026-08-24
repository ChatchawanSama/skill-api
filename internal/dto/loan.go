package dto

import "time"

type ApplyLoanRequest struct {
	FullName      string  `json:"fullName" validate:"required,min=2,max=255"`
	MonthlyIncome float64 `json:"monthlyIncome" validate:"required,min=5000,max=5000000"`
	LoanAmount    float64 `json:"loanAmount" validate:"required,min=1000,max=5000000"`
	LoanPurpose   string  `json:"loanPurpose" validate:"required,oneof=education home car business personal"`
	Age           int     `json:"age" validate:"required,gt=0"`
	PhoneNumber   string  `json:"phoneNumber" validate:"required,len=10,numeric"`
	Email         string  `json:"email" validate:"required,email"`
}

type ApplyLoanResponse struct {
	ApplicationID string    `json:"applicationId"`
	Eligible      bool      `json:"eligible"`
	Reason        string    `json:"reason"`
	Timestamp     time.Time `json:"timestamp"`
}

type LoanApplication struct {
	ApplicationID string    `json:"applicationId"`
	FullName      string    `json:"fullName"`
	MonthlyIncome float64   `json:"monthlyIncome"`
	LoanAmount    float64   `json:"loanAmount"`
	LoanPurpose   string    `json:"loanPurpose"`
	Age           int       `json:"age"`
	PhoneNumber   string    `json:"phoneNumber"`
	Email         string    `json:"email"`
	Eligible      bool      `json:"eligible"`
	Reason        string    `json:"reason"`
	Timestamp     time.Time `json:"timestamp"`
}

type ListLoansQuery struct {
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
	Eligible *bool  `query:"eligible"`
	Purpose  string `query:"purpose"`
}

type ListLoansResponse struct {
	Applications []LoanApplication `json:"applications"`
	Page         int               `json:"page"`
	TotalPages   int               `json:"totalPages"`
}
