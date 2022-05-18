package service

import (
	"context"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	coreErrors "gitlab.com/pos-alfa-microservices-go/core/errors"
	coreLog "gitlab.com/pos-alfa-microservices-go/core/log"

	"microservices/product-receiver/pkg/model"

	"github.com/pkg/errors"
)

type Service interface {
	Create(context.Context, *model.product) (*model.product, error)
}

type ServiceImpl struct {
	messagePublisher rabbitmq.MessagePublisher
}

func NewServiceImpl(messagePublisher rabbitmq.MessagePublisher) Service {
	return &ServiceImpl{
		messagePublisher: messagePublisher,
	}
}

func (s ServiceImpl) Create(ctx context.Context, product *model.Ticket) (*model.product, error) {
	if product.email == nil {
		return nil, &coreErrors.ValidationError{Message: "invalid email. email is required"}
	}

	if err := s.messagePublisher.Publish("ticket-pending", product); err != nil {
		return nil, errors.Wrap(err, "fail to publish ticket")
	}

	coreLog.Logger.Infof("ticket publish. %v", product)

	return product, nil
}
