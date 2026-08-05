package rest

import (
	"net/http"

	"github.com/fusic/skill-api/internal/services"
	"github.com/labstack/echo/v4"
)

type LoanHandler interface {
	Apply(c echo.Context) error
	GetByID(c echo.Context) error
	List(c echo.Context) error
}

type loanHandler struct {
	svc services.LoanService
}

func NewLoanHandler(svc services.LoanService) LoanHandler {
	return &loanHandler{svc: svc}
}

// TODO: implement. Bind models.ApplyLoanRequest, validate, call h.svc.Apply.
// Success: 200 + ApplyLoanResponse. Bad request: 400 + {message:"Invalid request body", reason:...}.
func (h *loanHandler) Apply(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}

// TODO: implement. Read :applicationId, call h.svc.GetByID.
// Success: 200 + LoanApplication. Not found: 404 + {message:"Loan application not found", reason:...}.
func (h *loanHandler) GetByID(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}

// TODO: implement. Bind query (page, limit, eligible, purpose), call h.svc.List.
// Success: 200 + ListLoansResponse.
func (h *loanHandler) List(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
