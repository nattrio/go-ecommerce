package appinfoHandlers

import (
	"github.com/nattrio/go-ecommerce/config"
	"github.com/nattrio/go-ecommerce/modules/appinfo/appinfoUsecases"
)

type IAppInfoHandler interface {
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
