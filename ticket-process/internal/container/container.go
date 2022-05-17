package container

import (
	"microservices/ticket-process/internal/repository"

	databaseService "microservices/ticket-process/internal/service/message"
	message "microservices/ticket-process/internal/service/message"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	"gitlab.com/pos-alfa-microservices-go/core/database"

	"gitlab.com/pos-alfa-microservices-go/core/config"
)

type Container struct {
	AppConfig *config.AppConfig

	Service message.Service
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
	repository := repository.NewRepositoryImpl(databaseManager)
	c.Service = databaseService.NewServiceImpl(repository)

	rabbitClient, err := rabbitmq.StartRabbitClient(c.AppConfig)
	if err != nil {
		return err
	}

	messagePublisher := rabbitmq.NewRabbitPublisher(rabbitClient)
	c.Service = message.NewServiceImpl(messagePublisher)

	return nil
}
