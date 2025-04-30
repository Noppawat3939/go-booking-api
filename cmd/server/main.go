package main

import (
	"go_booking/cmd/config"
	"go_booking/internal/adapters/handler/router"
	"go_booking/internal/db"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	db, _ := db.ConnectDB()

	app := fiber.New()
	router.InitialRoutes(app, db)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Starting Fiber server on port", port)
	app.Listen(":8080")
}
