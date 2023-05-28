package middlewareUsecases

import (
	"github.com/nattrio/go-ecommerce/modules/middleware"
	"github.com/nattrio/go-ecommerce/modules/middleware/middlewareRepositories"
)

type IMiddlewareUsecase interface {
	FindAccessToken(userId, AccessToken string) bool
	FindRole() ([]*middleware.Role, error)
}

type middlewareUsecase struct {
	middlewareRepository middlewareRepositories.IMiddlewareRepository
}

func MiddlewareUsecase(middlewareRepository middlewareRepositories.IMiddlewareRepository) IMiddlewareUsecase {
	return &middlewareUsecase{
		middlewareRepository: middlewareRepository,
	}
}

func (u *middlewareUsecase) FindAccessToken(userId, AccessToken string) bool {
	return u.middlewareRepository.FindAccessToken(userId, AccessToken)
}

func (u *middlewareUsecase) FindRole() ([]*middleware.Role, error) {
	roles, err := u.middlewareRepository.FindRole()
	if err != nil {
		return nil, err
	}
	return roles, nil
}
