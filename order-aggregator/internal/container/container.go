package container

import (
	"microservices/order-aggregator/internal/client"
	"microservices/order-aggregator/internal/config"
	"microservices/order-aggregator/internal/message"
	"microservices/order-aggregator/internal/service"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/pos-alfa-microservices-go/core/auth"
	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	coreConfig "gitlab.com/pos-alfa-microservices-go/core/config"
	coreClient "gitlab.com/pos-alfa-microservices-go/core/http/client"
)

type Container struct {
	AppConfig *config.AppConfig

	ProcessOrder service.ProcessOrder
}

func NewContainer(appConfig *config.AppConfig) *Container {
	return &Container{
		AppConfig: appConfig,
	}
}

func (c *Container) Start() error {
	rabbitClient, err := rabbitmq.StartRabbitClient(&coreConfig.AppConfig{
		RabbitMQ: c.AppConfig.RabbitMQ,
	})
	if err != nil {
		return err
	}

	httpClient := http.Client{Timeout: time.Duration(1) * time.Second}
	restClient := coreClient.NewRestClient(httpClient, true)

	customerClient := client.NewHttpCustomerClient(restClient, c.AppConfig.Customers.URL)
	productClient := client.NewHttpProductClient(restClient, c.AppConfig.Products.URL)
	tokenManager := auth.NewJWTTokenManager(&coreConfig.AppConfig{JWT: c.AppConfig.JWT})

	messagePublisher := rabbitmq.NewRabbitPublisher(rabbitClient)
	c.ProcessOrder = service.NewProcessOrderImpl(messagePublisher, customerClient, productClient, tokenManager)

	receivedQueueName := "orders-received"
	ordersProcessor := message.NewMessageOrderProcessor(receivedQueueName, rabbitClient, c.ProcessOrder)
	if err := ordersProcessor.StartConsume(); err != nil {
		return errors.Wrap(err, "fail on start consume for "+receivedQueueName)
	}

	return nil
}
