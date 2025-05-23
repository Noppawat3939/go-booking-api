package http

import (
	"go_booking/internal/adapters/handler/dto"
	"go_booking/internal/adapters/mapper"
	d "go_booking/internal/core/domain"
	"go_booking/internal/core/port"
	"go_booking/internal/core/util"
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
		return ErrorResponse(c, fiber.StatusBadRequest, d.InvalidBodyMsg)
	}

	hotel := mapper.ToCreateHotelReq(req)

	if err := h.svc.CreateHotel(&hotel); err != nil {
		return ErrorResponse(c, fiber.StatusConflict, err.Error())
	}

	return SuccessResponse(c, d.CreatedDataMsg, mapper.ToCreateHotelRes(&hotel))
}

func (h *HotelHandler) GetHotels(c *fiber.Ctx) error {
	_, limit, offset := util.GetPaginationParams(c)

	hotels, err := h.svc.FindAll(limit, offset)
	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, d.DataNotFoundMsg)
	}

	return SuccessResponse(c, d.GettedDataMsg, mapper.ToListHotelRes(hotels))
}

func (h *HotelHandler) GetHotel(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return ErrorResponse(c, fiber.StatusConflict, d.ErrConflict.Error())
	}

	hotel, err := h.svc.FindOneByID(id)

	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, d.DataNotFoundMsg)
	}

	return SuccessResponse(c, d.GettedDataMsg, mapper.ToHotelRes(hotel))
}

func (h *HotelHandler) UpdateByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return ErrorResponse(c, fiber.StatusConflict, d.ErrConflict.Error())
	}

	var req dto.UpdateHotelRequest

	if err := c.BodyParser(&req); err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, d.InvalidBodyMsg)
	}

	hotel := mapper.ToUpdateHotelReq(req)

	err = h.svc.UpdateByID(id, &hotel)

	if err != nil {
		return ErrorResponse(c, fiber.StatusNotFound, d.ErrConflict.Error())
	}

	return SuccessResponse(c, d.UpdatedDataMsg, nil)
}
