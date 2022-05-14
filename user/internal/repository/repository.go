package repository

import (
	"context"
	"microservices/user/pkg/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"gitlab.com/pos-alfa-microservices-go/core/database"
)

type Repository interface {
	FindById(context.Context, string) (*model.User, error)
	FindByLogin(context.Context, string) (*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
}

type RepositoryImpl struct {
	databaseManager database.DatabaseManager
}

func NewRepositoryImpl(databaseManager database.DatabaseManager) Repository {
	return &RepositoryImpl{databaseManager: databaseManager}
}

func (r RepositoryImpl) Create(ctx context.Context, user *model.User) (*model.User, error) {
	sql := "insert into \"user\" (login, password) values ($1, $2) returning id"

	var id uuid.UUID
	err := r.databaseManager.QueryRow(ctx, sql, user.Login, user.Password).Scan(&id)

	if err != nil {
		return nil, err
	}

	user.Id = id

	return user, nil
}

func (r RepositoryImpl) FindById(ctx context.Context, id string) (*model.User, error) {
	return r.findBy(ctx, "id", id, false)

}

func (r RepositoryImpl) FindByLogin(ctx context.Context, login string) (*model.User, error) {
	return r.findBy(ctx, "login", login, true)
}

func (r RepositoryImpl) findBy(ctx context.Context, paramName, paramValue string, withPassword bool) (*model.User, error) {
	sql := "select id, login, password from \"user\" u where " + paramName + " = $1"

	var id uuid.UUID
	var login string
	var password string

	err := r.databaseManager.QueryRow(ctx, sql, paramValue).Scan(&id, &login, &password)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, nil
		}

		return nil, err
	}

	user := &model.User{
		Id:    id,
		Login: login,
	}

	if withPassword {
		user.Password = password
	}

	return user, nil
}
