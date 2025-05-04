package http

import (
	"go_booking/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	svc port.HotelService
}

func NewHotelHandler(svc port.HotelService) *HotelHandler {
	return &HotelHandler{svc}
}

func (h *HotelHandler) CreateHotel(ctx *fiber.Ctx) error {
	return nil
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
