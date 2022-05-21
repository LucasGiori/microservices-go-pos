package handler

import (
	"context"
	service "microservices/ticket-process/internal/service/database"
	message "microservices/ticket-process/internal/service/message"
	"microservices/ticket-process/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Create(c echo.Context) error
	GetById(c echo.Context) error
	Update(c echo.Context) error
}

type HandlerImpl struct {
	service service.ServiceDatabase
	message message.ServiceMessage
}

func NewHandlerImpl(service service.ServiceDatabase, message message.ServiceMessage) Handler {
	return &HandlerImpl{
		service: service,
		message: message,
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

	h.message.Create(context.Background(), model)

	return c.JSON(http.StatusAccepted, model)
}

func (h HandlerImpl) GetById(c echo.Context) error {
	id := c.Param("id")
	ticket, err := h.service.FindById(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &model.Response{
			Message: "Ticket not found",
		})
	}

	return c.JSON(http.StatusOK, ticket)
}

func (h HandlerImpl) Update(c echo.Context) error {
	id := c.Param("id")
	request := model.Ticket{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ticket, _ := h.service.FindById(context.Background(), id)
	if ticket == nil {
		return c.JSON(http.StatusNotFound, &model.Response{
			Message: "Ticket not found",
		})
	}

	model, err := h.service.Update(context.Background(), &request)
	if err != nil {
		return err
	}

	h.message.Create(context.Background(), model)

	return c.JSON(http.StatusAccepted, model)
}
