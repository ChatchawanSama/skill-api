package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fusic/skill-api/internal/config"
	"github.com/fusic/skill-api/internal/pkg/configurer"
	"github.com/fusic/skill-api/internal/pkg/logger"
	"github.com/fusic/skill-api/internal/pkg/mysql_database"
	"github.com/fusic/skill-api/internal/repositories"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/fusic/skill-api/internal/handlers/rest"
	"github.com/fusic/skill-api/internal/services"
)

func main() {
	// create root context
	ctx, _ := context.WithCancel(context.Background())

	// load config
	conf := configurer.LoadConfig()

	if err := config.SetTimeZone(conf.Server.TimeZone); err != nil {
		log.Fatal("fail set timezone")
	}
	// load secret env
	configurer.LoadSecret(&conf.Secret)

	fmt.Println("Config loaded:", conf)

	// database
	db, err := mysql_database.Connect(&mysql_database.MysqlConfig{
		Name:                conf.DB.Name,
		SSLMode:             conf.DB.SSLMode,
		MaxOpenConns:        &conf.DB.MaxOpenConnection,
		MaxIdleConns:        &conf.DB.MaxIdleConnection,
		ConnMaxLifetimeHour: conf.DB.ConnectionMaxLifetime,
		MysqlHost:           conf.Secret.DBHost,
		MysqlPort:           conf.Secret.DBPort,
		MysqlUser:           conf.Secret.DBUser,
		MysqlPassword:       conf.Secret.DBPassword,
		Loc:                 time.Local,
	})

	if err != nil {
		logger.Fatal(ctx, "fail connect database", zap.Error(err))
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			logger.Warn(ctx, "fail get sql.DB", zap.Error(err))
			return
		}
		if err := sqlDB.Close(); err != nil {
			logger.Warn(ctx, "fail close database", zap.Error(err))
		}
	}()

	// if err := db.AutoMigrate(&models.LoanApplication{}); err != nil {
	// 	logger.Fatal(ctx, "fail to migrate loan applications", zap.Error(err))
	// }
	// TODO : Init repositories, services, handlers, and start the HTTP server.
	e := echo.New()
	e.Validator = rest.NewCustomValidator()

	loanRepository := repositories.NewLoanRepository(db)
	loanService := services.NewLoanService(loanRepository)
	loanHandler := rest.NewLoanHandler(loanService)

	healthCheckHandler := rest.NewHealthCheckHandler()

	httpServer := rest.NewHttpServer(
		conf,
		e,
		healthCheckHandler,
		loanHandler,
	)

	if err := httpServer.Start(conf.Server.Address); err != nil {
		logger.Info(ctx, "fail to start HTTP server", zap.Error(err))
	}

}
