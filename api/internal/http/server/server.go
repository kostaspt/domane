package server

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	elog "github.com/labstack/gommon/log"
	"github.com/ziflex/lecho/v2"

	handler "github.com/kostaspt/domane/api/internal/http/handler/v1"
)

func Start() error {
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Logger = lecho.New(
		os.Stdout,
		lecho.WithLevel(elog.DEBUG),
		lecho.WithTimestamp(),
		lecho.WithCaller(),
	)
	e.HideBanner = true

	// Routes
	h := handler.New()

	e.GET("/", h.RootIndex)
	e.POST("/domains/extensions", h.DomainsExtensions)

	// Start server
	return e.Start(":4000")
}
