package http

import (
	"microservices/user/internal/container"
	"microservices/user/internal/handler"

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

	userHandler := handler.NewHandlerImpl(r.container.Service)
	authHandler := handler.NewAuthHandler(r.container.AuthManager)

	e.GET("/health", healhCheck.Check)

	users := e.Group("/users")
	users.POST("", userHandler.Create)

	auth := e.Group("/auth")
	auth.POST("/login", authHandler.Login)

	return e
}
