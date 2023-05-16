package usersUsecases

import (
	"github.com/nattrio/go-ecommerce/config"
	"github.com/nattrio/go-ecommerce/modules/users/usersRepositories"
)

type IUsersUsecase interface {
}

type usersUsecase struct {
	cfg             config.IConfig
	usersRepository usersRepositories.IUsersRepository
}

func UsersUsecase(cfg config.IConfig, usersRepository usersRepositories.IUsersRepository) IUsersUsecase {
	return &usersUsecase{
		cfg:             cfg,
		usersRepository: usersRepository,
	}
}
