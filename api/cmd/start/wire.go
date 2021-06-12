// +build wireinject

package main

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/kostaspt/domane/api/config"
	"github.com/kostaspt/domane/api/internal/net/rest"
	"github.com/kostaspt/domane/api/internal/net/server"
	"github.com/kostaspt/domane/api/internal/net/ws"
)

func initApp(context.Context, *config.Config) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			rest.ProviderSet,
			server.ProviderSet,
			ws.ProviderSet,
			newApp,
		),
	)
}
