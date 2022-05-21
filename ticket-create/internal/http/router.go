package http

import (
	"microservices/ticket-create/internal/container"
	"microservices/ticket-create/internal/handler"

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
	handler := handler.NewHandlerImpl(r.container.ServiceImplMessage)
	healhCheck := server.NewDefautlHealhCheck()

	e.GET("/health", healhCheck.Check)

	ticket := e.Group("/ticket-create")
	ticket.POST("", handler.Create, jwtMiddleware)
	ticket.GET("/:id", handler.GetById, jwtMiddleware)

	return e
}
