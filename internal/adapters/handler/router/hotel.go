package router

import (
	"go_booking/internal/adapters/handler/http"
	"go_booking/internal/adapters/repository"
	"go_booking/internal/core/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HotelRouter(router fiber.Router, db *gorm.DB) {
	// initialized Hotel
	repo := repository.NewHotelRepository(db)
	service := service.NewHotelService(repo)
	handler := http.NewHotelHandler(service)

	r := router.Group("/hotel")

	r.Post("/", handler.CreateHotel)
	r.Get("/", handler.GetHotels)
	r.Get("/:id", handler.GetHotel)
}
