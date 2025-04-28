package controllers

import (
	"go_booking/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(s *services.UserService) *UserController {
	return &UserController{service: s}
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	users, err := uc.service.GetAllUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not get users"})
	}

	return c.JSON(users)
}
