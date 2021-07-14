package controller

import (
	"net/http"

	"github.com/PR-Developers/server-health-monitor/internal/api/service"
	"github.com/labstack/echo/v4"
)

type HealthController struct {
	service *service.HealthService
}

func NewHealthController() *HealthController {
	return &HealthController{
		service: service.NewHealthService(),
	}
}

// GetAllHealthData returns all health data for a given organization
func (controller *HealthController) GetAllHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.GetHealth())
}

// GetHealthByServerId returns all health data for a given organizations server
func (controller *HealthController) GetHealthByServerId(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.GetHealthByServerId())
}

// AddHealth adds health data for a given organizations server
func (controller *HealthController) PostHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.AddHealth())
}
