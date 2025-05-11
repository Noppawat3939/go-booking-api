package mapper

import (
	"go_booking/internal/adapters/handler/dto"
	"go_booking/internal/core/domain"
)

func ToCreateHotelReq(r dto.CreateHotelRequest) domain.Hotel {
	return domain.Hotel{
		Name:        r.Name,
		Description: r.Description,
		Address:     r.Address,
		City:        r.City,
		State:       r.State,
		Country:     r.Country,
		PostCode:    r.PostCode,
		PhoneNumber: r.PhoneNumber,
		Email:       r.Email,
		Website:     r.Website,
		Active:      true,
	}
}

func ToCreateHotelRes(h *domain.Hotel) dto.CreateHotelResponse {
	return dto.CreateHotelResponse{
		ID:          h.ID,
		Name:        h.Name,
		Description: h.Description,
		Address:     h.Address,
		City:        h.City,
		State:       h.State,
		Country:     h.Country,
		PostCode:    h.PostCode,
		PhoneNumber: h.PhoneNumber,
		Email:       h.Email,
		Website:     h.Website,
	}
}

func ToListHotelRes(hotels []*domain.Hotel) []dto.Hotel {
	var res []dto.Hotel

	for _, h := range hotels {
		res = append(res, dto.Hotel{
			ID:          h.ID,
			Name:        h.Name,
			Description: h.Description,
			Address:     h.Address,
			City:        h.City,
			State:       h.State,
			Country:     h.Country,
			PostCode:    h.PostCode,
			PhoneNumber: h.PhoneNumber,
			Email:       h.Email,
			Website:     h.Website,
			CreatedAt:   h.CreatedAt,
			UpdatedAt:   h.UpdatedAt,
		})
	}

	return res
}

func ToHotelRes(h *domain.Hotel) dto.Hotel {
	return dto.Hotel{
		ID:          h.ID,
		Name:        h.Name,
		Description: h.Description,
		Address:     h.Address,
		City:        h.City,
		State:       h.State,
		Country:     h.Country,
		PostCode:    h.PostCode,
		PhoneNumber: h.PhoneNumber,
		Email:       h.Email,
		Website:     h.Website,
		CreatedAt:   h.CreatedAt,
		UpdatedAt:   h.UpdatedAt,
	}
}

func ToUpdateHotelReq(r dto.UpdateHotelRequest) domain.Hotel {
	return domain.Hotel{
		Name:        r.Name,
		Description: r.Description,
		Address:     r.Address,
		City:        r.City,
		State:       r.State,
		Country:     r.Country,
		PostCode:    r.PostCode,
		PhoneNumber: r.PhoneNumber,
		Email:       r.Email,
		Website:     r.Website,
		Active:      true,
	}
}
