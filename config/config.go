package config

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port             int    `env:"PORT" envDefault:"8081"`
	AuthGRPCEndpoint string `env:"AuthGRPCEndpoint" envDefault:"localhost:1323"`
}

func Load() (cfg Config) {
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	return
}
