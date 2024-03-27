package config

import "github.com/caarlos0/env/v10"

type environment string

func (e environment) String() string {
	return string(e)
}

// This doesn't restrict comparisons to environments to these options but it
// helps to encourage it
const (
	Local   environment = "local"
	Dev     environment = "dev"
	Staging environment = "staging"
	Prod    environment = "prod"
)

type config struct {
	// For example, "local", "dev", "prod"...
	Environment environment `env:"ENVIRONMENT" envDefault:"local"`
	// Where to store the log file
	LogFile string `env:"LOG_FILE" envDefault:"logs/log.log"`
	// What port to listen on
	Port int `env:"PORT" envDefault:"3000"`
}

var Config config

func ParseConfig() error {
	Config = config{}
	if err := env.Parse(&Config); err != nil {
		return err
	}
	return nil
}
