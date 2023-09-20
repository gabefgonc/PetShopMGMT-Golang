package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gabefgonc/PetShopMGMT/internal/app/infra/config"
)

type Router interface {
	Load(router *fiber.App)
}

func NewRouter(config *config.Config, healthRouter *HealthRouter) *fiber.App {
	cfg := fiber.Config{
		AppName:       "Pet Shop Management by @gabefgonc",
		CaseSensitive: true,
		Prefork:       config.Server.Prefork,
	}

	r := fiber.New(cfg)

	healthRouter.Load(r)

	return r
}
