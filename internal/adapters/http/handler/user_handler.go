package handler

import (
	"fmt"
	"go_booking/internal/core/user"

	"github.com/gofiber/fiber/v2"
)

type UserHanlder struct {
	service *user.Service
}

func NewUserHandler(service *user.Service) *UserHanlder {
	return &UserHanlder{service}
}

func (h *UserHanlder) Register(c *fiber.Ctx) error {
	user := new(user.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Invalid body request"})
	}

	fmt.Print(1, user)

	return c.JSON(user)
}
