package handler

import (
	"go_booking/internal/user/port"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *userHandler {
	return &userHandler{svc}
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "register",
	})
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body", "success": false})
	}

	token, err := h.svc.Login(req.Username, req.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	return c.JSON(fiber.Map{"access_token": token})
}

func (h *userHandler) GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "get user",
	})
}
