package repositories

import (
	"github.com/fusic/skill-api/internal/models"
	"gorm.io/gorm"
)

//go:generate mockery --name LoanRepository --output ./mocks

type LoanRepository interface {
	ApplyLoan(application models.LoanApplication) error
	GetLoanStatus(applicationId string) (models.LoanApplication, error)
	GetAllLoans(queryModel models.ListLoansQuery) ([]models.LoanApplication, int64, error)
}

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{db: db}
}

func (r *loanRepository) ApplyLoan(application models.LoanApplication) error {
	return r.db.Create(&application).Error
}

func (r *loanRepository) GetLoanStatus(applicationId string) (models.LoanApplication, error) {
	var application models.LoanApplication

	err := r.db.First(&application, "application_id = ?", applicationId).Error

	return application, err
}

func (r *loanRepository) GetAllLoans(query models.ListLoansQuery) ([]models.LoanApplication, int64, error) {
	applications := make([]models.LoanApplication, 0)
	var total int64

	dbQuery := r.db.Model(&models.LoanApplication{})

	if query.Eligible != nil {
		dbQuery = dbQuery.Where("eligible = ?", *query.Eligible)
	}

	if query.Purpose != "" {
		dbQuery = dbQuery.Where("loan_purpose = ?", query.Purpose)
	}

	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.Limit

	err := dbQuery.
		Order("timestamp DESC").
		Offset(offset).
		Limit(query.Limit).
		Find(&applications).
		Error

	if err != nil {
		return nil, 0, err
	}

	return applications, total, nil
}

// TODO: implement to support Apply, GetByID, List methods. Use MySQL database with GORM.
