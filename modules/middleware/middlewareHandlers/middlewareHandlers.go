package middlewareHandlers

import (
	"github.com/nattrio/go-ecommerce/modules/middleware/middlewareUsecases"
)

type IMiddlewareHandler interface {
}

type middlewareHandler struct {
	middlewareUsecase middlewareUsecases.IMiddlewareUsecase
}

func MiddlewareHandler(middlewareUsecase middlewareUsecases.IMiddlewareUsecase) IMiddlewareHandler {
	return &middlewareHandler{
		middlewareUsecase: middlewareUsecase,
	}
}
