package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Domain  string
	Version string
	Port    uint16
}

func New(port uint16) (c *Config, err error) {
	v := initViper()

	if port > 0 {
		v.Set("port", port)
	}

	if err = v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msg("No config file found in search paths, using default values")
		}
	}

	err = v.Unmarshal(&c)
	return
}

func initViper() *viper.Viper {
	v := viper.NewWithOptions(viper.KeyDelimiter("_"))

	v.SetConfigFile(".env")
	v.AutomaticEnv()

	viper.SetDefault("domain", "domane.io")
	viper.SetDefault("version", "v1.0")
	viper.SetDefault("port", 4000)

	return v
}
