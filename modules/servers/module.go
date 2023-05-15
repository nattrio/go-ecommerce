package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-ecommerce/modules/monitor/monitorHandlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	r fiber.Router // r is a fiber.Router
	s *server      // s is a pointer to a server
}

func InitModule(r fiber.Router, s *server) IModuleFactory {
	return &moduleFactory{
		r: r,
		s: s,
	}
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HeathCheck)
}
