package config

import (
	coreConfig "gitlab.com/pos-alfa-microservices-go/core/config"
)

type AppConfig struct {
	RabbitMQ  coreConfig.RabbitMQ
	Customers coreConfig.API
	Products  coreConfig.API
	JWT       coreConfig.JWT
}
