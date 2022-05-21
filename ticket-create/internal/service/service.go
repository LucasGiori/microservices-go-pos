package service

import (
	"context"

	"microservices/ticket-create/pkg/model"

	"github.com/google/uuid"
	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	coreLog "gitlab.com/pos-alfa-microservices-go/core/log"

	"github.com/pkg/errors"
)

type ServiceMessage interface {
	Create(context.Context, *model.Ticket) (*model.Ticket, error)
}

type ServiceImpl struct {
	messagePublisher rabbitmq.MessagePublisher
}

func NewServiceImpl(messagePublisher rabbitmq.MessagePublisher) ServiceMessage {
	return &ServiceImpl{
		messagePublisher: messagePublisher,
	}
}

func (s ServiceImpl) Create(ctx context.Context, ticket *model.Ticket) (*model.Ticket, error) {

	ticket.Id = uuid.New()
	ticket.Status = model.OPEN
	if err := s.messagePublisher.Publish("ticket-pending", ticket); err != nil {
		return nil, errors.Wrap(err, "fail to publish ticket")
	}

	coreLog.Logger.Infof("ticket publish. %v", ticket)

	return ticket, nil
}
