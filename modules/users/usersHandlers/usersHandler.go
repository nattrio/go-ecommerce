package usersHandlers

import (
	"github.com/nattrio/go-ecommerce/config"
	"github.com/nattrio/go-ecommerce/modules/users/usersUsecases"
)

type IUsersHandler interface {
}

type usersHandler struct {
	cfg          config.IConfig
	usersUsecase usersUsecases.IUsersUsecase
}

func UsersHandler() IUsersHandler {
	return &usersHandler{}
}
