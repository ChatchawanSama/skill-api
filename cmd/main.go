package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fusic/skill-api/config"
	"github.com/fusic/skill-api/internal/handlers/rest"
	"github.com/fusic/skill-api/internal/repositories/adaptor/mysql"
	"github.com/fusic/skill-api/internal/services"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	db, err := mysql.NewDB(cfg.Secrets.MySQLDSN())
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	repo := mysql.NewLoanRepository(db)
	svc := services.NewLoanService(repo)
	loanHandler := rest.NewLoanHandler(svc)
	healthHandler := rest.NewHealthCheckHandler()

	httpServer := rest.NewHttpServer(cfg, echo.New(), healthHandler, loanHandler)

	go func() {
		if err := httpServer.Start(cfg.Server.Address); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Server().Shutdown(ctx); err != nil {
		log.Printf("shutdown: %v", err)
	}
}
