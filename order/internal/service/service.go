package service

import (
	"context"
	"microservices/order/internal/repository"

	customErrors "gitlab.com/pos-alfa-microservices-go/core/errors"

	"time"

	"microservices/order/pkg/model"
)

type Service interface {
	Create(context.Context, *model.Order) (*model.Order, error)
}

type ServiceImpl struct {
	repository repository.Repository
}

func NewServiceImpl(repository repository.Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

func (r ServiceImpl) Create(ctx context.Context, order *model.Order) (*model.Order, error) {
	if err := order.ValidateToCreate(); err != nil {
		return nil, &customErrors.ValidationError{Message: err.Error()}
	}

	if order.Customer == nil {
		return nil, &customErrors.ValidationError{Message: "invalid customer"}
	}

	for _, i := range order.Items {
		if i.Product == nil {
			return nil, &customErrors.ValidationError{Message: "invalid product"}
		}
	}

	if order.DateTime == nil {
		now := time.Now()
		order.DateTime = &now
	}

	order.Status = model.CLOSED
	order.CalcTotal()
	return r.repository.Create(ctx, order)
}
