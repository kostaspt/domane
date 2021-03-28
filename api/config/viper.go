package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func SetupViper() {
	initViper()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal().Err(err)
		}
	}
}

func initViper() {
	viper.AutomaticEnv()

	viper.SetDefault("VERSION", "v1.0")
	viper.SetDefault("DOMAIN", "domane.io")
}
