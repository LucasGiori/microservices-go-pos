package container

import (
	"microservices/product/internal/repository"
	"microservices/product/internal/service"

	"gitlab.com/pos-alfa-microservices-go/core/config"
	"gitlab.com/pos-alfa-microservices-go/core/database"
)

const queueName = "orders-aggregate"

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

	repository := repository.NewRepositoryImpl(databaseManager)
	c.Service = service.NewServiceImpl(repository)

	return nil
}
