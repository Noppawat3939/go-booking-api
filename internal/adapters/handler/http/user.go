package http

import (
	"go_booking/internal/adapters/handler/dto"
	"go_booking/internal/adapters/mapper"
	d "go_booking/internal/core/domain"
	"go_booking/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{svc}
}

func (uh *UserHandler) Register(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, d.InvalidBodyMsg)
	}

	user := mapper.ToCreateUserReq(req)

	if err := uh.svc.Register(c, &user); err != nil {
		return ErrorResponse(c, fiber.StatusConflict, err.Error())
	}

	return SuccessResponse(c, d.CreatedDataMsg, mapper.ToCreateUserRes(&user))
}
