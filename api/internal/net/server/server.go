package server

import (
	"fmt"
	"os"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	elog "github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"github.com/ziflex/lecho/v2"

	HTTPHandler "github.com/kostaspt/domane/api/internal/net/http"
	WSHandler "github.com/kostaspt/domane/api/internal/net/websockets"
)

func Start() error {
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://domane.io",
		},
	}))
	e.Logger = lecho.New(
		os.Stdout,
		lecho.WithLevel(elog.DEBUG),
		lecho.WithTimestamp(),
		lecho.WithCaller(),
	)
	e.HideBanner = true

	// Prometheus
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	// HTTP
	http := HTTPHandler.New()
	e.GET("/", http.RootIndex)
	e.POST("/domains/extensions", http.DomainsExtensions)
	e.POST("/domains/similars", http.DomainsSimilars)

	// WebSockets
	ws := WSHandler.New()
	e.GET("/ws/availabilities", ws.Availabilities)

	// Start server
	return e.Start(fmt.Sprintf(":%d", viper.GetInt("port")))
}
