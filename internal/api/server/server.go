package server

import (
	"fmt"

	"github.com/PR-Developers/server-health-monitor/internal/api/router"
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

func New() Server {
	return &EchoServer{
		Instance: echo.New(),
	}
}

// Start the web server
func (s *EchoServer) Start() {
	e := s.Instance

	// Currently this server is only used for the core API so the logic below
	// is fine here. If we need to expand this to be used in multiple locations
	// the below can be done via first-class functions
	e.Pre(middleware.HTTPSRedirect())

	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logger.Instance().Logger().Writer(),
	}))

	router.Setup(e)

	port := utils.GetVariable(consts.API_PORT)
	port = fmt.Sprintf(":%s", port)

	certDir := utils.GetVariable(consts.CERT_DIR)
	e.Logger.Fatal(e.StartTLS(port,
		fmt.Sprintf("%s/%s", certDir, utils.GetVariable(consts.API_CERT)),
		fmt.Sprintf("%s/%s", certDir, utils.GetVariable(consts.API_KEY))))
}
