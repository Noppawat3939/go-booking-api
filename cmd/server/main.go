package main

import (
	"go_booking/cmd/config"
	"go_booking/internal/db"
	"go_booking/internal/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	db.Connect()

	app := fiber.New()

	routes.InitialRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Starting Fiber server on port", port)
	app.Listen(":8080")
}
