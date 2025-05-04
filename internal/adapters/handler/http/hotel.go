package http

import (
	"go_booking/internal/adapters/handler/dto"
	"go_booking/internal/adapters/mapper"
	"go_booking/internal/core/domain"
	"go_booking/internal/core/port"
	"strconv"

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

	rsp := mapper.ToCreateHotelRes(&hotel)

	return SuccessResponse(c, domain.CreatedDataMsg, rsp)
}

func (h *HotelHandler) GetHotels(c *fiber.Ctx) error {
	hotels, err := h.svc.FindAll(c)
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, domain.DataNotFoundMsg)
	}

	rsp := mapper.ToListHotelRes(hotels)

	return SuccessResponse(c, domain.GettedDataMsg, rsp)
}

func (h *HotelHandler) GetHotel(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return ErrorResponse(c, fiber.StatusConflict, domain.ErrConflict.Error())
	}

	hotel, err := h.svc.FindOneByID(id)

	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, domain.DataNotFoundMsg)
	}

	rsp := mapper.ToHotelRes(hotel)

	return SuccessResponse(c, domain.GettedDataMsg, rsp)
}
