package container

import (
	"microservices/customer-service-onboarding/internal/repository"
	"microservices/customer-service-onboarding/internal/service"
	"net/http"

	"gitlab.com/pos-alfa-microservices-go/core/config"
)

const queueName = "customer-service-opened"

type Container struct {
	AppConfig *config.AppConfig

	Service service.Service
}

func NewContainer(appConfig *config.AppConfig) *Container {
	return &Container{
		AppConfig: appConfig,
	}
}

func (c *Container) Start() error {

	repository := repository.NewRepositoryImpl(http.Client{})
	c.Service = service.NewServiceImpl(repository)

	return nil
}
