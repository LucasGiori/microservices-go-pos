package handler

import (
	"context"
	// service "microservices/ticket-process/internal/service/message" // ver para chamar database ao invés de message
	service "microservices/ticket-process/internal/service/database"
	"microservices/ticket-process/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"fmt"
)

type Handler interface {
	Create(c echo.Context) error
	GetById(c echo.Context) error
}

type HandlerImpl struct {
	service service.ServiceDatabase
}

func NewHandlerImpl(service service.ServiceDatabase) Handler {
	return &HandlerImpl{
		service: service,
	}
}

func (h HandlerImpl) Create(c echo.Context) error {
	request := model.Ticket{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fmt.Print("Aqui     :    ", request)
	model, err := h.service.Create(context.Background(), &request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, model)
}

func (h HandlerImpl) GetById(c echo.Context) error {
	id := c.Param("id")
	product, err := h.service.FindById(context.Background(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}
