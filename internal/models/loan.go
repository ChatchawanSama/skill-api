package models

import "time"

type LoanApplication struct {
	ApplicationID string    `gorm:"column:application_id;primaryKey;size:36"`
	FullName      string    `gorm:"column:full_name"`
	MonthlyIncome float64   `gorm:"column:monthly_income"`
	LoanAmount    float64   `gorm:"column:loan_amount"`
	LoanPurpose   string    `gorm:"column:loan_purpose"`
	Age           int       `gorm:"column:age"`
	PhoneNumber   string    `gorm:"column:phone_number"`
	Email         string    `gorm:"column:email"`
	Eligible      bool      `gorm:"column:eligible"`
	Reason        string    `gorm:"column:reason"`
	Timestamp     time.Time `gorm:"column:timestamp"`
}

type ListLoansQuery struct {
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
	Eligible *bool  `query:"eligible"`
	Purpose  string `query:"purpose"`
}

func (LoanApplication) TableName() string {
	return "loan_applications"
}
