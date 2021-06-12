package logger

import (
	"github.com/rs/zerolog/log"
)

func Setup(name, version string) {
	log.Logger = log.
		With().
		Caller().
		Stack().
		Str("service", name).
		Str("version", version).
		Logger()
}
