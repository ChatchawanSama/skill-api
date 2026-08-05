package rest

import (
	"github.com/fusic/skill-api/internal/config"
	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	config *config.AppConfig
	server *echo.Echo
	// TODO: Add handlers here, e.g. HealthCheckHandler, LoanHandler
}

func NewHttpServer(
	cfg *config.AppConfig,
	server *echo.Echo,
	// TODO: Dependency inject handlers here, e.g. HealthCheckHandler, LoanHandler
) *HttpServer {
	h := &HttpServer{
		config: cfg,
		server: server,
		// TODO: Initialize handlers here, e.g.: HealthCheckHandler, LoanHandler
	}
	h.InitRoutes()
	return h
}

func (h *HttpServer) InitRoutes() {
	e := h.server

	// TODO: Add routes here, e.g. health check, loan application, etc.

	// e.Use(middleware.Recover())
	// e.Use(middleware.Logger())

	v1 := e.Group("/api/v1")
	_ = v1
	// TODO: Add routes here, e.g. /loans
}

func (h *HttpServer) Start(address string) error {
	return h.server.Start(address)
}

func (h *HttpServer) Server() *echo.Echo {
	return h.server
}
