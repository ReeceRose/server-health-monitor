package router

import (
	"github.com/PR-Developers/server-health-monitor/internal/api/controller"

	"github.com/labstack/echo/v4"
)

// Setup handles all the public/private server routes
func Setup(e *echo.Echo) {
	// General setup
	health := controller.NewHealthController()
	host := controller.NewHostController()

	// Public routes

	// TODO: setup auth middleware

	// Private routes
	e.GET("/api/v1/health/", func(c echo.Context) error { return health.GetHealth(c) })
	e.GET("/api/v1/health/:agent-id", func(c echo.Context) error { return health.GetHealthByAgentId(c) })
	e.POST("/api/v1/health/", func(c echo.Context) error { return health.PostHealth(c) })
	e.GET("/api/v1/host/", func(c echo.Context) error { return host.GetHosts(c) })
	e.GET("/api/v1/host/:agent-id", func(c echo.Context) error { return host.GetHostById(c) })
	e.POST("/api/v1/host/", func(c echo.Context) error { return host.PostHost(c) })

	// Websockets
	// e.GET("/ws/v1/health/", func(c echo.Context) error { return health.GetHealth() })
}
