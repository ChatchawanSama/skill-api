package entities

import "time"

type LoanRow struct {
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
