package http

import (
	"microservices/customer/internal/container"
	"microservices/customer/internal/handler"

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

	customerHandler := handler.NewHandlerImpl(r.container.Service)
	healhCheck := server.NewDefautlHealhCheck()

	e.GET("/health", healhCheck.Check)

	customers := e.Group("/customers")
	customers.POST("", customerHandler.Create)
	customers.GET("/:id", customerHandler.GetById)

	return e
}
