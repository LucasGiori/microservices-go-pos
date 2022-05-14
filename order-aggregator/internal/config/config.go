package config

import (
	coreConfig "gitlab.com/pos-alfa-microservices-go/core/config"

	"github.com/spf13/viper"
)

func Start() (*AppConfig, error) {
	if err := coreConfig.Read(); err != nil {
		return nil, err
	}

	appConfig := AppConfig{}
	viper.Unmarshal(&appConfig)

	return &appConfig, nil
}
