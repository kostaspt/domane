package main

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/kostaspt/domane/api/config"
	"github.com/kostaspt/domane/api/pkg/logger"
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

	c, err := config.New()
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
