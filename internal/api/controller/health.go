package controller

import (
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/service"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/labstack/echo/v4"
)

// HealthController provides a health service to interact with
type HealthController struct {
	service service.IHealthService
}

// NewHealthController returns a new HealthController with the service/repository initialized
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

// GetHealthWS returns all health data via websockets
func (controller *HealthController) GetHealthWS(c echo.Context) error {
	// res := controller.service.GetHealth(
	// 	c.Response().Header().Get("X-Request-ID"),
	// )
	// return c.JSON(res.StatusCode, res)
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}
			time.Sleep(time.Second * 20)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

// GetHealthByAgentId returns all health data for an agent
func (controller *HealthController) GetHealthByAgentId(c echo.Context) error {
	res := controller.service.GetHealthByAgentID(
		c.Response().Header().Get("X-Request-ID"), c.Param("agent-id"),
	)
	return c.JSON(res.StatusCode, res)
}

// PostHealth adds health data for an agent
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
