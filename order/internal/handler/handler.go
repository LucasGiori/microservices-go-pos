package handler

import (
	"context"
	"microservices/order/internal/service"
	"microservices/order/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Create(c echo.Context) error
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
	request := model.Order{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	model, err := h.service.Create(context.Background(), &request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model)
}
