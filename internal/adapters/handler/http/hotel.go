package http

import (
	"go_booking/internal/adapters/handler/dto"
	"go_booking/internal/adapters/mapper"
	"go_booking/internal/core/domain"
	"go_booking/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	svc port.HotelService
}

func NewHotelHandler(svc port.HotelService) *HotelHandler {
	return &HotelHandler{svc}
}

func (h *HotelHandler) CreateHotel(c *fiber.Ctx) error {
	var req dto.CreateHotelRequest

	if err := c.BodyParser(&req); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, domain.InvalidBodyMsg)
	}

	hotel := mapper.ToCreateHotelReq(req)

	if err := h.svc.CreateHotel(&hotel); err != nil {
		return ErrorResponse(c, fiber.StatusConflict, err.Error())
	}

	rsp := mapper.ToCreateHotelRes(hotel)

	return SuccessResponse(c, domain.CreatedDataMsg, rsp)
}

func (h *HotelHandler) GetHotel(ctx *fiber.Ctx) error {
	return nil
}

func (h *HotelHandler) GetHotels(ctx *fiber.Ctx) error {
	hotels, err := h.svc.FindAll(ctx)
	if err != nil {
		return ErrorResponse(ctx, fiber.StatusNotFound, "")
	}

	return SuccessResponse(ctx, "", hotels)
}
