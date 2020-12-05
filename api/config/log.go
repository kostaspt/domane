package config

import "github.com/rs/zerolog/log"

func SetupLog() {
	log.Logger = log.With().Caller().Logger()
}
