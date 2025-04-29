package routes

import (
	"go_booking/internal/services"

	"github.com/gofiber/fiber/v2"
)

func HealthRoute(router fiber.Router) {

	r := router.Group("health")

	r.Get("/", services.HealthCheck)
}
