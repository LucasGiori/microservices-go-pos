package repository

import (
	"context"
	"errors"
	"microservices/customer/pkg/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"gitlab.com/pos-alfa-microservices-go/core/database"
)

type Repository interface {
	FindAll(context.Context) ([]*model.Customer, error)
	FindById(context.Context, string) (*model.Customer, error)
	Create(context.Context, *model.Customer) (*model.Customer, error)
}

type RepositoryImpl struct {
	databaseManager database.DatabaseManager
}

func NewRepositoryImpl(databaseManager database.DatabaseManager) Repository {
	return &RepositoryImpl{databaseManager: databaseManager}
}

func (r RepositoryImpl) FindAll(ctx context.Context) ([]*model.Customer, error) {
	customers := make([]*model.Customer, 0)

	sql := "select id, name, email from customer"
	rows, err := r.databaseManager.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var id uuid.UUID
		var email string
		var name string

		if err := rows.Scan(&id, &name, &email); err != nil {
			return nil, err
		}

		customers = append(customers, &model.Customer{
			Id:    id,
			Name:  name,
			Email: email,
		})
	}

	return customers, nil

}

func (r RepositoryImpl) FindById(ctx context.Context, idParam string) (*model.Customer, error) {
	sql := "select c.id, c.name, c.email from customer c where c.id = $1"

	var id uuid.UUID
	var name string
	var email string

	err := r.databaseManager.QueryRow(ctx, sql, idParam).Scan(&id, &name, &email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &model.Customer{
		Id:    id,
		Name:  name,
		Email: email,
	}, nil
}

func (r RepositoryImpl) Create(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	sql := "insert into customer (name, email) values ($1, $2) returning id"

	var id uuid.UUID
	err := r.databaseManager.QueryRow(ctx, sql,
		customer.Name, customer.Email).Scan(&id)

	if err != nil {
		return nil, err
	}

	customer.Id = id

	return customer, nil
}
