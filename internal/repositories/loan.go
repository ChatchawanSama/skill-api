package adaptor

import "gorm.io/gorm"

type LoanRepository interface {
}

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{db: db}
}

// TODO: implement to support Apply, GetByID, List methods. Use MySQL database with GORM.
