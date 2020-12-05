package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func SetupViper() {
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal().Err(err)
		}
	}
}
