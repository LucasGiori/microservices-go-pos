package service

import (
	"context"
	"fmt"

	"microservices/ticket-create/internal/client"
	"microservices/ticket-create/pkg/model"

	"gitlab.com/pos-alfa-microservices-go/core/auth"

	"github.com/google/uuid"
	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	coreErrors "gitlab.com/pos-alfa-microservices-go/core/errors"
	coreLog "gitlab.com/pos-alfa-microservices-go/core/log"

	"github.com/pkg/errors"
)

type Service interface {
	Create(context.Context, *model.Ticket) (*model.Ticket, error)
	FindById(ctx context.Context, ticketId string) (*model.Ticket, error)
}

type ServiceImpl struct {
	messagePublisher rabbitmq.MessagePublisher
	ticketClient     client.TicketClient
	tokenManager     auth.TokenManager
}

func NewServiceImpl(messagePublisher rabbitmq.MessagePublisher, ticketClient client.TicketClient, tokenManager auth.TokenManager) Service {
	return &ServiceImpl{
		messagePublisher: messagePublisher,
		ticketClient:     ticketClient,
		tokenManager:     tokenManager,
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

func (s ServiceImpl) FindById(ctx context.Context, ticketId string) (*model.Ticket, error) {
	ctx, err := s.tokenManager.AddSystemTokenInContext(ctx)
	if err != nil {
		return nil, err
	}

	ticket, err := s.validadeAndFindTicket(ctx, ticketId)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r ServiceImpl) validadeAndFindTicket(ctx context.Context, ticketId string) (*model.Ticket, error) {
	if ticketId == "" {
		return nil, &coreErrors.ValidationError{Message: "invalid ticket. ticketId is required"}
	}

	ticket, err := r.ticketClient.GetById(ctx, ticketId)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("fail to get ticketId: %v", ticketId))
	}

	return ticket, nil
}
