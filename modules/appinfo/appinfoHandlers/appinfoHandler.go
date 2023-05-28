package appinfoHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-ecommerce/config"
	"github.com/nattrio/go-ecommerce/modules/appinfo/appinfoUsecases"
	"github.com/nattrio/go-ecommerce/modules/entities"
	"github.com/nattrio/go-ecommerce/pkg/myshopauth"
)

type appinfoHandlersErrCode string

const (
	generateApiKeyErr appinfoHandlersErrCode = "appinfo-001"
	findCategoryErr   appinfoHandlersErrCode = "appinfo-002"
	addCategoryErr    appinfoHandlersErrCode = "appinfo-003"
	removeCategoryErr appinfoHandlersErrCode = "appinfo-004"
)

type IAppInfoHandler interface {
	GenerateApiKey(c *fiber.Ctx) error
}

type appInfoHandler struct {
	cfg            config.IConfig
	appinfoUsecase appinfoUsecases.IAppInfoUsecase
}

func AppinfoHandler(cfg config.IConfig, appinfoUsecase appinfoUsecases.IAppInfoUsecase) IAppInfoHandler {
	return &appInfoHandler{
		cfg:            cfg,
		appinfoUsecase: appinfoUsecase,
	}
}

func (h *appInfoHandler) GenerateApiKey(c *fiber.Ctx) error {
	apiKey, err := myshopauth.NewMyshopAuth(
		myshopauth.ApiKey,
		h.cfg.Jwt(),
		nil,
	)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(generateApiKeyErr),
			err.Error(),
		).Res()
	}
	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		&struct {
			Key string `json:"key"`
		}{
			Key: apiKey.SignToken(),
		},
	).Res()
}
