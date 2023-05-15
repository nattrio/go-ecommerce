package monitorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-ecommerce/config"
	"github.com/nattrio/go-ecommerce/modules/entities"
	"github.com/nattrio/go-ecommerce/modules/monitor"
)

type IMonitorHandler interface {
	HeathCheck(c *fiber.Ctx) error
}

type monitorHandler struct {
	cfg config.IConfig
}

func MonitorHandler(cfg config.IConfig) IMonitorHandler {
	return &monitorHandler{
		cfg: cfg,
	}
}

func (h *monitorHandler) HeathCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, res).Res()
}
