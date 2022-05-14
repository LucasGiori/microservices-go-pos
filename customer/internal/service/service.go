package service

import (
	"context"
	"microservices/customer/internal/repository"

	customErrors "gitlab.com/pos-alfa-microservices-go/core/errors"

	"microservices/customer/pkg/model"
)

type Service interface {
	FindAll(context.Context) ([]*model.Customer, error)
	FindById(context.Context, string) (*model.Customer, error)
	Create(context.Context, *model.Customer) (*model.Customer, error)
}

type ServiceImpl struct {
	repository repository.Repository
}

func NewServiceImpl(repository repository.Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

func (r ServiceImpl) FindAll(ctx context.Context) ([]*model.Customer, error) {
	return r.repository.FindAll(ctx)
}

func (r ServiceImpl) FindById(ctx context.Context, id string) (*model.Customer, error) {
	if id == "" {
		return nil, customErrors.ErrEmptyIdParam
	}

	return r.repository.FindById(ctx, id)
}

func (r ServiceImpl) Create(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	return r.repository.Create(ctx, customer)
}
