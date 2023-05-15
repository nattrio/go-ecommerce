package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-ecommerce/modules/middleware/middlewareHandlers"
	"github.com/nattrio/go-ecommerce/modules/middleware/middlewareRepositories"
	"github.com/nattrio/go-ecommerce/modules/middleware/middlewareUsecases"
	"github.com/nattrio/go-ecommerce/modules/monitor/monitorHandlers"
)

type IModuleFactory interface {
	MonitorModule()
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
