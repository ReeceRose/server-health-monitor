package server

import (
	"fmt"
	"net/http"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server interface {
	Start()
}

type EchoServer struct {
	Instance *echo.Echo
}

var (
	_ Server = (*EchoServer)(nil)
)

func New() *EchoServer {
	return &EchoServer{
		Instance: echo.New(),
	}
}

func (s *EchoServer) Start() {
	e := s.Instance
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Server Health Monitor API")
	})

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logger.Instance().GenericLogger.Writer(),
	}))

	port := utils.GetVariable(consts.API_PORT, "")
	port = fmt.Sprintf(":%s", port)

	e.Logger.Fatal(e.Start(port))
}
