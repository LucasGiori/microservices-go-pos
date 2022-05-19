package container

import (
	message "microservices/ticket-create/internal/service"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"

	"gitlab.com/pos-alfa-microservices-go/core/config"
)

type Container struct {
	AppConfig *config.AppConfig

	ServiceImplMessage message.ServiceMessage
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
	c.ServiceImplMessage = message.NewServiceImpl(messagePublisher)

	return nil
}
