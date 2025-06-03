package handler

import (
	"go_booking/internal/user/model"
	"go_booking/internal/user/port"

	"github.com/gofiber/fiber/v2"
)

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Active   bool   `json:"active"`
}

type userHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *userHandler {
	return &userHandler{svc}
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	var req registerReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body", "success": false})
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
		Active:   req.Active,
	}

	if err := h.svc.Register(&user); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"success": false, "message": "failed register new user"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "created new user",
	})
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	var req loginReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body", "success": false})
	}

	token, err := h.svc.Login(req.Username, req.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	return c.JSON(fiber.Map{"success": true, "data": fiber.Map{"access_token": token}})
}

func (h *userHandler) GetUser(c *fiber.Ctx) error {
	id, _ := c.Locals("user_id").(string)

	user, err := h.svc.GetUserByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found", "success": true})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": nil,
		"data":    user,
	})
}
