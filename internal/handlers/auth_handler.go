package handlers

import (
	"go_booking/internal/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) LoginUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"success": true, "message": "user is logged-in"})
}

func (h *AuthHandler) RegisterUser(c *fiber.Ctx) error {
	return c.Status(201).JSON(fiber.Map{"success": true, "message": "created a new user successfully"})
}
