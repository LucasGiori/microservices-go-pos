package client

import (
	"context"
	"encoding/json"
	"microservices/order-aggregator/pkg/model"
	"strings"

	"gitlab.com/pos-alfa-microservices-go/core/http/client"
)

type ProductClient interface {
	GetByIds(ctx context.Context, ids []string) (map[string]*model.Product, error)
}

type HttpProductClient struct {
	httpClient client.HttpClient
	url        string
}

func NewHttpProductClient(httpClient client.HttpClient, url string) ProductClient {
	return &HttpProductClient{
		httpClient: httpClient,
		url:        url,
	}
}

func (c HttpProductClient) GetByIds(ctx context.Context, ids []string) (map[string]*model.Product, error) {
	body, err := c.httpClient.Get(ctx, c.url+"/products?ids="+strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	var products []*model.Product
	if err := json.Unmarshal(body, &products); err != nil {
		return nil, err
	}

	productById := make(map[string]*model.Product, 0)
	for _, p := range products {
		productById[p.Id.String()] = p
	}

	return productById, nil
}
