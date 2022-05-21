package container

import (
	"microservices/ticket-create/internal/client"
	message "microservices/ticket-create/internal/service"
	"net/http"
	"time"

	"gitlab.com/pos-alfa-microservices-go/core/auth"
	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"

	"gitlab.com/pos-alfa-microservices-go/core/config"
	coreConfig "gitlab.com/pos-alfa-microservices-go/core/config"
	coreClient "gitlab.com/pos-alfa-microservices-go/core/http/client"
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

	httpClient := http.Client{Timeout: time.Duration(1) * time.Second}
	restClient := coreClient.NewRestClient(httpClient, true)

	TicketClient := client.NewHttpticketClient(restClient, c.AppConfig.Ticket.URL)
	tokenManager := auth.NewJWTTokenManager(&coreConfig.AppConfig{JWT: c.AppConfig.JWT})

	messagePublisher := rabbitmq.NewRabbitPublisher(rabbitClient)
	c.ServiceImplMessage = message.NewServiceImpl(messagePublisher)

	return nil
}
