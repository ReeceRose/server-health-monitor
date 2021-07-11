package main

import (
	"fmt"
	"net/http"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log := logger.Instance()
	log.Info("Server Health Monitor API")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Server Health Monitor API")
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	port := utils.GetVariable(consts.API_PORT, "")
	port = fmt.Sprintf(":%s", port)

	e.Logger.Fatal(e.Start(port))
}
