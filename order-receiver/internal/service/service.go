package service

import (
	"context"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	coreErrors "gitlab.com/pos-alfa-microservices-go/core/errors"
	coreLog "gitlab.com/pos-alfa-microservices-go/core/log"

	"microservices/order-receiver/pkg/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Service interface {
	Create(context.Context, *model.Order) (*model.Order, error)
}

type ServiceImpl struct {
	messagePublisher rabbitmq.MessagePublisher
}

func NewServiceImpl(messagePublisher rabbitmq.MessagePublisher) Service {
	return &ServiceImpl{
		messagePublisher: messagePublisher,
	}
}

func (s ServiceImpl) Create(ctx context.Context, order *model.Order) (*model.Order, error) {
	if order.Customer == nil || order.Customer.Id == uuid.Nil {
		return nil, &coreErrors.ValidationError{Message: "invalid customer. customerId is required"}
	}

	for _, i := range order.Items {
		if i.Product == nil || i.Product.Id == uuid.Nil {
			return nil, &coreErrors.ValidationError{Message: "invalid product. productId is required"}
		}
	}

	if err := s.messagePublisher.Publish("orders-received", order); err != nil {
		return nil, errors.Wrap(err, "fail to publish order")
	}

	coreLog.Logger.Infof("order publish. %v", order)

	return order, nil
}
