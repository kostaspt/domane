package rest

import (
	"github.com/google/wire"
	"github.com/kostaspt/domane/api/config"
)

var ProviderSet = wire.NewSet(NewHandler)

type Handler struct {
	config *config.Config
}

func NewHandler(c *config.Config) *Handler {
	return &Handler{
		config: c,
	}
}
