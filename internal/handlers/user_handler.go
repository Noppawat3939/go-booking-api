package handlers

import (
	"go_booking/internal/models"
	"go_booking/internal/services"

	"github.com/gofiber/fiber/v2"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UserHanlder struct {
	service *services.UserService
}

func NewUserHanlder(s *services.UserService) *UserHanlder {
	return &UserHanlder{service: s}
}

func (h *UserHanlder) CreateUser(c *fiber.Ctx) error {
	req := new(CreateUserRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user := &models.User{
		Username: req.Username,
		Role:     req.Role,
	}

	if err := h.service.CreateUser(user); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "data": user})
}

func (h *UserHanlder) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not get users"})
	}

	return c.JSON(fiber.Map{"success": true, "data": users})
}
