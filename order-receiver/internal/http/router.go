package http

import (
	"microservices/order-receiver/internal/container"
	"microservices/order-receiver/internal/handler"

	"gitlab.com/pos-alfa-microservices-go/core/http/server"

	"github.com/labstack/echo/v4"
)

type ServerRouter struct {
	container *container.Container
}

func NewRouter(c *container.Container) server.Router {
	return &ServerRouter{
		container: c,
	}
}

func (r *ServerRouter) Create() *echo.Echo {
	e := server.NewCoreEcho()

	jwtMiddleware := server.ValidateJWTMiddleware(r.container.AppConfig)
	handler := handler.NewHandlerImpl(r.container.Service)
	healhCheck := server.NewDefautlHealhCheck()

	e.GET("/health", healhCheck.Check)

	models := e.Group("/orders")
	models.POST("", handler.Create, jwtMiddleware)

	return e
}
