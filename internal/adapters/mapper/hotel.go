package mapper

import (
	"go_booking/internal/adapters/handler/dto"
	"go_booking/internal/core/domain"
)

func ToCreateHotelReq(req dto.CreateHotelRequest) domain.Hotel {
	return domain.Hotel{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		City:        req.City,
		State:       req.State,
		Country:     req.Country,
		PostCode:    req.PostCode,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Website:     req.Website,
		Active:      true,
	}
}

func ToCreateHotelRes(hotel domain.Hotel) dto.CreateHotelResponse {
	return dto.CreateHotelResponse{
		ID:          hotel.ID,
		Name:        hotel.Name,
		Description: hotel.Description,
		Address:     hotel.Address,
		City:        hotel.City,
		State:       hotel.State,
		Country:     hotel.Country,
		PostCode:    hotel.PostCode,
		PhoneNumber: hotel.PhoneNumber,
		Email:       hotel.Email,
		Website:     hotel.Website,
	}
}
