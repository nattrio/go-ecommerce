package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-ecommerce/modules/appinfo/appinfoHandlers"
	"github.com/nattrio/go-ecommerce/modules/appinfo/appinfoRepositories"
	"github.com/nattrio/go-ecommerce/modules/appinfo/appinfoUsecases"
	"github.com/nattrio/go-ecommerce/modules/files/filesHandlers"
	"github.com/nattrio/go-ecommerce/modules/files/filesUsecases"
	"github.com/nattrio/go-ecommerce/modules/middleware/middlewareHandlers"
	"github.com/nattrio/go-ecommerce/modules/middleware/middlewareRepositories"
	"github.com/nattrio/go-ecommerce/modules/middleware/middlewareUsecases"
	"github.com/nattrio/go-ecommerce/modules/monitor/monitorHandlers"
	"github.com/nattrio/go-ecommerce/modules/users/usersHandlers"
	"github.com/nattrio/go-ecommerce/modules/users/usersRepositories"
	"github.com/nattrio/go-ecommerce/modules/users/usersUsecases"
)

type IModuleFactory interface {
	MonitorModule()
	UsersModule()
	AppInfoModule()
	FileModule()
}

type moduleFactory struct {
	r   fiber.Router                          // r is a fiber.Router
	s   *server                               // s is a pointer to a server
	mid middlewareHandlers.IMiddlewareHandler // mid is a middlewareHandlers
}

func InitModule(r fiber.Router, s *server, mid middlewareHandlers.IMiddlewareHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewareHandlers.IMiddlewareHandler {
	repository := middlewareRepositories.MiddlewareRepository(s.db)
	usecase := middlewareUsecases.MiddlewareUsecase(repository)
	handler := middlewareHandlers.MiddlewareHandler(s.cfg, usecase)
	return handler
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HeathCheck)
}

func (m *moduleFactory) UsersModule() {
	repository := usersRepositories.UsersRepository(m.s.db)
	usecase := usersUsecases.UsersUsecase(m.s.cfg, repository)
	handler := usersHandlers.UsersHandler(m.s.cfg, usecase)

	// /v1/users/
	router := m.r.Group("/users")

	router.Post("/signup", m.mid.ApiKeyAuth(), handler.SignUpCustomer)
	router.Post("/signin", m.mid.ApiKeyAuth(), handler.SignIn)
	router.Post("/refresh", m.mid.ApiKeyAuth(), handler.RefreshPassport)
	router.Post("/signout", m.mid.ApiKeyAuth(), handler.SignOut)
	router.Post("/signup-admin", m.mid.JwtAuth(), m.mid.Authorize(2), handler.SignUpAdmin)

	router.Get("/:user_id", m.mid.JwtAuth(), m.mid.ParamsCheck(), handler.GetUserProfile)
	router.Get("/admin/secret", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateAdminToken)

}

func (m *moduleFactory) AppInfoModule() {
	repository := appinfoRepositories.AppinfoRepository(m.s.db)
	usecase := appinfoUsecases.AppinfoUsecase(repository)
	handler := appinfoHandlers.AppinfoHandler(m.s.cfg, usecase)

	router := m.r.Group("/appinfo")

	router.Post("/categories", m.mid.JwtAuth(), m.mid.Authorize(2), handler.AddCategory)

	router.Get("/categories", m.mid.ApiKeyAuth(), handler.FindCategory)
	router.Get("/apikey", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateApiKey)

	router.Delete("/:category_id/categories", m.mid.JwtAuth(), m.mid.Authorize(2), handler.RemoveCategory)
}

func (m *moduleFactory) FileModule() {
	usecase := filesUsecases.FilesUsecase(m.s.cfg)
	handler := filesHandlers.FilesHandler(m.s.cfg, usecase)

	router := m.r.Group("/files")

	router.Post("/upload", m.mid.JwtAuth(), m.mid.Authorize(2), handler.UploadFiles)
	router.Delete("/delete", m.mid.JwtAuth(), m.mid.Authorize(2), handler.DeleteFile)
}
