package rest

import (
	"github.com/fusic/skill-api/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	config             *config.Config
	server             *echo.Echo
	healthCheckHandler HealthCheckHandler
	loanHandler        LoanHandler
}

func NewHttpServer(
	cfg *config.Config,
	server *echo.Echo,
	healthCheckHandler HealthCheckHandler,
	loanHandler LoanHandler,
) *HttpServer {
	h := &HttpServer{
		config:             cfg,
		server:             server,
		healthCheckHandler: healthCheckHandler,
		loanHandler:        loanHandler,
	}
	h.InitRoutes()
	return h
}

func (h *HttpServer) InitRoutes() {
	e := h.server

	e.GET("/health", h.healthCheckHandler.HealthCheck)
	e.GET("/ready", h.healthCheckHandler.ReadinessCheck)

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	v1 := e.Group("/api/v1")
	v1.POST("/loans", h.loanHandler.Apply)
	v1.GET("/loans", h.loanHandler.List)
	v1.GET("/loans/:applicationId", h.loanHandler.GetByID)
}

func (h *HttpServer) Start(address string) error {
	return h.server.Start(address)
}

func (h *HttpServer) Server() *echo.Echo {
	return h.server
}
