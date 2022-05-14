package handler

import (
	"context"
	"microservices/product/internal/service"
	"microservices/product/pkg/model"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Create(c echo.Context) error
	GetAll(c echo.Context) error
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
	request := model.Product{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	model, err := h.service.Create(context.Background(), &request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model)
}

func (h HandlerImpl) GetAll(c echo.Context) error {
	idsParam := c.QueryParam("ids")
	var ids []string
	if idsParam != "" {
		ids = strings.Split(idsParam, ",")
	}

	products, err := h.service.FindAll(context.Background(), ids)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}

func (h HandlerImpl) GetById(c echo.Context) error {
	id := c.Param("id")
	product, err := h.service.FindById(context.Background(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}
