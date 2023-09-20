package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gabefgonc/PetShopMGMT/internal/app/infra/httpapi/controllers"
)

type HealthRouter struct {
	controller *controllers.HealthController
}

func (h *HealthRouter) Load(r *fiber.App) {
	r.Get("/hello", h.controller.Answer)
}

func NewHealthRouter(controller *controllers.HealthController) *HealthRouter {
	return &HealthRouter{
		controller: controller,
	}
}
