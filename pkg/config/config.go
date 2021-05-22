package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	ServerPort     string `envconfig:"SERVER_PORT" default:"80"`
	DataSourceName string `envconfig:"DATA_SOURCE_NAME"`
	LogLevel       string `envconfig:"LOG_LEVEL" default:"debug"`
}

func New() (*Configuration, error) {
	var c Configuration
	err := envconfig.Process("myapp", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
