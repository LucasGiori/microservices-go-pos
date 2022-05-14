package container

import (
	"microservices/order/internal/message"
	"microservices/order/internal/repository"
	"microservices/order/internal/service"

	"github.com/pkg/errors"
	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	"gitlab.com/pos-alfa-microservices-go/core/config"
	"gitlab.com/pos-alfa-microservices-go/core/database"
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
	pool, err := database.StartPool(c.AppConfig)
	if err != nil {
		return err
	}

	databaseManager := database.NewDatabaseManagerImpl(pool)
	rabbitClient, err := rabbitmq.StartRabbitClient(c.AppConfig)
	if err != nil {
		return err
	}

	repository := repository.NewRepositoryImpl(databaseManager)
	c.Service = service.NewServiceImpl(repository)

	queueName := "orders-pending"
	ordersProcessor := message.NewMessageOrderProcessor(queueName, rabbitClient, c.Service)
	if err := ordersProcessor.StartConsume(); err != nil {
		return errors.Wrap(err, "fail on start consume for "+queueName)
	}

	return nil
}
