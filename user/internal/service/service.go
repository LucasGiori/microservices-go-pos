package service

import (
	"context"
	"fmt"
	"microservices/user/internal/hash"
	"microservices/user/internal/repository"

	customErrors "gitlab.com/pos-alfa-microservices-go/core/errors"

	"microservices/user/pkg/model"
)

type Service interface {
	Create(context.Context, *model.User) (*model.User, error)
	FindByLogin(context.Context, string) (*model.User, error)
}

type ServiceImpl struct {
	repository repository.Repository
}

func NewServiceImpl(repository repository.Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

func (r ServiceImpl) FindByLogin(ctx context.Context, login string) (*model.User, error) {
	return r.repository.FindByLogin(ctx, login)
}

func (s ServiceImpl) Create(ctx context.Context, user *model.User) (*model.User, error) {
	u, err := s.repository.FindByLogin(ctx, user.Login)
	if err != nil {
		return nil, err
	}

	if u != nil {
		return nil, &customErrors.ValidationError{
			Message: fmt.Sprintf("user %s already exists", u.Login),
		}
	}

	hashPassword, err := hash.NewHash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashPassword

	createdUser, err := s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser.Password = ""

	return createdUser, nil
}
