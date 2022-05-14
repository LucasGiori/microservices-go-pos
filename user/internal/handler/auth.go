package handler

import (
	"microservices/user/internal/service"
	"microservices/user/pkg/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Login(c echo.Context) error
}

type HttpAuthHandler struct {
	authManager service.AuthManager
}

func NewAuthHandler(authManager service.AuthManager) AuthHandler {
	return &HttpAuthHandler{
		authManager: authManager,
	}
}

func (h HttpAuthHandler) Login(c echo.Context) error {
	authRequest := model.AuthRequest{}
	if err := c.Bind(&authRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := h.authManager.Login(&authRequest)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, token)
}
