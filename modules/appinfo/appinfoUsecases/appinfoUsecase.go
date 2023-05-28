package appinfoUsecases

import "github.com/nattrio/go-ecommerce/modules/appinfo/appinfoRepositories"

type IAppInfoUsecase interface {
}

type appInfoUsecase struct {
	appinfoRepository appinfoRepositories.IAppInfoRepository
}

func AppinfoUsecase(appinfoRepository appinfoRepositories.IAppInfoRepository) IAppInfoUsecase {
	return &appInfoUsecase{
		appinfoRepository: appinfoRepository,
	}
}
