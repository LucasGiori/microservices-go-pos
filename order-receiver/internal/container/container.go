package container

import (
	"microservices/order-receiver/internal/service"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"

	"gitlab.com/pos-alfa-microservices-go/core/config"
)

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
	rabbitClient, err := rabbitmq.StartRabbitClient(c.AppConfig)
	if err != nil {
		return err
	}

	messagePublisher := rabbitmq.NewRabbitPublisher(rabbitClient)
	c.Service = service.NewServiceImpl(messagePublisher)

	return nil
}
