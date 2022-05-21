package container

import (
	"microservices/ticket-process/internal/repository"

	"github.com/pkg/errors"

	consumer "microservices/ticket-process/internal/message"
	databaseService "microservices/ticket-process/internal/service/database"
	message "microservices/ticket-process/internal/service/message"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	"gitlab.com/pos-alfa-microservices-go/core/database"

	"gitlab.com/pos-alfa-microservices-go/core/config"
)

type Container struct {
	AppConfig *config.AppConfig

	ServiceImplMessage  message.ServiceMessage
	ServiceImplDatabase databaseService.ServiceDatabase
}

func NewContainer(appConfig *config.AppConfig) *Container {
	return &Container{
		AppConfig: appConfig,
	}
}

func (c *Container) Start() error {
	pool, err := database.StartPool(c.AppConfig)
	if err != nil {
		return err
	}

	databaseManager := database.NewDatabaseManagerImpl(pool)
	if err != nil {
		return err
	}

	repository := repository.NewRepositoryImpl(databaseManager)
	c.ServiceImplDatabase = databaseService.NewServiceImpl(repository)

	rabbitClient, err := rabbitmq.StartRabbitClient(c.AppConfig)
	if err != nil {
		return err
	}

	messagePublisher := rabbitmq.NewRabbitPublisher(rabbitClient)
	c.ServiceImplMessage = message.NewServiceImpl(messagePublisher)

	receivedQueueName := "ticket-pending"

	ticketsProcessor := consumer.NewMessageTicketProcessor(receivedQueueName, rabbitClient, c.ServiceImplDatabase, c.ServiceImplMessage)
	if err := ticketsProcessor.StartConsume(); err != nil {
		return errors.Wrap(err, "fail on start consume for "+receivedQueueName)
	}

	return nil
}
