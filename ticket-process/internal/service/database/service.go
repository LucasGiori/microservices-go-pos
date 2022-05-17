package service

import (
	"context"
	"microservices/ticket-process/internal/repository"

	"github.com/google/uuid"
	customErrors "gitlab.com/pos-alfa-microservices-go/core/errors"

	"microservices/ticket-process/pkg/model"
)

type Service interface {
	FindById(context.Context, uuid.UUID) (*model.Ticket, error)
	Create(context.Context, *model.Ticket) (*model.Ticket, error)
}

type ServiceImpl struct {
	repository repository.Repository
}

func NewServiceImpl(repository repository.Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

func (r ServiceImpl) FindById(ctx context.Context, ticketId uuid.UUID) (*model.Ticket, error) {
	if ticketId == "" {
		return nil, customErrors.ErrEmptyIdParam
	}

	return r.repository.FindById(ctx, ticketId)
}

func (r ServiceImpl) Create(ctx context.Context, ticket *model.Ticket) (*model.Ticket, error) {
	return r.repository.Create(ctx, ticket)
}
