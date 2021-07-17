package controller

import (
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

// NOTE: Currently Agent-ID is currently being used for both orginization and agent ids.
// Eventually this value (ie, orginization id) will be pulled from the auth. token.

// GetHealthByServerId returns all health data for a given organizations server
func (controller *HealthController) GetHealthByServerId(c echo.Context) error {
	res := controller.service.GetHealth(c.Param("agent-id"))
	return c.JSON(res.StatusCode, res)
}

// AddHealth adds health data for a given organizations server
func (controller *HealthController) PostHealth(c echo.Context) error {
	res := controller.service.GetHealth(c.Request().Header.Get("Agent-ID"))
	return c.JSON(res.StatusCode, res)
}
