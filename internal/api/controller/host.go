package controller

import (
	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/service"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/labstack/echo/v4"
)

// HostController provides a host service to interact with
type HostController struct {
	service service.IHostService
}

// NewHostController returns a new HostController with the service/repository initialized
func NewHostController() *HostController {
	return &HostController{
		service: service.NewHostService(repository.NewHostRepository(), repository.NewHealthRepository()),
	}
}

// GetHost returns all hosts associated with an account
func (controller *HostController) GetHosts(c echo.Context) error {
	res := controller.service.GetHosts(
		c.Response().Header().Get("X-Request-ID"),
	)
	return c.JSON(res.StatusCode, res)
}

// GetHostById returns information about the host
func (controller *HostController) GetHostById(c echo.Context) error {
	res := controller.service.GetHostByID(
		c.Response().Header().Get("X-Request-ID"), c.Param("agent-id"),
	)
	return c.JSON(res.StatusCode, res)
}

// PostHost adds host data or updates existing host
func (controller *HostController) PostHost(c echo.Context) error {
	host := new(types.Host)

	if err := c.Bind(host); err != nil {
		return c.JSON(400, types.StandardResponse{
			StatusCode: 400,
			Error:      "failed to bind host data",
			Success:    false,
			Data:       nil,
		})
	}

	res := controller.service.AddHost(
		c.Response().Header().Get("X-Request-ID"),
		c.Request().Header.Get("Agent-ID"), host,
	)
	return c.JSON(res.StatusCode, res)
}
