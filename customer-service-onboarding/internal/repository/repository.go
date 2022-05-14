package repository

import (
	"context"
	"microservices/customer-service-onboarding/pkg/model"
	"net/http"
)

type Repository interface {
	Create(context.Context, *model.Model) (*model.Model, error)
}

type RepositoryImpl struct {
	client http.Client
}

func NewRepositoryImpl(client http.Client) Repository {
	return &RepositoryImpl{client: client}
}

func (r RepositoryImpl) Create(ctx context.Context, order *model.Model) (*model.Model, error) {
	return nil, nil
}
