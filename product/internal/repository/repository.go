package repository

import (
	"context"
	"microservices/product/pkg/model"

	"github.com/google/uuid"

	pgtypeExt "github.com/jackc/pgtype/ext/shopspring-numeric"
	"gitlab.com/pos-alfa-microservices-go/core/database"
)

type Repository interface {
	Create(context.Context, *model.Product) (*model.Product, error)
	FindAll(context.Context) ([]*model.Product, error)
	FindById(context.Context, string) (*model.Product, error)
	FindByIds(context.Context, []string) ([]*model.Product, error)
}

type RepositoryImpl struct {
	databaseManager database.DatabaseManager
}

func NewRepositoryImpl(databaseManager database.DatabaseManager) Repository {
	return &RepositoryImpl{databaseManager: databaseManager}
}

func (r RepositoryImpl) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	sql := "insert into product (name, value) values ($1, $2) returning id"

	var id uuid.UUID
	err := r.databaseManager.QueryRow(ctx, sql, product.Name, product.Value).Scan(&id)

	if err != nil {
		return nil, err
	}

	product.Id = id

	return product, nil
}

func (r RepositoryImpl) FindAll(ctx context.Context) ([]*model.Product, error) {
	return r.findBy(ctx, "")
}

func (r RepositoryImpl) FindByIds(ctx context.Context, ids []string) ([]*model.Product, error) {
	return r.findBy(ctx, "WHERE id = ANY($1)", ids)
}

func (r RepositoryImpl) findBy(ctx context.Context, sqlCondition string, args ...interface{}) ([]*model.Product, error) {
	products := make([]*model.Product, 0)

	sql := "select id, name, \"value\" from product p " + sqlCondition
	rows, err := r.databaseManager.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var id uuid.UUID
		var name string
		var value pgtypeExt.Numeric

		if err := rows.Scan(&id, &name, &value); err != nil {
			return nil, err
		}

		products = append(products, &model.Product{
			Id:    id,
			Name:  name,
			Value: &value.Decimal,
		})
	}

	return products, nil
}

func (r RepositoryImpl) FindById(ctx context.Context, paramId string) (*model.Product, error) {
	sql := "select id, name, \"value\" from product p where id = $1"

	var id uuid.UUID
	var name string
	var value pgtypeExt.Numeric

	err := r.databaseManager.QueryRow(ctx, sql, paramId).Scan(&id, &name, &value)
	if err != nil {
		return nil, err
	}

	return &model.Product{
		Id:    id,
		Name:  name,
		Value: &value.Decimal,
	}, nil
}
