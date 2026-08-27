package rest

import (
	"fmt"
	"net/http"

	"github.com/fusic/skill-api/internal/dto"
	"github.com/fusic/skill-api/internal/entity"
	"github.com/fusic/skill-api/internal/services"
	"github.com/labstack/echo/v4"
)

type LoanHandler interface {
	ApplyLoan(c echo.Context) error
	GetLoanStatus(c echo.Context) error
	GetAllLoans(c echo.Context) error
}

type loanHandler struct {
	loanService services.LoanService
}

func NewLoanHandler(loanService services.LoanService) LoanHandler {
	return &loanHandler{
		loanService: loanService,
	}
}

func (h *loanHandler) ApplyLoan(c echo.Context) error {
	var requestDTO dto.ApplyLoanRequest

	if err := c.Bind(&requestDTO); err != nil {
		return c.JSON(http.StatusBadRequest, dto.AppErrorResponse{
			Message: "Invalid request body",
			Reason:  "Generic error",
		})
	}

	if err := c.Validate(&requestDTO); err != nil {
		return c.JSON(http.StatusBadRequest, dto.AppErrorResponse{
			Message: "Invalid request body",
			Reason:  validationReason(err),
		})
	}

	requestEntity := entity.ApplyLoanRequest{
		FullName:      requestDTO.FullName,
		MonthlyIncome: requestDTO.MonthlyIncome,
		LoanAmount:    requestDTO.LoanAmount,
		LoanPurpose:   requestDTO.LoanPurpose,
		Age:           requestDTO.Age,
		PhoneNumber:   requestDTO.PhoneNumber,
		Email:         requestDTO.Email,
	}

	responseService, err := h.loanService.ApplyLoan(requestEntity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.AppErrorResponse{
			Message: "Unable to process loan application",
			Reason:  err.Error(),
		})
	}

	responseDTO := dto.ApplyLoanResponse{
		ApplicationID: responseService.ApplicationID,
		Eligible:      responseService.Eligible,
		Reason:        responseService.Reason,
		Timestamp:     responseService.Timestamp,
	}

	return c.JSON(http.StatusOK, responseDTO)
}

func (h *loanHandler) GetLoanStatus(c echo.Context) error {
	applicationId := c.Param("applicationId")

	responseEntity, err := h.loanService.GetLoanStatus(applicationId)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.AppErrorResponse{
			Message: "Loan Application not found",
			Reason:  "applicationId not found: " + applicationId,
		})
	}

	responseDTO := dto.LoanApplication{
		ApplicationID: responseEntity.ApplicationID,
		FullName:      responseEntity.FullName,
		MonthlyIncome: responseEntity.MonthlyIncome,
		LoanAmount:    responseEntity.LoanAmount,
		LoanPurpose:   responseEntity.LoanPurpose,
		Age:           responseEntity.Age,
		PhoneNumber:   responseEntity.PhoneNumber,
		Email:         responseEntity.Email,
		Eligible:      responseEntity.Eligible,
		Reason:        responseEntity.Reason,
		Timestamp:     responseEntity.Timestamp,
	}
	return c.JSON(http.StatusOK, responseDTO)
}

func (h *loanHandler) GetAllLoans(c echo.Context) error {
	var queryDTO dto.ListLoansQuery

	if err := c.Bind(&queryDTO); err != nil {
		return c.JSON(http.StatusBadRequest, dto.AppErrorResponse{
			Message: "Invalid query parameters",
			Reason:  err.Error(),
		})
	}
	fmt.Println(queryDTO)

	queryEntity := entity.ListLoansQuery{
		Page:     queryDTO.Page,
		Limit:    queryDTO.Limit,
		Eligible: queryDTO.Eligible,
		Purpose:  queryDTO.Purpose,
	}

	responseEntity, err := h.loanService.GetAllLoans(queryEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.AppErrorResponse{
			Message: "Unable to retrieve loan applications",
			Reason:  err.Error(),
		})
	}
	applicationsDTO := make([]dto.LoanApplication, 0, len(responseEntity.Applications))

	for _, applicationEntity := range responseEntity.Applications {
		applicationDTO := dto.LoanApplication{
			ApplicationID: applicationEntity.ApplicationID,
			FullName:      applicationEntity.FullName,
			MonthlyIncome: applicationEntity.MonthlyIncome,
			LoanAmount:    applicationEntity.LoanAmount,
			LoanPurpose:   applicationEntity.LoanPurpose,
			Age:           applicationEntity.Age,
			PhoneNumber:   applicationEntity.PhoneNumber,
			Email:         applicationEntity.Email,
			Eligible:      applicationEntity.Eligible,
			Reason:        applicationEntity.Reason,
			Timestamp:     applicationEntity.Timestamp,
		}
		applicationsDTO = append(applicationsDTO, applicationDTO)
	}

	responseDTO := dto.ListLoansResponse{
		Applications: applicationsDTO,
		Page:         responseEntity.Page,
		TotalPages:   responseEntity.TotalPages,
	}

	return c.JSON(http.StatusOK, responseDTO)
}

// TODO: implement. Bind models.ApplyLoanRequest, validate, call h.svc.Apply.
// Success: 200 + ApplyLoanResponse. Bad request: 400 + {message:"Invalid request body", reason:...}.

// TODO: implement. Read :applicationId, call h.svc.GetByID.
// Success: 200 + LoanApplication. Not found: 404 + {message:"Loan application not found", reason:...}.

// TODO: implement. Bind query (page, limit, eligible, purpose), call h.svc.List.
// Success: 200 + ListLoansResponse.
