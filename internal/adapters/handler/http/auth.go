package http

import (
	"go_booking/internal/adapters/handler/dto"
	"go_booking/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	svc port.AuthService
}

func NewAuthHandler(svc port.AuthService) *AuthHandler {
	return &AuthHandler{svc}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return ErrorResponse(c, fiber.StatusUnauthorized, "invalid credentials")
	}

	token, err := h.svc.Login(c, req.Username, req.Password)

	if err != nil {
		return ErrorResponse(c, fiber.StatusUnauthorized, "invalid username or password")
	}

	return SuccessResponse(c, "logged-in is success", token)
}
