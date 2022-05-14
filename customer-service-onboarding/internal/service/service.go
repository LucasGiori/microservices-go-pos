package service

import (
	"context"
	"microservices/customer-service-onboarding/internal/repository"

	"microservices/customer-service-onboarding/pkg/model"
)

type Service interface {
	Create(context.Context, *model.Model) (*model.Model, error)
}

type ServiceImpl struct {
	repository repository.Repository
}

func NewServiceImpl(repository repository.Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

func (s ServiceImpl) Create(ctx context.Context, model *model.Model) (*model.Model, error) {
	return s.repository.Create(ctx, model)
}
