package container

import (
	"microservices/ticket-process/internal/repository"
	service "microservices/ticket-process/internal/service/message"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"

	"gitlab.com/pos-alfa-microservices-go/core/config"
	"gitlab.com/pos-alfa-microservices-go/core/database"
)

type Container struct {
	AppConfig *config.AppConfig

	MessageService service.MessageService
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
	c.MessageService = service.NewServiceImpl(repository)

	rabbitClient, err := rabbitmq.StartRabbitClient(c.AppConfig)
	if err != nil {
		return err
	}

	messagePublisher := rabbitmq.NewRabbitPublisher(rabbitClient)
	c.MessageService = service.NewServiceImpl(messagePublisher)

	return nil
}
