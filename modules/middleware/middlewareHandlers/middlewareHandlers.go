package middlewareHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nattrio/go-ecommerce/config"
	"github.com/nattrio/go-ecommerce/modules/middleware/middlewareUsecases"
)

type IMiddlewareHandler interface {
	Cors() fiber.Handler
}

type middlewareHandler struct {
	cfg               config.IConfig
	middlewareUsecase middlewareUsecases.IMiddlewareUsecase
}

func MiddlewareHandler(cfg config.IConfig, middlewareUsecase middlewareUsecases.IMiddlewareUsecase) IMiddlewareHandler {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}

func (h *middlewareHandler) Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}
