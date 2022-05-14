package message

import (
	"context"
	"encoding/json"
	"fmt"

	"microservices/order/internal/service"
	"microservices/order/pkg/model"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	coreLog "gitlab.com/pos-alfa-microservices-go/core/log"
)

type OrderProcessor interface {
	StartConsume() error
}

type MessageOrderProcessor struct {
	queue        string
	client       *rabbitmq.RabbitClient
	orderService service.Service
}

func NewMessageOrderProcessor(queue string, client *rabbitmq.RabbitClient, orderService service.Service) OrderProcessor {
	return &MessageOrderProcessor{
		queue:        queue,
		client:       client,
		orderService: orderService,
	}
}

func (p MessageOrderProcessor) StartConsume() error {
	consumer := rabbitmq.NewRabbitConsumer(p.client, "orders-consumer")
	return consumer.Consume(p.queue, func(body []byte) error {
		bodyOrder := model.Order{}
		if err := json.Unmarshal(body, &bodyOrder); err != nil {
			return fmt.Errorf("fail to unmarshal order %w", err)
		}

		order, err := p.orderService.Create(context.Background(), &bodyOrder)
		if err != nil {
			return fmt.Errorf("fail to Create order %w", err)
		}

		coreLog.Logger.Infof("order processed: %v", order)

		return nil
	})
}
