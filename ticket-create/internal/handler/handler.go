package handler

import (
	"context"
	"microservices/ticket-create/internal/service"
	"microservices/ticket-create/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Create(c echo.Context) error
}

type HandlerImpl struct {
	service service.ServiceMessage
}

func NewHandlerImpl(service service.ServiceMessage) Handler {
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
