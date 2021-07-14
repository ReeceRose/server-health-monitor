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

// NOTE: Currently Agent-UUID is currently being used for both orginization and agent ids.
// Eventually this value will be pulled from the auth. token.

// GetHealthByServerId returns all health data for a given organizations server
func (controller *HealthController) GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.GetHealth(c.Request().Header.Get("Agent-UUID")))
}

// AddHealth adds health data for a given organizations server
func (controller *HealthController) PostHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, controller.service.AddHealth(c.Request().Header.Get("Agent-UUID")))
}
