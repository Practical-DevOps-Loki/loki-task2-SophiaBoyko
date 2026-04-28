package main

import (
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"webapp/pkg/config"
)

func main() {
	port := config.GetEnv("PORT", "3000")

	// Create log directory and file
	if err := os.MkdirAll("/var/log/webapp", 0755); err != nil {
		// fallback: log to stdout only
	}

	logFile, err := os.OpenFile("/var/log/webapp/access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	e := echo.New()
	e.HideBanner = true

	if err == nil {
		// Write logs to both file and stdout
		e.Logger.SetOutput(io.MultiWriter(os.Stdout, logFile))
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: io.MultiWriter(os.Stdout, logFile),
		}))
	} else {
		e.Use(middleware.Logger())
	}

	e.Use(middleware.Recover())

	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		return c.File("public/views/webapp.html")
	})

	e.Start(":" + port)
}