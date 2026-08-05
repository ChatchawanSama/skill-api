package rest

type LoanHandler interface {
}

type loanHandler struct {
}

func NewLoanHandler() LoanHandler {
	return &loanHandler{}
}

// TODO: implement. Bind models.ApplyLoanRequest, validate, call h.svc.Apply.
// Success: 200 + ApplyLoanResponse. Bad request: 400 + {message:"Invalid request body", reason:...}.

// TODO: implement. Read :applicationId, call h.svc.GetByID.
// Success: 200 + LoanApplication. Not found: 404 + {message:"Loan application not found", reason:...}.

// TODO: implement. Bind query (page, limit, eligible, purpose), call h.svc.List.
// Success: 200 + ListLoansResponse.
