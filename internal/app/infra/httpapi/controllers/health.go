package controllers

import "github.com/gofiber/fiber/v2"

type HealthController struct{}

func (h *HealthController) Answer(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

func NewHealthController() *HealthController {
	return &HealthController{}
}
