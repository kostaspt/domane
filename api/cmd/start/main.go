package main

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/kostaspt/domane/api/config"
	"github.com/kostaspt/domane/api/pkg/logger"
	"github.com/spf13/pflag"
)

var (
	Name    = "domane-api"
	Version string
)

func newApp(srv *http.Server) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Server(srv),
	)
}

func main() {
	logger.Setup(Name, Version)

	var port uint16
	pflag.Uint16VarP(&port, "port", "p", 4000, "Server's port")

	c, err := config.New(port)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a, cleanup, err := initApp(ctx, c)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err = a.Run(); err != nil {
		panic(err)
	}
}
