package controller

import (
	"strconv"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/service"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

// HealthController provides a health service to interact with
type HealthController struct {
	service     service.IHealthService
	hostService service.IHostService
}

// NewHealthController returns a new HealthController with the service/repository initialized
func NewHealthController() *HealthController {
	healthService := service.NewHealthService(repository.NewHealthRepository(), repository.NewHostRepository())
	return &HealthController{
		service:     healthService,
		hostService: service.NewHostService(repository.NewHostRepository(), healthService),
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

// GetLatestHealthDataForAgentByAgentId returns all health data for an agent since a given time
func (controller *HealthController) GetLatestHealthDataForAgentByAgentId(c echo.Context) error {
	since, err := strconv.Atoi(c.Param("since"))
	if err != nil {
		return c.JSON(400, types.StandardResponse{
			StatusCode: 400,
			Error:      "failed to parse since timestamp",
			Success:    false,
			Data:       nil,
		})
	}

	res := controller.service.GetLatestHealthDataByAgentID(c.Response().Header().Get("X-Request-ID"),
		c.Param("agent-id"),
		int64(since),
	)
	return c.JSON(200, types.StandardResponse{
		Data:       res,
		StatusCode: 200,
		Success:    true,
	})
}

// GetLatestHealthDataForAgents returns all health data for all agents since a given time
func (controller *HealthController) GetLatestHealthDataForAgents(c echo.Context) error {
	since, err := strconv.Atoi(c.Param("since"))
	if err != nil {
		return c.JSON(400, types.StandardResponse{
			StatusCode: 400,
			Error:      "failed to parse since timestamp",
			Success:    false,
			Data:       nil,
		})
	}

	res := controller.service.GetLatestHealthDataForAgents(c.Response().Header().Get("X-Request-ID"),
		int64(since),
	)
	return c.JSON(200, types.StandardResponse{
		Data:       res,
		StatusCode: 200,
		Success:    true,
	})
}

// GetHealthWS returns all health data via websockets
func (controller *HealthController) GetHealthWS(c echo.Context) error {
	ws_delay := utils.GetVariable(consts.DATA_WEBSOCKET_DELAY)
	delay, err := strconv.Atoi(ws_delay)
	if err != nil {
		delay = 30
	}
	time.Sleep(time.Second * time.Duration(delay))

	log := logger.Instance()
	requestID := c.Response().Header().Get("X-Request-ID")
	lastCheck := utils.GetMinimumLastHealthPacketTime(time.Now(), 2)

	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			res := controller.service.GetLatestHealthDataForAgents(requestID, lastCheck)
			websocket.JSON.Send(ws, res)
			if !res.Success {
				log.Error(res.Error)
			}

			lastCheck = time.Now().UTC().UnixNano()
			time.Sleep(time.Second * time.Duration(delay))
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
