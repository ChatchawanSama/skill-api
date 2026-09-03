package repositories

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fusic/skill-api/internal/constant"
	"github.com/fusic/skill-api/internal/models"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDatabase(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, sqlMock, err := sqlmock.New()
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = sqlDB.Close()
	})

	gormDB, err := gorm.Open(
		mysql.New(mysql.Config{
			Conn:                      sqlDB,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{},
	)
	require.NoError(t, err)

	return gormDB, sqlMock
}

func TestLoanRepository_ApplyLoan_Success(t *testing.T) {
	gormDB, sqlMock := setupMockDatabase(t)
	repository := NewLoanRepository(gormDB)

	application := models.LoanApplication{
		ApplicationID: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		FullName:      "John Doe",
		MonthlyIncome: 10000,
		LoanAmount:    120000,
		LoanPurpose:   "home",
		Age:           20,
		PhoneNumber:   "0812345678",
		Email:         "john@example.com",
		Eligible:      true,
		Reason:        constant.ReasonEligible,
		Timestamp: time.Date(
			2026, time.September, 3,
			10, 30, 0, 0,
			time.Local,
		),
	}

	expectedSQL := regexp.QuoteMeta(
		"INSERT INTO `loan_applications` " +
			"(`application_id`,`full_name`,`monthly_income`,`loan_amount`," +
			"`loan_purpose`,`age`,`phone_number`,`email`,`eligible`,`reason`,`timestamp`) " +
			"VALUES (?,?,?,?,?,?,?,?,?,?,?)",
	)

	sqlMock.ExpectBegin()

	sqlMock.
		ExpectExec(expectedSQL).
		WithArgs(
			application.ApplicationID,
			application.FullName,
			application.MonthlyIncome,
			application.LoanAmount,
			application.LoanPurpose,
			application.Age,
			application.PhoneNumber,
			application.Email,
			application.Eligible,
			application.Reason,
			application.Timestamp,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	sqlMock.ExpectCommit()

	err := repository.ApplyLoan(application)

	require.NoError(t, err)
	require.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestLoanRepository_ApplyLoan_Error(t *testing.T) {
	gormDB, sqlMock := setupMockDatabase(t)
	repository := NewLoanRepository(gormDB)

	application := models.LoanApplication{
		ApplicationID: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		FullName:      "John Doe",
		MonthlyIncome: 10000,
		LoanAmount:    120000,
		LoanPurpose:   "home",
		Age:           20,
		PhoneNumber:   "0812345678",
		Email:         "john@example.com",
		Eligible:      true,
		Reason:        constant.ReasonEligible,
		Timestamp: time.Date(
			2026, time.September, 3,
			10, 30, 0, 0,
			time.Local,
		),
	}

	expectedError := errors.New("failed to insert loan application")

	expectedSQL := regexp.QuoteMeta(
		"INSERT INTO `loan_applications` " +
			"(`application_id`,`full_name`,`monthly_income`,`loan_amount`," +
			"`loan_purpose`,`age`,`phone_number`,`email`,`eligible`,`reason`,`timestamp`) " +
			"VALUES (?,?,?,?,?,?,?,?,?,?,?)",
	)

	sqlMock.ExpectBegin()

	sqlMock.
		ExpectExec(expectedSQL).
		WithArgs(
			application.ApplicationID,
			application.FullName,
			application.MonthlyIncome,
			application.LoanAmount,
			application.LoanPurpose,
			application.Age,
			application.PhoneNumber,
			application.Email,
			application.Eligible,
			application.Reason,
			application.Timestamp,
		).
		WillReturnError(expectedError)

	sqlMock.ExpectRollback()

	err := repository.ApplyLoan(application)

	require.ErrorIs(t, err, expectedError)
	require.NoError(t, sqlMock.ExpectationsWereMet())
}
