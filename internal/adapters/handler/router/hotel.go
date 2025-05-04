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
	hotelRepo := repository.NewHotelRepository(db)
	hotelService := service.NewHotelService(hotelRepo)
	hotelHandler := http.NewHotelHandler(hotelService)

	tokenService := service.NewTokenService()

	r := router.Group("/hotel")

	r.Use(middleware.JwtMiddleware(tokenService))

	r.Post("/", hotelHandler.CreateHotel)
	r.Get("/", hotelHandler.GetHotels)
	r.Get("/:id", hotelHandler.GetHotel)
}
