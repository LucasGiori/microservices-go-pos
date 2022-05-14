package service

import (
	"context"
	"errors"
	"microservices/product/internal/repository"

	"microservices/product/pkg/model"

	coreErrors "gitlab.com/pos-alfa-microservices-go/core/errors"
)

type Service interface {
	Create(context.Context, *model.Product) (*model.Product, error)
	FindAll(context.Context, []string) ([]*model.Product, error)
	FindById(context.Context, string) (*model.Product, error)
	FindByIds(context.Context, []string) (map[string]*model.Product, error)
}

type ServiceImpl struct {
	repository repository.Repository
}

func NewServiceImpl(repository repository.Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

func (s ServiceImpl) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	return s.repository.Create(ctx, product)
}

func (s ServiceImpl) FindAll(ctx context.Context, ids []string) ([]*model.Product, error) {
	if len(ids) > 0 {
		productById, err := s.repository.FindByIds(ctx, ids)
		if err != nil {
			return nil, errors.New("")
		}

		products := make([]*model.Product, 0)
		products = append(products, productById...)

		return products, nil
	}

	return s.repository.FindAll(ctx)
}

func (s ServiceImpl) FindById(ctx context.Context, id string) (*model.Product, error) {
	if id == "" {
		return nil, coreErrors.ErrEmptyIdParam
	}

	return s.repository.FindById(ctx, id)
}

func (s ServiceImpl) FindByIds(ctx context.Context, ids []string) (map[string]*model.Product, error) {
	if len(ids) == 0 {
		return nil, coreErrors.ErrEmptyIdParam
	}

	products, err := s.repository.FindByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	productById := make(map[string]*model.Product, 0)
	for _, p := range products {
		productById[p.Id.String()] = p
	}

	return productById, nil
}
