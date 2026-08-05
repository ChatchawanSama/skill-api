package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fusic/skill-api/internal/models"
)

type loanRepository struct {
	db *sql.DB
}

func NewLoanRepository(db *sql.DB) *loanRepository {
	return &loanRepository{db: db}
}

var errUnimplemented = errors.New("repository not implemented")

// TODO: implement. Insert loan application row, return stored record.
func (r *loanRepository) Apply(ctx context.Context, application models.LoanApplication) (models.LoanApplication, error) {
	return models.LoanApplication{}, errUnimplemented
}

// TODO: implement. Select loan row by applicationId. Return ErrLoanNotFound when missing.
func (r *loanRepository) GetByID(ctx context.Context, applicationID string) (models.LoanApplication, error) {
	return models.LoanApplication{}, errUnimplemented
}

// TODO: implement. Select loan rows with pagination + optional filters (eligible, purpose).
// Return applications slice and total page count.
func (r *loanRepository) List(ctx context.Context, query models.ListLoansQuery) ([]models.LoanApplication, int, error) {
	return nil, 0, errUnimplemented
}
