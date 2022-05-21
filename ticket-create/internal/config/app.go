package config

import (
	coreConfig "gitlab.com/pos-alfa-microservices-go/core/config"
)

type AppConfig struct {
	Ticket coreConfig.API
	JWT    coreConfig.JWT
}
