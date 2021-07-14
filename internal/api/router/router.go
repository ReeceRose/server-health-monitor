package router

import (
	"github.com/PR-Developers/server-health-monitor/internal/api/controller"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	// General setup
	health := controller.NewHealthController()

	// Public routes

	// TODO: setup auth middleware

	// Private routes
	e.GET("/api/v1/health/", func(c echo.Context) error { return health.GetAllHealth(c) })
	e.GET("/api/v1/health/:id", func(c echo.Context) error { return health.GetHealthByServerId(c) })
	e.POST("/api/v1/health/", func(c echo.Context) error { return health.PostHealth(c) })
}
