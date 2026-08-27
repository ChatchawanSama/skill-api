package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/fusic/skill-api/internal/dto"
	"github.com/fusic/skill-api/internal/entity"
	"github.com/fusic/skill-api/internal/services/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoanHandler_ApplyLoan(t *testing.T) {

	tests := []struct {
		name       string
		body       string //struct? any?
		mockSetup  func(*mocks.LoanService)
		wantStatus int
		wantBody   any
	}{
		{
			name:       "Malformed JSON (HTTP 400)",
			body:       `{"fullName":"Somkanit",}`,
			mockSetup:  nil,
			wantStatus: http.StatusBadRequest,
			wantBody: dto.AppErrorResponse{
				Message: "Invalid request body",
				Reason:  "Generic error",
			},
		},
		{
			name: "Missing required fields (HTTP 400)",
			body: `{
  						"monthlyIncome": 0,
  						"loanAmount": 100000,
  						"loanPurpose": "home",
  						"age": 25,
  						"phoneNumber": "0851234567",
  						"email": "demo@example.com"
					}`,
			mockSetup:  nil,
			wantStatus: http.StatusBadRequest,
			wantBody: dto.AppErrorResponse{
				Message: "Invalid request body",
				Reason:  "missing required fields: fullName, monthlyIncome",
			},
		},
		{
			name: "Invalid fullName (HTTP 400)",
			body: `{
  						"fullName": "S",
 						"monthlyIncome": 15000,
  						"loanAmount": 100000,
  						"loanPurpose": "home",
  						"age": 25,
  						"phoneNumber": "0851234567",
  						"email": "demo@example.com"
					}`,
			mockSetup:  nil,
			wantStatus: http.StatusBadRequest,
			wantBody: dto.AppErrorResponse{
				Message: "Invalid request body",
				Reason:  "fullName must be a valid fullName",
			},
		},
		{
			name: "Invalid monthlyIncome (HTTP 400)",
			body: `{
  						"fullName": "Somkanit Jitsanook",
 						"monthlyIncome": 4999,
  						"loanAmount": 100000,
  						"loanPurpose": "home",
  						"age": 25,
  						"phoneNumber": "0851234567",
  						"email": "demo@example.com"
					}`,
			mockSetup:  nil,
			wantStatus: http.StatusBadRequest,
			wantBody: dto.AppErrorResponse{
				Message: "Invalid request body",
				Reason:  "monthlyIncome must be a valid monthlyIncome",
			},
		},
		{
			name: "Invalid loanAmount (HTTP 400)",
			body: `{
  						"fullName": "Somkanit Jitsanook",
 						"monthlyIncome": 15000,
  						"loanAmount": 999,
  						"loanPurpose": "home",
  						"age": 25,
  						"phoneNumber": "0851234567",
  						"email": "demo@example.com"
					}`,
			mockSetup:  nil,
			wantStatus: http.StatusBadRequest,
			wantBody: dto.AppErrorResponse{
				Message: "Invalid request body",
				Reason:  "loanAmount must be a valid loanAmount",
			},
		},
		{
			name: "Invalid loanPurpose (HTTP 400)",
			body: `{
  						"fullName": "Somkanit Jitsanook",
 						"monthlyIncome": 15000,
  						"loanAmount": 100000,
  						"loanPurpose": "holiday",
  						"age": 25,
  						"phoneNumber": "0851234567",
  						"email": "demo@example.com"
					}`,
			mockSetup:  nil,
			wantStatus: http.StatusBadRequest,
			wantBody: dto.AppErrorResponse{
				Message: "Invalid request body",
				Reason:  "loanPurpose must be a valid loanPurpose",
			},
		},
		{
			name: "Invalid age (HTTP 400)",
			body: `{
  						"fullName": "Somkanit Jitsanook",
 						"monthlyIncome": 15000,
  						"loanAmount": 100000,
  						"loanPurpose": "home",
  						"age": -1,
  						"phoneNumber": "0851234567",
  						"email": "demo@example.com"
					}`,
			mockSetup:  nil,
			wantStatus: http.StatusBadRequest,
			wantBody: dto.AppErrorResponse{
				Message: "Invalid request body",
				Reason:  "age must be a valid age",
			},
		},
		{
			name: "Invalid phoneNumber (HTTP 400)",
			body: `{
  						"fullName": "Somkanit Jitsanook",
 						"monthlyIncome": 15000,
  						"loanAmount": 100000,
  						"loanPurpose": "home",
  						"age": 25,
  						"phoneNumber": "085123",
  						"email": "demo@example.com"
					}`,
			mockSetup:  nil,
			wantStatus: http.StatusBadRequest,
			wantBody: dto.AppErrorResponse{
				Message: "Invalid request body",
				Reason:  "phoneNumber must be a valid phoneNumber",
			},
		},
		{
			name: "Invalid email (HTTP 400)",
			body: `{
  						"fullName": "Somkanit Jitsanook",
 						"monthlyIncome": 15000,
  						"loanAmount": 100000,
  						"loanPurpose": "home",
  						"age": 25,
  						"phoneNumber": "0851234567",
  						"email": "not-an-email"
					}`,
			mockSetup:  nil,
			wantStatus: http.StatusBadRequest,
			wantBody: dto.AppErrorResponse{
				Message: "Invalid request body",
				Reason:  "email must be a valid email",
			},
		},
		{
			name: "Valid successful (HTTP 200)",
			body: `{
  						"fullName": "Somkanit Jitsanook",
 						"monthlyIncome": 15000,
  						"loanAmount": 100000,
  						"loanPurpose": "home",
  						"age": 25,
  						"phoneNumber": "0851234567",
  						"email": "demo@example.com"
					}`,
			mockSetup: func(mockService *mocks.LoanService) {
				// expectedRequest := entity.ApplyLoanRequest{
				// 	FullName:      "Somkanit Jitsanook",
				// 	MonthlyIncome: 15000,
				// 	LoanAmount:    100000,
				// 	LoanPurpose:   "home",
				// 	Age:           25,
				// 	PhoneNumber:   "0851234567",
				// 	Email:         "demo@example.com",
				// }
				serviceResponse := entity.ApplyLoanResponse{
					ApplicationID: "test-application-id",
					Eligible:      true,
					Reason:        "Eligible under base rules",
					Timestamp:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
				}
				mockService.
					On("ApplyLoan", mock.Anything).
					Return(serviceResponse, nil).
					Once()
			},
			wantStatus: http.StatusOK,
			wantBody: dto.ApplyLoanResponse{
				ApplicationID: "test-application-id",
				Eligible:      true,
				Reason:        "Eligible under base rules",
				Timestamp:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Valid unsuccessful (HTTP 500)",
			body: `{
  						"fullName": "Somkanit Jitsanook",
 						"monthlyIncome": 15000,
  						"loanAmount": 100000,
  						"loanPurpose": "home",
  						"age": 25,
  						"phoneNumber": "0851234567",
  						"email": "demo@example.com"
					}`,
			mockSetup: func(mockService *mocks.LoanService) {
				expectedRequest := entity.ApplyLoanRequest{
					FullName:      "Somkanit Jitsanook",
					MonthlyIncome: 15000,
					LoanAmount:    100000,
					LoanPurpose:   "home",
					Age:           25,
					PhoneNumber:   "0851234567",
					Email:         "demo@example.com",
				}
				mockService.
					On("ApplyLoan", expectedRequest).
					Return(
						entity.ApplyLoanResponse{},
						errors.New("database error"),
					).
					Once()
			},
			wantStatus: http.StatusInternalServerError,
			wantBody: dto.AppErrorResponse{
				Message: "Unable to process loan application",
				Reason:  "database error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mocks.NewLoanService(t)
			h := NewLoanHandler(mockService)

			if tt.mockSetup != nil {
				tt.mockSetup(mockService)
			}

			e := echo.New()
			e.Validator = NewCustomValidator()

			req := httptest.NewRequest(
				http.MethodPost,
				"/api/v1/loans",
				strings.NewReader(tt.body),
			)
			req.Header.Set(
				echo.HeaderContentType,
				echo.MIMEApplicationJSON,
			)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			handlerErr := h.ApplyLoan(c)
			assert.NoError(t, handlerErr)
			assert.Equal(t, tt.wantStatus, rec.Code)

			expectedBody, marshalErr := json.Marshal(tt.wantBody)
			assert.NoError(t, marshalErr)
			assert.JSONEq(t, string(expectedBody), rec.Body.String())
			t.Log("expectedBody :", expectedBody)
			t.Log("rec:", rec.Body)
			t.Log("recString:", rec.Body.String())

			//difference btw Equal and JSONEq
		})
	}
}
