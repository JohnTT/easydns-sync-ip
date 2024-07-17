# Instructions

Ensure the following environment variables are set:
```
type Config struct {
	Host          string `env:"EASYDNS_HOST,required"`
	Token         string `env:"EASYDNS_TOKEN,required"`
	Key           string `env:"EASYDNS_KEY,required"`
	TickerSeconds uint   `env:"TickerSeconds" envDefault:"30"`
}
```
