package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Host          string `env:"EASYDNS_HOST,required"`
	Token         string `env:"EASYDNS_TOKEN,required"`
	Key           string `env:"EASYDNS_KEY,required"`
	TickerSeconds uint   `env:"TickerSeconds" envDefault:"30"`
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
