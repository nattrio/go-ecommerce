package middlewareUsecases

import (
	"github.com/nattrio/go-ecommerce/modules/middleware/middlewareRepositories"
)

type IMiddlewareUsecase interface {
}

type middlewareUsecase struct {
	middlewareRepository middlewareRepositories.IMiddlewareRepository
}

func MiddlewareUsecase(middlewareRepository middlewareRepositories.IMiddlewareRepository) IMiddlewareUsecase {
	return &middlewareUsecase{
		middlewareRepository: middlewareRepository,
	}
}
