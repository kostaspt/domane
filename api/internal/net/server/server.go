package server

import (
	"fmt"
	"time"

	transphttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/kostaspt/domane/api/config"
	"github.com/kostaspt/domane/api/internal/net/rest"
	"github.com/kostaspt/domane/api/internal/net/ws"
	"github.com/kostaspt/domane/api/pkg/logger"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ProviderSet = wire.NewSet(NewServer)

func NewServer(c *config.Config, rh *rest.Handler, wh *ws.Handler) *transphttp.Server {
	opts := []transphttp.ServerOption{
		transphttp.Timeout(10 * time.Second),
		transphttp.Logger(logger.NewWrapper()),
	}

	if c.Port > 0 {
		opts = append(opts, transphttp.Address(fmt.Sprintf("[::]:%d", c.Port)))
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://" + c.Domain,
			"https://" + c.Domain,
		},
	}))

	// Prometheus
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	// Rest
	e.GET("/", rh.RootIndex)
	e.POST("/domains/extensions", rh.DomainsExtensions)
	e.POST("/domains/similars", rh.DomainsSimilars)

	// WebSockets
	e.GET("/ws/whois", wh.Whois)

	srv := transphttp.NewServer(opts...)

	srv.HandlePrefix("/", e)

	return srv
}
