package servers

import (
	"github.com/gofiber/fiber/v2"
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

	router.Post("/signup", handler.SignUpCustomer)
	router.Post("/signin", handler.SignIn)
	router.Post("/refresh", handler.RefreshPassport)
	router.Post("/signout", handler.SignOut)
	router.Post("/signup-admin", handler.SignUpAdmin)

	router.Get("/:user_id", m.mid.JwtAuth(), m.mid.ParamsCheck(), handler.GetUserProfile)
	router.Get("/admin/secret", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateAdminToken)

}
