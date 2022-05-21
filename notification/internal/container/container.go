package container

import (
	"github.com/pkg/errors"
	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	"gitlab.com/pos-alfa-microservices-go/core/config"
	"microservices/notification/internal/message"
)

type Container struct {
	AppConfig *config.AppConfig
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

	queueName := "notification"

	notificationProcessor := message.NewNotificationConsumerProcessor(queueName, rabbitClient)
	if err := notificationProcessor.StartConsume(); err != nil {
		return errors.Wrap(err, "fail on start send notification for "+queueName)
	}

	return nil
}
