package service

import (
	"fmt"

	"gitlab.com/pos-alfa-microservices-go/core/auth"
	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"

	coreErrors "gitlab.com/pos-alfa-microservices-go/core/errors"

	"context"
	"microservices/order-aggregator/internal/client"
	"microservices/order-aggregator/pkg/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ProcessOrder interface {
	Exec(context.Context, *model.Order) (*model.Order, error)
}

type ProcessOrderImpl struct {
	messagePublisher rabbitmq.MessagePublisher
	customerClient   client.CustomerClient
	productClient    client.ProductClient
	tokenManager     auth.TokenManager
}

func NewProcessOrderImpl(messagePublisher rabbitmq.MessagePublisher, customerClient client.CustomerClient, productClient client.ProductClient, tokenManager auth.TokenManager) ProcessOrder {
	return &ProcessOrderImpl{
		messagePublisher: messagePublisher,
		customerClient:   customerClient,
		productClient:    productClient,
		tokenManager:     tokenManager,
	}
}

func (r ProcessOrderImpl) Exec(ctx context.Context, order *model.Order) (*model.Order, error) {
	ctx, err := r.tokenManager.AddSystemTokenInContext(ctx)
	if err != nil {
		return nil, err
	}

	customer, err := r.validadeAndFindCustomer(ctx, order)
	if err != nil {
		return nil, err
	}

	productsMap, err := r.validadeAndFindProducts(ctx, order)
	if err != nil {
		return nil, err
	}

	for _, i := range order.Items {
		product := productsMap[i.Product.Id.String()]
		if product == nil {
			return nil, &coreErrors.ValidationError{Message: "invalid product. productId not found: " + i.Product.Id.String()}
		}

		i.Product = product
	}

	order.Customer = customer
	order.Status = model.OPEN

	if err := r.messagePublisher.Publish("orders-pending", order); err != nil {
		return nil, errors.Wrap(err, "fail to publish order")
	}

	return order, nil

}

func (r ProcessOrderImpl) validadeAndFindCustomer(ctx context.Context, order *model.Order) (*model.Customer, error) {
	if order.Customer == nil || order.Customer.Id == uuid.Nil {
		return nil, &coreErrors.ValidationError{Message: "invalid customer. customerId is required"}
	}

	customer, err := r.customerClient.GetById(ctx, order.Customer.Id.String())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("fail to get customer id: %v", order.Id))
	}

	if customer == nil || customer.Id == uuid.Nil {
		return nil, &coreErrors.ValidationError{Message: "invalid customer."}
	}

	return customer, nil
}

func (r ProcessOrderImpl) validadeAndFindProducts(ctx context.Context, order *model.Order) (map[string]*model.Product, error) {
	ids := make([]string, 0)
	for _, i := range order.Items {
		if i.Product == nil || i.Product.Id == uuid.Nil {
			return nil, &coreErrors.ValidationError{Message: "invalid product. productId is required"}
		}
		ids = append(ids, i.Product.Id.String())
	}

	products, err := r.productClient.GetByIds(ctx, ids)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("fail to get products ids: %v", ids))
	}

	return products, nil
}
