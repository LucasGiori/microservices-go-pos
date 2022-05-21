package http

import (
	"microservices/notification/internal/container"
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

	healhCheck := server.NewDefautlHealhCheck()

	e.GET("/health", healhCheck.Check)

	return e
}
