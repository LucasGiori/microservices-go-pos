package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	service "microservices/ticket-process/internal/service/database"
	message "microservices/ticket-process/internal/service/message"
	"microservices/ticket-process/pkg/model"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	coreLog "gitlab.com/pos-alfa-microservices-go/core/log"
)

type TicketProcessor interface {
	StartConsume() error
}

type MessageTicketProcessor struct {
	queue            string
	client           *rabbitmq.RabbitClient
	service          service.ServiceDatabase
	messagePublisher message.ServiceMessage
}

func NewMessageTicketProcessor(queue string, client *rabbitmq.RabbitClient, service service.ServiceDatabase, messagePublisher message.ServiceMessage) TicketProcessor {
	return &MessageTicketProcessor{
		queue:            queue,
		client:           client,
		service:          service,
		messagePublisher: messagePublisher,
	}
}

func (p MessageTicketProcessor) StartConsume() error {
	consumer := rabbitmq.NewRabbitConsumer(p.client, "ticket-pending")
	return consumer.Consume(p.queue, func(body []byte) error {
		bodyOrder := model.Ticket{}
		if err := json.Unmarshal(body, &bodyOrder); err != nil {
			return fmt.Errorf("fail to unmarshal ticket %w", err)
		}

		ticket, err := p.service.Create(context.Background(), &bodyOrder)
		if err != nil {
			return fmt.Errorf("fail to process ticket %w", err)
		}

		p.messagePublisher.Create(context.Background(), ticket)

		coreLog.Logger.Infof("ticket processed: %v", ticket)

		return nil
	})
}
