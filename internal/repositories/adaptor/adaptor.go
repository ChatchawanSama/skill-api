package adaptor

//go:generate mockery --name LoanRepository --output ./mocks

import (
	"context"

	"github.com/fusic/skill-api/internal/models"
)

type LoanRepository interface {
	Apply(ctx context.Context, application models.LoanApplication) (models.LoanApplication, error)
	GetByID(ctx context.Context, applicationID string) (models.LoanApplication, error)
	List(ctx context.Context, query models.ListLoansQuery) ([]models.LoanApplication, int, error)
}
