package client

import (
	"context"
	"encoding/json"
	"microservices/order-aggregator/pkg/model"

	"gitlab.com/pos-alfa-microservices-go/core/http/client"
)

type CustomerClient interface {
	GetById(ctx context.Context, id string) (*model.Customer, error)
}

type HttpCustomerClient struct {
	httpClient client.HttpClient
	url        string
}

func NewHttpCustomerClient(httpClient client.HttpClient, url string) CustomerClient {
	return &HttpCustomerClient{
		httpClient: httpClient,
		url:        url,
	}
}

func (c HttpCustomerClient) GetById(ctx context.Context, id string) (*model.Customer, error) {
	body, err := c.httpClient.Get(ctx, c.url+"/customers/"+id)
	if err != nil {
		return nil, err
	}

	var customer model.Customer
	if err := json.Unmarshal(body, &customer); err != nil {
		return nil, err
	}

	return &customer, nil
}
