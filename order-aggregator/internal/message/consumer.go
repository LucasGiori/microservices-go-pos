package message

import (
	"context"
	"encoding/json"
	"fmt"
	"microservices/order-aggregator/internal/service"
	"microservices/order-aggregator/pkg/model"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	coreLog "gitlab.com/pos-alfa-microservices-go/core/log"
)

type OrderProcessor interface {
	StartConsume() error
}

type MessageOrderProcessor struct {
	queue        string
	client       *rabbitmq.RabbitClient
	processOrder service.ProcessOrder
}

func NewMessageOrderProcessor(queue string, client *rabbitmq.RabbitClient, processOrder service.ProcessOrder) OrderProcessor {
	return &MessageOrderProcessor{
		queue:        queue,
		client:       client,
		processOrder: processOrder,
	}
}

func (p MessageOrderProcessor) StartConsume() error {
	consumer := rabbitmq.NewRabbitConsumer(p.client, "orders-aggregator")
	return consumer.Consume(p.queue, func(body []byte) error {
		bodyOrder := model.Order{}
		if err := json.Unmarshal(body, &bodyOrder); err != nil {
			return fmt.Errorf("fail to unmarshal order %w", err)
		}

		order, err := p.processOrder.Exec(context.Background(), &bodyOrder)
		if err != nil {
			return fmt.Errorf("fail to process order %w", err)
		}

		coreLog.Logger.Infof("order aggregate processed: %v", order)

		return nil
	})
}
