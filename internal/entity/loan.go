package entity

import "time"

type ApplyLoanRequest struct {
	FullName      string
	MonthlyIncome float64
	LoanAmount    float64
	LoanPurpose   string
	Age           int
	PhoneNumber   string
	Email         string
}

type ApplyLoanResponse struct {
	ApplicationID string
	Eligible      bool
	Reason        string
	Timestamp     time.Time
}

type LoanApplication struct {
	ApplicationID string
	FullName      string
	MonthlyIncome float64
	LoanAmount    float64
	LoanPurpose   string
	Age           int
	PhoneNumber   string
	Email         string
	Eligible      bool
	Reason        string
	Timestamp     time.Time
}

type ListLoansQuery struct {
	Page     int
	Limit    int
	Eligible *bool
	Purpose  string
}

type ListLoansResponse struct {
	Applications []LoanApplication
	Page         int
	TotalPages   int
}
