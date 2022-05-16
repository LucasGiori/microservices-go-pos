package repository

import (
	"context"
	"microservices/ticket-process/pkg/model"

	"gitlab.com/pos-alfa-microservices-go/core/database"
)

type Repository interface {
	Create(context.Context, *model.Ticket) (*model.Ticket, error)
}

type RepositoryImpl struct {
	databaseManager database.DatabaseManager
}

func NewRepositoryImpl(databaseManager database.DatabaseManager) Repository {
	return &RepositoryImpl{databaseManager: databaseManager}
}

func (r RepositoryImpl) Create(ctx context.Context, order *model.Ticket) (*model.Ticket, error) {
	return nil, nil
}
