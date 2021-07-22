package controller

import (
	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/service"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/labstack/echo/v4"
)

type HealthController struct {
	service *service.HealthService
}

func NewHealthController() *HealthController {
	return &HealthController{
		service: service.NewHealthService(repository.NewHealthRepository()),
	}
}

// NOTE: Currently Agent-ID is currently being used for both orginization and agent ids.
// Eventually this value (ie, orginization id) will be pulled from the auth. token.

// GetHealth returns all health data
func (controller *HealthController) GetHealth(c echo.Context) error {
	res := controller.service.GetHealth(
		c.Response().Header().Get("X-Request-ID"),
	)
	return c.JSON(res.StatusCode, res)
}

// GetHealthByAgentId returns all health data for an agent
func (controller *HealthController) GetHealthByAgentId(c echo.Context) error {
	res := controller.service.GetHealthByAgentID(
		c.Response().Header().Get("X-Request-ID"), c.Param("agent-id"),
	)
	return c.JSON(res.StatusCode, res)
}

// AddHealth adds health data for an agent
func (controller *HealthController) PostHealth(c echo.Context) error {
	health := new(types.Health)

	if err := c.Bind(health); err != nil {
		return c.JSON(400, types.StandardResponse{
			StatusCode: 400,
			Error:      "failed to bind health data",
			Success:    false,
			Data:       nil,
		})
	}

	res := controller.service.AddHealth(
		c.Response().Header().Get("X-Request-ID"),
		c.Request().Header.Get("Agent-ID"), health,
	)
	return c.JSON(res.StatusCode, res)
}
