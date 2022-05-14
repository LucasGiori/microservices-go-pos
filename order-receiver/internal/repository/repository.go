package repository

import (
	"context"
	"microservices/order-receiver/pkg/model"

	"gitlab.com/pos-alfa-microservices-go/core/database"
)

type Repository interface {
	Create(context.Context, *model.Model) (*model.Model, error)
}

type RepositoryImpl struct {
	databaseManager database.DatabaseManager
}

func NewRepositoryImpl(databaseManager database.DatabaseManager) Repository {
	return &RepositoryImpl{databaseManager: databaseManager}
}

func (r RepositoryImpl) Create(ctx context.Context, order *model.Model) (*model.Model, error) {
	return nil, nil
}
