package client

import (
	"context"
	"encoding/json"
	"microservices/ticket-create/pkg/model"

	"gitlab.com/pos-alfa-microservices-go/core/http/client"
)

type TicketClient interface {
	GetById(ctx context.Context, id string) (*model.Ticket, error)
}

type HttpticketClient struct {
	httpClient client.HttpClient
	url        string
}

func NewHttpticketClient(httpClient client.HttpClient, url string) TicketClient {
	return &HttpticketClient{
		httpClient: httpClient,
		url:        url,
	}
}

func (c HttpticketClient) GetById(ctx context.Context, id string) (*model.Ticket, error) {
	body, err := c.httpClient.Get(ctx, c.url + id)
	if err != nil {
		return nil, err
	}

	var Ticket model.Ticket
	if err := json.Unmarshal(body, &Ticket); err != nil {
		return nil, err
	}

	return &Ticket, nil
}
