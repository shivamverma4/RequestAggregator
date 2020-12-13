package config

import (
	"github.com/op/go-logging"

	"requestaggregator/internal/configuration"
)

var config configuration.Config

const namespace = "requestaggregator"

func init() {
	config = configuration.Config{
		Namespace:          namespace,
		Deployment:         configuration.DEBUG,
		LogLevel:           logging.INFO,
		LogFilePath:        "../internal/logs/requestaggregator.log",
		RequestLogFilePath: "../internal/logs/request.log",
		Port:               ApiPort,
	}
}

func GetConfig() *configuration.Config {
	return &config
}
