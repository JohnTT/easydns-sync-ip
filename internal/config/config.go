package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Host           string `env:"EASYDNS_HOST,required"`
	Token          string `env:"EASYDNS_TOKEN,required"`
	Key            string `env:"EASYDNS_KEY,required"`
	Domain         string `env:"EASYDNS_DOMAIN,required"`
	RecordHost     string `env:"EASYDNS_RECORD_HOST,required"`
	UpdateInterval uint   `env:"UPDATE_INTERVAL" envDefault:"600"`
}

var cfg Config

func init() {
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
}

func Get() *Config {
	return &cfg
}
