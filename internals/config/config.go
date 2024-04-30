package config

import (
	"github.com/caarlos0/env/v10"
	log "github.com/robowealth-mutual-fund/stdlog"
)

type Config struct {
	AppName      string `env:"APP_NAME" envDefault:"blueprint-service"`
	Environment  string `env:"ENV" envDefault:"develop"`
	PlatformName string `env:"PLATFORM_NAME" envDefault:"Blueprint Service"`

	Database   Database
	Server     Server
	Trace      Trace
	Connection Connection
	Redis      Redis
}

func New() Config {
	c := Config{}

	if err := env.Parse(&c); err != nil {
		log.Error("Load configuration failed", err)
		panic(err)
	}

	return c
}
