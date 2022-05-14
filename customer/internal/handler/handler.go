package handler

import (
	"context"
	"microservices/customer/internal/service"
	"microservices/customer/pkg/model"
	"net/http"

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
	request := model.Customer{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	model, err := h.service.Create(context.Background(), &request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model)
}

func (h HandlerImpl) GetById(c echo.Context) error {
	id := c.Param("id")
	customer, err := h.service.FindById(context.Background(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, customer)
}
