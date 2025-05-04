package router

import (
	"go_booking/internal/adapters/handler/http"
	"go_booking/internal/adapters/handler/http/middleware"
	"go_booking/internal/adapters/repository"
	"go_booking/internal/core/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HotelRouter(router fiber.Router, db *gorm.DB) {
	// initialized Hotel
	rp := repository.NewHotelRepository(db)
	hs := service.NewHotelService(rp)
	hh := http.NewHotelHandler(hs)

	ts := service.NewTokenService()

	r := router.Group("/hotel")

	r.Post("/", middleware.JwtMiddleware(ts), hh.CreateHotel)
	r.Get("/", hh.GetHotels)
	r.Get("/:id", hh.GetHotel)
}
