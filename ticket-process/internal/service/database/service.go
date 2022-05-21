package service

import (
	"context"
	"microservices/ticket-process/internal/repository"

	customErrors "gitlab.com/pos-alfa-microservices-go/core/errors"

	"microservices/ticket-process/pkg/model"
)

type ServiceDatabase interface {
	FindById(context.Context, string) (*model.Ticket, error)
	Create(context.Context, *model.Ticket) (*model.Ticket, error)
	Update(context.Context, *model.Ticket) (*model.Ticket, error)
}

type ServiceImplDatabase struct {
	repository repository.Repository
}

func NewServiceImpl(repository repository.Repository) ServiceDatabase {
	return &ServiceImplDatabase{
		repository: repository,
	}
}

func (r ServiceImplDatabase) FindById(ctx context.Context, ticketId string) (*model.Ticket, error) {
	if ticketId == "" {
		return nil, customErrors.ErrEmptyIdParam
	}

	return r.repository.FindById(ctx, ticketId)
}

func (r ServiceImplDatabase) Create(ctx context.Context, ticket *model.Ticket) (*model.Ticket, error) {
	return r.repository.Create(ctx, ticket)
}

func (r ServiceImplDatabase) Update(ctx context.Context, ticket *model.Ticket) (*model.Ticket, error) {
	return r.repository.Update(ctx, ticket)
}
