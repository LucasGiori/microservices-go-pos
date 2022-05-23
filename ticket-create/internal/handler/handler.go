package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"microservices/ticket-create/internal/service"
	"microservices/ticket-create/pkg/model"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Create(c echo.Context) error
	GetById(c echo.Context) error
}

type HandlerImpl struct {
	service service.Service
}

func NewHandlerImpl(service service.Service) Handler {
	return &HandlerImpl{
		service: service,
	}
}

func (h HandlerImpl) Create(c echo.Context) error {
	request := model.Ticket{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	model, err := h.service.Create(context.Background(), &request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, model)
}

func (h HandlerImpl) GetById(c echo.Context) error {
	id := c.Param("id")

	redisDatabase := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	ticketCached, err := redisDatabase.Get(context.Background(), id).Result()

	if err == nil && err != redis.Nil {
		ticketCached := []byte(ticketCached)

		ticketStruct := model.Ticket{}

		if err := json.Unmarshal(ticketCached, &ticketStruct); err != nil {
			return fmt.Errorf("fail to unmarshal ticket %w", err)
		}

		return c.JSON(http.StatusOK, ticketStruct)
	}

	ticket, err := h.service.FindById(context.Background(), id)

	if err != nil || ticket.Status == "" {
		return c.JSON(http.StatusBadRequest, &model.Response{
			Message: "Ticket not found",
		})
	}

	return c.JSON(http.StatusOK, ticket)
}
