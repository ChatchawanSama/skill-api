package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheckHandler interface {
	HealthCheck(c echo.Context) error
	ReadinessCheck(c echo.Context) error
}

type healthCheckHandler struct{}

func NewHealthCheckHandler() HealthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) HealthCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *healthCheckHandler) ReadinessCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}