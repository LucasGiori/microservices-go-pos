package http

import (
	"microservices/product/internal/container"
	"microservices/product/internal/handler"

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
	productHandler := handler.NewHandlerImpl(r.container.Service)
	healhCheck := server.NewDefautlHealhCheck()

	e.GET("/health", healhCheck.Check)

	products := e.Group("/products")
	products.POST("", productHandler.Create, jwtMiddleware)
	products.GET("", productHandler.GetAll, jwtMiddleware)
	products.GET("/:id", productHandler.GetById, jwtMiddleware)

	return e
}
